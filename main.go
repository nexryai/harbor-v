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
	fmt.Println(" ⌵\n ⌵\n===Stdout===")


	args := strings.Split(arg, " ")

	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	err := command.Run()


	if err != nil {
		msg := "Failed to execute " + cmd + "\n 	>>> " + err.Error()
		msgErr(msg)
		os.Exit(1)
	}

	fmt.Printf("\x1b[0m%s\n", "\n")
}

func execInContainer(cmd string, container string){
	execArg := "-D " + "/var/lib/machines/" + container + " " + cmd
	execCmd("systemd-nspawn", execArg)
}


func main() {
	msgInfo("Welcome to Harbor-V !")

	if os.Geteuid() != 0 {
		msgErr("harbor-v MUST be run as root !!!")
		os.Exit(1)
	}

	flag.Parse()
	var dist = flag.Arg(0)
	var containerName = flag.Arg(1)
	var username = flag.Arg(2)
	var mvInterface = flag.Arg(3)

	if len(mvInterface) == 0 {
		msgErr("There are not enough arguments. \nusage: ./harbor debian.bullseye [container_name] [username] [network_interface_name_for_container]")
		os.Exit(1)
	}

	if !strings.Contains(dist, ".") {
		msgErr("The first argument must be in the form of \"[distribution-name].[distribution-version]\" \neg. \"debian.bullseye\" ")
		os.Exit(1)
	}

	arr := strings.Split(dist, ".")
	var distName = arr[0]
	var distVersion = arr[1]

	//var nspawnPath = "/var/lib/machines"

	fmt.Println("Distribution:", distName, " Version:", distVersion, " container_name:", containerName, " User:", username, " interface:", mvInterface)

	if (distName == "debian") {
		write_NetConf(containerName, mvInterface)
		buildDebian(distVersion, containerName, username)
	}

	configNetworkd(containerName, mvInterface)

}



func write_NetConf(containerName string, mvInterface string){

	msgInfo("Generate a nspawn config file...")

	var configFile_path = "/etc/systemd/nspawn/" + containerName + ".nspawn"
	var configFile = "[Network]\nMACVLAN=" + mvInterface

	file, err := os.Create(configFile_path)
	if err != nil {
		var msg = "Failed to create " + configFile_path
		msgErr(msg)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(configFile)
	if err != nil {
		var msg = "Failed to write to " + configFile_path
		msgErr(msg)
		os.Exit(1)
	}

}

func configNetworkd(containerName string, mvInterface string){

	msgInfo("Generate a systemd-networkd config file...")

	var configFile_path = "/var/lib/machines/" + containerName + "/etc/systemd/network/mv-" + mvInterface + ".network"
	var configFile = "[Match]\nName=mv-" + mvInterface + "\n\n[Network]\nDHCP=yes"

	file, err := os.Create(configFile_path)
	if err != nil {
		var msg = "Failed to create " + configFile_path
		msgErr(msg)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(configFile)
	if err != nil {
		var msg = "Failed to write to " + configFile_path
		msgErr(msg)
		os.Exit(1)
	}

}

func buildDebian(debianVer string, containerName string, username string) {
	msgInfo("Build debian container...")

	machineDir := "/var/lib/machines/" + containerName

	if err := os.Mkdir(machineDir, 0755); err != nil {
		msg := "Failed to create directory for container \n 	>>>" + err.Error()
        msgErr(msg)
		os.Exit(1)
    }

	execArg := "--arch=amd64 --include=systemd,dbus,openssh-server,ufw,sudo " + debianVer + " " + machineDir + " https://ftp.riken.jp/Linux/debian/debian/"
	execCmd("debootstrap", execArg)

	execInContainer("passwd", containerName)

	fmt.Println("Set up a root account. Please enter your password.")
	command := "adduser " + username
	execInContainer(command, containerName)

	fmt.Println("Set up your account. Please enter password.")
	command = "gpasswd -a " + username + " sudo"
	execInContainer(command, containerName)

	command = "systemctl enable systemd-networkd"
	execInContainer(command, containerName)

}