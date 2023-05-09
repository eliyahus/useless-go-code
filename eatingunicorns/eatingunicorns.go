// Package eatingunicorns is a package for eating unicorns.
package eatingunicorns

// Unicorn is a struct for unicorns.
type Unicorn struct {
	first string
}

// NewUnicorn creates a new unicorn.
func NewUnicorn(first string) Unicorn {
	return Unicorn{
		first: first,
	}
}

// Eatable returns true if a unicorn can be eaten.
func Eatable(u *Unicorn) bool {
	return u != nil
}

// Name returns the firstname of the unicorn.
func (u Unicorn) Name() string {
	return u.first
}
