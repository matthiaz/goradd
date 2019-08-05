package column

import (
	"context"
	"github.com/goradd/goradd/pkg/page/control"
)

// TexterColumn is a table column that lets you use a CellTexter to specify the content of the cells in the column.
// One convenient way to use this is to define a CellText function on the
// parent object and pass it as the CellTexter. If your CellTexter is going to output html instead of raw text, call
// SetIsHtml() on the column after creating it.
type TexterColumn struct {
	control.ColumnBase
}

// NewTexterColumn creates a new column with a custom cell texter.
func NewTexterColumn(texter CellTexter) *TexterColumn {
	i := TexterColumn{}
	i.Init(texter)
	return &i
}

func (c *TexterColumn) Init(texter CellTexter) {
	c.ColumnBase.Init(c)
	c.SetCellTexter(texter)
}

// TexterColumnCreator creates a column that uses a CellTexter to get the content of each cell.
type TexterColumnCreator struct {
	// Texter returns the text that should go in each cell. Pass a string control id, or a CellTexter.
	Texter interface{}
	// Title is the title at the top of the column
	Title string
	control.ColumnOptions
}

func (c TexterColumnCreator) Create(ctx context.Context, parent control.TableI) control.ColumnI {
	if c.Texter == nil {
		panic("a Texter is required")
	}
	var texter CellTexter
	if s,ok := c.Texter.(string); ok {
		texter = GetCellTexter(parent, s)
	} else {
		texter = c.Texter.(CellTexter)
	}

	col := NewTexterColumn(texter)
	col.SetTitle(c.Title)
	col.ApplyOptions(ctx, parent, c.ColumnOptions)
	return col
}