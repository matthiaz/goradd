package panels

import (
	"context"
	"github.com/goradd/goradd/pkg/datetime"
	"github.com/goradd/goradd/pkg/page"
	"github.com/goradd/goradd/pkg/page/action"
	. "github.com/goradd/goradd/pkg/page/control"
	"github.com/goradd/goradd/pkg/url"
	"github.com/goradd/goradd/test/browsertest"
	"github.com/goradd/goradd/web/examples/controls"
)



type TextboxPanel struct {
	Panel
}

func NewTextboxPanel(ctx context.Context, parent page.ControlI) {
	p := &TextboxPanel{}
	p.Panel.Init(p, parent, "textboxPanel")

	p.Panel.AddControls(ctx,
		FormFieldCreator{
			ID:    "plainText-ff",
			Label: "Plain Text",
			Child: TextboxCreator{
				ID:        "plainText",
				SaveState: true,
			},
		},
		FormFieldCreator{
			ID:    "multiText-ff",
			Label: "Multi Text",
			Child: TextboxCreator{
				ID:        "multiText",
				SaveState: true,
				RowCount:  2,
			},
		},
		FormFieldCreator{
			ID:    "intText-ff",
			Label: "Integer Text",
			Child: IntegerTextboxCreator{
				ID:        "intText",
			},
		},
		FormFieldCreator{
			ID:    "floatText-ff",
			Label: "Float Text",
			Child: FloatTextboxCreator{
				ID:        "floatText",
			},
		},
		FormFieldCreator{
			ID:    "emailText-ff",
			Label: "Email Text",
			Child: EmailTextboxCreator{
				ID:        "emailText",
			},
		},
		FormFieldCreator{
			ID:    "passwordText-ff",
			Label: "Password",
			Child: TextboxCreator{
				ID:        "passwordText",
				Type:TextboxTypePassword,
			},
		},
		FormFieldCreator{
			ID:    "searchText-ff",
			Label: "Search",
			Child: TextboxCreator{
				ID:        "searchText",
				Type:TextboxTypeSearch,
			},
		},
		FormFieldCreator{
			ID:    "dateTimeText-ff",
			Label: "U.S. Date-time",
			Child: DateTextboxCreator{
				ID:        "dateTimeText",
			},
		},
		FormFieldCreator{
			ID:    "dateText-ff",
			Label: "Euro Date",
			Child: DateTextboxCreator{
				ID:        "dateText",
				Format: datetime.EuroDate,
			},
		},
		FormFieldCreator{
			ID:    "timeText-ff",
			Label: "U.S. Time",
			Child: DateTextboxCreator{
				ID:        "timeText",
				Format: datetime.UsTime,
			},
		},
		ButtonCreator{
			ID: "ajaxButton",
			Text: "Submit Ajax",
			SubmitAction:action.Ajax("textboxPanel", ButtonSubmit),
		},
		ButtonCreator{
			ID: "serverButton",
			Text: "Submit Server",
			SubmitAction:action.Server("textboxPanel", ButtonSubmit),
		},

	)
}

func (p *TextboxPanel) Action(ctx context.Context, a page.ActionParams) {
	switch a.ID {
	case ButtonSubmit:
	}
}

func init() {
	browsertest.RegisterTestFunction("Textbox Ajax Submit", testTextboxAjaxSubmit)
	browsertest.RegisterTestFunction("Textbox Server Submit", testTextboxServerSubmit)
	controls.RegisterPanel("textbox", "Textboxes", NewTextboxPanel, 2)
}

func testTextboxAjaxSubmit(t *browsertest.TestForm) {
	testTextboxSubmit(t, "ajaxButton")

	t.Done("Complete")
}

func testTextboxServerSubmit(t *browsertest.TestForm) {
	testTextboxSubmit(t, "serverButton")

	t.Done("Complete")
}

