package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/build"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdBuild(project model.Project) {
	if p.rank < model.RankSquire || p.rank > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	if p.OwnSystem().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.OwnSystem().Name())
		return
	}

	template, ok := getBuildTemplate(project)
	if !ok {
		p.Outputm(text.BUILD_UNKNOWN_PROJECT)
		return
	}
	if p.rank < template.LowRank || p.rank > template.HighRank {
		p.Outputm(text.BUILD_WRONG_RANK)
		return
	}
	if p.IsPromoCharacter() {
		p.Outputm(text.BUILD_WRONG_RANK)
		return
	}
	if p.BuildProject != build.NOTHING {
		p.Outputm(text.BUILD_ALREADY_ACTIVE)
		return
	}
	if !startBuild(p, template) {
		return
	}
	p.Save(database.SaveNow)

	// Start the build timer.
	p.resumeConstruction()
}
