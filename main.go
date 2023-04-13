package main

import (
	"1/t1"
	"bufio"
	"bytes"
)

func main() {
	s := bytes.Buffer{}
	s.WriteString(`3 3 12
DISABLE 1 2
DISABLE 2 1
DISABLE 3 3
GETMAX
RESET 1
RESET 2
DISABLE 1 2
DISABLE 1 3
DISABLE 2 2
GETMAX
RESET 3
GETMIN`)
	scanner := bufio.NewScanner(&s)
	t1.Kolya(scanner)
}
