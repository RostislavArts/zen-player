package ui

import (
	"time"

	"zen-player/player"

	"github.com/gdamore/tcell/v2"
)

type UI struct {
	screen tcell.Screen
	style tcell.Style
	page int
}

func NewUI() *UI {
	return &UI{}
}

func (u *UI) InitScreen() error {
	var err error

	u.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset)
	u.screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}

	err = u.screen.Init()
	if err != nil {
		return err
	}

	u.screen.SetStyle(u.style)
	u.screen.Clear()

	return nil
}

func (u *UI) EventLoop(p *player.Player) {
	for {
		ev := u.screen.PollEvent()

		switch tev := ev.(type) {
		case *tcell.EventResize:
			u.screen.Sync()
		case *tcell.EventKey:
			switch tev.Rune() {
			case 'b':
				p.NoiseOff()
			case 'p':
				p.Pause()
			case 'n':
				p.NextTrack()
			case 'a':
				p.SeekLeft()
			case 's':
				p.SeekRight()
			case 'z':
				p.ChangeNoiseVolume(-5)
			case 'x':
				p.ChangeNoiseVolume(+5)
			case 'q':
				return
			}
		}
	}
}

func (u *UI) StartTicker(p *player.Player) {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			u.screen.Clear()
			u.drawPlayer(p)
			u.drawSongList(p)
			u.screen.Show()
		}
	}()
}

func (u *UI) Cleanup() {
	u.screen.Fini()
}

