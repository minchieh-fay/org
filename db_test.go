package main

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	fmt.Println("123")
	t.Log("yes")
	a := dbe.getUserByLoginID("admin1")
	t.Log(a)
	fmt.Println("123-------------", a)
}
