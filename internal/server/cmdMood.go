package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allow players to set a few words of description to add to their xxx has
// entered message.
func (p *Player) CmdMood(mood string) {
	if mood == "" {
		if p.Mood() == "" {
			p.Outputm(text.MoodNotSet)
		} else {
			p.Outputm(text.Mood, p.Mood())
		}
		return
	}
	if len(mood) >= model.MOOD_SIZE {
		p.Outputm(text.MoodTooLong, model.MOOD_SIZE-1)
		return
	}
	if mood[0] == '/' || mood[0] == '>' {
		p.Outputm(text.MoodBadLeader)
		return
	}
	p.SetMood(mood)
	p.Outputm(text.MoodSet, p.Mood(), p.Name())
	p.Save(database.SaveNow)
}

// Unsets the player's current mood text.
func (p *Player) CmdMoodOff() {
	if p.Mood() == "" {
		p.Outputm(text.MoodNotSet)
		return
	}
	p.SetMood("")
	p.Outputm(text.MoodCleared)
	p.Save(database.SaveNow)
}
