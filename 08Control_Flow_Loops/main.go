//Control Flow: Loops
// For Loop
// Nested Loop Control
// Loop control statements like Break, Continue and GoTo

package main

import "fmt"

func main() {
	// 1* For Loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 2* Nested Loop Control
	// Multiplication Table
	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= 5; j++ {
	// 		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	// 	}
	// 	fmt.Println()
	// }
	// Loop control Statements
	// A Break
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// B Continue
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// C GoTo
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 		if i == 5 {
	// 			goto end
	// 		}
	// 	}
	// end:
	// 	fmt.Println("Loop ended")

	// Second example of Goto
	// 	i := 0
	// start:
	// 	if i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 		goto start
	// 	}
	// Infinite Loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }
	// Basic way ?
start:
	fmt.Println("This is an infinite loop!")
	goto start
}

// Basic
// 10 PRINT "This is an infinite loop!"
// 20 GOTO 10
