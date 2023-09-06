//WORKING WITH POINTERS

package main

import (
	"fmt"
)

var name = "Dennis" //Global variable; can be declared and not be used. Doesn't work the same way for local variables.

func main() {

	songname := "dosed_rhcp_dwnld.mp4"
	fmt.Printf("/Music/%s \n", songname)
	renameFile(&songname)
	fmt.Println("File renamed to", songname)

}

func renameFile(songname *string) string {
	*songname = "Dosed.mp4"
	return *songname

}