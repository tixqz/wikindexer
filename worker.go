package main

import (
	"os"
	"sync"
	"time"
)

type Dispatcher struct {
	level        int
	wg           sync.WaitGroup
	workersCount int
	sleepTime    time.Duration
	freeWorkers  chan *Worker
	articlesPool chan *ArticleNode
	found        chan *ArticleNode
	interrupt    chan os.Signal
}

func NewDispatcher(workersCount int, sleepTime time.Duration, found chan *ArticleNode, interrupt chan os.Signal) *Dispatcher {
	d := &Dispatcher{
		level:        0,
		workersCount: workersCount,
		sleepTime:    sleepTime,
		found:        found,
		interrupt:    interrupt,
	}

	d.freeWorkers = make(chan *Worker)

	for i := 0; i <= workersCount; i++ {
		worker := NewWorker()
		d.freeWorkers <- worker
	}

	return d
}

func (d *Dispatcher) Run() {
	for {
		select {
		case w := <-d.freeWorkers:
			d.wg.Add(1)
			go w.Run()
		case <-d.interrupt:
			d.Stop()
			return
		}
	}
}

func (d *Dispatcher) UpdateLevel() {
	d.level += 1
}

func (d *Dispatcher) Submit(article *ArticleNode) {
	d.articlesPool <- article
}

func (d *Dispatcher) Stop() {
	d.wg.Wait()
	close(d.freeWorkers)
	close(d.found)
	close(d.interrupt)
}

type Worker struct {
	dis    *Dispatcher
	input  chan *ArticleNode
	out    chan *ArticleNode
	failed chan<- *ArticleNode
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) Run(sleep time.Duration) {
	defer w.dis.wg.Done()
	article := <-w.input
	node, _ := article.LoadNode()
	parsedPages, _ := ParseAllLinks(node)
	pool := NewLinksPool(parsedPages)
	pool.CleanStartFromPool()

	hasTarget := pool.VerifyTarget()
	if hasTarget {

	}

	for nextTitle, nextUrl := range pool.pages {
		NewArticleNode(nextUrl, nextTitle)
	}
	time.Sleep(sleep)
}
