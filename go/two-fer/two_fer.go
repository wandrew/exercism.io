// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith reads a string input, checks if it is empty, if it is, assigns "you" to name and outputs a string
func ShareWith(name string) string {
	// Mentor feedback: it's really only looking for 0
	if len(name) == 0 {
		name = "you"
	}

	// Mentor feedback, using specific types for Sprintf protects against undesired behavior if this code is refactored.
	return fmt.Sprintf("One for %s, one for me.", name)
}
