// Package twofer provides function ShareWith to 
// generate "one for you/<name>, one for me." string
package twofer

import "fmt"

// ShareWith accepts <name> string and returns string 
// in the format "one for <name>, one for me." if the 
// name is empty then returns the same string but with 
// "you" instead of <name>
func ShareWith(name string) string {
    msgTemplate := "One for %s, one for me."
    if len(name) != 0 {
        return fmt.Sprintf(msgTemplate, name)
    }
	return fmt.Sprintf(msgTemplate, "you")
}
