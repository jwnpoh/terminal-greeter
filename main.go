package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"math/rand"
	"os/user"
	"time"
)

func main() {
	// time.Now().Hour() gets the current hour
	// user.Current().username() gets username

	u, _ := user.Current()

	rand.Seed(time.Now().UnixNano())

	var msg string

	switch h := time.Now().Hour(); {
	case h >= 18: // 18 to 23
		msg = "\rGood evening, %s!\n"
	case h >= 12: // 12 to 18
		msg = "\rGood afternoon, %s!\n"
	case h >= 6: // 6 to 11
		msg = "\rGood morning, %s!\n"
	default: // 0 to 5
		msg = "\rGood morning, %s! Why are you not sleeping??\n"
	}

	r := rand.Intn((6))
	switch r {
	case 1:
		fmt.Print(aurora.Sprintf(aurora.Red(msg), aurora.Red(u.Username)))
	case 2:
		fmt.Print(aurora.Sprintf(aurora.Green(msg), aurora.Green(u.Username)))
	case 3:
		fmt.Print(aurora.Sprintf(aurora.Blue(msg), aurora.Blue(u.Username)))
	case 4:
		fmt.Print(aurora.Sprintf(aurora.Yellow(msg), aurora.Yellow(u.Username)))
	case 5:
		fmt.Print(aurora.Sprintf(aurora.Magenta(msg), aurora.Magenta(u.Username)))
	case 6:
		fmt.Print(aurora.Sprintf(aurora.Cyan(msg), aurora.Cyan(u.Username)))
	case 0:
		fmt.Print(aurora.Sprintf(aurora.White(msg), aurora.White(u.Username)))
	}
}
