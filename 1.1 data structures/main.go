package main

import (
	"fmt"
	"sort"
)

func main(){

	// pointer()
	// arrays()
	// slices()
	// maps()
	// structs()

}

func arrays(){

	//arrays cant be used to order or modifyed durring run time
	// an array size must be declared
	var colors[3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int {5,3,1,2,4}
	fmt.Println(numbers)
	
	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}

// changes to a slice effects its parent array as well as all its sibeling slices
func slices(){
	// an slices is an array just without its size declared
	var colors = [] string {"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	// to remove items from a slice you use:
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// using make([]int, 5 ,5) the slice is limited to 5 items leaving out the last number removes the limiter make([]int, 5)
	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 146
	fmt.Println(numbers)
	
	numbers = append(numbers, 236)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}

// maps are reference types. So changes will only effect there function scope ( there soft copys )
func maps(){

	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "Califronia"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")

	states["NY"] = "New York"
	fmt.Println(states)

	for key, value := range states {
		fmt.Printf("%v: %v\n", key, value)
	}

	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}

	fmt.Println(keys)
	// sort the keys
	sort.Strings(keys)
	fmt.Println(keys)


	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}


func structs(){

	// structs can be seen as classes only like other langues the cannot inhert properties or methods the are all inviduals

	poodle := Dog{"Poodle", 12}
	fmt.Println(poodle)
	fmt.Printf("%v\n", poodle)
	fmt.Printf("Breed:%v\nWieght:%v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 10
	fmt.Printf("Breed:%v\nWieght:%v\n", poodle.Breed, poodle.Weight)
}

type Dog struct {
	// by using a Uppercase starting letter a variable becomes part of the global scope "its exported"
	Breed string
	Weight int
}