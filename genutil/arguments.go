package genutil

import (
	"os"
	"strings"
)

func App() string {
	return os.Args[0]
}

func NonFlagArgs() (args []string) {
	return nonFlagArgs(os.Args[1:])
}

func nonFlagArgs(argsIn []string) (argsOut []string) {
	for _, argIn := range argsIn {
		if !strings.HasPrefix(argIn, "-") {
			argsOut = append(argsOut, argIn)
		}
	}
	return
}
