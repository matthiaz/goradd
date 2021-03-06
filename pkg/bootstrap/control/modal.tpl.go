//** This file was code generated by got. DO NOT EDIT. ***

package control

import (
	"bytes"
	"context"
	"html"
)

func (m *Modal) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`    <div class="modal-dialog" role="document">
        <div class="modal-content">
`)

	m.titleBar.AddClass("modal-header")
	m.titleBar.Draw(ctx, buf)

	buf.WriteString(`            <div class="modal-body">
`)

	l := len(m.Children())
	if l > 2 {
		for _, child := range m.Children() {
			if child.ID() != m.titleBar.ID() && child.ID() != m.buttonBar.ID() {
				child.Draw(ctx, buf)
			}
		}
	} else {

		buf.WriteString(`<p>`)

		buf.WriteString(html.EscapeString(m.Text()))

		buf.WriteString(` </p>`)

	}

	buf.WriteString(`
            </div>
`)

	m.buttonBar.AddClass("modal-footer")
	m.buttonBar.Draw(ctx, buf)

	buf.WriteString(`        </div>
    </div>

`)

	return
}

func (d *TitleBar) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {
	if d.Title != "" {

		buf.WriteString(`     <h5 id="`)

		buf.WriteString(d.Parent().ID())

		buf.WriteString(`_title" class="modal-title">`)

		buf.WriteString(d.Title)

		buf.WriteString(`</h5>
`)

	}
	if d.HasCloseBox {

		buf.WriteString(`    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
      <span aria-hidden="true">&times;</span>
    </button>
`)

	}
	return
}
