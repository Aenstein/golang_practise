package pets

type Dog struct {
	Animal
}

func (c *Dog) Eat(amount uint8) (uint8, error) {
	if c.IsSleep {
		return 0, &ActionError{Name: c.Name, Reason: "it's sleep"}
	}

	if amount > 15 {
		return 0, newError("Dog can not eat that much", nil)
	}

	return amount, nil
}

func (c *Dog) Walk() string {
	return "Dog is walking!"
}
