package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

const (
	backgroundColor = 0x111111
)

var (
	window *ui.Window
)

// Colors for different teams
var teamColors = []uint32{
	0xFF0000,
	0x00FF00,
	0x0000FF,
	0xFFFF00,
	0xFF00FF,
	0x00FFFF,
}

var obstacleColor = mkSolidBrush(0x000000, 0xff)

// helper to quickly set a brush color
func mkSolidBrush(color uint32, alpha uint32) *ui.DrawBrush {
	brush := new(ui.DrawBrush)
	brush.Type = ui.DrawBrushTypeSolid
	component := uint8((color >> 16) & 0xFF)
	brush.R = float64(component) / 255
	component = uint8((color >> 8) & 0xFF)
	brush.G = float64(component) / 255
	component = uint8(color & 0xFF)
	brush.B = float64(component) / 255
	brush.A = float64(alpha) / 255
	return brush
}

type battleGridAreaHandler struct {
	bi *BattleInfo
}

func (ah *battleGridAreaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	// Draw the background
	path := ui.DrawNewPath(ui.DrawFillModeWinding)
	defer path.Free()
	path.AddRectangle(0, 0, dp.AreaWidth, dp.AreaHeight)
	path.End()

	dp.Context.Fill(path, mkSolidBrush(backgroundColor, 0xFF))

	// Draw the cells

	gx, gy := ah.bi.Size()
	cw, ch := dp.AreaWidth/float64(gx), dp.AreaHeight/float64(gy)

	cellAlpha := func(cell *Cell) uint32 {
		const minAlpha = 100
		return minAlpha + uint32(float64(cell.Power)*(255-minAlpha)/255)
	}

	for x := 0; x < gx; x++ {
		for y := 0; y < gy; y++ {
			x1, y1 := float64(x)*cw, float64(y)*ch

			if ah.bi.GetObstacle(x, y) {
				path := ui.DrawNewPath(ui.DrawFillModeWinding)
				path.AddRectangle(x1, y1, cw, ch)
				path.End()

				dp.Context.Fill(path, obstacleColor)

				path.Free()
			} else if cell := ah.bi.GetCell(x, y); cell != nil {

				path := ui.DrawNewPath(ui.DrawFillModeWinding)
				path.AddRectangle(x1, y1, cw, ch)
				path.End()

				dp.Context.Fill(path, mkSolidBrush(teamColors[cell.Team%len(teamColors)], cellAlpha(cell)))

				path.Free()
			}
		}
	}
}

// Events are ignored
func (ah *battleGridAreaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {}
func (ah *battleGridAreaHandler) MouseCrossed(a *ui.Area, left bool)           {}
func (ah *battleGridAreaHandler) DragBroken(a *ui.Area)                        {}
func (ah *battleGridAreaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// q for quit
	if ke.Key == 'q' {
		window.Destroy()
		ui.Quit()
	}
	return false
}

type statsAreaHandler struct {
	teamStrings []*ui.AttributedString
}

func NewStatsAreaHandler(bi *BattleInfo) *statsAreaHandler {
	ah := &statsAreaHandler{}
	for i, t := range bi.Teams {
		as := ui.NewAttributedString(t.Name())
		ah.teamStrings = append(ah.teamStrings, as)
	}
	return ah
}
func (ah *statsAreaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String: nil,
		Width:  100,
		Align:  ui.DrawTextAlignRight,
	})
}
func (ah *statsAreaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent)            {}
func (ah *statsAreaHandler) MouseCrossed(a *ui.Area, left bool)                      {}
func (ah *statsAreaHandler) DragBroken(a *ui.Area)                                   {}
func (ah *statsAreaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) { return false }

func createUI(bi *BattleInfo) func() {
	return func() {
		ah := battleGridAreaHandler{
			bi: bi,
		}
		a := ui.NewArea(&ah)
		vbox := ui.NewVerticalBox()
		vbox.Append(a, true)

		window = ui.NewWindow("Life Simulator", 640, 480, true)
		window.SetChild(vbox)
		window.OnClosing(func(w *ui.Window) bool {
			w.Destroy()
			ui.Quit()
			return false
		})
		ui.OnShouldQuit(func() bool {
			window.Destroy()
			return true
		})

		window.Show()
	}
}
