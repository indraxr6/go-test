package main

import "fmt"

type Address struct {
	City, Prov, Region string
}

func updateRegion(region *Address) {
	region.Region = "Indonesia"
}

func main() {
	addr1 := Address{"Surabaya", "Jawa", "Indo"}
	addr2 := &addr1 //pass value

	addr2.City = "Jakarta"

	//pass reference
	*addr2 = Address{"Malang", "JawaTimur", "Indo"} //forcing all var who refer same memory to follow duplicate from pointer var (addr2)
	addr2 = &Address{"Malang", "JawaTimur", "Indo"} //make new Addr as separate var

	fmt.Println(addr1)
	fmt.Println(addr2)

	//make with fresh new data from struct
	var newAdress *Address = new(Address)
	newAdress.City = "Bogay"
	newAdress.Prov = "JawaBarat"
	newAdress.Region = "Indo"
	fmt.Println(newAdress)

	var locate = Address{
		City:   "Sorong",
		Prov:   "Papua",
		Region: "WOYY",
	}
	updateRegion(&locate)
	fmt.Println(locate)

}
