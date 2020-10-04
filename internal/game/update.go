package game

import (
	"time"

	"github.com/brianseitel/mudder/internal/positions"
)

func ticker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				updateViolence()
			}
		}
	}()
}

func updateViolence() {
	ch := gameWorld.Character
	if ch.Fighting == nil {
		return
	}

	victim := ch.Fighting

	if ch.Position >= positions.POS_RESTING && ch.CurrentRoom == victim.CurrentRoom {
		multiHit(ch, victim)
	} else {
		stopFighting(ch, victim)
	}

	if victim.Fighting == nil {
		return
	}

	if victim.Position > positions.POS_RESTING && ch.CurrentRoom == victim.CurrentRoom {
		multiHit(victim, ch)
	} else {
		stopFighting(victim, ch)
	}
}
