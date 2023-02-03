package view

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/one2nc/cloud-lens/internal/ui"
	"github.com/rs/zerolog/log"
)

type EC2 struct {
	ResourceViewer
}

// NewPod returns a new viewer.
func NewEC2(resource string) ResourceViewer {
	var e EC2
	e.ResourceViewer = NewBrowser(resource)
	e.AddBindKeysFn(e.bindKeys)
	e.GetTable().SetEnterFn(e.describeInstace)
	return &e
}

func (e *EC2) bindKeys(aa ui.KeyActions) {
	aa.Add(ui.KeyActions{
		ui.KeyShiftI:    ui.NewKeyAction("Sort Instance-Id", e.GetTable().SortColCmd("Instance-Id", true), true),
		ui.KeyShiftS:    ui.NewKeyAction("Sort Instance-State", e.GetTable().SortColCmd("Instance-State", true), true),
		ui.KeyShiftT:    ui.NewKeyAction("Sort Instance-Type", e.GetTable().SortColCmd("Instance-Type", true), true),
		ui.KeyShiftL:    ui.NewKeyAction("Sort Launch-Time", e.GetTable().SortColCmd("Launch-Time", true), true),
		ui.KeyShiftM:    ui.NewKeyAction("Sort Monitoring-State", e.GetTable().SortColCmd("Monitoring-State", true), true),
		ui.KeyShiftP:    ui.NewKeyAction("Sort Public-DNS", e.GetTable().SortColCmd("Public-DNS", true), true),
		tcell.KeyEscape: ui.NewKeyAction("Back", e.App().PrevCmd, true),
		tcell.KeyEnter:  ui.NewKeyAction("View", e.enterCmd, true),
	})
}

func (e *EC2) describeInstace(app *App, model ui.Tabular, resource string) {
	log.Info().Msg(fmt.Sprintf("TODO: describe: %v", resource))
	// if err := app.inject(co); err != nil {
	// 	app.Flash().Err(err)
	// }
}

func (e *EC2) enterCmd(evt *tcell.EventKey) *tcell.EventKey {
	instanceId := e.GetTable().GetSelectedItem()
	e.App().Flash().Info("Instance-Id: " + instanceId)

	return nil
}
