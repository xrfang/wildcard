package main

import (
	"fmt"
	"github.com/xrfang/wildcard"
)
func main() {
	p1 := wildcard.Pattern("hello*")
    p2 := wildcard.Pattern("hello*!")
	p3 := wildcard.Pattern("?hello")
	s1 := "hello world"
	s2 := "hello world!"
	s3 := "Xhello"
	fmt.Printf("%q matches %q: %v\n", p1, s1, p1.Match(s1))
	fmt.Printf("%q matches %q: %v\n", p1, s2, p1.Match(s2))
	fmt.Printf("%q matches %q: %v\n", p1, s3, p1.Match(s3))
	fmt.Printf("%q matches %q: %v\n", p2, s1, p2.Match(s1))
	fmt.Printf("%q matches %q: %v\n", p2, s2, p2.Match(s2))
	fmt.Printf("%q matches %q: %v\n", p2, s3, p2.Match(s3))
	fmt.Printf("%q matches %q: %v\n", p3, s1, p3.Match(s1))
	fmt.Printf("%q matches %q: %v\n", p3, s2, p3.Match(s2))
	fmt.Printf("%q matches %q: %v\n", p3, s3, p3.Match(s3))
	if p3.Match("hello") {
		fmt.Println("'?' matches zero character")
	} else {
		fmt.Println("'?' does NOT match zero character")
	}
}