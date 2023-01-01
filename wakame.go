package wakame

// Node represents a HTML node in a tree structure
type Node struct {
	// HTML tag of the node
	Tag string
	// Inner text of the node
	Content string
	// Parent node
	Parent *Node
	// Children nodes
	Children []*Node
	// Map of attributes of the HTML node, with the attribute name as the key and the attribute value as the value
	Attributes map[string]interface{}
}

// FindAll returns a slice of pointers to nodes that have the specified tag and attributes.
// If tag is an empty string, nodes with any tag will be returned. If attributes is nil or an empty map,
// nodes with any attributes will be returned. If both tag and attributes are specified, only nodes that
// have the specified tag and attributes will be returned.
func (n *Node) FindAll(tag string, attributes map[string]interface{}) []*Node {
	nodes := []*Node{}

	if n.Tag == tag && n.hasAttributes(attributes) {
		nodes = append(nodes, n)
	}

	for _, child := range n.Children {
		nodes = append(nodes, child.FindAll(tag, attributes)...)
	}

	return nodes
}

// hasAttributes checks if a Node has all the attributes specified in the provided map
// If the value of an attribute in the map is an empty string, it only checks if the attribute is present in the Node
// Otherwise, it checks if the attribute is present and has the same value as in the map
// Returns true if the Node has all the specified attributes, false otherwise
func (n *Node) hasAttributes(attributes map[string]interface{}) bool {
	for key, val := range attributes {
		nodeVal := n.Attributes[key]

		// Check only if key exist
		if val == "" {
			if nodeVal == nil {
				return false
			}
		} else {
			if nodeVal == nil || nodeVal != val {
				return false
			}
		}
	}

	return true
}
