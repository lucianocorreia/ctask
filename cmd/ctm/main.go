package main

import (
	"fmt"
	"os"

	"github.com/lucianocorreia/ctm/pkg/commands"
)

func main() {
	if err := commands.RootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
