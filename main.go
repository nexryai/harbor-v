package main

import (
	"os"
	"os/exec"
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

func execCmd(cmd string, arg string){
	fmt.Printf("\x1b[34m%s\n", "[EXEC]")
	fmt.Println("Command:", cmd, arg)

	out, err := exec.Command(cmd, arg).Output()


	if err != nil {
		msg := "Failed to execute " + cmd + "\n 	>>> " + err.Error()
		msgErr(msg)
	} else {
		fmt.Println(" ⌵\n ⌵\n ⌵")
		msg := "===Stdout===\n" + (string(out))
		msgInfo(msg)
	}


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

	if (distName == "debian") {
		buildDebian(distVersion, username)
	}

}

func buildDebian(debianVer string, username string) {
	msgInfo("Build debian container...")
	execCmd("zypper", "refresh")
}