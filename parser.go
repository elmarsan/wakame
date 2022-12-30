package wakame

import (
	"fmt"
	"strings"
)

type Node struct {
	Tag      string
	Content  string
	Parent   *Node
	Children []*Node
}

func ParseHTML(html string) *Node {
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\t", "")

	root := &Node{
		Parent:   nil,
		Children: nil,
	}
	nodePtr := &root

	i := 0
	tag := ""
	content := ""

	for i < len(html) {
		// Tag begin
		if html[i] == '<' && html[i+1] != '/' {
			i++

			for html[i] != '>' {
				tag += string(html[i])
				i++
			}

			if len((*nodePtr).Tag) == 0 {
				(*nodePtr).Tag = tag
				fmt.Printf("Parent <%s> created\n", (*nodePtr).Tag)

			} else if len((*nodePtr).Tag) > 0 && len((*nodePtr).Content) == 0 {
				child := &Node{
					Parent:   *nodePtr,
					Children: nil,
					Tag:      tag,
				}

				(*nodePtr).Children = append((*nodePtr).Children, child)
				fmt.Printf("Child <%s> of <%s> added\n", tag, (*nodePtr).Tag)
			}

			tag = ""
		}

		// Content begin
		if html[i] != '<' && html[i] != '>' && html[i] != '/' {
			if html[i-1] == '>' && html[i-2] != '/' {
				for html[i] != '<' {
					content += string(html[i])
					i++
				}

				if len((*nodePtr).Tag) == 0 {
					(*nodePtr).Content = content
					fmt.Printf("Setting parent content <%s>%s<%s>\n", (*nodePtr).Tag, (*nodePtr).Content, (*nodePtr).Tag)

				} else if len((*nodePtr).Tag) > 0 && len((*nodePtr).Content) == 0 {
					child := (*nodePtr).Children[len((*nodePtr).Children)-1]
					child.Content = content

					fmt.Printf("Setting child content <%s>%s<%s>\n", (*child).Tag, (*child).Content, (*child).Tag)
				}

				content = ""
			}
		}

		i++
	}

	return root
}
