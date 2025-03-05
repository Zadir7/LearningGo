package cmd

import (
	"fmt"
	"os"
)

/*

 */

func Execute() {
	//ENV Variables
	var env []string
	env = os.Environ() //slice!
	os.Setenv("NEWVAR", "val")
	fmt.Println(env[0]) // NEWVAR=val
	newvar, _ := os.LookupEnv("NEWVAR")
	fmt.Printf("%s", newvar) // val
}
