package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var message string
	if len(os.Args) < 2 {
		message = " Gopher says Go!"
	} else {
		message = " " + os.Args[1]
	}

	rand.Seed(time.Now().UnixNano())

	cyan := getColor(14)
	yellow := getColor(215)
	white := getColor(15)
	black := getColor(0)
	reset := getReset()

	colors := []string{cyan, yellow, white}
	selectedColor := colors[rand.Intn(len(colors))]

	fmt.Printf("\n%s +------------------------+\n", white)
	fmt.Printf("%s |%s      ´.-::::::-.´%s      |\n", white, cyan, white)
	fmt.Printf("%s |%s  .:-::::::::::::::-:.%s  |\n", white, cyan, white)
	fmt.Printf("%s |%s  ´_::%s:    ::    :%s::_´%s  |\n", white, cyan, white, cyan, white)
	fmt.Printf("%s |%s   .:%s( ^   :: ^   )%s:.%s   |\n", white, cyan, white, cyan, white)
	fmt.Printf("%s |%s   ´::%s:   %s(%s..%s)%s   :%s::.%s   |\n", white, cyan, white, yellow, black, yellow, white, cyan, white)
	fmt.Printf("%s |%s   ´:::::::%sUU%s:::::::%s´%s   |\n", white, cyan, white, cyan, white, white)
	fmt.Printf("%s |%s   .::::::::::::::::.%s   |\n", white, cyan, white)
	fmt.Printf("%s |%s   O%s::::::::::::::::%sO%s   |\n", white, yellow, cyan, yellow, white)
	fmt.Printf("%s |%s   -::::::::::::::::- %s  |\n", white, cyan, white)
	fmt.Printf("%s |%s   ´::::::::::::::::´ %s  |\n", white, cyan, white)
	fmt.Printf("%s |%s    .::::::::::::::. %s   |\n", white, cyan, white)
	fmt.Printf("%s |%s      oO%s:::::::%sOo%s       |\n", white, yellow, cyan, yellow, white)
	fmt.Printf("%s +------------------------+\n", white)
	fmt.Printf("%s", reset)
	fmt.Printf("%s%s\n", selectedColor, message)
}

func getColor(color int) string {
	cmd := exec.Command("tput", "setaf", fmt.Sprint(color))
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func getReset() string {
	cmd := exec.Command("tput", "sgr0")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}
