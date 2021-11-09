package trie

type Node struct {
    children map[rune]*Node
    value    interface{}
    word     bool
}

type MyTrie struct {
    size int
    root *Node
}

func NewMyTrie() *MyTrie {
    return &MyTrie{
        size: 0,
        root: nil,
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
    m.root = nil
}

func (m *MyTrie) Get(key string) interface{} {
    node := m.getNode(key)
    if node != nil && node.word {
        return node.value
    } else {
        return nil
    }
}

func (m *MyTrie) Contains(key string) bool {
    node := m.getNode(key)
    return node != nil && node.word
}

func (m *MyTrie) Add(key string, value interface{}) interface{} {
    m.keyCheck(key)
    
    if m.root == nil {
        m.root = &Node{
            children: make(map[rune]*Node),
            value: nil,
            word: false,
        }
    }
    
    node := m.root
    keyChars := []rune(key)
    length := len(keyChars)
    
    for i := 0; i < length; i++ {
        keyChar := keyChars[i]
        
        // 判断children是否为空
        emptyChildren := node.children == nil
        var childNode *Node
        if emptyChildren {
            childNode = nil
        } else {
            childNode = node.children[keyChar]
        }
        if childNode == nil {
            childNode = &Node{}
            if emptyChildren {
                node.children = make(map[rune]*Node)
            }
            node.children[keyChar] = childNode
        }
        node = childNode
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
    return m.getNode(prefix) != nil
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
        if node == nil || node.children == nil || len(node.children) == 0 {
            return nil
        }
        keyChar := keyChars[i]
        node = node.children[keyChar]
    }
    return node
}

func (m *MyTrie) keyCheck(key string) {
    if key == "" {
        panic("key must not be empty")
    }
}