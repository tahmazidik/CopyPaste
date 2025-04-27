package editor

import (
	"bufio"
	"os"
)

type TextBuffer struct {
	lines []string
}

func (tb *TextBuffer) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tb.lines = append(tb.lines, scanner.Text())
	}

	//tb.lines = append(tb.lines, "") //Пустая строчка в конце файла
	return scanner.Err()
}

func (tb *TextBuffer) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range tb.lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func (tb *TextBuffer) RemoveLines(start, count int) {
	if start < 0 || start >= len(tb.lines) || count <= 0 {
		return
	}

	end := start + count
	if end > len(tb.lines) {
		end = len(tb.lines)
	}

	tb.lines = append(tb.lines[:start], tb.lines[end:]...)
}

func (tb *TextBuffer) InsertLines(pos int, newlines []string) {
	if pos < 0 {
		pos = 0
	}
	if pos > len(tb.lines) {
		pos = len(tb.lines)
	}

	tb.lines = append(tb.lines[:pos], append(newlines, tb.lines[pos:]...)...)
}

type Editor struct {
	bufer       TextBuffer
	cursor      int
	anchor      int
	ShiftActive bool
	clipboard   []string
}

func NewEditor() *Editor {
	return &Editor{
		bufer:       TextBuffer{lines: make([]string, 0)},
		cursor:      0,
		anchor:      0,
		ShiftActive: false,
		clipboard:   make([]string, 0),
	}
}

func (e *Editor) LoadFile(filename string) error {
	return e.bufer.LoadFromFile(filename)
}

func (e *Editor) SaveFile(filename string) error {
	return e.bufer.SaveToFile(filename)
}

func (e *Editor) ProcessCommand(cmd string) {
	switch cmd {
	case "Down":
		if e.cursor < len(e.bufer.lines)-1 {
			e.cursor++
		}
	case "Up":
		if e.cursor > 0 {
			e.cursor--
		}
	case "Shift":
		e.ShiftActive = true
		e.anchor = e.cursor
	case "Ctrl+X":
		e.cut()
	case "Ctrl+V":
		e.paste()
	}
}

func (e *Editor) getSelection() (int, int) {
	start, end := e.anchor, e.cursor
	if start > end {
		start, end = end, start
	}

	if end >= len(e.bufer.lines)-1 {
		end = len(e.bufer.lines) - 2
	}

	return start, end
}

func (e *Editor) cut() {
	if e.ShiftActive {
		start, end := e.getSelection()
		if start >= len(e.bufer.lines)-1 {
			e.ShiftActive = false
			return
		}

		e.clipboard = make([]string, end-start+1)
		copy(e.clipboard, e.bufer.lines[start:end+1])
		e.bufer.RemoveLines(start, end-start+1)
		e.cursor = start
	} else {
		if e.cursor >= len(e.bufer.lines)-1 {
			return
		}
		e.clipboard = []string{(e.bufer.lines[e.cursor])}
		e.bufer.RemoveLines(e.cursor, 1)
	}
	e.ShiftActive = false
}

func (e *Editor) paste() {
	if len(e.clipboard) == 0 {
		return
	}

	if e.ShiftActive {
		start, end := e.getSelection()
		e.bufer.RemoveLines(start, end-start+1)
		e.bufer.InsertLines(start, e.clipboard)
		e.cursor = start + len(e.clipboard)
	} else {
		e.bufer.InsertLines(e.cursor, e.clipboard)
		e.cursor += len(e.clipboard)
	}
	e.ShiftActive = false
}
