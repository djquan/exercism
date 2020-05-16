package erratum

//Use opens a resource, calls Frob, and then closes the resource
func Use(o ResourceOpener, input string) (err error) {
	var r Resource

	for r, err = o(); err != nil; r, err = o() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer r.Close()

	defer func() {
		if rec := recover(); rec != nil {
			if val, ok := rec.(FrobError); ok {
				r.Defrob(val.defrobTag)
			}

			err = rec.(error)
		}
	}()

	r.Frob(input)

	return nil
}
