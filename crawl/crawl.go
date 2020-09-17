package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	// "io/ioutil"
	"log"
	"net/http"
	// 	"regexp"
	// 	"strings"
)

func download(url string, c chan string) {

	res, ok := http.Get(url)
	if nil != ok {
		return
	}
	defer res.Body.Close()
	// str, ok := ioutil.ReadAll(res.Body)
	// if ok != nil {
	// 	return
	// }
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := doc.Find("img#validateImage").Attr("src")

	// fmt.Printf("%s", t)
	// t := doc.Text()
	fmt.Printf("from links: %s\n ", t)

}

// var href_reg *regexp.Regexp
// var hrefs_been_found map[string]int
// var hrefs_undone []string

// func get_all_href(url string) []string {
// 	var ret []string
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ret
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	hrefs := href_reg.FindAllString(string(body), -1)
// 	// fmt.Println(hrefs)
// 	for _, v := range hrefs {
// 		str := strings.Split(v, "\"")[1]
// 		// fmt.Printf("%s \n", str)
// 		link_list := strings.Split(str, "/")
// 		// fmt.Println(len(link_list), link_list, "0:"+link_list[0])
// 		if len(link_list) == 3 && link_list[1] == "item" {
// 			ret = append(ret, "http://www.bootstrapmb.com"+str)
// 		}
// 	}
// 	return ret
// }
// func init_global_var() {
// 	href_pattern := "href=\"(.+?)\""
// 	href_reg = regexp.MustCompile(href_pattern)
// 	hrefs_been_found = make(map[string]int)
// }
// func is_href_been_found(href string) bool {
// 	_, ok := hrefs_been_found[href]
// 	return ok
// }
// func add_hrefs_to_undone_list(hrefs []string) {
// 	for _, value := range hrefs {
// 		ok := is_href_been_found(value)
// 		if !ok {
// 			fmt.Printf("new url:(%s)\n", value)
// 			hrefs_undone = append(hrefs_undone, value)
// 			hrefs_been_found[value] = 1
// 		} else {
// 			hrefs_been_found[value]++
// 		}
// 	}
// }

func Test() {
	// url := "https://ks.sac.net.cn/sac/manage/yzm_xvcode.jsp"
	url := "http://person.amac.org.cn/pages/login/exam_query.html"
	c := make(chan string)
	// for i := 0; i < 100; i++ {
	// 	download(url)
	// }
	go func() {
		for i := 0; i < 10; i++ {
			download(url, c)
		}
		close(c)

	}()
	for range c {
		<-c
	}

	// init_global_var()
	// var pos = 0
	// var urls = []string{"http://www.bootstrapmb.com/muban?page=162"}
	// add_hrefs_to_undone_list(urls)
	// for {
	// 	if pos >= len(hrefs_undone) {
	// 		break
	// 	}
	// 	if len(hrefs_undone) > 100 {
	// 		break
	// 	}
	// 	url := hrefs_undone[0]
	// 	hrefs_undone = hrefs_undone[1:]
	// 	hrefs := get_all_href(url)
	// 	add_hrefs_to_undone_list(hrefs)
	// }
}
