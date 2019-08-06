package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

var i = 1
var initUrls []string

type Info struct {
	Src   string `bson:"src"`
	Url   string `bson:"url"`
	Title string `bson:"title"`
}

var wg sync.WaitGroup

func main() {
	t1 := time.Now()

	wg.Add(2)

	ch := make(chan Info, 10)
	go getMeta(ch)
	go downloadPic(ch)

	wg.Wait()

	t2 := time.Now()
	fmt.Println("user time:", t2.Sub(t1))
}
func getMeta(ch chan Info) {
	for p := 1; p <= 20; p++ {
		url := fmt.Sprintf("https://www.zcool.com.cn?p=%d", p)
		initUrls = append(initUrls, url)
	}

	for _, url := range initUrls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("err: ", err)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("err: ", err)
		}

		doc.Find(".card-box").Each(func(i int, s *goquery.Selection) {
			src, _ := s.Find(".card-img-hover img").Attr("src")
			title := s.Find(".card-info .card-info-type").Text()
			url, _ := s.Find("a").Attr("href")
			ch <- Info{
				Src:   src,
				Title: title,
				Url:   url,
			}
		})
		resp.Body.Close()
	}
	close(ch)
	wg.Done()
}
func downloadPic(ch chan Info) {
	var dir = "pic1"
	os.Remove(dir)
	os.Mkdir(dir, os.ModePerm)

	for v := range ch {
		fmt.Println("download ... ", v.Title)

		//res, _ := http.Get(v.Src)
		//fileName := fmt.Sprintf("./%s/%d.png", dir, i)
		//file, _ := os.Create(fileName)
		//io.Copy(file, res.Body)
		//i++
	}
	wg.Done()
}
