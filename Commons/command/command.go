package command

type Command interface {
	IsValid() bool
	Execute() (interface{}, error)
}
