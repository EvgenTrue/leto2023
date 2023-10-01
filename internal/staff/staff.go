package staff

type Chief struct {
	Name string
}


func (c *Chief) GetName() string {
	return c.Name
}
func (c *Chief) Cook() string {
	return " "
}
func (c *Chief) Order() string {
	return c.Name
}

type Waiter struct {
	Name string
}
func (w *Waiter) GetName() string {
	return w.Name
}

func (w *Waiter) Order() string {
	return " "
}