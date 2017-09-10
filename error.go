package tackle

type ApplicationError struct {
	Err error `json:"error,omitempty"`
}

func (ae ApplicationError) Error() error { println("fdsf"); return ae.Err }
