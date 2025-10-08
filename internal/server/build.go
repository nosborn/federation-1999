package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/build"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	// Explorer builds.
	ProjectLink model.Project = 1 + iota

	// Planetary investment builds.
	ProjectEducation
	ProjectEnergy
	ProjectHealth
	ProjectInfra
	ProjectSecurity

	// Duke puzzle builds.
	ProjectUpside
	ProjectDownside
	ProjectAccel
	ProjectMatProc
	ProjectC3
	ProjectMatTrans
	ProjectTimeMach
)

type BuildComponent struct {
	Index    model.Commodity
	Quantity int32
}

type BuildTemplate struct {
	IDProject   model.Project
	VocProject  int // TODO: see if this is used
	Name        string
	LowRank     model.Rank
	HighRank    model.Rank
	Components  []BuildComponent
	Labour      int32
	Cash        int32
	Duration    int32 // time to build (seconds)
	TimerPeriod int32 // update timer frequency (seconds)
}

var BuildTemplates = []BuildTemplate{
	// Explorer builds.
	{
		build.LINK,
		1108, // TODO: see if this is used
		"Link",
		model.RankExplorer, model.RankExplorer,
		[]BuildComponent{
			{model.CommodityPowerPacks, 3850},
			{model.CommodityRads, 3700},
			{model.CommodityHypnotapes, 850},
			{model.CommoditySensAmps, 1435},
			{model.CommodityKatydidics, 1985},
			{model.CommoditySims, 645},
		},
		0,
		0,    // price is covered by planet purchase
		3600, // 1 hour
		36,   // 36 seconds (1%)
	},

	// Plenetary investment builds.
	{
		build.EDUCATION,
		1102, // TODO: see if this is used
		"Education",
		model.RankSquire, model.RankBaron,
		[]BuildComponent{
			{model.CommodityVidi, 2100},
			{model.CommodityTools, 870},
			{model.CommodityRNA, 1810},
			{model.CommoditySpices, 400},
			{model.CommodityStock, 500},
			{model.CommodityStudios, 4370},
			{model.CommodityArtifacts, 1900},
		},
		130,
		400000000, // 400,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.ENERGY,
		1101, // TODO: see if this is used
		"Energy",
		model.RankSquire, model.RankBaron,
		[]BuildComponent{
			{model.CommodityVidi, 400},
			{model.CommodityAntiMatter, 1070},
			{model.CommodityPowerPacks, 2000},
			{model.CommodityLubOils, 1650},
			{model.CommodityNitros, 930},
			{model.CommodityRads, 4900},
			{model.CommodityKatydidics, 1570},
			{model.CommodityUnivators, 500},
		},
		70,
		400000000, // 400,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.HEALTH,
		1105, // TODO: see if this is used
		"Health",
		model.RankSquire, model.RankBaron,
		[]BuildComponent{
			{model.CommodityPharms, 5600},
			{model.CommodityRNA, 2530},
			{model.CommodityStock, 1630},
			{model.CommodityFurs, 1200},
			{model.CommodityHypnotapes, 3000},
		},
		160,
		400000000, // 400,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.INFRA,
		1104, // TODO: see if this is used
		"Infra",
		model.RankSquire, model.RankBaron,
		[]BuildComponent{
			{model.CommodityElectros, 900},
			{model.CommodityGenerators, 1700},
			{model.CommodityPolymers, 2000},
			{model.CommodityWoods, 1000},
			{model.CommodityNickel, 2100},
			{model.CommodityAlloys, 5300},
			{model.CommodityLibraries, 700},
		},
		60,
		400000000, // 400,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.SECURITY,
		1106, // TODO: see if this is used
		"Security",
		model.RankSquire, model.RankBaron,
		[]BuildComponent{
			{model.CommodityWeapons, 3600},
			{model.CommodityVidi, 1500},
			{model.CommodityDroids, 1390},
			{model.CommodityPropellants, 450},
			{model.CommodityNitros, 3340},
			{model.CommodityUnivators, 800}, // ?? spec says Libraries ??
			{model.CommoditySims, 700},
		},
		200,
		400000000, // 400,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},

	// Duke puzzle builds.
	{
		build.UPSIDE,
		1069, // TODO: see if this is used
		"Upside",
		model.RankBaron, model.RankDuke,
		[]BuildComponent{
			{model.CommodityDroids, 5000},
			{model.CommodityAlloys, 40000},
			{model.CommodityCrystals, 6500},
			{model.CommodityMonopoles, 5000},
			{model.CommodityGold, 12000},
			{model.CommodityXmetals, 2500},
			{model.CommodityControllers, 17000},
			{model.CommodityGAsChips, 5000},
			{model.CommodityMasers, 10500},
			{model.CommodityElectros, 20000},
			{model.CommodityTools, 4750},
		},
		250,
		350000000, // 350,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.DOWNSIDE,
		1070, // TODO: see if this is used
		"Downside",
		model.RankBaron, model.RankDuke,
		[]BuildComponent{
			{model.CommodityAlloys, 25000},
			{model.CommodityLanzariK, 35000},
			{model.CommodityAlloys, 15000},
			{model.CommodityControllers, 5000},
			{model.CommodityGenerators, 20000},
			{model.CommodityNickel, 22000},
			{model.CommodityMechParts, 8750},
		},
		300,
		250000000, // 250,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.ACCEL,
		1072, // TODO: see if this is used
		"BeV",
		model.RankBaron, model.RankBaron,
		[]BuildComponent{
			{model.CommodityMonopoles, 20000},
			{model.CommodityAlloys, 10000},
			{model.CommodityXmetals, 9000},
			{model.CommodityPolymers, 22500},
			{model.CommodityBioChips, 7500},
			{model.CommodityGAsChips, 6000},
			{model.CommodityMasers, 3250},
			{model.CommodityGold, 10000},
		},
		400,
		650000000, // 650,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.MATPROC,
		1073, // TODO: see if this is used
		"MatProc",
		model.RankBaron, model.RankBaron,
		[]BuildComponent{
			{model.CommodityPolymers, 50000},
			{model.CommodityMechParts, 22500},
			{model.CommodityControllers, 10000},
			{model.CommoditySynths, 23500},
			{model.CommodityDroids, 7500},
			{model.CommodityNickel, 12500},
			{model.CommodityAlloys, 3250},
		},
		100,
		625000000, // 625,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.C3,
		1074, // TODO: see if this is used
		"C3",
		model.RankBaron, model.RankBaron,
		[]BuildComponent{
			{model.CommodityAlloys, 5000},
			{model.CommodityNickel, 17500},
			{model.CommodityPolymers, 15500},
			{model.CommodityGold, 30000},
			{model.CommodityMonopoles, 20000},
			{model.CommodityGAsChips, 5500},
			{model.CommodityBioChips, 25000},
			{model.CommodityVidi, 11500},
			{model.CommodityLanzariK, 6250},
			{model.CommoditySensAmps, 7250},
		},
		150,
		175000000, // 175,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
	{
		build.MATTRANS,
		1071, // TODO: see if this is used
		"MatTrans",
		model.RankBaron, model.RankDuke,
		[]BuildComponent{
			{model.CommodityAlloys, 10000},
			{model.CommodityNickel, 10000},
			{model.CommodityPolymers, 50000},
			{model.CommodityBioChips, 30000},
			{model.CommoditySynths, 15000},
			{model.CommodityCrystals, 17500},
		},
		250,
		1350000000, // 1,350,000,000 IG
		3600,       // 1 hour
		36,         // 36 seconds (1%)
	},
	{
		build.TIMEMACH,
		1075, // TODO: see if this is used
		"HGW",
		model.RankBaron, model.RankBaron,
		[]BuildComponent{
			{model.CommodityCrystals, 50000},
			{model.CommodityXmetals, 10000},
			{model.CommodityNickel, 5000},
			{model.CommodityAlloys, 6500},
			{model.CommodityMonopoles, 21500},
			{model.CommodityBioChips, 13500},
			{model.CommodityLanzariK, 6500},
			{model.CommodityRNA, 16500},
		},
		500,
		800000000, // 800,000,000 IG
		3600,      // 1 hour
		36,        // 36 seconds (1%)
	},
}

func getBuildTemplate(id model.Project) (*BuildTemplate, bool) {
	for i := range BuildTemplates {
		if BuildTemplates[i].IDProject == id {
			return scaleBuildTemplate(&BuildTemplates[i]), true
		}
	}
	return nil, false
}

// func getBuildTemplateById(id int) (*BuildTemplate, bool) {
// 	for i := range BuildTemplates {
// 		if BuildTemplates[i].IDProject == model.Project(id) {
// 			return scaleBuildTemplate(&BuildTemplates[i]), true
// 		}
// 	}
// 	return nil, false
// }

func scaleBuildTemplate(template *BuildTemplate) *BuildTemplate {
	scaled := *template
	if global.TestFeaturesEnabled {
		scaled.Components = make([]BuildComponent, len(template.Components))
		copy(scaled.Components, template.Components)
		for i := range scaled.Components {
			scaled.Components[i].Quantity /= 100
		}
		scaled.Labour /= 10
		scaled.Cash /= 10
		scaled.Duration /= 10
	}
	return &scaled
}

func startBuild(player *Player, template *BuildTemplate) bool {
	if player.ownSystem == nil || player.ownSystem.IsClosed() {
		player.Outputm(text.BUILD_PLANET_CLOSED)
		return false
	}

	// Find the player's mega-warehouse.
	storage := player.storage
	if storage == nil || storage.Warehouse[0] == nil {
		player.Outputm(text.NO_MEGA_WAREHOUSE)
		return false
	}
	warehouse := storage.Warehouse[0]
	debug.Check(warehouse.Planet == player.ownSystem.Name())

	/* Scale the build if necessary */

	var scale int32
	switch template.IDProject {
	case build.UPSIDE:
		scale = 100 - player.Facilities[model.DK_UPSIDE]
	case build.DOWNSIDE:
		scale = 100 - player.Facilities[model.DK_DOWNSIDE]
	case build.ACCEL:
		scale = 100 - player.Facilities[model.DK_ACCEL]
	case build.MATPROC:
		scale = 100 - player.Facilities[model.DK_MATPROC]
	case build.C3:
		scale = 100 - player.Facilities[model.DK_C3]
	case build.MATTRANS:
		scale = 100 - player.Facilities[model.DK_MATTRANS]
	case build.TIMEMACH:
		scale = 100 - player.Facilities[model.DK_TIMEMACH]
	default:
		scale = 100
	}
	if scale <= 0 {
		player.Outputm(text.BUILD_FACILITY_INTACT, template.Name)
		return false
	}
	// if scale < 100 {
	// 	// TODO:
	// 	// memcpy(&bldtRepair, pbldt, sizeof(bldtRepair));
	// 	// for (size_t index = 0; index < COMPONENTS; ++index) {
	// 	// 	BLDC *pbldc = &bldtRepair.abldc[index];
	// 	// 	if (pbldc->quantity > 0) {
	// 	// 		pbldc->quantity = (pbldc->quantity * nScale) / 100;
	// 	// 	}
	// 	// }
	// 	// bldtRepair.labour = (bldtRepair.labour * nScale) / 100;
	// 	// bldtRepair.cash = (bldtRepair.cash * nScale) / 100;
	// 	// bldtRepair.tDuration = (bldtRepair.tDuration * nScale) / 100;
	// 	// template = &repair
	// }

	// See if there's enough raw material in the warehouse.
	hasStock := true
	for i := range template.Components {
		required := template.Components[i].Quantity
		if required < 0 {
			continue
		}
		for j := range 20 {
			bay := &warehouse.Bay[j]
			if bay.Type == template.Components[i].Index && bay.Quantity > 0 {
				required -= bay.Quantity
			}
		}
		if required > 0 {
			hasStock = false
		}
	}
	if !hasStock {
		player.Outputm(text.BUILD_NO_STOCK)
		return false
	}
	if player.ownSystem.Balance() < template.Cash {
		player.Outputm(text.NO_TREASURY_FUNDS)
		return false
	}

	planet := player.ownSystem.Planets()[0]
	if planet.population-2 < template.Labour { // Need 2 left in order to breed!
		player.Outputm(text.BUILD_LABOUR_SHORTAGE)
		return false
	}

	switch template.IDProject {
	case build.ENERGY, build.EDUCATION, build.HEALTH, build.INFRA, build.SECURITY:
		investment := planet.education + planet.energy + planet.health + planet.infstr + planet.intSecurity
		if investment >= int32(planet.level)*10 || investment >= 50 {
			player.Outputm(text.BUILD_INVESTMENT_LIMIT)
			return false
		}
	default: // make linter happy
	}

	for i := range template.Components {
		required := template.Components[i].Quantity
		if required < 0 {
			continue
		}
		for j := range 20 {
			bay := &warehouse.Bay[j]
			if bay.Type == template.Components[i].Index && bay.Quantity > 0 {
				if bay.Quantity <= required {
					required -= bay.Quantity
					bay = &model.Cargo{}
				} else {
					bay.Quantity -= required
					// required = 0
					break
				}
			}
		}
	}

	player.BuildProject = template.IDProject
	player.BuildDuration = template.Duration
	player.BuildElapsed = 0

	planet.Expenditure(template.Cash)
	planet.population -= template.Labour

	// FIXME: online players only
	for _, company := range AllCompanies() {
		company.StopFactories(planet)
	}

	player.Outputm(text.BUILD_STARTED, template.Name)

	log.Printf("%s build started for %s", template.Name, player.Name())
	return true
}

// StartLinkBuild is really part of the planet ordering process rather than a
// build process, so any messages should reflect that.
func StartLinkBuild(player *Player, planetCost int32) bool {
	// TODO:
	return false
}
