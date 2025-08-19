package player

import (
	"time"

	"github.com/gopxl/beep/speaker"
)

func (p *Player) Pause() {
	speaker.Lock()
	defer speaker.Unlock()

	if p.trackPanel.Ctrl != nil {
		p.trackPanel.Ctrl.Paused = !p.trackPanel.Ctrl.Paused
		p.TrackPaused = !p.TrackPaused
	}
}

func (p *Player) NoiseOff() {
	speaker.Lock()
	defer speaker.Unlock()

	if p.noisePanel.Noises != nil {
		for _, ctrl := range p.noisePanel.Noises {
			ctrl.Ctrl.Paused = !ctrl.Ctrl.Paused
		}
		p.NoisePaused = !p.NoisePaused
	}
}

func (p *Player) NextTrack() {
	speaker.Lock()
	defer speaker.Unlock()

	if p.trackPanel.Ctrl != nil {
		p.trackPanel.Ctrl.Paused = true
	}
	if p.trackPanel.Streamer != nil {
		_ = p.trackPanel.Streamer.Close()
	}

	done <- true
}

func (p *Player) SeekLeft() {
	speaker.Lock()
	defer speaker.Unlock()

	if p.trackPanel.Streamer != nil {
		newPos := p.trackPanel.Streamer.Position() - p.trackPanel.SampleRate.N(time.Second)
		newPos = max(newPos, 0)
		p.trackPanel.Streamer.Seek(newPos)
	}
}

func (p *Player) SeekRight() {
	speaker.Lock()
	defer speaker.Unlock()

	if p.trackPanel.Streamer != nil {
		newPos := p.trackPanel.Streamer.Position() + p.trackPanel.SampleRate.N(time.Second)
		newPos = min(newPos, p.trackPanel.Streamer.Len() - 1)
		p.trackPanel.Streamer.Seek(newPos)
	}
}

