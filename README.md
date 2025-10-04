# Federation: 1999

This is a recreation[^1] of the Federation space fantasy game as it was in 1999.

Federation is the incarnation of the game that came after [Federation II][wiki] and before [Federation 2][wiki]. Obviously.

[wiki]: https://en.wikipedia.org/wiki/Federation_II

[^1]: Mostly.

## Component Parts

### modemd

This didn't exist in 1999, it's a modern addition that limits connections to the game to 1999-era dial-up modem speeds. There's no for this but it amuses me. modemd is the first point of contact when connecting to the game and it launches a telnetd for each connection.

### telnetd

Originally a 4.4BSD telnetd, modified to launch the login program without root privileges. No longer present in the source tree because modern Linux lets us achieve the same thing with a standard telnetd.

### login

Originally a 4.4BSD login, modified to authenticate against the IBGames account database, and now reimplemented in Go. On successful authentication it launches perivale.

### perivale

Also know as "the driver", this relays player input to the game and game output to the player, fiddling with text along the way, and performing rate-limiting and flow-control on player input. This is the modernised but otherwise original C version, although there is a Go version in the repository. Single-threaded with precise control of network I/O doesn't really play to Go's strengths, and the only way to get it to work properly in Go was to write an exect replica of the C code.

The name is an homage to the original Compunet Federation which was housed in an industrial estate in idyllic Perivale, in West London, giving rise to the term "Perivale Relay Station".

### fedtpd

The game engine itself. This was a single-threaded program running on a single CPU in 1999  and that pattern is followed in the Go version by largely eschewing Go's concurrency features. Most of the interesting high-performance features that let AOL Federation survive 1000+ concurrent players had been removed by 1999 as they were no longer required, but the Go version does retain the custom persona "database" and can read 1999-era persona files.[^2]

[^2]: This could be construed as implying the continued existence of one or more 1999-era persona file(s).

### workbench

The editor for player planets and a run-time checker for player planets to sanity check them before the game engine tried to load them. This is the original C code; an incomplete Go version is also present in the repository. The Go version currently implements only the checker part, and can read 1999-era workbench files.[^3]

[^3]: This could be construed as implying the continued existence of one or more 1999-era workbench file(s) (and therefore player planet(s)).
