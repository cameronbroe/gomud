package main

import (
	"fmt"
	"github.com/cameronbroe/gomud/pkg"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Hello, this is gomud")
	router := pkg.NewRouter()
	router.AddRoute(regexp.MustCompile("help"), func() {
		fmt.Println("Got help command!")
	}).AddRoute(regexp.MustCompile("attack [a-z]+"), func() {
		fmt.Println("Got attack command!")
	})
	bot := pkg.NewBot(os.Getenv("DISCORD_TOKEN")).InstallRouter(router)

	bot.Start()
}
