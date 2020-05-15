package erratum

//Use opens a resource, calls Frob, and then closes the resource
func Use(o ResourceOpener, input string) (err error) {
	var r Resource

	if r, err = o(); err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}

	defer func() {
		if rec := recover(); rec != nil {
			if val, ok := rec.(FrobError); ok {
				r.Defrob(val.defrobTag)
			}

			err = rec.(error)
		}

		r.Close()
	}()

	r.Frob(input)

	return
}
