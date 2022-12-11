package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	result := strings.Contains("I love Go Programming!", "lovex")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO PyTHON jaVA"))
	p(strings.ToUpper("GO PyTHON jaVA"))

	p("go" == "go")
	p("GO" == "go")
	p(strings.ToLower(("GO")) == strings.ToLower("go"))
	p(strings.EqualFold("GO", "go"))
}
