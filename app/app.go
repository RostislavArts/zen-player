package app

import (
	"log"

	"zen-player/ui"
	"zen-player/player"
)

func Run() {
	var err error
	p := player.NewPlayer()
	u := ui.NewUI()

	err = u.InitScreen()
	if err != nil {
		log.Println(err)
		return
	}
	defer u.Cleanup()

	u.StartTicker(p)

	err = p.Init()
	if err != nil {
		u.Cleanup()
		log.Println(err)
		return
	}

	err = p.Start()
	if err != nil {
		u.Cleanup()
		log.Println(err)
		return
	}

	u.EventLoop(p)
}

