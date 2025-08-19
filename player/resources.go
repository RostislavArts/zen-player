package player

import _ "embed"

//go:embed resources/rain.mp3
var rainNoise []byte

//go:embed resources/river.mp3
var riverNoise []byte

//go:embed resources/night.mp3
var nightNoise []byte

//go:embed resources/thunder.mp3
var thunderNoise []byte

//go:embed resources/wind.mp3
var windNoise []byte

//go:embed resources/sea-waves.mp3
var seaNoise []byte

var noiseMap = map[string][]byte{
	"rain":  rainNoise,
	"river": riverNoise,
	"night": nightNoise,
	"thunder": thunderNoise,
	"wind":  windNoise,
	"sea": seaNoise,
}

