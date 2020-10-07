/*Package twofer provides string manipulation
Ex. If parameter Name is entered as Amit then output should be
"One for Alex, one for me" OR
"One for you, one for me"
*/
package twofer

// ShareWith does basic string manipulation
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return "One for " + name + ", one for me."

}
