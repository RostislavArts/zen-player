package player

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	audiopanel "zen-player/audio_panel"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/flac"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func (p *Player) playLooped(data []byte) error {
	reader := bytes.NewReader(data)
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		return err
	}

	buf := beep.NewBuffer(format)
	buf.Append(streamer)
	looped := beep.Loop(-1, buf.Streamer(0, buf.Len()))

	// Volume works only with last noise. Need to change that
	ctrl := &beep.Ctrl{Streamer: looped, Paused: false}
	volume := &effects.Volume{Streamer: ctrl, Base: 2}

	p.noisePanel.Noises = append(p.noisePanel.Noises, audiopanel.Noise{Ctrl: ctrl, Volume: volume})

	speaker.Play(volume)
	return nil
}

func (p *Player) playFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	var streamer beep.StreamSeekCloser

	switch strings.ToLower(filepath.Ext(path)) {
	case ".mp3":
		streamer, _, err = mp3.Decode(file)
	case ".wav":
		streamer, _, err = wav.Decode(file)
	case ".flac":
		streamer, _, err = flac.Decode(file)
	}
	if err != nil {
		return err
	}

	ctrl := &beep.Ctrl{Streamer: streamer, Paused: false}
	p.trackPanel.Streamer = streamer
	p.trackPanel.Ctrl = ctrl

	done = make(chan bool)

	speaker.Play(beep.Seq(p.trackPanel.Ctrl,
	beep.Callback(func() {
		done <- true
	})))

	return nil
}

