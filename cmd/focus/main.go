//go:build windows

package main

import (
	"fmt"
	"os"
	"strconv"

	"focus/internal/session"
	"focus/internal/ui"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "status":
			s, err := session.Load()
			if err != nil {
				fmt.Println("No active focus session.")
				return
			}
			rem := s.Remaining()
			if rem <= 0 {
				fmt.Println("Session finished.")
				return
			}
			m := int(rem.Minutes())
			sec := int(rem.Seconds()) % 60
			fmt.Printf("Remaining: %dm %02ds\n", m, sec)
			return

		case "stop":
			if err := session.ForceStop(); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Session stopped. Hosts restored.")
			}
			return

		default:
			_, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Usage: focus [minutes | status | stop]")
				return
			}
		}
	}

	if !ui.IsAdmin() {
		ui.RelaunchAsAdmin()
		return
	}

	ui.RunDashboard()
}
