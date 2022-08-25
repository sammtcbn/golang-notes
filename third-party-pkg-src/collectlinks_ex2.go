package main

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
)

func download(url string) {
    resp, _ := http.Get(url)
    links := collectlinks.All(resp.Body)
    for _, link := range links {
        fmt.Println(link)
    }
}

func main() {
    url := "https://www.freecodecamp.org/"
    download (url)
}

/* result:

$ go run collectlinks_ex2.go
/learn
/news
/forum
/signin
/donate
/news/about/
https://www.linkedin.com/school/free-code-camp/people/
https://github.com/freeCodeCamp/
/news/shop/
/news/support/
/news/sponsors/
/news/academic-honesty-policy/
/news/code-of-conduct/
/news/privacy-policy/
/news/terms-of-service/
/news/copyright-policy/
/news/2019-web-developer-roadmap/
/news/best-python-tutorial/
/news/understanding-flexbox-everything-you-need-to-know-b4013d4dc9af/
/news/best-javascript-tutorial/
/news/python-example/
/news/best-html-html5-tutorial/
/news/linux-command-line-bash-tutorial/
/news/javascript-example/
/news/best-git-tutorial/
/news/best-react-javascript-tutorial/
/news/best-java-8-tutorial/
/news/the-best-linux-tutorials/
/news/best-css-and-css3-tutorial/
/news/the-best-jquery-examples/
/news/best-sql-database-tutorial/
/news/css-example-css3/
/news/react-examples-reactjs/
/news/best-angular-tutorial-angularjs/
/news/the-best-bootstrap-examples/
/news/the-ultimate-guide-to-ssh-setting-up-ssh-keys/
/news/best-wordpress-tutorial/
/news/the-best-php-examples/

*/
