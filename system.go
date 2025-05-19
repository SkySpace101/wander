package main

import (
	"fmt"
	"log"

	"github.com/zcalusic/sysinfo"
)

func main() {
	si := sysinfo.SysInfo{}
	if err := si.GetSysInfo(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("System Info: %+v\n", si)
}
