package wakame

import (
	"regexp"
	"strings"
)

// ParseHTML parses a HTML string and returns a tree structure of Node objects representing the HTML nodes
func ParseHTML(html string) *Node {
	// Remove newlines and tabs from HTML string
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\t", "")

	// Initialize root node and pivot node
	root := &Node{
		Parent:   nil,
		Children: nil,
	}
	rootInit := false
	pivot := &root

	// Initialize variables for parsing
	i := 0
	tag := ""
	content := ""

	// Loop through each character in HTML string
	for i < len(html) {
		// Tag begin
		if html[i] == '<' && html[i+1] != '/' {
			i++

			for html[i] != '>' {
				tag += string(html[i])
				i++
			}

			// Check if tag is self-closing
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

			// Create new child node for pivot node
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
