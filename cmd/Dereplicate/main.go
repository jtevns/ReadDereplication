package main

import (
	"os"
)

func main() {
	fastqparse.parseSingle(os.Args[1])
}
