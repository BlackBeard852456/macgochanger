package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func execCommand(name string, attributes []string) {
	cmd := exec.Command(name, attributes...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("err")
	}
}

func main() {
	//var netWorkInterfaceName string = getNetworkInterfaceName()
	nameOfNetworkInterface := flag.String("i", "eth0", "Nom de l'interface réseau sur lequel appliqué les changements")
	newMacAddress := flag.String("m", "", "Nouvelle addresse Mac")
	flag.Parse()
	execCommand("sudo", []string{"ifconfig", *nameOfNetworkInterface, "down"})
	execCommand("sudo", []string{"ifconfig", *nameOfNetworkInterface, "hw", "ether", *newMacAddress})
	execCommand("sudo", []string{"ifconfig", *nameOfNetworkInterface, "up"})
}
