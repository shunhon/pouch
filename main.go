package main

import (
	"fmt"

	"github.com/shunhon/pouch/modules"
)

func loadModules() []modules.IModule {
	moduleList := []modules.IModule{}
	moduleList = append(moduleList, &modules.Cwd{})
	// Cmd must be last module
	moduleList = append(moduleList, &modules.Cmd{})
	for i := 0; i < len(moduleList); i++ {
		moduleList[i].Init()
	}
	return moduleList
}

func main() {
	prompt := "autoload -Uz colors\n"
	prompt += "colors\n"
	prompt += "PROMPT=\"\n"
	moduleList := loadModules()
	for i := 0; i < len(moduleList); i++ {
		module := moduleList[i]
		out := module.GetOutput()
		prompt += module.AddColor(out)
		prompt += " "
	}
	prompt += "\""
	fmt.Println(prompt)
}
