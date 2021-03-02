package main

import (
	"fmt"
	"io/ioutil"
    "net/http"
    "log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client
//http.Request is a data structure that represents the client HTTP request. 
//r.URL.Path is the path component of the request URL
// [1:] means "create a sub-slice of Path from the 1st character to the end.
//func handler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    //Path is re-sliced with [len("/view/"):] to drop the leading "/view/" component of the request path
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}


//http package to handle all requests to the web root ("/") with handler.
func main() {
    http.HandleFunc("/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    //ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
    log.Fatal(http.ListenAndServe(":8080", nil))
}