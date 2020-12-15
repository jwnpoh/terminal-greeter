package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	// time.Now().Hour() gets the current hour
	// user.Current().username() gets username

	u, _ := user.Current()

	switch h := time.Now().Hour(); {
	case h >= 18: // 18 to 23
		fmt.Printf("Good evening, %s!\n", u.Username)
	case h >= 12: // 12 to 18
		fmt.Printf("Good afternoon, %s!\n", u.Username)
	case h >= 6: // 6 to 11
		fmt.Printf("Good morning, %s!\n", u.Username)
	default: // 0 to 5
		fmt.Printf("Good morning, %s! Why are you not sleeping??\n", u.Username)
	}
}
