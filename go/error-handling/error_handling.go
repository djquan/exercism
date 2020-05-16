package erratum

//Use opens a resource, calls Frob, and then closes the resource
func Use(o ResourceOpener, input string) (err error) {
	r, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		r, err = o()
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
