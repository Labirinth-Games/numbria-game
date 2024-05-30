package utils

import (
	"fmt"
	"strings"
)

func SystemSay(message string) {
	fmt.Printf("\n\t\033[0;31mSistema: %s\033[0m\n\n", message)
}

func NPCSay(message string) {
	fmt.Printf("\n\t\033[0;33mNPC: %s\033[0m\n\n", message)
}

func UniverseSay(message ...string) {
	if len(message) == 1 {
		fmt.Printf("\n\t%s\n", message[0])
		return
	}

	if len(message) > 1 {
		if strings.Contains(message[0], "%s") {
			m := fmt.Sprintf(message[0], message[1])
			fmt.Printf("\n\t%s\n", m)

			return
		}

		fmt.Printf("\n\t%s\n", message[0])
	}
}

func SpaceBlank() {
	fmt.Println("")
}
