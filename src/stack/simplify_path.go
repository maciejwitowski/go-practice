package stack

import (
	"strings"
)

/*
Given an absolute path for a Unix-style file system, which begins with a slash '/', transform this path into its simplified canonical path.
In Unix-style file system context, a single period '.' signifies the current directory, a double period ".." denotes moving up one directory level, and multiple slashes such as "//" are interpreted as a single slash. In this problem, treat sequences of periods not covered by the previous rules (like "...") as valid names for files or directories.
The simplified canonical path should adhere to the following rules:

It must start with a single slash '/'.
Directories within the path should be separated by only one slash '/'.
It should not end with a slash '/', unless it's the root directory.
It should exclude any single or double periods used to denote current or parent directories.
Return the new path.

Example 1:
	Input: path = "/home/"
	Output: "/home"
	Explanation:
		The trailing slash should be removed.
*/

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")

	s := stack{}

	for _, p := range parts {
		switch p {
		case "", ".":
			continue
		case "..":
			s.pop()
		default:
			s.push(p)
		}
	}

	var b strings.Builder

	for _, el := range s.elements {
		b.WriteString("/" + el)
	}

	if b.Len() == 0 {
		b.WriteString("/")
	}
	return b.String()
}

type stack struct {
	elements []string
}

func (s *stack) push(el string) {
	s.elements = append(s.elements, el)
}

func (s *stack) pop() {
	if len(s.elements) == 0 {
		return
	}
	s.elements = s.elements[:len(s.elements)-1]
	return
}
