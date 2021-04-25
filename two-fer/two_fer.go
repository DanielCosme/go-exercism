// Package twofer helps sharing
package twofer

import "fmt"

// ShareWith will share with you or me
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
