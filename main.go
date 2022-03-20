// toxiquote demonstrates an issue with rendering "''" as "”" (this won't
// seem like such a big deal until you look at the source code).
//
// It gets worse though, since it also happens in <pre> blocks where users
// may well be copying text to use.
//
//  $ go run github.com/kortschak/toxiquote
//  ''
//
package main

import "fmt"

func main() {
	fmt.Println("''")
}
