package main

import "fmt"

//deklarasi interface
type Room interface {
	dataRoom() // interface memiliki kontrak function
}

type Luxury struct {
	Name, Location string
	Price          int32
}

type FirstClass struct {
	Name, Location string
	Price          int32
}

//deklarasi struct function dengan interface 1
func (luxury Luxury) dataRoom() {
	fmt.Println("Name :", luxury.Name)
	fmt.Println("Location :", luxury.Location)
	fmt.Println("Price : $", luxury.Price, "/ Night")

}

//deklarasi struct function dengan interface 2
func (firstclass FirstClass) dataRoom() {
	fmt.Println("Name :", firstclass.Name)
	fmt.Println("Location :", firstclass.Location)
	fmt.Println("Price : Rp.", firstclass.Price, "/ Night")

}

func main() {
	// pengisian data struct 1
	lx := Luxury{
		Name:     "Clinton Hall",
		Location: "Boston",
		Price:    980,
	}

	// pengisian data struct 2
	fs := FirstClass{
		Name:     "JW Marriot",
		Location: "New York",
		Price:    3200,
	}

	//menampilkan data
	lx.dataRoom()
	fmt.Println(" ")
	fs.dataRoom()
}
