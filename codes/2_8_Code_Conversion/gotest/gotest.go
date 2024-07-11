package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://jpub.tistory.com/category/출간%20전%20책%20소식"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	titles := []string{}
	excerpts := []string{}

	document.Find(".post-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		excerpt := s.Find(".excerpt").Text()

		// 정규표현식을 이용해 100자가 넘을 경우, 97자만 남기고 뒤에 ...을 붙입니다
		re := regexp.MustCompile("^(.{97}).*")
		excerpt_new := re.ReplaceAllString(excerpt, "$1...")

		titles = append(titles, strings.TrimSpace(title))
		excerpts = append(excerpts, strings.TrimSpace(excerpt_new))
	})

	fmt.Println("Title\tExcerpt")

	for i := 0; i < len(titles); i++ {
		fmt.Printf("%s\t%s\n", titles[i], excerpts[i])
	}
}
