package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/browser"
	"os"
	"strings"
)

func main() {

	q := question()
	se := searchEngine()

	q = strings.ReplaceAll(q, " ", "+")

	var searchEngine string

	switch se {
	case "duckduckgo":
		sr := fmt.Sprint("http://"+se+".com", "/?q=")
		searchEngine = sr
	case "duck":
		sr := fmt.Sprint("http://"+se+".com", "/?q=")
		searchEngine = sr
	case "duckduckgo.com":
		sr := fmt.Sprint("http://" + se + "/?q=")
		searchEngine = sr
	case "google":
		sr := fmt.Sprint("http://"+se+".com", "/search?q=")
		searchEngine = sr
	case "google.com":
		sr := fmt.Sprint("http://" + se + "/search?q=")
		searchEngine = sr
	default:
		fmt.Printf("'%v' is not supported yet.", se)
	}

	// used for debugging
	// fmt.Println(find)

	// open url
	find := fmt.Sprint(searchEngine, q)
    err:=browser.OpenURL(find)
    if err!=nil{
        fmt.Println("error when opening browser: ", err)
    }

}

func question() string {
	fmt.Println("enter your question: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()
}

// TODO: allow specification of where results come from
// func resultsFrom() string {
// 	fmt.Println("enter the name of the sites you wish to get results from: ")
// 	scan := bufio.NewScanner(os.Stdin)
// 	scan.Scan()
// 	return scan.Text()
// }

func searchEngine() string {
	pse,present := os.LookupEnv("FAV_SE")
	if !present {
		fmt.Println("enter the name of your preffered search engine: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		return scan.Text()
	} else {
		return pse
	}
}
