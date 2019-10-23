package hello

import "fmt"

// Interface ...
type Interface interface {
	SayHello(FirstName string) string
}

// Hello ...
type Hello struct {
	LastName string
}

var h *Hello
var hi Interface

// New ...
func New(i Interface) Interface {
	hi = i
	if hello, ok := i.(*Hello); ok {
		h = hello
	}
	return hi
}

// SayHello ...
func (h *Hello) SayHello(firstName string) string {
	r := fmt.Sprintf("Hello, %s %s!", h.LastName, firstName)
	r = r + h.SayWorld()
	fmt.Println(r)
	return r
}

// SayWorld ...
func (h *Hello) SayWorld() string {
	return " world!"
}

// SayHello ...
func SayHello(firstName string) string {
	return hi.SayHello(firstName)
}
