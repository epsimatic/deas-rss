package rss

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gobuffalo/pop/v5"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

func ParseURL(url string) error {
	log.Printf("Fetching URL: %s", url)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return errors.Wrapf(err, "parse '%s'", url)
	}
	fmt.Println(feed.Title)

	articles, err := convertFeed(feed, url)
	if err != nil {
		return errors.Wrapf(err, "convert '%s'", url)
	}

	tx, err := pop.Connect("development")
	if err != nil {
		return errors.Wrap(err, "connect to DB")
	}

	for _, a := range articles {
		//spew.Dump(a)
		log.Printf("Storing item '%s'", a.Guid)
		if err := a.ValidateAndUpsertByGUID(tx); err != nil {
			return errors.Wrapf(err, "store '%s'", a.Guid)
		}
	}
	return nil
}


func parseString(feedData string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(feedData)
	fmt.Println(feed.Title)
	spew.Dump(feed)
	return err
}
