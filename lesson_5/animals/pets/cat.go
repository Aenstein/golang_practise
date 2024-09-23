package pets

type Cat struct {
	Animal
}

func (c *Cat) Eat(amount uint8) (uint8, error) {
	if c.IsSleep {
		return 0, &ActionError{Name: c.Name, Reason: "it's sleep"}
	}
	
	if amount > 5 {
		return 0, newError("Cat can not eat that much", nil)
	}

	return amount, nil
}

func (c *Cat) Walk() string {
	return "Cat is walking!"
}
