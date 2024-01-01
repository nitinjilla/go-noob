// Method Declaration with pointers

package main

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {

	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}

}

func main() {
	sv := NewSemanticVersion(1, 1, 1)
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	sv.IncrementPatch()
	fmt.Println(sv.String())
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)

}

func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
	fmt.Println(&sv.major)
}

func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
	fmt.Println(&sv)
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
	fmt.Println(&sv)
}

/*1.1.1
sv.IncrementMajor() 0xc00001a018
sv.IncrementMajor() 0xc00001a018
sv.IncrementMajor() 0xc00001a018
4.1.1*/


/*1.1.1
sv.IncrementMajor() 0xc0000a2000
sv.IncrementMajor() 0xc0000a2000
sv.IncrementMajor() 0xc0000a2000
sv.IncrementMinor() 0xc00009a020
sv.IncrementMinor() 0xc00009a028
sv.IncrementPatch() 0xc00009a030
sv.IncrementPatch() 0xc00009a038
4.3.3

*/