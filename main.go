package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func getNetworkInterfaceName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entre le nom de l'interface réseau :")
	scanner.Scan()
	return scanner.Text()
}

func execCommand(name string, attributes []string) {
	cmd := exec.Command(name, attributes...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("err")
	}
}

func main() {
	//var netWorkInterfaceName string = getNetworkInterfaceName()
	execCommand("sudo", []string{"ifconfig", "enp5s0", "down"})
	execCommand("sudo", []string{"ifconfig", "enp5s0", "hw", "ether", "00:11:22:33:44:55:66"})
	execCommand("sudo", []string{"ifconfig", "enp5s0", "up"})
}
