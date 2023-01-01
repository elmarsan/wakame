package wakame

import "strings"

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

	rootInit := false
	pivot := &root

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

			if !rootInit {
				(*pivot).Tag = tag
				rootInit = true
			} else {
				child := &Node{
					Parent:   *pivot,
					Children: nil,
					Tag:      tag,
				}

				(*pivot).Children = append((*pivot).Children, child)
				pivot = &child
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

				(*pivot).Content = content
				content = ""
			}
		}

		// Tag end
		if html[i] == '<' && html[i+1] == '/' {
			i += 2

			for html[i] != '>' {
				tag += string(html[i])
				i++
			}

			parent := (*pivot).Parent
			pivot = &parent
			tag = ""
		}

		i++
	}

	return root
}
