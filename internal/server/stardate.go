package server

import (
	"fmt"
	"time"

	"github.com/nosborn/federation-1999/internal/text"
)

var (
	starday  int64
	startime int64
)

func init() {
	now := time.Now().Unix()
	seconds := now % 86400              // FIXME: gameStartTime & SECS_IN_A_DAY
	starday = 200501 + (now / 86400)    // FIXME: (gameStartTime / SECS_IN_A_DAY)
	startime = (seconds * 1000) / 86400 // FIXME: SECS_IN_A_DAY

	// time.AfterFunc(seconds*time.Second, startimeTimerProc)
}

func CurrentStardate() string {
	return fmt.Sprintf("%d:%03d", starday, startime)
}

func stardate() string {
	return text.Msg(text.STARDATE, CurrentStardate())
}

// func startimeTimerProc() {
// 	now := time.Now().Unix() // FIXME: Transaction theTransaction
//
// 	wasStarday := starday
//
// 	starday = 200501 + (now / 86400)          // FIXME: theTransaction.dayNumber()
// 	startime = ((now % 86400) * 1000) / 86400 // FIXME: ((theTransaction.time() % SECS_IN_A_DAY) * 1000) / SECS_IN_A_DAY
//
// 	if starday == wasStarday {
// 		return
// 	}
//
// 	log.Printf("Starting day %d", now/86400) // FIXME: theTransaction.dayNumber());
//
// 	// End of day processing.
//
// 	// time_t highTouristTime = 1;  // Excludes anything with zero minutes
// 	// vector<System*> t4Contenders;
// 	//
// 	// for (SystemList::const_iterator iter = systemList.begin();
// 	// iter != systemList.end();
// 	// iter++)
// 	// {
// 	// 	time_t touristTime = (*iter)->endOfDay();
// 	//
// 	// 	if (touristTime > highTouristTime) {
// 	// 		highTouristTime = touristTime;
// 	// 		t4Contenders.clear();
// 	// 	}
// 	//
// 	// 	if (touristTime == highTouristTime) {
// 	// 		t4Contenders.push_back(*iter);
// 	// 	}
// 	// }
//
// 	// Pick the T4 winner.
//
// 	// System* t4Winner = NULL;
// 	//
// 	// if (t4Contenders.size() > 0) {
// 	// 	t4Winner = t4Contenders[Random(1, t4Contenders.size()) - 1];
// 	// 	log.Printf("Awarding tourism trophy to %s", t4Winner.Name());
// 	//
// 	// 	const char* t4Announce = message(mnTOURISM_AWARD_ANNOUNCEMENT,
// 	// 	t4Winner->name());
// 	//
// 	// 	for (GPCI iter = g_playing.begin(); iter != g_playing.end(); iter++) {
// 	// 		if (iter->second->sCommsOff()) {
// 	// 			continue;
// 	// 		}
// 	//
// 	// 		if (iter->second->currentSystem()->isHidden()) {
// 	// 			continue;
// 	// 		}
// 	//
// 	// 		iter->second->Nsoutput(t4Announce);
// 	// 		iter->second->FlushOutput();
// 	// 	}
//
// 	// Start of day processing.
//
// 	// for (SystemList::iterator iter = systemList.begin();
// 	// iter != systemList.end();
// 	// iter++)
// 	// {
// 	// 	(*iter)->startOfDay(theTransaction.dayNumber(), (*iter == t4Winner));
// 	// }
//
// 	// debug.Trace("Stardate %i:%03i", starday, startime)
//
// 	time.AfterFunc(86400*time.Second, startimeTimerProc) // Tcl_CreateTimerHandler(SECS_IN_A_DAY, startimeTimerProc, NULL)
// }
