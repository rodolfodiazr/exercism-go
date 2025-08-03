// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb returns a list of proverbs given the set of words.
func Proverb(rhyme []string) []string {
    lastIndex := len(rhyme) - 1
    output :=[]string{}
	for i := range rhyme {
        if i == lastIndex {
            output = append(output, "And all for the want of a " + rhyme[0] + ".")
            continue
        }

        output = append(output, "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost.")
    }
    return output
}
