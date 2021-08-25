package grifts

import (
	"fmt"

	"deas_rss/rss"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

// Ex: https://www.feedforall.com/sample.xml
// Ex: http://allpozitive.ru/feed/
// Ex: http://humanstory.ru/good-news?format=feed&amp;type=atom

var _ = grift.Namespace("rss", func() {
	grift.Desc("fetch", "Fetch RSS feed with given URL")
	grift.Add("fetch", func(c *grift.Context) error {
		if len(c.Args) >= 1 {
			for _, url := range c.Args {
				if err := rss.ParseURL(url) ; err != nil {
					return errors.Wrapf(err, "grit fetch RSS '%s'", url)
				}

			}
			return nil
		}

		fmt.Printf("Usage: buffalo task rss:fetch {URL} [URL]...\n")
		// TODO: fetch every URL; update all
		return nil
	})
})
