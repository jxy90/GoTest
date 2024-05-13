package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func Test_intro(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	defer cancel()
	r := Search(ctx, "golang")
	if r.Err != nil {
		log.Println(r.Err)
	} else {
		defer r.Body.Close()
		fmt.Fprintf(os.Stdout, "%v\n %v\n", r.Source, r.Status)
		io.Copy(os.Stdout, r.Body)
	}
}

func Search(ctx context.Context, keyword string) *Result {
	bingUrl := bingSearchURL(keyword)
	sogouUrl := sougouSearchURL(keyword)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bingRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, bingUrl.String(), nil)
	if err != nil {
		log.Panic(err)
	}
	sogouRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, sogouUrl.String(), nil)
	if err != nil {
		log.Panic(err)
	}

	ch := make(chan *Result)
	go func() {
		ch <- NewResult(http.DefaultClient.Do(bingRequest))
	}()
	go func() {
		ch <- NewResult(http.DefaultClient.Do(sogouRequest))
	}()
	r := <-ch
	cancel()
	//如果第二个goroutine回来了，要做到管生管养
	failure := <-ch
	if failure.Body != nil {
		//读取并抛弃这个body。还要关闭
		io.Copy(io.Discard, failure.Body)
		failure.Body.Close()
	}
	return r
}

func bingSearchURL(keyword string) *url.URL {
	u, err := url.Parse("https://bing.com/search")
	if err != nil {
		log.Panic(err)
	}
	q := make(url.Values)
	q.Add("q", keyword)
	u.RawQuery = q.Encode()
	return u
}

func sougouSearchURL(keyword string) *url.URL {
	u, err := url.Parse("https://sogou.com/web")
	if err != nil {
		log.Panic(err)
	}
	q := make(url.Values)
	q.Add("query", keyword)
	u.RawQuery = q.Encode()
	return u
}

// The result of a search request.
type Result struct {
	Source string
	Status string
	Body   io.ReadCloser
	Err    error
}

func NewResult(r *http.Response, err error) *Result {
	if err != nil {
		return &Result{Err: err}
	}
	return &Result{
		Source: r.Request.URL.Host,
		Status: r.Status,
		Body:   r.Body,
	}
}
