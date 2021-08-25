package rss

import (
	"fmt"
	"time"

	"deas_rss/models"
	"github.com/gobuffalo/nulls"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

func convertFeed(feed *gofeed.Feed, feedURL string) (articles models.Articles, err error) {
	// TODO: work out what to do when feedUrl != feed.Link

	if len(feed.Items) == 0 {
		return nil, errors.New("no articles in feed")
	}
	for _, item := range feed.Items {
		articles = append(articles, convertFeedItem(*item, feedURL))
	}

	return articles, nil
}

func convertFeedItem(item gofeed.Item, feedURL string) models.Article {
	return models.Article{
		FeedURL:     feedURL,
		Guid:        itemGUID(item, feedURL),
		Title:       item.Title,       // TODO: change model and convert "" → nil
		Description: item.Description, // TODO: change model and convert "" → nil
		Content:     item.Content,     // TODO: change model and convert "" → nil
		Link:        item.Link,        // TODO: change model and convert "" → nil
		ImageURL:    itemImageURL(item),
		PublishedAt: itemPublishedAt(item),
	}
}

func itemGUID(item gofeed.Item, feedURL string) string {
	if item.GUID != "" {
		return item.GUID
	}
	if item.Link != "" {
		return item.Link
	}
	return fmt.Sprintf("%s-%s", feedURL, item.Published)
}

func itemPublishedAt(item gofeed.Item) time.Time {
	if item.PublishedParsed != nil {
		return *item.PublishedParsed
	}
	return time.Now()
}

func itemImageURL(item gofeed.Item) nulls.String {
	if item.Image != nil && item.Image.URL != "" {
		nulls.NewString(item.Image.URL)
	}
	// TODO: Try to find an <img> tag in Description or Content
	return nulls.String{}
}
