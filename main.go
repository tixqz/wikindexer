package main

import (
	"errors"
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	WIKI_DOMAIN     string = "https://en.wikipedia.org"
	TEST_START_PAGE string = "Hellsing"
	TEST_END_PAGE   string = "Evangelion"
)

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
	startingLink := WIKI_DOMAIN + "/wiki/" + TEST_START_PAGE

	article, err := htmlquery.LoadURL(startingLink)
	if err != nil {
		fmt.Printf("Can't load wiki page. Error: %v", err)
	}

	fmt.Println(ParseAllLinks(article))
}

func GetArticle() {

}

func ParseAllLinks(doc *html.Node) (map[string]string, error) {
	refs := make(map[string]string)
	res, err := htmlquery.QueryAll(doc, "//*[@id='mw-content-text']/div[1]/p/a")
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("no results after quering")
	}
	for _, node := range res {
		// Every link element on wiki page equals this
		// <a href="/wiki/<Article>" title="Article">Article</a>
		refs[node.Attr[1].Val] = node.Attr[0].Val
	}

	return refs, nil
}
