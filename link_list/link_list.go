package link_list

// 定义节点结构
type Node struct {
    Data interface{}
    Next *Node
}

type LinkList struct {
    Head   *Node
    Rear *Node
    Length int
}

// 初始化链表, 相当于整个空链表
func (link *LinkList) Init() *LinkList {
    link.Head.Next = nil
    link.Rear = link.Head
    link.Length = 0
    return link
}

// 新建一个链表
func NewLinkList() *LinkList {
    linkList := new(LinkList)
    linkList.Head = new(Node)
    linkList.Length = 0
    linkList.Init()
    return linkList
}

// 获取链表长度
func (link *LinkList) GetLength() int {
    return link.Length
}


// 在链表尾部插入元素
func (link *LinkList) PushBack(v interface{}) *Node{
    link.lazyInit()
    node := &Node{
        Data: v,
        Next:nil,
    }
    link.Rear.Next = node
    link.Rear = node
    link.Length++
    return node
}

// 在链表头部插入元素
func (link *LinkList) PushFront(v interface{}) *Node {
    link.lazyInit()
    node := &Node{
        Data: v,
        Next: link.Head.Next,
    }
    link.Head.Next = node
    link.Length++
    return node
}



// 获取第一个元素
func (link *LinkList) Front() *Node {
    if link.Length == 0 {
        return nil
    }
    return link.Head.Next
}

// 获取最后元素
func (link *LinkList) Back() *Node {
    if link.Length == 0 {
        return nil
    }
    return link.Rear
}

func (link *LinkList) lazyInit() {
    if link.Head.Next == nil {
        link.Init()
    }
}
