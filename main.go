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
	TEST_TARGET_PAGE string = "Neon Genesis Evangelion"
)

var foundTarget chan *ArticleNode

type WikiArticleTree struct {
}

type ArticleNode struct {
	url      string
	title    string
	previous *ArticleNode
}

func NewArticleNode(url, title string, previous *ArticleNode) *ArticleNode {
	return &ArticleNode{
		title:    title,
		previous: previous,
	}
}

type LinksPool struct {
	pages map[string]string
}

func NewPagesPool(pages map[string]string) *LinksPool {
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
	targetLink := WIKI_DOMAIN + "/wiki/" + TEST_TARGET_PAGE

	if !CheckStartAndTargetPagesExist(startingLink, targetLink) {
		fmt.Println("Wiki page for one of objects does not exist.")
	}

	if !CheckStartAndTargetPagesNotSame(startingLink, targetLink) {
		fmt.Println("Starting and target are same wiki pages.")
	}

	foundTarget = make(chan *ArticleNode, 1)
	defer close(foundTarget)

	go FindTarget(startingLink, TEST_START_PAGE, nil)

	targetNode, ok := <-foundTarget

	if ok {
		BuildPathToTarget(targetNode)
	} else {
		fmt.Printf("Didn't found target: %s", TEST_TARGET_PAGE)
	}
}

// FindTarget is the main function for finding target article.
func FindTarget(url, title string, prev *ArticleNode) {
	currentArticle := NewArticleNode(url, title, prev)
	node, _ := htmlquery.LoadURL(WIKI_DOMAIN + url)
	parsedPages, _ := ParseAllLinks(node)
	pool := NewPagesPool(parsedPages)
	pool.CleanStartFromPool()

	hasTarget := pool.VerifyTarget()
	if hasTarget {
		foundTarget <- currentArticle
		return
	}

	for nextTitle, nextUrl := range pool.pages {
		go FindTarget(nextUrl, nextTitle, currentArticle)
	}
}

// ParseAllLinks gets all links from article's body and map titles of links with url paths.
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

func BuildPathToTarget(node *ArticleNode) {

}

func CheckStartAndTargetPagesNotSame(startingLink, targetLink string) bool {
	return true
}

func CheckStartAndTargetPagesExist(start, target string) bool {
	fmt.Printf("Checking wiki articles for %s and %s exist...", start, target)
	return wikiExists(start) && wikiExists(target)
}

func wikiExists(url string) bool {
	return true
}
