package bstprint

import (
	"fmt"
	"strings"
)

var length = 2
var rightAppend = "┌" + strings.Repeat("-", length)
var leftAppend = "└" + strings.Repeat("-", length)
var blankAppend = strings.Repeat(" ", length+1)
var lineAppend = "|" + strings.Repeat(" ", length)

type InorderPrinter struct {
	BinaryTreeInfo
}

func (p *InorderPrinter) PrintString() string {
	var b strings.Builder
	b.WriteString(p.printString(p.Root(), "", "", ""))
    s := b.String()
	return s[:b.Len() - 1]
}

func (p *InorderPrinter) printString(
	node any,
	nodePrefix string,
	leftPrefix string,
	rightPrefix string) string {
    
	left := p.Left(node)
	right := p.Right(node)
	s := fmt.Sprint(p.String(node))

	length := len(s)
	if length%2 == 0 {
		length--
	}

	length >>= 1

	nodeString := ""
	if right != nil {
		rightPrefix += strings.Repeat(" ", length)
		nodeString += p.printString(right, rightPrefix+rightAppend, rightPrefix+lineAppend, rightPrefix+blankAppend)
	}
	nodeString += nodePrefix + s + "\n"
	if left != nil {
		leftPrefix += strings.Repeat(" ", length)
		nodeString += p.printString(left, leftPrefix+leftAppend, leftPrefix+blankAppend, leftPrefix+lineAppend)
	}
	return nodeString
}
