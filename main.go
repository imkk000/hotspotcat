package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	var servFlag string
	flag.StringVar(&servFlag, "services", "dhcpcd.service,hostapd.service,dnsmasq.service", "list of services")
	flag.Parse()

	services := strings.Split(servFlag, ",")

	fmt.Println("service start")
	time.Sleep(10 * time.Second)
	fmt.Printf("services: %v\n", services)

	for _, s := range services {
		check(s)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done
	fmt.Println("service stop")
}

func check(name string) {
	restartService(name)
	time.Sleep(2 * time.Second)
	for checkService(name) != nil {
		restartService(name)
	}
	fmt.Printf("check %s: %v\n", name, checkService(name))
}

func checkService(name string) error {
	return exec.
		Command("systemctl", "status", name).
		Run()
}

func restartService(name string) error {
	return exec.
		Command("systemctl", "restart", name).
		Run()
}
