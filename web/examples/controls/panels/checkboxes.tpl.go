//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"
)

func (ctrl *CheckboxPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>Checkboxes and Radio Buttons</h1>
<p>
Checkboxes and RadioButtons create html input tags with a type of "checkbox" or "radio".
They are used for setting true/false values, or for selecting from a group of items.
`)

	buf.WriteString(`Goradd&#39;s`)

	buf.WriteString(` code generation uses checkboxes to set binary values in a database, and
uses a checkbox list to setup relationships between records in a one-to-many relationship.
</p>
<p>
See also the CheckboxList and RadioList controls for working with groups of these controls,
and the CheckboxColumn for working with a column of checkboxes in an html table.
</p>
<h2>Labels</h2>
<p>
Dealing with labels and checkboxes can be tricky. Html uses a `)

	buf.WriteString(`&lt;label&gt;`)

	buf.WriteString(` tag to associate text
with a checkbox. The problem is that `)

	buf.WriteString(`&lt;label&gt;`)

	buf.WriteString(` tags are also used in html to associate titles
with input elements like textboxes and select lists. What if you want both a title, and a checkbox
label? You need two labels!
</p>
<p>
The problem is compounded by the many ways a label can be placed next to a checkbox. A label can
come before or after a checkbox, and also wrap around the checkbox. Different CSS and JS frameworks
require different placement of the label to correctly do their styling. In fact,
<a href="http://www.getbootstrap.com">Bootstrap</a>, a popular CSS framework, changed its placement
requirement between versions 3 and 4.
</p>
<p>
Goradd gives you the flexibility to place your labels where you want, and also gives you
a global setting to default the label placement across your whole site. The Goradd Bootstrap
support will automatically do the correct thing for Bootstrap.
</p>
<p>
In the examples below, view the source html of the page to see how Goradd deals with labels on
checkboxes.
</p>
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("checkbox1-ff").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("checkbox2-ff").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`


<h2>Radio Buttons</h2>
Radio buttons are generally part of a group of buttons with the purpose of allowing the user
to select only one item from the group. You choose whatever group name you want and assign it to
each button. The browser will then make sure only one gets selected.
<fieldset style="width:200px;">
<legend>Place</legend>
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("radio1").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("radio2").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("radio3").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("infoPanel").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
</fieldset>

`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("ajaxButton").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("serverButton").ProcessAttributeString(``).Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	return
}
