package lib

func GoFunctions(f ...func() (err error)) (err error) {
	ch := make(chan error)
	for i := 0; i < len(f); i++ {
		i := i
		go func() {
			errr := f[i]()
			if errr != nil {
				ch <- errr
				return
			}
			ch <- nil
		}()
	}
	for i := 0; i < len(f); i++ {
		err = <-ch
		if err != nil {
			return
		}
	}
	return
}
