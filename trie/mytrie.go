package trie

type Node struct {
    parent   *Node
    children map[rune]*Node
    character rune
    value    interface{}
    word     bool
}

func NewNode(parent *Node) *Node {
    return &Node{
        parent: parent,
        children: nil,
        value: nil,
        word: false,
    }
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
        m.root = NewNode(nil)
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
            childNode = &Node{
                parent: node,
            }
            childNode.character = keyChar
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
    // 找到最后一个节点
    node := m.getNode(key)
    // 如果不是单词结尾，不用做任何处理
    if node == nil || !node.word {
        return nil
    }
    m.size--
    
    oldValue := node.value
    
    // 如果还有子节点
    if node.children != nil && len(node.children) != 0 {
        node.word = false
        node.value = nil
        return oldValue
    }
    
    // 没有子节点
    var parent *Node
    delete(node.parent.children, node.character)
    for node.parent != nil{
        parent = node.parent
        delete(parent.children, node.character)
        if len(parent.children) != 0 {
            break
        }
        node = parent
        parent = node.parent
    }
    return oldValue
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
