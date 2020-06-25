package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

func notify(task string) {

	snooze := 0
	for snooze < 2 {
		err := beeep.Notify("URGENT️️️️️️️️️️️📢", task, "/home/joel/Pictures/wallpapers/4.png")
		if err != nil {
			panic(err)
		}

		for i := 0; i < 20; i++ {
			err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			time.Sleep(time.Millisecond * 150)
			time.Sleep(time.Millisecond * 150)
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(time.Minute)
		snooze++
	}
}
