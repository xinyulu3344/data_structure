package trie

import (
    "fmt"
    "testing"
)

func TestA(t *testing.T) {
    myTrie := NewMyTrie()
    myTrie.Add("cat", 1)
    myTrie.Add("dog", 2)
    myTrie.Add("catalog", 3)
    myTrie.Add("cast", 4)
    myTrie.Add("小码哥", 5)
    fmt.Println(myTrie.Size() == 5)        // true
    fmt.Println(myTrie.Get("cat") == 1)    // true
    fmt.Println(myTrie.Add("cat", 2) == 1) // true
    fmt.Println(myTrie.Add("cat", 3) == 2) // true
    fmt.Println(myTrie.Get("小码哥") == 5)    // true
    
    fmt.Println(myTrie.StartWith("c"))     // true
    fmt.Println(myTrie.StartWith("ca"))    // true
    fmt.Println(myTrie.StartWith("cat"))   // true
    fmt.Println(!myTrie.StartWith("cate")) // true
    fmt.Println(!myTrie.StartWith("hehe")) // true
    fmt.Println(myTrie.StartWith("小"))     // true
    
    fmt.Println(!myTrie.Contains("小码")) // true
    fmt.Println(myTrie.Contains("小码哥")) // true
    
    myTrie.Clear()
    
    fmt.Println(myTrie.Size() == 0)      // true
    fmt.Println(!myTrie.Contains("小码哥")) // true
    myTrie.Add("小码哥", 5)
    fmt.Println(myTrie.Contains("小码哥")) // true
    fmt.Println(myTrie.Get("小码哥") == 5) // true
    
    myTrie.Add("cat", 1)
    myTrie.Add("dog", 2)
    myTrie.Add("catalog", 3)
    myTrie.Add("cast", 4)
    myTrie.Add("小码哥", 5)
    fmt.Println(myTrie.Remove("cat") == 1)     // true
    fmt.Println(myTrie.Remove("catalog") == 3) // true
    fmt.Println(myTrie.Remove("cast") == 4)    // true
    fmt.Println(myTrie.Size() == 2)            // true
    fmt.Println(myTrie.StartWith("小"))         // true
    fmt.Println(myTrie.StartWith("do"))        // true
    fmt.Println(!myTrie.StartWith("c"))        // true
}
