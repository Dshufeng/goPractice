package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var i = 1
var initUrls []string

type Info struct {
	Src   string `bson:"src"`
	Url   string `bson:"url"`
	Title string `bson:"title"`
}

var sliceInfo []Info

func main() {
	t1 := time.Now()
	picInfo := getMeta()
	downloadPic(picInfo)

	t2 := time.Now()
	fmt.Println("user time:", t2.Sub(t1))
}
func getMeta() []Info {
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
			sliceInfo = append(sliceInfo, Info{
				Src:   src,
				Title: title,
				Url:   url,
			})
		})
		resp.Body.Close()
	}
	return sliceInfo
}
func downloadPic(picInfo []Info) {
	var dir = "pic"
	os.Remove(dir)
	os.Mkdir(dir, os.ModePerm)

	for _, v := range picInfo {
		fmt.Println("download ... ", v.Title)
		//
		//res, _ := http.Get(v.Src)
		//fileName := fmt.Sprintf("./%s/%d.png", dir, i)
		//file, _ := os.Create(fileName)
		//io.Copy(file, res.Body)
		//i++
	}
}
