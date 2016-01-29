package main

import (
	"fmt"

	"github.com/pkg/profile"
	"github.com/technopoetic/iindex"
)

var index map[string][]int

func main() {
	p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	// initialize representation
	index, _ = iindex.MakeIndex("/home/rhibbitts/Dropbox/Notes/WorkNotes")
	// run  user interface
	ui()
	p.Stop()
}

func ui() {
	fmt.Println(len(index), "words indexed in", len(iindex.Indexed), "files")
	fmt.Println("enter single words to search for")
	fmt.Println("enter a blank line when done")
	var word string
	for {
		fmt.Print("search word: ")
		wc, _ := fmt.Scanln(&word)
		if wc == 0 {
			return
		}
		switch dl := iindex.Index[word]; len(dl) {
		case 0:
			fmt.Println("no match")
		case 1:
			fmt.Println("one match:")
			fmt.Println("   ", iindex.Indexed[dl[0]].File, iindex.Indexed[dl[0]].Title)
		default:
			fmt.Println(len(dl), "matches:")
			for _, d := range dl {
				fmt.Println("   ", iindex.Indexed[d].File, iindex.Indexed[d].Title)
			}
		}
	}
}
