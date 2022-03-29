package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
)

func msgInfo(msg string){
	fmt.Fprintln(os.Stdout, msg)
}

func msgErr(err string){
	fmt.Println("\n")
	fmt.Printf("\x1b[31m%s\n", "===ERROR===")
	fmt.Fprintln(os.Stderr, err)
	fmt.Printf("\x1b[0m%s\n", "\n")

}

func main() {
	msgInfo("Welcome to harbor-v !")

	if os.Geteuid() != 0 {
		msgErr("harbor-v MUST be run as root !!!")
		os.Exit(1)
	}

	flag.Parse()
	var dist = flag.Arg(0)
	var username = flag.Arg(1)

	if !strings.Contains(dist, ".") {
		msgErr("The first argument must be in the form of \"[distribution-name].[distribution-version]\" \neg. \"debian.bullseye\" ")
		os.Exit(1)
	}

	arr := strings.Split(dist, ".")
	var distName = arr[0]
	var distVersion = arr[1]

	fmt.Println("Distribution:", distName, " Version:", distVersion, " User:", username)
}

func buildDebian(dist string, username string) {
	fmt.Println("Build debian container...")
}