package main

import "fmt"

func main() {

	//One thing we can notice from its output is it sorts in alphabetical order of the map key
	//but not so much when we print it using the range loop
	//Maps are reference types, any changes that are made to the maps, makes an actual change to map, not on a copy
	//Good pratice to specify the size (especially if it is a big map)
	//Maps are not safe for concurrency

	newMap := make(map[string]string)
	newMap["Name"] = "Dennis"
	newMap["Surname"] = "Reynolds"
	newMap["Occupation"] = "Bartender"
	newMap["Education"] = "UPenn"

	// Prints info with above 4 keys

	fmt.Println("\n", newMap)

	//Output: map[Education:Univeristy of Pennsylvania Name:Dennis Occupation:Bartender Surname:Reynolds]

	// Adding a new key:value to the map

	newMap["Hometown"] = "Pennsylvania"

	fmt.Println("\n", newMap, "\n")

	//Output: map[Education:Univeristy of Pennsylvania Hometown:Pennsylvania Name:Dennis Occupation:Bartender Surname:Reynolds]

	//Just trying to emulate the describe command from SQL

	desc(newMap)

	//Deleting a key:value from a map

	delete(newMap, "Occupation")

	fmt.Println("\n", newMap)

	//Output: map[Education:Univeristy of Pennsylvania Hometown:Pennsylvania Name:Dennis Surname:Reynolds]

}

// How to pass a map to a function?

func desc(newMap map[string]string) {
	for mk, _ := range newMap {
		fmt.Println(mk)

	}

}

/* Output for desc(), the order of the map key alters
Occupation
Education
Hometown
Name
Surname
*/
