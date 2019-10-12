package main

import (
	"flag"
	"fmt"
)

func main() {
	wordptr := flag.String("word", "defaultword", "a string")
	numptr := flag.Int("number", 100, "an int")
	boolptr := flag.Bool("fork", true, "boolean value")
	var svar string
	flag.StringVar(&svar, "svar", "default_word", "a string") // Another way of using flags.
	flag.Parse()
	fmt.Println("word: ", *wordptr)
	fmt.Println("age: ", *numptr)
	fmt.Println("flag: ", *boolptr)
	fmt.Println("svar:", svar)
	fmt.Println("Other Arguments passed:", flag.Args())
}

// gostudy[master !?+] $ ./cli_flags -word="Arun Venkatram" -number=27 -svar="from bangalore" -fork=true testing other arguments
// word:  Arun Venkatram
// number:  27
// flag:  true
// svar: from bangalore
// Other Arguments passed: [testing other arguments]
// gostudy[master !?+] $

// gostudy[master !?+] $ ./cli_flags -h
// Usage of ./cli_flags:
//   -fork
//         boolean value (default true)
//   -number int
//         an int (default 100)
//   -svar string
//         a string (default "default_word")
//   -word string
//         a string (default "defaultword")
