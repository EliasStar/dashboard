package display

type Result string

func (r Result) IsOK() bool {
	return true
}
