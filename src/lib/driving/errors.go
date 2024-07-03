package driving

type err struct {
	name string
}

func (err) isError() {}

func ErrNotFound() err {
	return err{"ErrNotFound"}
}

func ErrBadRequest() err {
	return err{"ErrBadRequest"}
}

func ErrDriven() err {
	return err{"ErrDriven"}
}
