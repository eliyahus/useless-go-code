// Package eatingunicorns is a package for eating unicorns
package eatingunicorns

// Unicorn is a struct for unicorns
type Unicorn struct {
	first   string
	age     int
	eatable bool
}

// NewUnicorn creates a new unicorn. It decides whether it's eatable based on the age provided
func NewUnicorn(first string, age int) Unicorn {
	return Unicorn{
		first:   first,
		age:     age,
		eatable: age > 0,
	}
}

// Eatable returns true if a unicorn can be eaten
func Eatable(u Unicorn) bool {
	return u.eatable
}
