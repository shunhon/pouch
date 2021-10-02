package modules

type Cmd struct {
	Module
}

func (m *Cmd) Init() {
	m.color = "green"
}

func (m *Cmd) GetOutput() string {
	return "\n➾ "
}
