package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

func notify() {

	snooze := 0
	for snooze < 2 {

		err := beeep.Notify("Title", "Time To Eat", "/home/joel/Pictures/wallpapers/4.png")
		if err != nil {
			panic(err)
		}

		for i := 0; i < 20; i++ {
			err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			time.Sleep(time.Millisecond * 120)
			time.Sleep(time.Millisecond * 120)
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(time.Minute)
		snooze++
	}
}
