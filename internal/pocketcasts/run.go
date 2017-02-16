// +build ignore

package main

import (
	"log"

	"github.com/hdonnay/podcasts/internal/pocketcasts"
)

func main() {
	log.SetFlags(log.LUTC | log.Lshortfile | log.LstdFlags)

	c, err := pocketcasts.New("hdonnay@gmail.com", "tgolTsJUgEa8")
	if err != nil {
		log.Fatal(err)
	}
	a, err := c.All()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(a)
}
