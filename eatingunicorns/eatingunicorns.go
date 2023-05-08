// Package eatingunicorns is a package for eating unicorns
package eatingunicorns

// Unicorn is a struct for unicorns
type Unicorn struct {
	first   string
	age     int
	eatable bool
}

// EatUnicorn returns true if a unicorn can be eaten
func EatUnicorn(u Unicorn) bool {
	return u.eatable
}
