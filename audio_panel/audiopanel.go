package audiopanel

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
)

type TrackPanel struct {
	Format     beep.Format
	SampleRate beep.SampleRate
	Streamer   beep.StreamSeekCloser
	Ctrl       *beep.Ctrl
}

type NoisePanel struct {
	Format     beep.Format
	SampleRate beep.SampleRate
	Noises     []Noise
}

type Noise struct {
	Ctrl   *beep.Ctrl
	Volume *effects.Volume
}

