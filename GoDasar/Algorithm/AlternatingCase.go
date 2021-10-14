package main

import (
	"fmt"
	"strings"
)

func ToAlternatingCase(str string) string  {
	temp := strings.Split(str,"")

	var res string
	for i:=0; i <= len(temp)-1;i++ {

		upper := strings.ToUpper(temp[i])
		lower := strings.ToLower(temp[i])

		if temp[i] == upper {
		    res += lower
		} else if temp[i] == lower {
		  res += upper
		} else {
		 res += temp[i]
		}

	}
	return res


}
 func main()  {

fmt.Print(ToAlternatingCase("1a2b3c4d5e"))
}

//Expect(ToAlternatingCase("hello world")).To(Equal("HELLO WORLD"))
//Expect(ToAlternatingCase("HELLO WORLD")).To(Equal("hello world"))
//Expect(ToAlternatingCase("hello WORLD")).To(Equal("HELLO world"))
//Expect(ToAlternatingCase("HeLLo WoRLD")).To(Equal("hEllO wOrld"))
//Expect(ToAlternatingCase("12345")).To(Equal("12345"))
//Expect(ToAlternatingCase("1a2b3c4d5e")).To(Equal("1A2B3C4D5E"))
//Expect(ToAlternatingCase("String.prototype.toAlternatingCase")).To(Equal("sTRING.PROTOTYPE.TOaLTERNATINGcASE"))
//Expect(ToAlternatingCase(ToAlternatingCase("Hello World"))).To(Equal("Hello World"))
//Expect(ToAlternatingCase("altERnaTIng cAsE")).To(Equal("ALTerNAtiNG CaSe"))
//Expect(ToAlternatingCase("ALTerNAtiNG CaSe")).To(Equal("altERnaTIng cAsE"))
//Expect(ToAlternatingCase("altERnaTIng cAsE <=> ALTerNAtiNG CaSe")).To(Equal("ALTerNAtiNG CaSe <=> altERnaTIng cAsE"))
//Expect(ToAlternatingCase("ALTerNAtiNG CaSe <=> altERnaTIng cAsE")).To(Equal("altERnaTIng cAsE <=> ALTerNAtiNG CaSe"))
