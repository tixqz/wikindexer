package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	WIKI_DOMAIN     string = "https://en.wikipedia.org"
	TEST_START_PAGE string = "Hellsing"
	TEST_END_PAGE   string = "Evangelion"
)

type PageDocument *html.Node

type WikiNode struct {
	ID           int
	Title        string
	PreviousPage *WikiNode
	Last         bool
}

type ArticleNode struct {
	links   []*LinksPool
	level   int
	hasLast bool
}

func (an *ArticleNode) NewArticleNode(level int) *ArticleNode {
	return &ArticleNode{
		level: level,
	}
}

type LinksPool struct {
	pages map[string]string
}

func (lp *LinksPool) NewPagesPool(level int) *LinksPool {
	return &LinksPool{}
}

func main() {
	article, err := htmlquery.LoadURL(WIKI_DOMAIN + TEST_START_PAGE)
	if err != nil {
		fmt.Printf("Can't load wiki page. Error: %v", err)
	}

	fmt.Println(ParseAllLinks(article))
}

func ParseAllLinks(doc PageDocument) (map[string]string, error) {
	refs := make(map[string]string)

	res, err := htmlquery.QueryAll(doc, "//a[@href]")
	if err != nil {
		return nil, err
	}
	for _, node := range res {
		fmt.Print(node)
	}

	return refs, nil
}

func keepOnlyPageContent(doc PageDocument) {}
