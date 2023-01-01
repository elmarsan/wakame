package wakame

import (
	"regexp"
	"strings"
)

type Node struct {
	Tag        string
	Content    string
	Parent     *Node
	Children   []*Node
	Attributes map[string]interface{}
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

			selfClosing := tag[len(tag)-1:] == "/"
			if selfClosing {
				tag = strings.TrimSpace(tag[:len(tag)-1])
			}

			// Extract tag type
			tagTypeRe := regexp.MustCompile(`<(\w+).*>`)
			tagType := tagTypeRe.FindStringSubmatch("<" + tag + ">")[1]

			// Extract attributes of tag
			attributesRe := regexp.MustCompile(`(?P<name>\w+)\s*=\s*"(?P<value>[^"]*)"`)
			rawAttributes := attributesRe.FindAllStringSubmatch(tag, -1)

			attributes := map[string]interface{}{}
			for _, attr := range rawAttributes {
				key := attr[1]
				val := attr[2]
				attributes[key] = val
			}

			if !rootInit {
				(*pivot).Tag = tagType
				(*pivot).Attributes = attributes
				rootInit = true
			} else {
				child := &Node{
					Parent:     *pivot,
					Children:   nil,
					Tag:        tagType,
					Attributes: attributes,
				}

				(*pivot).Children = append((*pivot).Children, child)

				if !selfClosing {
					pivot = &child
				}
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
