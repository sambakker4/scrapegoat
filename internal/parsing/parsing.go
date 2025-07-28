package parsing

import (
	"errors"
	"fmt"
	"strings"
)

type Document struct {
	Body     Element
	MetaData map[string]string
	Title    string
}

type Element struct {
	ChildElements []Element
	Class         string
	ID            string
}

func ParseHTML(html string) (Document, error) {
	headStart := strings.Index(html, "<head>") + len("<head>")
	headEnd := strings.Index(html, "</head>")
	if headStart == -1 || headEnd == -1 {
		return Document{}, errors.New("Unable to parse out the head of the document")
	}

	head := html[headStart:headEnd]
	metaData, title, err := ParseHTMLHead(head)
	if err != nil {
		return Document{}, err
	}

	bodyStart := strings.Index(html, "<body>") + len("<body>")
	bodyEnd := strings.Index(html, "</body>")
	if bodyStart == -1 || bodyEnd == -1 {
		return Document{}, errors.New("Unable to parse out body of the document")
	}
	
	body := html[bodyStart:bodyEnd]
	parsedBody, err := ParseHTMlBody(body)
	if err != nil {
		return Document{}, err
	}

	return Document{
		Title: title,
		MetaData: metaData,
		Body: parsedBody,
	}, nil
}

func ParseHTMLBody(body string) (Element, error) {
	return Element{}, nil
}

func ParseHTMLHead(head string) (metaData map[string]string, title string, err error) {
	return
}
