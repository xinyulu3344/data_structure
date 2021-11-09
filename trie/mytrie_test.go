package trie

import (
    "fmt"
    "testing"
)

func TestA(t *testing.T) {
    myTrie := NewMyTrie()
    fmt.Println(myTrie.Add("cat", 1)) // <nil>
    myTrie.Add("dog", 2)
    myTrie.Add("catalog", 3)
    myTrie.Add("cast", 4)
    myTrie.Add("小码哥", 5)
    fmt.Println(myTrie.Size())            // 5
    fmt.Println(myTrie.Get("cat"))        // 1
    fmt.Println(myTrie.Add("cat", 2))     // 1
    fmt.Println(myTrie.Add("cat", 3))     // 2
    fmt.Println(myTrie.Get("小码哥"))        // 5
    
    fmt.Println(myTrie.StartWith("c"))    // true
    fmt.Println(myTrie.StartWith("ca"))   // true
    fmt.Println(myTrie.StartWith("cat"))  // true
    fmt.Println(myTrie.StartWith("cate")) // false
    fmt.Println(myTrie.StartWith("hehe")) // false
    fmt.Println(myTrie.StartWith("小"))    // true
    
    fmt.Println(myTrie.Contains("小码")) // false
    fmt.Println(myTrie.Contains("小码哥")) // true
    
    myTrie.Clear()
    
    fmt.Println(myTrie.Size()) // 0
    fmt.Println(myTrie.Contains("小码哥")) // false
    myTrie.Add("小码哥", 5)
    fmt.Println(myTrie.Contains("小码哥")) // true
    fmt.Println(myTrie.Get("小码哥")) // 5
}
