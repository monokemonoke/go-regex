package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/monokemonoke/go-regex/regex"
)

func main() {
	for {
		res, ok := scanPattern()
		if !ok {
			break
		}

		fmt.Println(regex.Search(res[0], res[1]))
	}
}

func scanPattern() ([2]string, bool) {
	var res [2]string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("pattern > ")
	if ok := scanner.Scan(); !ok {
		return res, false
	}
	res[0] = scanner.Text()

	fmt.Printf("text    > ")
	if ok := scanner.Scan(); !ok {
		return res, false
	}
	res[1] = scanner.Text()

	return res, true
}
