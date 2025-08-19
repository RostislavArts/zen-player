package player

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"
	"path/filepath"

	"zen-player/fsutil"
	"zen-player/utils"
	ap "zen-player/audio_panel"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/wav"
	"github.com/gopxl/beep/flac"
	"github.com/gopxl/beep/speaker"
)

type Player struct {
	noisePanel ap.NoisePanel
	trackPanel ap.TrackPanel

	noiseList []string
	TrackList []string
	CurrentTrack string
	CurrentTrackIndex int

	TrackPaused bool
	NoisePaused bool
}

var done chan bool
var flags = utils.ParseFlags()

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Init() error {
	var err error

	p.TrackList, err = fsutil.ParseFilesFromDir(flags.Path)
	if err != nil {
		return err
	}

	if flags.Shuffle {
		utils.ShuffleSlice(&p.TrackList)
	}

	if len(p.TrackList) == 0 {
		return fmt.Errorf("No tracks found")
	}
	firstFile, err := os.Open(p.TrackList[0])
	if err != nil {
		return err
	}
	defer firstFile.Close()

	var format beep.Format

	switch strings.ToLower(filepath.Ext(p.TrackList[0])) {
	case ".mp3":
		_, format, err = mp3.Decode(firstFile)
	case ".wav":
		_, format, err = wav.Decode(firstFile)
	case ".flac":
		_, format, err = flac.Decode(firstFile)
	}
	if err != nil {
		return err
	}

	p.trackPanel.Format = format
	p.trackPanel.SampleRate = format.SampleRate

	p.noisePanel.Format = format
	p.noisePanel.SampleRate = format.SampleRate

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return nil
}

func (p *Player) Start() error {
	noises := flags.NoiseList

	if len(noises) > 0 {
		for _, name := range noises {
			data, ok := noiseMap[name]
			if !ok {
				log.Printf("Unknown noise: %s", name)
				continue
			}
			if err := p.playLooped(data); err != nil {
				return err
			}
		}
	} else {
		err := p.playLooped(rainNoise)
		if err != nil {
			return err
		}
	}

	go func() {
		for {
			for _, track := range p.TrackList {
				p.CurrentTrack = filepath.Base(track)

				err := p.playFile(track)
				if err != nil {
					log.Println("Error playing track:", err)
					continue
				}
				p.CurrentTrackIndex++
				<-done
			}

			p.CurrentTrackIndex = 0

			if !flags.Loop {
				break
			}
		}
	}()

	return nil
}

func (p *Player) GetPosition() time.Duration {
	if p.trackPanel.Streamer == nil {
		return 0
	}

	streamer := p.trackPanel.Streamer
	position := p.trackPanel.SampleRate.D(streamer.Position())
	return position.Round(time.Second)
}

func (p *Player) GetTrackLength() time.Duration {
	if p.trackPanel.Streamer == nil {
		return 0
	}

	streamer := p.trackPanel.Streamer
	position := p.trackPanel.SampleRate.D(streamer.Len())
	return position.Round(time.Second)
}

