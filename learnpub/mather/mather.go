package mather

type Runner interface {
	Run() int
}

type Walker interface {
	Walk() int
}

type Track interface {
	Runner
	Walker
}

type Athlete struct {
	Speed int
	Steps int
}

func (a Athlete) Run() int {
	return a.Speed * 2
}

func (a Athlete) Walk() int {
	return a.Steps * 1
}
