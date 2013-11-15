package main

import (
	"bytes"
	"strings"
)

type Table struct {
	Rows    [][]string
	Options *TableOptions

	MaxColWidth int

	numColumns   int
	columnsWidth []int
}

type TableOptions struct {
	Padding      int
	Header       []string
	MaxColWidth  int
}

var defaultTableOptions = &TableOptions{
	Padding: 1,
	Header: nil,
	MaxColWidth: 50,
}

func NewTable(options *TableOptions) *Table {
	t := &Table{
		Options: options,
	}

	if t.Options == nil {
		t.Options = defaultTableOptions
	}

	if t.HasHeader() {
		t.AddRow(t.Options.Header)
	}

	if t.Options.MaxColWidth > 0 {
		t.MaxColWidth = t.Options.MaxColWidth
	} else {
		t.MaxColWidth = 50
	}

	return t
}

func (t *Table) HasHeader() bool {
	return t.Options.Header != nil
}

func (t *Table) AddRow(row []string) {
	if len(t.Rows) == 0 {
		t.numColumns = len(row)

		for i := 0; i < t.numColumns; i++ {
			t.columnsWidth = append(t.columnsWidth, 1)
		}
	}

	for j, col := range row {
		if len(col) > t.columnsWidth[j] {
			if t.MaxColWidth != 0 && len(col) > t.MaxColWidth {
				t.columnsWidth[j] = t.MaxColWidth
			} else {
				t.columnsWidth[j] = len(col)
			}
		}

		// tabs ruin our wrapping, rewrite them to 4 spaces
		row[j] = strings.Replace(col, "\t", "    ", -1)
	}

	lastCol := t.numColumns - 1
	if t.MaxColWidth > 0 && len(row[lastCol]) > t.MaxColWidth {
		wrapped := []string{ "", "", row[lastCol][t.MaxColWidth:] }
		row[lastCol] = row[lastCol][0:t.MaxColWidth]
		t.Rows = append(t.Rows, row)
		t.AddRow(wrapped)
	} else {
		t.Rows = append(t.Rows, row)
	}
}

func (t *Table) Render() string {
	// allocate a 1k byte buffer
	bb := make([]byte, 0, 1024)
	buf := bytes.NewBuffer(bb)

	buf.WriteString(t.separatorLine() + "\n")

	for i, row := range t.Rows {
		for j, _ := range row {
			buf.WriteString(t.getCell(i, j))
		}

		buf.WriteRune('\n')

		// below header
		if i == 0 && t.HasHeader() {
			buf.WriteString(t.separatorLine() + "\n")
		}
	}

	buf.WriteString(t.separatorLine() + "\n")

	return buf.String()
}

func (t *Table) separatorLine() string {
	sep := "+"
	for _, w := range t.columnsWidth {
		sep += strings.Repeat("-", w + 2 * t.Options.Padding)
		sep += "+"
	}
	return sep
}

func (t *Table) getCell(row, col int) string {
	cellContent := t.Rows[row][col]

	if len(cellContent) > t.MaxColWidth {
		cellContent = cellContent[0:t.MaxColWidth]
	}

	spacePadding := strings.Repeat(" ", t.Options.Padding)
	cellPadding := strings.Repeat(" ", t.columnsWidth[col] - len(cellContent))

	cell := cellContent + cellPadding

	if col == 0 {
		cell = "|" + spacePadding + cell + spacePadding + "|"
	} else {
		cell = spacePadding + cell + spacePadding + "|"
	}

	return cell
}
