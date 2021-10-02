package modules

type Module struct {
	color string
}

func (m *Module) AddColor(str string) string {
	return "%F{" + m.color + "}" + str + "%f"
}

type IModule interface {
	Init()
	GetOutput() string
	AddColor(str string) string
}
