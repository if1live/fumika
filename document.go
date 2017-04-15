package fumika

// https://github.com/if1live/haru/blob/master/common/document.go

import (
	"strings"

	"golang.org/x/net/html"
)

func GetElementByID(n *html.Node, id string) *html.Node {
	for _, a := range n.Attr {
		if a.Key == "id" {
			if a.Val == id {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found := GetElementByID(c, id)
		if found != nil {
			return found
		}
	}
	return nil
}

func GetElementByClassName(n *html.Node, classname string) *html.Node {
	for _, a := range n.Attr {
		if a.Key == "class" {
			classes := strings.Split(a.Val, " ")
			for _, val := range classes {
				if val == classname {
					return n
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found := GetElementByClassName(c, classname)
		if found != nil {
			return found
		}
	}
	return nil
}

func GetElementsByClassName(n *html.Node, classname string) []*html.Node {
	retval := []*html.Node{}
	return getElementsByClassName_r(n, classname, retval)
}
func getElementsByClassName_r(n *html.Node, classname string, retval []*html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			classes := strings.Split(a.Val, " ")
			for _, val := range classes {
				if val == classname {
					retval = append(retval, n)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retval = getElementsByClassName_r(c, classname, retval)
	}
	return retval
}

func GetElementsByTagName(n *html.Node, tag string) []*html.Node {
	retval := []*html.Node{}
	return getElementsByTagName_r(n, tag, retval)
}

func getElementsByTagName_r(n *html.Node, tag string, retval []*html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		retval = append(retval, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retval = getElementsByTagName_r(c, tag, retval)
	}
	return retval
}
