package arinc

import (
	"fmt"
	"testing"
)

func TestLoadRunways(t *testing.T) {

	a, err := LoadRunways("arincdata-full.dat")
	if err != nil {
		t.Error(err)
		return
	}

	// for _, r := range a {
	// 	fmt.Printf("%#v\n", r)
	// }
	fmt.Println("total records ::::::::::::::::", len(a))

}
