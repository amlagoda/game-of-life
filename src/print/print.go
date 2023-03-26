package print

import (
	"fmt"
	"strings"
)

func Printed(gen [][]bool, num int, cnt int) string {
	prt := strings.Builder{}
	prt.WriteString(fmt.Sprintf("Generation: %d\n", num))
	prt.WriteString(fmt.Sprintf("Alive: %d\n\n", cnt))

	for i := range gen {
		for j := range gen[i] {
			if gen[i][j] {
				prt.WriteString("O")
			} else {
				prt.WriteString(" ")
			}
		}
		prt.WriteString("\n")
	}

	return prt.String()
}
