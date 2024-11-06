package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jvcosta-dev/goclijokes/internal"
)

func main() {
	argType := flag.String("type", "Any", "Joke type")
	argLang := flag.String("lang", "en", "Response language")
	argBlacklistFlags := flag.String("blacklist", "", "Blacklist of Flags that you dont want to see jokes about")
	argRange := flag.String("range", "", "Id range ex: 1-10")
	argAmount := flag.String("amount", "1", "Amount of jokes")

	flag.Parse()

	jokes, err := internal.GetJokes(*argType, *argLang, *argBlacklistFlags, *argRange, *argAmount)
	if err != nil {
		log.Fatal(err)
	}

	for i, joke := range jokes.Jokes {
		fmt.Printf("Joke #%d\n", i+1)
		fmt.Printf("%-50s\n", joke.Setup)
		fmt.Printf("%-50s\n", joke.Delivery)
		fmt.Println()
	}

}