// testTextboxSubmit does a variety of submits using the given button. We use this to double check the various
// results we might get after a submission, as well as nsure that the ajax and server submits produce
// the same results.
func testTextboxSubmit(t *browsertest.TestForm, btnName string) {
	var myUrl = url.NewBuilder(controlsFormPath).SetValue("control", "textbox").SetValue("testing", 1).String()
	f := t.LoadUrl(myUrl)
	btn := f.Page().GetControl(btnName)

	t.ChangeVal("plainText", "me")
	t.ChangeVal("multiText", "me\nyou")
	t.ChangeVal("intText", "me")
	t.ChangeVal("floatText", "me")
	t.ChangeVal("emailText", "me")
	t.ChangeVal("dateTimeText", "me")
	t.ChangeVal("dateText", "me")
	t.ChangeVal("timeText", "me")

	t.Click(btn)

	t.AssertEqual("me", t.ControlValue("plainText"))
	t.AssertEqual("me\nyou", t.ControlValue("multiText"))
	t.AssertEqual("me", t.ControlValue("intText"))
	t.AssertEqual("me", t.ControlValue("floatText"))
	t.AssertEqual("me", t.ControlValue("emailText"))
	t.AssertEqual("me", t.ControlValue("dateTimeText"))
	t.AssertEqual("me", t.ControlValue("dateText"))
	t.AssertEqual("me", t.ControlValue("timeText"))

	t.AssertEqual(true, t.HasClass("intText-ff", "error"))
	t.AssertEqual(true, t.HasClass("floatText-ff", "error"))
	t.AssertEqual(true, t.HasClass("emailText-ff", "error"))
	t.AssertEqual(true, t.HasClass("dateText-ff", "error"))
	t.AssertEqual(true, t.HasClass("timeText-ff", "error"))
	t.AssertEqual(true, t.HasClass("dateTimeText-ff", "error"))

	plainText := GetTextbox(f,"plainText")
	intText := GetIntegerTextbox(f, "intText")
	floatText := GetFloatTextbox(f, "floatText")
	emailText := GetEmailTextbox(f, "emailText")
	dateText := GetDateTextbox(f, "dateText")
	timeText := GetDateTextbox(f, "timeText")
	dateTimeText := GetDateTextbox(f, "dateTimeText")

	plainText.SetInstructions("Sample instructions")
	t.ChangeVal("intText", 5)
	t.ChangeVal("floatText", 6.7)
	t.ChangeVal("emailText", "me@you.com")
	t.ChangeVal("dateText", "19/2/2018")
	t.ChangeVal("timeText", "4:59 am")
	t.ChangeVal("dateTimeText", "2/19/2018 4:23 pm")

	t.Click(btn)

	t.AssertEqual(5, intText.Int())
	t.AssertEqual(6.7, floatText.Float64())
	t.AssertEqual("me@you.com", emailText.Text())
	t.AssertEqual(datetime.NewDateTime("19/2/2018", datetime.EuroDate), dateText.Date())
	t.AssertEqual(datetime.NewDateTime("4:59 am", datetime.UsTime), timeText.Date())
	t.AssertEqual(datetime.NewDateTime("2/19/2018 4:23 pm", datetime.UsDateTime), dateTimeText.Date())

	t.AssertEqual(false, t.HasClass("intText-ff", "error"))
	t.AssertEqual(false, t.HasClass("floatText-ff", "error"))
	t.AssertEqual(false, t.HasClass("emailText-ff", "error"))
	t.AssertEqual(false, t.HasClass("dateText-ff", "error"))
	t.AssertEqual(false, t.HasClass("timeText-ff", "error"))
	t.AssertEqual(false, t.HasClass("dateTimeText-ff", "error"))
	t.AssertEqual("Sample instructions", t.InnerHtml("plainText-ff_inst"))

	t.AssertEqual("plainText-ff_lbl plainText", t.ControlAttribute("plainText", "aria-labelledby"))

	// Test SaveState
	f = t.LoadUrl(myUrl)
	plainText = GetTextbox(f, "plainText")
	t.AssertEqual("me", plainText.Text())
}
