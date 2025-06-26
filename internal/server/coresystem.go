package server

import (
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
)

type CoreSystem struct {
	System
}

func (s *CoreSystem) DrinkEvent(player *Player, object *Object) bool {
	debug.Precondition(player != nil)
	debug.Precondition(object != nil)

	if object.Number() == sol.ObWHOOSH {
		player.DrinkWHOOSH(object)
		return true
	}
	return false
}

func (s *CoreSystem) EndOfDay() int32 {
	return 0
}

func (s *CoreSystem) Expenditure(_ int32) {
	// Do nothing. We don't care about expenditure in core systems.
}

// Planet*
// CoreSystem::guessPlanet( unsigned locationNo ) const
// {
//    const Location* theLocation = findLocation(locationNo);
//
//    if (theLocation == 0 || theLocation->isSpace()) {
//       return 0;
//    }
//
//    for (PlanetData::const_iterator iter = m_planets.begin();
//         iter != m_planets.end();
//         iter++)
//    {
//       if (locationNo == (*iter)->exchange ||
//           locationNo == (*iter)->hospital ||
//           locationNo == (*iter)->landing)
//       {
//          return *iter;
//       }
//    }
//
//    if (m_planets.size() == 1) {
//       return m_planets[0];
//    }
//
//    return 0;
// }

// unsigned
// CoreSystem::hospital() const
// {
//    dbgTrace("CoreSystem::hospital");
//
//    for (PlanetData::const_iterator iter = m_planets.begin();
//         iter != m_planets.end();
//         iter++)
//    {
//       dbgTrace("%s %u", (*iter)->name(), (*iter)->hospital);
//
//       if ((*iter)->hospital > 0) {
//          const Location* theLocation = findLocation((*iter)->hospital);
//
//          if (theLocation != 0) {
//             dbgTrace("Returning %u", (*iter)->hospital);
//             return (*iter)->hospital;
//          }
//       }
//    }
//
//    dbgTrace("Not found");
//    return 0;
// }

func (s *CoreSystem) Income(_ int32, _ bool) {
	// Do nothing. We don't care about income in core systems.
}

// bool
// CoreSystem::isCapital() const
// {
//    for (PlanetData::const_iterator iter = m_planets.begin();
//         iter != m_planets.end();
//         iter++)
//    {
//       if ((*iter)->m_level == levelCapital) {
//          return true;
//       }
//    }
//
//    return false;
// }

// unsigned
// CoreSystem::landingLocationNo( unsigned orbitLocationNo ) const
// {
//    dbgTrace("CoreSystem::landingLocationNo(%u)", orbitLocationNo);
//
//    for (PlanetData::const_iterator iter = m_planets.begin();
//         iter != m_planets.end();
//         iter++)
//    {
//       dbgTrace("%s=%u", (*iter)->name(), (*iter)->orbit);
//
//       if ((*iter)->orbit == orbitLocationNo) {
//          dbgTrace("Returning %u", (*iter)->landing);
//          return (*iter)->landing;
//       }
//    }
//
//    dbgTrace("Not found");
//    return 0;
// }

func (s *CoreSystem) LoaderQueuePosition() int {
	return 0
}

// unsigned
// CoreSystem::orbitLocationNo( unsigned landingLocationNo ) const
// {
//    for (PlanetData::const_iterator iter = m_planets.begin();
//         iter != m_planets.end();
//         iter++)
//    {
//       if (landingLocationNo == (*iter)->landing) {
//          return (*iter)->orbit;
//       }
//    }
//
//    return 0;
// }

// Planet*
// CoreSystem::planet() const
// {
//    if (m_planets.size() != 1) {
//       IB_log("PANIC: planet() called for %s", m_name.c_str());
//       abort();
//    }
//
//    dbgCheck(m_planets.size() == 1);
//    return m_planets[0];
// }

func (s *CoreSystem) Save(_ database.SaveWhen) {
	// Do nothing.
}

func (s *CoreSystem) StartOfDay(_ int32, _ bool) {
	// Do nothing.
}
