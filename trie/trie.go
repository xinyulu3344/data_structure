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
    if n == nil {
        return nil
    } else {
        return n.value
    }
}

func (t *Trie) Contains(key string) bool {
    return t.getNodeByKey(key) != nil
}

func (t *Trie) Add(key string, value any) any {
	panic("not implemented") // TODO: Implement
}

func (t *Trie) Remove(key string) any {
	panic("not implemented") // TODO: Implement
}

func (t *Trie) StartWith(prefix string) bool {
	panic("not implemented") // TODO: Implement
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
        if n != nil || n.children == nil || len(n.children) == 0 {
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