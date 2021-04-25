package display

type DisplayRst struct {
	Error error
}

func (d DisplayRst) OK() bool {
	return d.Error == nil
}

func (d DisplayRst) Err() error {
	return d.Error
}
