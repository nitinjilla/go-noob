// Method Declaration

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
	sv = sv.IncrementMajor()
	sv = sv.IncrementMajor()
	sv = sv.IncrementMajor()
	fmt.Println(sv.String())
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)

}

func (sv SemanticVersion) IncrementMajor() SemanticVersion {
	sv.major += 1
	fmt.Println(&sv.major)
	return sv

}

func (sv SemanticVersion) IncrementMinor() SemanticVersion {
	sv.minor += 1
	fmt.Println(&sv)
	return sv

}

func (sv SemanticVersion) IncrementPatch() SemanticVersion {
	sv.patch += 1
	fmt.Println(&sv)
	return sv

}

/*
func String(sv SemanticVersion) string {

	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)

}*/
