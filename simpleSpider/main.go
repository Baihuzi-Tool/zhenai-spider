package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code is ", resp.StatusCode)
	}
	//e := determineEncoding(resp.Body)
	//reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printCityAndUrl(all)
}
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
func printCityAndUrl(content []byte) {
	re := regexp.MustCompile(`<a [^href]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]*)</a>`)
	submatch := re.FindAllSubmatch(content, -1)
	for _, m := range submatch {
		fmt.Printf("City:%s Url:%s\n", m[2], m[1])
	}
	fmt.Println("len", len(submatch))
}
