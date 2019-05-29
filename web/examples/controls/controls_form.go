package controls

import (
	"context"
	"github.com/goradd/goradd/pkg/page"
	. "github.com/goradd/goradd/pkg/page/control"
	"github.com/goradd/goradd/pkg/page/control/data"
	"github.com/goradd/goradd/pkg/url"
	"sort"
)

const ControlsFormPath = "/goradd/examples/controls.g"
const ControlsFormId = "ControlsForm"

const (
	TestButtonAction = iota + 1
)

type ControlsForm struct {
	FormBase
	list 	 		*UnorderedList
	detail 		  *Panel
}

type createFunction func(ctx context.Context, parent page.ControlI)
type controlEntry struct {
	key string
	name string
	f createFunction
	order int
}

var controls []controlEntry

func NewControlsForm(ctx context.Context) page.FormI {
	f := &ControlsForm{}
	f.Init(ctx, f, ControlsFormPath, ControlsFormId)
	f.AddRelatedFiles()

	f.list = NewUnorderedList(f, "panelList")
	f.list.SetDataProvider(f)
	f.detail = NewPanel(f, "detailPanel")

	return f
}

func (f *ControlsForm) LoadControls(ctx context.Context) {
	var createF createFunction
	if id, ok := page.GetContext(ctx).FormValue("control"); ok {
		for _,c := range controls {
			if c.key == id {
				createF = c.f
			}
		}
	}

	if createF == nil {
		createF = controls[0].f
	}

	createF(ctx, f.detail)
}

func (f *ControlsForm) BindData(ctx context.Context, s data.DataManagerI) {
	sort.Slice(controls, func(i,j int) bool {
		return controls[i].order < controls[j].order
	});
	pageContext := page.GetContext(ctx)
	for _,c := range controls {
		item := f.list.AddItem(c.name, c.key)
		a := url.
			NewBuilderFromUrl(*pageContext.URL).
			SetValue("control", c.key).
			String()
		item.SetAnchor(a)
	}
}

func RegisterPanel(key string,
		name string,
		f createFunction,
		order int) {
	controls = append(controls, controlEntry{key, name, f, order})
}

func init() {
	page.RegisterPage(ControlsFormPath, NewControlsForm, ControlsFormId)
}

