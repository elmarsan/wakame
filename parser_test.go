package wakame

import "testing"

func TestDepthLevelOne(t *testing.T) {
	html := `
		<div>
			<p>Hello world</p>
			<p>Goodbye world</p>
			<p>Wakame prototype</p>
		</div>
	`

	root := ParseHTML(html)

	if root.Tag != "div" {
		t.Error("Wrong root tag")
	}

	if len(root.Children) != 3 {
		t.Error("Wrong children len")
	}

	for _, child := range root.Children {
		if child.Tag != "p" {
			t.Error("Wrong child tag")
		}
	}

	if root.Children[0].Content != "Hello world" {
		t.Error("Wrong child content")
	}

	if root.Children[1].Content != "Goodbye world" {
		t.Error("Wrong child content")
	}

	if root.Children[2].Content != "Wakame prototype" {
		t.Error("Wrong child content")
	}
}
