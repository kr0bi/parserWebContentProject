package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

var stackPersona = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "table", "tbody", "tr", "td", "table", "tbody", "tr"}
var stackAnno = []string{"html", "body", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "div", "table", "tbody", "tr", "td"}
var temp = []string{"table", "tbody", "tr", "td", "table", "tbody", "tr"}

func main() {
	/*data, _ := os.Open("/home/daniele/Programming/parserWebContentProject/main/pagina.html")

	doc, err := html.Parse(data)*/
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	first, last, nome := false, false, true
	var fir, las, nom *bool
	fir, las, nom = &first, &last, &nome

	var n, v []string

	var nomi *[]string = &n
	var voti *[]string = &v

	findStuff(nil, doc, fir, las, nom, nomi, voti)

	for i := 0; i < len(v); i++ {
		fmt.Printf("%-25s\t%s\n", n[i], v[i])
	}
}

func findStuff(stack []string, n *html.Node, first, second, nome *bool, nomi, voti *[]string) {
	if n.Type == html.TextNode {
		if !(*first && *second) {
			if !*first {
				if strings.Contains(n.Data, "2018-19") {
					*first = true
				}
			} else {
				if strings.Contains(n.Data, "2017-18") {
					*second = true
				} else {
					stack = append(stack, n.Data)
					if *nome {
						if n.Data != "\n" {
							*nomi = append(*nomi, n.Data)
							*nome = false
						}
					} else {
						if n.Data != "\n" {
							*voti = append(*voti, n.Data)
							*nome = true
						}

					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findStuff(stack, c, first, second, nome, nomi, voti)
	}
}