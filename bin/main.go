package main

import (
	cont "github.com/expproletariy/pip-timers-service/container"
	"os"
)

func main() {
	proc := cont.NewTimeSessionProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)
}
