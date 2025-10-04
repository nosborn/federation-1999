package goods

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	FirstCommodity = model.CommodityGAsChips
	LastCommodity  = model.CommoditySims
)

const (
	MAX_GOODS = 52
)

type InputCommodity struct {
	Commodity model.Commodity
	Quantity  int32
}

type TradeGoods struct { // trading goods basic parameters
	Name      string
	BasePrice uint32
	Group     model.CommodityGroup
	Labour    uint32           // labour required for 100 tons output
	Message   [4]text.MsgNum   // message number describing facility
	Inputs    []InputCommodity // inputs required for 100 tons output
}

var GoodsArray = [MAX_GOODS]TradeGoods{
	{
		"GAs-chips",              // name
		450,                      // basePrice
		model.GroupTechnological, // type
		60,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_FabricationPlant,
			text.Factory_FabricationPlant_Loc,
			text.Factory_FabricationPlant_Closed,
			text.Factory_FabricationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 60},
			{model.CommodityNitros, 10},
		},
	},
	{
		"Bio-chips",              // name
		650,                      // basePrice
		model.GroupTechnological, // type
		100,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_FabricationPlant,
			text.Factory_FabricationPlant_Loc,
			text.Factory_FabricationPlant_Closed,
			text.Factory_FabricationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityRNA, 60},
			{model.CommodityMonopoles, 10},
		},
	},
	{
		"Masers",                 // name
		470,                      // basePrice
		model.GroupTechnological, // type
		100,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityElectros, 10},
			{model.CommodityAntiMatter, 20},
			{model.CommodityRads, 10},
		},
	},
	{
		"Weapons",                // name
		420,                      // basePrice
		model.GroupTechnological, // type
		20,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityNickel, 50},
			{model.CommodityAlloys, 50},
		},
	},
	{
		"Vidi",                   // name
		380,                      // basePrice
		model.GroupTechnological, // type
		50,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityElectros, 50},
			{model.CommodityPowerPacks, 10},
		},
	},
	{
		"Electros",               // name
		250,                      // basePrice
		model.GroupTechnological, // type
		20,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_FabricationPlant,
			text.Factory_FabricationPlant_Loc,
			text.Factory_FabricationPlant_Closed,
			text.Factory_FabricationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityNickel, 30},
			{model.CommodityAlloys, 20},
		},
	},
	{
		"Tools",                  // name
		350,                      // basePrice
		model.GroupTechnological, // type
		50,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityRads, 10},
			{model.CommodityNickel, 30},
		},
	},
	{
		"Synths",                 // name
		560,                      // basePrice
		model.GroupTechnological, // type
		50,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityGAsChips, 50},
			{model.CommodityAntiMatter, 20},
		},
	},
	{
		"Droids",                 // name
		700,                      // basePrice
		model.GroupTechnological, // type
		50,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_FabricationPlant,
			text.Factory_FabricationPlant_Loc,
			text.Factory_FabricationPlant_Closed,
			text.Factory_FabricationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityBioChips, 50},
			{model.CommodityAntiMatter, 20},
		},
	},
	{
		"Anti-Matter",            // name
		480,                      // basePrice
		model.GroupTechnological, // type
		40,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ResearchFacility,
			text.Factory_ResearchFacility_Loc,
			text.Factory_ResearchFacility_Closed,
			text.Factory_ResearchFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityGAsChips, 30},
			{model.CommodityMonopoles, 20},
		},
	},
	{
		"PowerPacks",             // name
		680,                      // basePrice
		model.GroupTechnological, // type
		50,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityMasers, 30},
			{model.CommodityRads, 50},
		},
	},
	{
		"Controllers",            // name
		710,                      // basePrice
		model.GroupTechnological, // type
		50,                       // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityBioChips, 50},
			{model.CommodityElectros, 40},
		},
	},
	{
		"Generators",          // name
		670,                   // basePrice
		model.GroupIndustrial, // type
		70,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityElectros, 30},
			{model.CommodityTools, 20},
			{model.CommodityPolymers, 10},
			{model.CommodityLubOils, 10},
			{model.CommodityAlloys, 100},
		},
	},
	{
		"Polymers",            // name
		200,                   // basePrice
		model.GroupIndustrial, // type
		20,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ProcessingPlant,
			text.Factory_ProcessingPlant_Loc,
			text.Factory_ProcessingPlant_Closed,
			text.Factory_ProcessingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 30},
			{model.CommodityXmetals, 10},
		},
	},
	{
		"Lub-Oils",            // name
		560,                   // basePrice
		model.GroupIndustrial, // type
		10,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_Refinery,
			text.Factory_Refinery_Loc,
			text.Factory_Refinery_Closed,
			text.Factory_Refinery_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 80},
		},
	},
	{
		"Pharms",              // name
		630,                   // basePrice
		model.GroupIndustrial, // type
		20,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_Laboratory,
			text.Factory_Laboratory_Loc,
			text.Factory_Laboratory_Closed,
			text.Factory_Laboratory_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityRNA, 100},
			{model.CommodityXmetals, 10},
		},
	},
	{
		"Petros",              // name
		400,                   // basePrice
		model.GroupIndustrial, // type
		50,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_Refinery,
			text.Factory_Refinery_Loc,
			text.Factory_Refinery_Closed,
			text.Factory_Refinery_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPharms, 10},
			{model.CommodityPetros, 40},
		},
	},
	{
		"RNA",                 // name
		380,                   // basePrice
		model.GroupIndustrial, // type
		40,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPharms, 20},
			{model.CommodityXmetals, 20},
		},
	},
	{
		"Propellants",         // name
		260,                   // basePrice
		model.GroupIndustrial, // type
		10,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 20},
			{model.CommodityNitros, 40},
		},
	},
	{
		"Explosives",          // name
		470,                   // basePrice
		model.GroupIndustrial, // type
		60,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityElectros, 20},
			{model.CommodityNitros, 50},
			{model.CommodityNickel, 20},
		},
	},
	{
		"LanzariK",            // name
		700,                   // basePrice
		model.GroupIndustrial, // type
		50,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPolymers, 40},
			{model.CommodityRNA, 20},
			{model.CommodityXmetals, 50},
			{model.CommodityAlloys, 20},
		},
	},
	{
		"Nitros",              // name
		230,                   // basePrice
		model.GroupIndustrial, // type
		40,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ProcessingPlant,
			text.Factory_ProcessingPlant_Loc,
			text.Factory_ProcessingPlant_Closed,
			text.Factory_ProcessingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 30},
		},
	},
	{
		"Munitions",           // name
		250,                   // basePrice
		model.GroupIndustrial, // type
		20,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityExplosives, 30},
			{model.CommodityAlloys, 10},
		},
	},
	{
		"MechParts",           // name
		420,                   // basePrice
		model.GroupIndustrial, // type
		20,                    // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityTools, 10},
			{model.CommodityAlloys, 100},
			{model.CommodityGold, 10},
		},
	},
	{
		"Cereals",               // name
		270,                     // basePrice
		model.GroupAgricultural, // type
		60,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_FarmComplex,
			text.Factory_FarmComplex_Loc,
			text.Factory_FarmComplex_Closed,
			text.Factory_FarmComplex_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityMechParts, 10},
			{model.CommodityCereals, 10},
			{model.CommodityTextiles, 10},
		},
	},
	{
		"Woods",                 // name
		700,                     // basePrice
		model.GroupAgricultural, // type
		100,                     // labour
		[4]text.MsgNum{ // message
			text.Factory_Sawmill,
			text.Factory_Sawmill_Loc,
			text.Factory_Sawmill_Closed,
			text.Factory_Sawmill_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityLubOils, 10},
			{model.CommodityPetros, 20},
			{model.CommodityMechParts, 10},
			{model.CommodityWoods, 20},
			{model.CommodityTextiles, 10},
		},
	},
	{
		"Hides",                 // name
		400,                     // basePrice
		model.GroupAgricultural, // type
		60,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_ProcessingPlant,
			text.Factory_ProcessingPlant_Loc,
			text.Factory_ProcessingPlant_Closed,
			text.Factory_ProcessingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityMechParts, 10},
			{model.CommodityTextiles, 10},
			{model.CommodityMeat, 10},
			{model.CommodityStock, 10},
		},
	},
	{
		"Textiles",              // name
		300,                     // basePrice
		model.GroupAgricultural, // type
		40,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_Mill,
			text.Factory_Mill_Loc,
			text.Factory_Mill_Closed,
			text.Factory_Mill_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityWoods, 10},
			{model.CommodityTextiles, 10},
			{model.CommodityFurs, 10},
		},
	},
	{
		"Meat",                  // name
		410,                     // basePrice
		model.GroupAgricultural, // type
		40,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_RanchComplex,
			text.Factory_RanchComplex_Loc,
			text.Factory_RanchComplex_Closed,
			text.Factory_RanchComplex_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityTextiles, 10},
			{model.CommodityMeat, 10},
			{model.CommodityStock, 100},
		},
	},
	{
		"Spices",                // name
		650,                     // basePrice
		model.GroupAgricultural, // type
		80,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_FarmComplex,
			text.Factory_FarmComplex_Loc,
			text.Factory_FarmComplex_Closed,
			text.Factory_FarmComplex_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityMechParts, 10},
			{model.CommodityWoods, 10},
			{model.CommodityTextiles, 10},
			{model.CommoditySpices, 10},
			{model.CommodityFruit, 30},
		},
	},
	{
		"Fruit",                 // name
		420,                     // basePrice
		model.GroupAgricultural, // type
		90,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_FarmComplex,
			text.Factory_FarmComplex_Loc,
			text.Factory_FarmComplex_Closed,
			text.Factory_FarmComplex_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityLubOils, 10},
			{model.CommodityPetros, 10},
			{model.CommodityMechParts, 10},
			{model.CommodityTextiles, 10},
			{model.CommodityFruit, 10},
		},
	},
	{
		"Soya",                  // name
		230,                     // basePrice
		model.GroupAgricultural, // type
		40,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_HydroponicsLab,
			text.Factory_HydroponicsLab_Loc,
			text.Factory_HydroponicsLab_Closed,
			text.Factory_HydroponicsLab_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommoditySpices, 10},
			{model.CommoditySoya, 10},
		},
	},
	{
		"Stock",                 // name
		210,                     // basePrice
		model.GroupAgricultural, // type
		50,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_RanchComplex,
			text.Factory_RanchComplex_Loc,
			text.Factory_RanchComplex_Closed,
			text.Factory_RanchComplex_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityCereals, 20},
		},
	},
	{
		"Furs",                  // name
		650,                     // basePrice
		model.GroupAgricultural, // type
		60,                      // labour
		[4]text.MsgNum{ // message
			text.Factory_ProcessingPlant,
			text.Factory_ProcessingPlant_Loc,
			text.Factory_ProcessingPlant_Closed,
			text.Factory_ProcessingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPetros, 10},
			{model.CommodityMechParts, 10},
			{model.CommodityTextiles, 20},
			{model.CommodityStock, 10},
		},
	},
	{
		"Rads",            // name
		540,               // basePrice
		model.GroupMining, // type
		90,                // labour
		[4]text.MsgNum{ // message
			text.Factory_DeepMine,
			text.Factory_DeepMine_Loc,
			text.Factory_DeepMine_Closed,
			text.Factory_DeepMine_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityMasers, 10},
			{model.CommodityTools, 20},
			{model.CommodityPowerPacks, 20},
		},
	},
	{
		"Nickel",          // name
		370,               // basePrice
		model.GroupMining, // type
		80,                // labour
		[4]text.MsgNum{ // message
			text.Factory_OpenCastMine,
			text.Factory_OpenCastMine_Loc,
			text.Factory_OpenCastMine_Closed,
			text.Factory_OpenCastMine_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityLubOils, 10},
			{model.CommodityNitros, 10},
		},
	},
	{
		"Xmetals",         // name
		480,               // basePrice
		model.GroupMining, // type
		100,               // labour
		[4]text.MsgNum{ // message
			text.Factory_Laboratory,
			text.Factory_Laboratory_Loc,
			text.Factory_Laboratory_Closed,
			text.Factory_Laboratory_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityMasers, 20},
			{model.CommodityPolymers, 50},
		},
	},
	{
		"Crystals",        // name
		650,               // basePrice
		model.GroupMining, // type
		200,               // labour
		[4]text.MsgNum{ // message
			text.Factory_Laboratory,
			text.Factory_Laboratory_Loc,
			text.Factory_Laboratory_Closed,
			text.Factory_Laboratory_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPowerPacks, 10},
			{model.CommodityLubOils, 20},
		},
	},
	{
		"Alloys",          // name
		180,               // basePrice
		model.GroupMining, // type
		90,                // labour
		[4]text.MsgNum{ // message
			text.Factory_Furnace,
			text.Factory_Furnace_Loc,
			text.Factory_Furnace_Closed,
			text.Factory_Furnace_Open,
		},
		[]InputCommodity{}, // inputs
	},
	{
		"Gold",            // name
		600,               // basePrice
		model.GroupMining, // type
		200,               // labour
		[4]text.MsgNum{ // message
			text.Factory_DeepMine,
			text.Factory_DeepMine_Loc,
			text.Factory_DeepMine_Closed,
			text.Factory_DeepMine_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityExplosives, 30},
		},
	},
	{
		"Monopoles",       // name
		720,               // basePrice
		model.GroupMining, // type
		300,               // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPharms, 10},
			{model.CommodityNitros, 10},
		},
	},
	{
		"Hypnotapes",       // name
		430,                // basePrice
		model.GroupLeisure, // type
		40,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_DuplicationPlant,
			text.Factory_DuplicationPlant_Loc,
			text.Factory_DuplicationPlant_Closed,
			text.Factory_DuplicationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityCrystals, 30},
			{model.CommodityGold, 30},
		},
	},
	{
		"Studios",          // name
		390,                // basePrice
		model.GroupLeisure, // type
		70,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_ManufacturingPlant,
			text.Factory_ManufacturingPlant_Loc,
			text.Factory_ManufacturingPlant_Closed,
			text.Factory_ManufacturingPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityElectros, 40},
			{model.CommodityHides, 20},
		},
	},
	{
		"SensAmps",         // name
		520,                // basePrice
		model.GroupLeisure, // type
		80,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityCrystals, 10},
			{model.CommodityGold, 10},
		},
	},
	{
		"Games",            // name
		210,                // basePrice
		model.GroupLeisure, // type
		30,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_DuplicationPlant,
			text.Factory_DuplicationPlant_Loc,
			text.Factory_DuplicationPlant_Closed,
			text.Factory_DuplicationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityBioChips, 10},
			{model.CommodityHides, 10},
		},
	},
	{
		"Artifacts",        // name
		730,                // basePrice
		model.GroupLeisure, // type
		370,                // labour
		[4]text.MsgNum{ // message
			text.Factory_Excavation,
			text.Factory_Excavation_Loc,
			text.Factory_Excavation_Closed,
			text.Factory_Excavation_Open,
		},
		[]InputCommodity{},
	},
	{
		"Katydidics",       // name
		610,                // basePrice
		model.GroupLeisure, // type
		100,                // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityAntiMatter, 20},
			{model.CommodityPolymers, 40},
			{model.CommodityGold, 20},
		},
	},
	{
		"Musiks",           // name
		470,                // basePrice
		model.GroupLeisure, // type
		200,                // labour
		[4]text.MsgNum{ // message
			text.Factory_DuplicationPlant,
			text.Factory_DuplicationPlant_Loc,
			text.Factory_DuplicationPlant_Closed,
			text.Factory_DuplicationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityWoods, 10},
			{model.CommodityAlloys, 50},
		},
	},
	{
		"Libraries",        // name
		680,                // basePrice
		model.GroupLeisure, // type
		80,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_DuplicationPlant,
			text.Factory_DuplicationPlant_Loc,
			text.Factory_DuplicationPlant_Closed,
			text.Factory_DuplicationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityPolymers, 20},
			{model.CommodityWoods, 30},
			{model.CommodityTextiles, 40},
		},
	},
	{
		"Holos",            // name
		260,                // basePrice
		model.GroupLeisure, // type
		40,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_FabricationPlant,
			text.Factory_FabricationPlant_Loc,
			text.Factory_FabricationPlant_Closed,
			text.Factory_FabricationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityBioChips, 20},
		},
	},
	{
		"Univators",        // name
		570,                // basePrice
		model.GroupLeisure, // type
		90,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_ProductionFacility,
			text.Factory_ProductionFacility_Loc,
			text.Factory_ProductionFacility_Closed,
			text.Factory_ProductionFacility_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityMonopoles, 10},
		},
	},
	{
		"Sims",             // name
		310,                // basePrice
		model.GroupLeisure, // type
		50,                 // labour
		[4]text.MsgNum{ // message
			text.Factory_FabricationPlant,
			text.Factory_FabricationPlant_Loc,
			text.Factory_FabricationPlant_Closed,
			text.Factory_FabricationPlant_Open,
		},
		[]InputCommodity{ // inputs
			{model.CommodityAntiMatter, 20},
			{model.CommodityGold, 10},
		},
	},
}
