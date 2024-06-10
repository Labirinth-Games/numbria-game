package utils

import (
	"fmt"
	"strings"
	"time"
)

func SystemDialog(message string) {
	fmt.Printf("\033[0;31m[Sistema]: %s\033[0m\n", message)
}

func SystemMultiplyDialog(value []string, delay time.Duration) {
	for _, text := range value {
		SystemDialog(text)
		time.Sleep(delay * time.Millisecond)
	}
}

func NPCSay(message string) {
	fmt.Printf("\t\033[0;33mNPC: %s\033[0m\n", message)
}

func NarrationDialog(message string) {
	fmt.Printf("\t%s\n", message)
}

func NarrationMultiplyDialog(value []string, delay time.Duration) {
	for _, text := range value {
		NarrationDialog(text)
		time.Sleep(delay * time.Millisecond)
	}
}

func SpaceBlank() {
	fmt.Println("")
}

func DisplaySession(message string) {
	display := fmt.Sprintf(`
	========================================
		%s
	========================================
	`, message)

	fmt.Printf("\t\033[0;31m%s\033[0m\n", display)
}

func DisplayCommandIndicator(isBattle bool) {
	if isBattle {
		fmt.Print("\n[Modo Batalha]: ")
		return
	}

	fmt.Print("\n[Mundo]: ")
}

func TableItem(w string, size int) string {
	remaindSpace := size - len(w)
	var space []string

	for i := 0; i < remaindSpace; i++ {
		space = append(space, " ")
	}

	return w + strings.Join(space, "")
}
