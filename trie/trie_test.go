package trie

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Clear()
	fmt.Printf("trie.Contains(\"Hello 世界！\"): %v\n", trie.Contains("Hello 世界！"))
	fmt.Printf("trie.Contains(\"世界 Hello!\"): %v\n", trie.Contains("世界 Hello!"))
	fmt.Printf("trie.Get(\"Hello 世界！\"): %v\n", trie.Get("Hello 世界！"))
	fmt.Printf("trie.Get(\"世界 Hello!\"): %v\n", trie.Get("世界 Hello!"))
	fmt.Printf("trie.Remove(\"世界 Hello!\"): %v\n", trie.Remove("世界 Hello!"))
	fmt.Printf("trie.Size(): %v\n", trie.Size())
	fmt.Printf("trie.IsEmpty(): %v\n", trie.IsEmpty())
	fmt.Printf("trie.StartWith(\"世界 Hello!\"): %v\n", trie.StartWith("世界 Hello!"))
	fmt.Printf("trie.Add(\"世界 Hello!\", 18): %v\n", trie.Add("世界 Hello!", 18))
	fmt.Printf("trie.Add(\"Hello 世界！\", 28): %v\n", trie.Add("Hello 世界！", 28))
	fmt.Printf("trie.Add(\"Hello 世界！\", 29): %v\n", trie.Add("Hello 世界！", 29))
	fmt.Printf("trie.Size(): %v\n", trie.Size())
	fmt.Printf("trie.Contains(\"世界 Hello!\"): %v\n", trie.Contains("世界 Hello!"))
	fmt.Printf("trie.Contains(\"Hello 世界！\"): %v\n", trie.Contains("Hello 世界！"))
	fmt.Printf("trie.Get(\"世界 Hello!\"): %v\n", trie.Get("世界 Hello!"))
	fmt.Printf("trie.Get(\"Hello 世界！\"): %v\n", trie.Get("Hello 世界！"))
	fmt.Printf("trie.StartWith(\"Hell\"): %v\n", trie.StartWith("Hell"))
}

func Test1(t *testing.T) {
	trie := NewTrie()
	assert(t, !trie.Contains("Hello 世界！"))
	assert(t, !trie.Contains("世界 Hello!"))
	assert(t, trie.Get("Hello 世界！") == nil)
	assert(t, trie.Get("世界 Hello") == nil)
	assert(t, trie.Remove("世界 Hello!") == nil)
	assert(t, trie.Size() == 0)
	assert(t, trie.IsEmpty())
	assert(t, !trie.StartWith("世界 Hello!"))
	assert(t, trie.Add("世界 Hello!", 18) == nil)
	assert(t, trie.Add("Hello 世界！", 28) == nil)
	assert(t, trie.Add("Hello 世界！", 29) == 28)
	assert(t, trie.Size() == 2)
	assert(t, trie.Contains("Hello 世界！"))
	assert(t, trie.Contains("世界 Hello!"))
	assert(t, trie.Get("Hello 世界！") == 29)
	assert(t, trie.Get("世界 Hello!") == 18)
	assert(t, trie.StartWith("Hell"))
	assert(t, trie.Remove("Hello 世界！") == 29)
	assert(t, trie.Size() == 1)
	assert(t, !trie.Contains("Hello 世界！"))
}

func assert(t *testing.T, ok bool) {
	if !ok {
		debug.PrintStack()
		t.FailNow()
	}
}