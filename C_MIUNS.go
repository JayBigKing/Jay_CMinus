package main

import (
	"../CMinus/CMIUNS"
)

func main() {
	var c CMIUNS.C_MIUNS
	c.Init("codingFile2.txt")
	c.Scan()
	if c.Parse() == true {
		CMIUNS.PrintTree("parseTree")
	}

}
