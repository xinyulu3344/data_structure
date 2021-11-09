package trie

type Node struct {
    children map[rune]*Node
    value    interface{}
    word     bool
}

func (n *Node) getChildren() map[rune]*Node {
    if n.children == nil {
        n.children = make(map[rune]*Node)
    }
    return n.children
}

type MyTrie struct {
    size int
    root *Node
}

func NewMyTrie() *MyTrie {
    return &MyTrie{
        size: 0,
        root: &Node{
            children: make(map[rune]*Node),
            value: nil,
            word: false,
        },
    }
}

func (m *MyTrie) Size() int {
    return m.size
}

func (m *MyTrie) IsEmpty() bool {
    return m.size == 0
}

func (m *MyTrie) Clear() {
    m.size = 0
    n := m.root.getChildren()
    for k := range n {
        delete(n, k)
    }
}

func (m *MyTrie) Get(key string) interface{} {
    node := m.getNode(key)
    if node != nil {
        return node.value
    } else {
        return nil
    }
}

func (m *MyTrie) Contains(key string) bool {
    return m.getNode(key) != nil
}

func (m *MyTrie) Add(key string, value interface{}) interface{} {
    m.keyCheck(key)
    
    node := m.root
    keyChars := []rune(key)
    length := len(keyChars)
    
    for i := 0; i < length; i++ {
        keyChar := keyChars[i]
        if _, ok := node.getChildren()[keyChar]; !ok {
            childNode := &Node{}
            node.getChildren()[keyChar] = childNode
        }
        node = node.getChildren()[keyChar]
    }
    if node.word { // 已经存在这个单词
        // 覆盖
        oldValue := node.value
        node.value = value
        return oldValue
    }
    // 新增一个单词
    node.word = true
    node.value = value
    m.size++
    return nil
}

func (m *MyTrie) Remove(key string) interface{} {
    panic("implement me")
}

func (m *MyTrie) StartWith(prefix string) bool {
    m.keyCheck(prefix)
    
    node := m.root
    prefixChars := []rune(prefix)
    length := len(prefixChars)
    
    for i := 0; i < length; i++ {
        prefixChar := prefixChars[i]
        if n, ok := node.getChildren()[prefixChar]; !ok {
            return false
        } else {
            node = n
        }
    }
    return true
}

func (m *MyTrie) getNode(key string) *Node {
    if m.root == nil {
        return nil
    }
    m.keyCheck(key)
    node := m.root
    keyChars := []rune(key)
    length := len(keyChars)
    for i := 0; i < length; i++ {
        keyChar := keyChars[i]
        if n, ok := node.getChildren()[keyChar]; !ok {
            return nil
        } else {
            node = n
        }
    }
    if node.word {
        return node
    }
    return nil
}

func (m *MyTrie) keyCheck(key string) {
    if key == "" {
        panic("key must not be empty")
    }
}