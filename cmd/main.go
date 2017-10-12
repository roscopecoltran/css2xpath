package main

import (
	"fmt"
	"github.com/roscopecoltran/css2xpath"
)

func main() {
	ss := []string{
		"div",
		"> div",
		"div, > span",
		"div.foo",
		"div.foo.bar",
		"div#foo",
		"div#foo.bar#hux",
		"> div#foo.bar#hux",
		".bar",
		":first-child",
		"div:first-child",
		"div:nth-child(odd)",
		"div:nth-child(even)",
		"div:nth-child(2n + 1)",
		"div:nth-child(-3n-6)",
		"div:nth-of-type(5)",
		":nth-child(4)",
		"div :nth-child(2)",
		"div[a='b']",
		"div[a~='b']",
		"div[a|='b']",
		"div[a*='b']",
		"div[a ^= 'b' ]",
		"div   [ a $= 'b' ]",
		"> :only-of-type",
		"div[a='b']:first-of-type.foo", // generates xpath that behaves differently from the css
		"div.bar:not(#foo:first-child)",
	}
	//
	for _, s := range ss {
		fmt.Println("css: ", s, " >>>> xpath [GLOBAL]:", css2xpath.Convert(s, css2xpath.GLOBAL))
		fmt.Println("css: ", s, " >>>> xpath [LOCAL]:", css2xpath.Convert(s, css2xpath.LOCAL))
		fmt.Print("\n")
	}
}
