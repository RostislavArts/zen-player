download:
	echo "Downloading additional resources. Please wait"
	wget -O player/resources/night.mp3 https://cdn.pixabay.com/download/audio/2022/02/07/audio_51b5acd355.mp3?filename=night-ambience-17064.mp3
	wget -O player/resources/rain.mp3 https://cdn.pixabay.com/download/audio/2024/10/30/audio_42e6870f29.mp3?filename=calming-rain-257596.mp3
	wget -O player/resources/river.mp3 https://cdn.pixabay.com/download/audio/2024/10/01/audio_d9e2d28b63.mp3?filename=flowing-water-246403.mp3
	wget -O player/resources/sea.mp3 https://cdn.pixabay.com/download/audio/2023/10/02/audio_c695edda8e.mp3?filename=sea-waves-169411.mp3
	wget -O player/resources/thunder.mp3 https://cdn.pixabay.com/download/audio/2025/06/23/audio_5bbbf0c7f3.mp3?filename=dry-thunder-364468.mp3
	wget -O player/resources/wind.mp3 https://cdn.pixabay.com/download/audio/2021/12/24/audio_f392b97f60.mp3?filename=wind-blowing-sfx-12809.mp3

build:
	download
	echo "Building..."
	go build -o zen

install:
	build
	echo "Moving built app to /usr/bin (you need to execute it as sudo)"
	sudo mv zen /usr/bin

