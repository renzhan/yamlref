package yaml

import (
	"strings"
)

func findRef(n *node, refKey string) string {
	keys := strings.Split(refKey, ".")

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		curNode, ok := findKey(n, key)
		if ok && curNode != nil {
			n = curNode
			continue
		}
	}
	return n.value
}

func findKey(n *node, key string) (*node, bool) {
	if n.tag == key || n.value == key {
		return n, true
	} else {
		if n.children != nil && len(n.children) > 0 {
			if n.kind == mappingNode {
				for i := 0; i < len(n.children); i += 2 {
					_, ok := findKey(n.children[i], key)
					if ok {
						return n.children[i+1], true
					}
				}
			} else {
				for i := 0; i < len(n.children); i++ {
					return findKey(n.children[i], key)
				}
			}
		}
	}
	return nil, false
}
