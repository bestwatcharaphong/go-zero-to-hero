package formatter

type citizen struct{}

type Citizen struct {
	Firstname string
	firstname string //export
}

func (c *Citizen) PrintIDCard() { //export

}

func (c *Citizen) printIdCard() { //no export

}
