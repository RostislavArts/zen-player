package ui

import (
	"fmt"
	"path/filepath"

	"zen-player/player"

	"github.com/gdamore/tcell/v2"
)

func (u *UI) drawText(x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		u.screen.SetContent(col, row, r, nil, u.style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (u *UI) drawBox(x1, y1, x2, y2 int) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			u.screen.SetContent(col, row, ' ', nil, u.style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		u.screen.SetContent(col, y1, tcell.RuneHLine, nil, u.style)
		u.screen.SetContent(col, y2, tcell.RuneHLine, nil, u.style)
	}
	for row := y1 + 1; row < y2; row++ {
		u.screen.SetContent(x1, row, tcell.RuneVLine, nil, u.style)
		u.screen.SetContent(x2, row, tcell.RuneVLine, nil, u.style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		u.screen.SetContent(x1, y1, tcell.RuneULCorner, nil, u.style)
		u.screen.SetContent(x2, y1, tcell.RuneURCorner, nil, u.style)
		u.screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, u.style)
		u.screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, u.style)
	}
}

func (u *UI) drawPlayer(p *player.Player) {
	// Calculating coordinates to center everything
	scrW, _ := u.screen.Size()
	var boxW int = 48
	x1 := scrW / 2 - boxW / 2 // left edge
	x2 := scrW / 2 + boxW / 2 // right edge

	u.drawBox(x1, 0, x2, 5)
	u.drawText(x1 + 2, 1, x2 - 1, 1, fmt.Sprintf("Now Playing: %s ", p.CurrentTrack))
	u.drawText(x1 + 2, 2, x2 - 1, 2, fmt.Sprintf("(%v / %v)", p.GetPosition(), p.GetTrackLength()))

	if p.NoisePaused {
		u.drawText(x1 + 2, 3, x2 - 1, 3, fmt.Sprintf("[B]ackground Noise: OFF ([Z/X] Vol: %v%%)", p.GetNoiseVolume()))
	} else {
		u.drawText(x1 + 2, 3, x2 - 1, 3, fmt.Sprintf("[B]ackground Noise: ON  ([Z/X] Vol: %v%%)", p.GetNoiseVolume()))
	}

	if p.TrackPaused {
		u.drawText(x1 + 2, 4, x2 - 1, 4, "[P]lay   [N]ext  [A/S] Seek  [Q]uit")
	} else {
		u.drawText(x1 + 2, 4, x2 - 1, 4, "[P]ause  [N]ext  [A/S] Seek  [Q]uit")
	}
}

func (u *UI) drawSongList(p *player.Player) {
	scrW, scrH := u.screen.Size()
	var boxW = 48

	x1 := scrW / 2 - boxW / 2 // left edge
	x2 := scrW / 2 + boxW / 2 // right edge
	pageSize := scrH - 6

	for i, track := range p.TrackList {
		// need to change logic here
		if p.CurrentTrackIndex - u.page * pageSize > pageSize {
			u.page++
		} else if p.CurrentTrackIndex - u.page * pageSize < 0 {
			u.page = 0
		}

		if i - u.page * pageSize >= 0 {
			if i == p.CurrentTrackIndex - 1 {
				u.drawText(x1, 6 + i - u.page * pageSize, x2, 6 + i - u.page * pageSize, fmt.Sprintf(">> %s", filepath.Base(track)))
			} else {
				u.drawText(x1, 6 + i - u.page * pageSize, x2, 6 + i - u.page * pageSize, filepath.Base(track))
			}
		}
	}
}

