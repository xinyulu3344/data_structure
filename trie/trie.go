package trie

type ITrie interface {
    Size() int
    IsEmpty() bool
    Clear()
    Get(key string) any
    Contains(key string) bool
    Add(key string, value any) any
    Remove(key string) any
    StartWith(prefix string) bool
}

type Node struct {
    parent   *Node
    children map[rune]*Node
    character rune
    value    any  // 当word为true的时候，将对应的value存储到这里
    word     bool // 是否为单词的结尾（是否为一个完整的单词）
}

func NewNode(parent *Node) *Node {
    return &Node{
        parent: parent,
        children: nil,
        value: nil,
        word: false,
    }
}

type Trie struct {
    size int
    root *Node
}

func NewTrie() *Trie {
    return &Trie{
        root: nil,
    }
}

func (t *Trie) Size() int {
    return t.size
}

func (t *Trie) IsEmpty() bool {
    return t.size == 0
}

func (t *Trie) Clear() {
    t.size = 0
    t.root =nil
}

func (t *Trie) Get(key string) any {
    n := t.getNodeByKey(key)
    if n != nil && n.word {
        return n.value
    } else {
        return nil
    }
}

func (t *Trie) Contains(key string) bool {
    n := t.getNodeByKey(key)
    return n != nil && n.word
}

func (t *Trie) Add(key string, value any) any {
    t.keyCheck(key)
    if t.root == nil {
        t.root = NewNode(nil)
    }
    n := t.root
    keyChars := []rune(key)
    length := len(keyChars)
    for i := 0; i < length; i++ {
        keyChar := keyChars[i]

        // 判断children是否为空
        emptyChildren := n.children == nil
        var childNode *Node
        if emptyChildren {
            childNode = nil
        } else {
            childNode = n.children[keyChar]
        }
        if childNode == nil {
            childNode = &Node{
                parent: n,
            }
            childNode.character = keyChar
            if emptyChildren {
                n.children = make(map[rune]*Node)
            }
            n.children[keyChar] = childNode
        }
        n = childNode
    }
    if n.word { // 已经存在这个单词
        // 覆盖
        oldValue := n.value
        n.value = value
        return oldValue
    }
    n.word = true
    n.value = value
    t.size++
    return nil
}

func (t *Trie) Remove(key string) any {
    // 找到最后一个节点
    n := t.getNodeByKey(key)
    if n == nil || !n.word {
        return nil
    }
    t.size--
    oldValue := n.value

    // 如果还有子节点
    if n.children != nil && len(n.children) != 0 {
        n.word = false
        n.value = nil
        return oldValue
    }
    // 没有子节点
    var parent *Node
    delete(n.parent.children, n.character)
    for n.parent != nil {
        parent = n.parent
        delete(parent.children, n.character)
        if parent.word || len(parent.children) != 0 {
            break
        }
        n = parent
        parent = n.parent
    }
    return oldValue
}

func (t *Trie) StartWith(prefix string) bool {
    return t.getNodeByKey(prefix) != nil
}


func (t *Trie) getNodeByKey(key string) *Node {
    if t.root == nil {
        return nil
    }
    t.keyCheck(key)

    n := t.root
    keyChars := []rune(key)
    length := len(keyChars)
    for i := 0; i < length; i++ {
        if n == nil || n.children == nil || len(n.children) == 0 {
            return nil
        }
        keyChar := keyChars[i]
        n = n.children[keyChar]
    }
    return n
}

func (t *Trie) keyCheck(key string) {
    if key == "" {
        panic("key must not be empty")
    }
}