package main

import "fmt"

func main() {
	// cara deklarasi pertama
	var one, two, three string = "ini satu", "ini dua", "ini tiga"
	// cara kedua tanpa typedata (kayak python)
	four, five, six := 4, "lima", 6.6
	fmt.Println("ini pake cara println one: ", one, "two: "+two, "three:", three)
	fmt.Printf("ini pake printf four: %d  five: %s  six: %f\n", four, five, six)

	// predefined
	_ = "data ini tidak bisa ditampilkan dan tidak berguna :) kayak kamu ehehe"

	// ini pointer
	contoh_pointer := new(string)
	println("isi dari contoh_pointer", contoh_pointer)
	println("isi dari *contoh_pointer", *contoh_pointer)
	*contoh_pointer = "ini si dari *contoh_pointer ketika di assign nilai"
	println("isi dari *contoh_pointer", *contoh_pointer)
}
