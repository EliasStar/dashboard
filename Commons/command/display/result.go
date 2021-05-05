package display

type DisplayRst string

func (d DisplayRst) IsOK() bool {
	return true
}

func (d DisplayRst) Err() error {
	return nil
}
