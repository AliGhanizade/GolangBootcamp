package week_6

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)
var url string = "https://www.film2movie.asia/"
func getWebTitle(wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body := response.Body

	reTitle := regexp.MustCompile(`<title>(.*?)</title>`)
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		title := reTitle.FindStringSubmatch(line)
		if len(title) > 1 {
			fmt.Println("Title of the web page is:", title[1])

		}
	}

}
func getWebLink(wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body := response.Body

	reLink := regexp.MustCompile("href=\"(http[s]?://.*?)\"")
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		link := reLink.FindStringSubmatch(line)
		if len(link) > 1 {
			fmt.Println("Link of the web page is:",link[1])

		}
	}

}
