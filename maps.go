package main

import "fmt"

// map is used to store key-value pair.
func main() {
	age := make(map[string]int)
	age["arun"] = 27
	age["venkatram"] = 28
	fmt.Println("Arun's age is :", age["arun"])
	fmt.Println("stranger's age is :", age["unknownperson"]) // it returns 0 because age has no map for unknownperson
	// Similar to other go objects, maps also return 2 values
	arun, arunOK := age["arun"]
	venkat, venkatOK := age["venkat"]
	fmt.Println(arun, arunOK)
	fmt.Println(venkat, venkatOK)
	//We can also print length of a map
	fmt.Println("Length of age map is :", len(age))
	// Entries of a map can also be deleted.
	fmt.Println(age)
	delete(age, "venkatram")
	fmt.Println(age)
	// Adding entry again
	age["venkatram"] = 28
	// Output can also be formatted like awk
	for key, value := range age {
		fmt.Println(key, "==>", value)
	}

}
