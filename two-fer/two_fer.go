package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	if len(name) != 0 {
		return fmt.Sprintf("One for %s, one for me.", name)
	}

	return "One for you, one for me."
}
