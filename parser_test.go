package wakame

import (
	"testing"
)

type NodeTestCase struct {
	content  string
	tag      string
	children []*NodeTestCase
}

func (tc *NodeTestCase) CompareNode(t *testing.T, node *Node) {
	if tc.content != node.Content {
		t.Errorf("Content: %s, expected: %s", node.Content, tc.content)
	}

	if len(tc.children) != len(node.Children) {
		t.Errorf("Number of children: %d, expected: %d", len(tc.children), len(node.Children))
	}

	if tc.tag != node.Tag {
		t.Errorf("Tag: %s, expected: %s", node.Tag, tc.tag)
	}

	for i, child := range tc.children {
		child.CompareNode(t, node.Children[i])
	}
}

func TestDepthLevelOne(t *testing.T) {
	html := `
		<div>
			<p>Hello world</p>
			<p>Goodbye world</p>
			<p>Wakame prototype</p>
		</div>
	`

	expectedRoot := NodeTestCase{
		tag: "div",
		children: []*NodeTestCase{
			{
				tag:     "p",
				content: "Hello world",
			},
			{
				tag:     "p",
				content: "Goodbye world",
			},
			{
				tag:     "p",
				content: "Wakame prototype",
			},
		},
	}

	root := ParseHTML(html)
	expectedRoot.CompareNode(t, root)
}

func TestDepthLevelTwo(t *testing.T) {
	html := `
		<div>
			<div>
				<p>depth two</p>
			</div>
			<p>Hello world</p>
			<p>Goodbye world</p>
			<p>Wakame prototype</p>
		</div>
	`

	expectedRoot := NodeTestCase{
		tag: "div",
		children: []*NodeTestCase{
			{
				tag: "div",
				children: []*NodeTestCase{
					{
						tag:     "p",
						content: "depth two",
					},
				},
			},
			{
				tag:     "p",
				content: "Hello world",
			},
			{
				tag:     "p",
				content: "Goodbye world",
			},
			{
				tag:     "p",
				content: "Wakame prototype",
			},
		},
	}

	root := ParseHTML(html)
	expectedRoot.CompareNode(t, root)
}

func TestDepthLevelThree(t *testing.T) {
	html := `
		<div>
			<div>
				<p>depth two</p>
			</div>
			<p>
				<strong>Hello world</strong>
			</p>
			<p>Goodbye world</p>
			<div>
				<h1>Title</h1>
				<p>Wakame prototype</p>
			</div>

			<footer>
				<div>
					<ul>
						<li>Item 1</li>
						<li>Item 2</li>
						<li>Item 3</li>
					</ul>
				</div>
			</footer>
		</div>
	`

	expectedRoot := NodeTestCase{
		tag: "div",
		children: []*NodeTestCase{
			{
				tag: "div",
				children: []*NodeTestCase{
					{
						tag:     "p",
						content: "depth two",
					},
				},
			},
			{
				tag: "p",
				children: []*NodeTestCase{
					{
						tag:     "strong",
						content: "Hello world",
					},
				},
			},
			{
				tag:     "p",
				content: "Goodbye world",
			},
			{
				tag: "div",
				children: []*NodeTestCase{
					{
						tag:     "h1",
						content: "Title",
					},
					{
						tag:     "p",
						content: "Wakame prototype",
					},
				},
			},
			{
				tag: "footer",
				children: []*NodeTestCase{
					{
						tag: "div",
						children: []*NodeTestCase{
							{
								tag: "ul",
								children: []*NodeTestCase{
									{
										tag:     "li",
										content: "Item 1",
									},
									{
										tag:     "li",
										content: "Item 2",
									},
									{
										tag:     "li",
										content: "Item 3",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	root := ParseHTML(html)
	expectedRoot.CompareNode(t, root)
}

func TestSelfClosingTag(t *testing.T) {
	html := `
		<div>
			<img/>
			<p>Paragraph</p>
			<input />
			< br/>
			< br />
		</div>
	`

	expectedRoot := NodeTestCase{
		tag: "div",
		children: []*NodeTestCase{
			{
				tag: "img",
			},
			{
				tag:     "p",
				content: "Paragraph",
			},
			{
				tag: "input",
			},
			{
				tag: "br",
			},
			{
				tag: "br",
			},
		},
	}

	root := ParseHTML(html)
	t.Log(root.Children[0])
	expectedRoot.CompareNode(t, root)
}
