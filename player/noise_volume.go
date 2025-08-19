package player

import (
	"math"
)

func (p *Player) GetNoiseVolume() int {
	if p.noisePanel.Noises == nil {
		return 0
	}

	return DBToPercent(p.noisePanel.Noises[0].Volume.Volume)
}

func PercentToDB(percent float64) float64 {
	if percent <= 0 {
		return -40.0
	}
	if percent > 100 {
		percent = 100
	}
	return 20 * math.Log10(percent / 100.0)
}

func DBToPercent(db float64) int {
	if db <= -40.0 {
		return 0
	}
	p := math.Pow(10, db/20.0) * 100
	if p > 100 {
		p = 100
	}
	return int(p + 0.5)
}

func (p *Player) ChangeNoiseVolume(delta float64) {
	curPercent := float64(DBToPercent(p.noisePanel.Noises[0].Volume.Volume))
	newPercent := curPercent + delta
	if newPercent < 0 {
		newPercent = 0
	} else if newPercent > 100 {
		newPercent = 100
	}

	for _, vol := range p.noisePanel.Noises {
		vol.Volume.Volume = PercentToDB(newPercent)
	}
}

