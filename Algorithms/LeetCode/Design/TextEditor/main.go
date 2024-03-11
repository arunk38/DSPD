package main

import "fmt"

/*
 * https://leetcode.com/problems/design-a-text-editor/description/
 */

type TextEditor struct {
	text   []byte
	cursor int
}

func Constructor() TextEditor {
	return TextEditor{}
}

func (this *TextEditor) AddText(char string) {
	bytesToInsert := []byte(char)
	this.text = append(this.text[:this.cursor], append(bytesToInsert, this.text[this.cursor:]...)...)
	this.cursor += len(bytesToInsert)
}

func (this *TextEditor) DeleteText(k int) int {
	if this.cursor-k >= 0 {
		this.text = append(this.text[:this.cursor-k], this.text[this.cursor:]...)
		this.cursor -= k
		return k
	} else {
		this.text = this.text[this.cursor:]
		k = this.cursor
		this.cursor = 0
		return k
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (this *TextEditor) LeftSubString() string {
	startIndex := max(0, this.cursor-10)
	return string(this.text[startIndex:this.cursor])
}

func (this *TextEditor) CursorLeft(k int) string {
	this.cursor -= k
	if this.cursor < 0 {
		this.cursor = 0
	}

	return this.LeftSubString()
}

func (this *TextEditor) CursorRight(k int) string {
	this.cursor += k
	if this.cursor > len(this.text) {
		this.cursor = len(this.text)
	}

	return this.LeftSubString()
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */

func main() {
	editor := Constructor()
	editor.AddText("leetcode")
	fmt.Println(editor.DeleteText(4))
	editor.AddText("practice")
	fmt.Println(editor.CursorRight(3))
	fmt.Println(editor.CursorLeft(8))
	fmt.Println(editor.DeleteText(10))
	fmt.Println(editor.CursorLeft(2))
	fmt.Println(editor.CursorRight(6))
}
