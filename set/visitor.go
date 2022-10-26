package set

// Visit 遍历二叉树时执行的操作
// return 是否中止遍历，true终止，false不终止
type Visit func(e E) bool
