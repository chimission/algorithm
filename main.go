package main

import "algorithm/structure"

func main() {

	p := structure.InitDoubleLink(2)
	p.AddEndNode(3)
	p.AddFrontNode(1)
	p.Insert(4, 2)
	p.Print()

}
