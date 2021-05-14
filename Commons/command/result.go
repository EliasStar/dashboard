package command

type Result interface {
	IsOK() bool
}

func NewResultFromError(err error) Result {
	if err == nil {
		return OKRst{}
	}

	return ErrorRst(err.Error())
}

type OKRst struct{}

func (o OKRst) IsOK() bool {
	return true
}

type ErrorRst string

func (e ErrorRst) IsOK() bool {
	return false
}
