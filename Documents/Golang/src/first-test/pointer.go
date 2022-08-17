package main

import "fmt"

type Addr struct {
	City, Prov, Region string
}

func main() {
	addr1 := Addr{"Surabaya", "Jawa", "Indo"}
	addr2 := &addr1 //pass value

	addr2.City = "Jakarta"

	//pass reference
	*addr2 = Addr{"Malang", "JawaTimur", "Indo"} //forcing all var who refer same memory to follow duplicate from pointer var (addr2)
	addr2 = &Addr{"Malang", "JawaTimur", "Indo"} //make new Addr as separate var

	fmt.Println(addr1)
	fmt.Println(addr2)

	//make with fresh new data from struct
	var newAdress *Addr = new(Addr)
	newAdress.City = "Bogay"
	newAdress.Prov = "JawaBarat"
	newAdress.Region = "Indo"
	fmt.Println(newAdress)
}
