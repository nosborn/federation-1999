package server

func multiTalk(message string, omit *Player) {
	for _, p := range Players {
		if !p.IsPlaying() {
			continue
		}
		p.Nsoutput(message)
		p.FlushOutput()
	}
}
