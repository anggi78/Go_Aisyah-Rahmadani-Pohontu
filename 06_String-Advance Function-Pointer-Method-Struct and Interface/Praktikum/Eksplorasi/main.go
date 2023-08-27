package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	key := "zyxwvutsrqponmlkjihgfedcba"
	var nameEncode = ""

	for i := 0; i < len(s.name); i++ {
		if s.name[i] >= 'a' && s.name[i] <= 'z' {
			index := int(s.name[i] - 'a')
			nameEncode += string(key[index])
		} else {
			nameEncode += string(s.name[i])
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	key := "zyxwvutsrqponmlkjihgfedcba"
	var nameDecode = ""

	for i := 0; i < len(s.nameEncode); i++ {
		if s.nameEncode[i] >= 'a' && s.nameEncode[i] <= 'z' {
			index := strings.IndexRune(key, rune(s.nameEncode[i]))
			nameDecode += string(index + 'a')
		} else {
			nameDecode += string(s.nameEncode[i])
		}
	}
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		a.nameEncode = c.Encode()
		fmt.Print("\nEncode of student's name " + a.name + " is: " + a.nameEncode)
	} else if menu == 2 {
		fmt.Print("\nInput Encoded Name: ")
		fmt.Scan(&a.nameEncode)
		a.name = c.Decode()
		fmt.Print("\nDecode of encoded name " + a.nameEncode + " is: " + a.name)
	}
}
