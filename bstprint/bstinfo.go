package bstprint

type BinaryTreeInfo interface {
	Root() any
	Left(node any) any
	Right(node any) any
	String(node any) any
}
