package command

type DisplayCmd struct {
	URL string
}

func (d DisplayCmd) IsValid() bool {

	return false
}

func (d DisplayCmd) Execute() (interface{}, error) {

	return nil, nil
}
