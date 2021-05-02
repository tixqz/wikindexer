package main

import (
	"errors"
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	WIKI_DOMAIN      string = "https://en.wikipedia.org"
	TEST_START_PAGE  string = "Hellsing"
	TEST_TARGET_PAGE string = "Evangelion"
)

type WikiArticleTree struct {
}

type ArticleNode struct {
	previous *ArticleNode
	level    int
	isTarget bool
}

func (an *ArticleNode) NewArticleNode(previous *ArticleNode, level int, isTarget bool) *ArticleNode {
	return &ArticleNode{
		previous: previous,
		level:    level,
		isTarget: isTarget,
	}
}

type LinksPool struct {
	pages map[string]string
}

func (lp *LinksPool) NewPagesPool(pages map[string]string) *LinksPool {
	return &LinksPool{
		pages: pages,
	}
}

func (lp *LinksPool) VerifyTarget() bool {
	_, hasTarget := lp.pages[TEST_TARGET_PAGE]

	return hasTarget
}

func (lp *LinksPool) CleanStartFromPool() {
	delete(lp.pages, TEST_START_PAGE)
}

func main() {
	startingLink := WIKI_DOMAIN + "/wiki/" + TEST_START_PAGE

	article, err := htmlquery.LoadURL(startingLink)
	if err != nil {
		fmt.Printf("Can't load wiki page. Error: %v", err)
	}

	fmt.Println(ParseAllLinks(article))
}

func GetNextArticle() {

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

func CheckStartAndTargetPagesExist(start, target string) bool {
	fmt.Printf("Check wiki articles for %s and %s exist.", start, target)
	return wikiExists(start) && wikiExists(target)
}

func wikiExists(url string) bool {
	return true
}
