package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please Input Some Lowercase String:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputStr := scanner.Text()
		upperStr := strings.ToUpper(inputStr)
		fmt.Println(upperStr)
	}
}
