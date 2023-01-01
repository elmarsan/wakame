package wakame

import "testing"

func TestNodeFindAll(t *testing.T) {
	root := &Node{
		Tag: "div",
		Children: []*Node{
			{
				Tag:     "p",
				Content: "wakame is funny",
				Attributes: map[string]interface{}{
					"class": "mb-5",
				},
			},
			{
				Tag:     "p",
				Content: "crazy stuff here",
				Attributes: map[string]interface{}{
					"class": "mb-5",
				},
			},
			{
				Tag:     "p",
				Content: "more crazy stuff",
				Attributes: map[string]interface{}{
					"class": "mb-10",
				},

				Children: []*Node{
					{
						Tag:     "a",
						Content: "I'm stupid",
					},
				},
			},
		},
	}

	nodes := root.FindAll("p", map[string]interface{}{
		"class": "mb-5",
	})

	if len(nodes) != 2 {
		t.Errorf("Invalid number of nodes was retrieved: %d, expected 2", len(nodes))
	}

	expectedNodes := []NodeTestCase{
		{
			tag:     "p",
			content: "wakame is funny",
			attributes: map[string]interface{}{
				"class": "mb-5",
			},
		},
		{
			tag:     "p",
			content: "crazy stuff here",
			attributes: map[string]interface{}{
				"class": "mb-5",
			},
		},
	}

	for i, node := range expectedNodes {
		node.CompareNode(t, nodes[i])
	}
}

func TestNodeHasAttributes(t *testing.T) {
	node := &Node{
		Tag: "img",
		Attributes: map[string]interface{}{
			"src":   "https://draxe.com/wp-content/uploads/2018/10/WakameHeader.jpg",
			"class": "img-class",
			"alt":   "Wakame photo",
		},
	}

	t.Run("Node should has attributes when by attribute name", func(t *testing.T) {
		attributes := map[string]interface{}{
			"alt": "",
		}

		if !node.hasAttributes(attributes) {
			t.Errorf("Some attribute of: %#v is missing, node attributes: %#v\n", attributes, node.Attributes)
		}
	})

	t.Run("Node should has attributes when finding exactly attribute and attribute name", func(t *testing.T) {
		attributes := map[string]interface{}{
			"class": "img-class",
			"alt":   "",
		}

		if !node.hasAttributes(attributes) {
			t.Errorf("Some attribute of: %#v is missing, node attributes: %#v\n", attributes, node.Attributes)
		}
	})

	t.Run("Node should has NOT attributes when finding by attribute name", func(t *testing.T) {
		attributes := map[string]interface{}{
			"id": "wakame-img",
		}

		if node.hasAttributes(attributes) {
			t.Errorf("Some attribute of: %#v is present in node, node attributes: %#v\n", attributes, node.Attributes)
		}
	})
}
