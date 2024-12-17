package main

import "fmt"

func main() {
	h := NewHyperLogLog(3)

	i := 49498
	h.Ingest(int32(i))
	fmt.Println(h.Cardinality())
	h.Print()

	fmt.Println("---")

	i = 18597
	h.Ingest(int32(i))
	fmt.Println(h.Cardinality())
	h.Print()
}
