package command

type LaunchCmd struct {
	Executable string
	Arguments  []string
}

func (l LaunchCmd) IsValid() bool {

	return false
}

func (l LaunchCmd) Execute() (interface{}, error) {

	return nil, nil
}
