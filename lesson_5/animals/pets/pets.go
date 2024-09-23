package pets

import "fmt"

type Eater interface {
	Eat(amount uint8) (uint8, error)
}

type Walker interface {
	Walk() string
}

type Named interface {
	GetName() string
}

type EaterWalker interface {
	Eater
	Walker
	Named
}

type Animal struct {
	Name    string
	Age     uint8
	Weight  uint8
	IsSleep bool
}

func (c *Animal) GetName() string {
	return c.Name
}

type ActionError struct {
	Name   string
	Reason string
}

func (e *ActionError) Error() string {

	return fmt.Sprintf("%s can not perfom the action: %s", e.Name, e.Reason)
}

func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return fmt.Errorf(msg)
}
