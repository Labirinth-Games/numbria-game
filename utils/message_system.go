package utils

import (
	"fmt"
	"strings"
	"time"
)

func SystemDialog(message string) {
	fmt.Printf("\t\033[0;31mSistema: %s\033[0m\n\n", message)
}

func SystemMultiplyDialog(value []string, delay time.Duration) {
	for _, text := range value {
		SystemDialog(text)
		time.Sleep(delay * time.Millisecond)
	}
}

func NPCSay(message string) {
	fmt.Printf("\n\t\033[0;33mNPC: %s\033[0m\n\n", message)
}

func NarrationMultiplyDialog(value []string, delay time.Duration) {
	for _, text := range value {
		NarrationDialog(text)
		time.Sleep(delay * time.Millisecond)
	}
}

func NarrationDialog(message ...string) {
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
