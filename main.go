package main

import (
	"log"

	"github.com/probeldev/aerospace-autokill/aerospace"
	"github.com/probeldev/aerospace-autokill/processinfo"
)

func main() {
	windows, err := aerospace.GetAllWindows()
	if err != nil {
		log.Println(err)
	}

	log.Println(windows)

	process, err := processinfo.GetProcessByName("Chrome")
	if err != nil {
		log.Println(err)
	}

	log.Println(process)

	pidsForKill := []string{}

	for _, p := range process {
		isSetWndows := false
		for _, w := range windows {
			if p.PID == w.Pid {
				isSetWndows = true
			}
		}

		if !isSetWndows {
			pidsForKill = append(pidsForKill, p.PID)
		}
	}

	log.Println(pidsForKill)
}
