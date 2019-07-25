package util

import (
	"bufio"
	"os"
)

var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() string {
	line, err := inputReader.ReadString('\n')
	Check(err)
	return line
}