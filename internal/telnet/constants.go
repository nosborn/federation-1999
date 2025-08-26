package telnet

const (
	IAC   = byte(255) // interpret as command:
	DONT  = byte(254) // you are not to use option
	DO    = byte(253) // please, you use option
	WONT  = byte(252) // I won't use option
	WILL  = byte(251) // I will use option
	SB    = byte(250) // interpret as subnegotiation
	GA    = byte(249) // you may reverse the line
	EL    = byte(248) // erase the current line
	EC    = byte(247) // erase the current character
	AYT   = byte(246) // are you there
	AO    = byte(245) // abort output--but let prog finish
	IP    = byte(244) // interrupt process--permanently
	BREAK = byte(243) // break
	DM    = byte(242) // data mark--for connect. cleaning
	NOP   = byte(241) // nop
	SE    = byte(240) // end sub negotiation
	EOR   = byte(239) // end of record (transparent mode)
	ABORT = byte(238) // Abort process
	SUSP  = byte(237) // Suspend process
	EOF   = byte(236) // End of file: xEOF in arpa/telnet.h
)

const (
	SYNCH = 242 // for telfunc calls
)

const ( // telnet options
	TELOPT_BINARY         = byte(0)   // 8-bit data path
	TELOPT_ECHO           = byte(1)   // echo
	TELOPT_RCP            = byte(2)   // prepare to reconnect
	TELOPT_SGA            = byte(3)   // suppress go ahead
	TELOPT_NAMS           = byte(4)   // approximate message size
	TELOPT_STATUS         = byte(5)   // give status
	TELOPT_TM             = byte(6)   // timing mark
	TELOPT_RCTE           = byte(7)   // remote controlled transmission and echo
	TELOPT_NAOL           = byte(8)   // negotiate about output line width
	TELOPT_NAOP           = byte(9)   // negotiate about output page size
	TELOPT_NAOCRD         = byte(10)  // negotiate about CR disposition
	TELOPT_NAOHTS         = byte(11)  // negotiate about horizontal tabstops
	TELOPT_NAOHTD         = byte(12)  // negotiate about horizontal tab disposition
	TELOPT_NAOFFD         = byte(13)  // negotiate about formfeed disposition
	TELOPT_NAOVTS         = byte(14)  // negotiate about vertical tab stops
	TELOPT_NAOVTD         = byte(15)  // negotiate about vertical tab disposition
	TELOPT_NAOLFD         = byte(16)  // negotiate about output LF disposition
	TELOPT_XASCII         = byte(17)  // extended ascic character set
	TELOPT_LOGOUT         = byte(18)  // force logout
	TELOPT_BM             = byte(19)  // byte macro
	TELOPT_DET            = byte(20)  // data entry terminal
	TELOPT_SUPDUP         = byte(21)  // supdup protocol
	TELOPT_SUPDUPOUTPUT   = byte(22)  // supdup output
	TELOPT_SNDLOC         = byte(23)  // send location
	TELOPT_TTYPE          = byte(24)  // terminal type
	TELOPT_EOR            = byte(25)  // end or record
	TELOPT_TUID           = byte(26)  // TACACS user identification
	TELOPT_OUTMRK         = byte(27)  // output marking
	TELOPT_TTYLOC         = byte(28)  // terminal location number
	TELOPT_3270REGIME     = byte(29)  // 3270 regime
	TELOPT_X3PAD          = byte(30)  // X.3 PAD
	TELOPT_NAWS           = byte(31)  // window size
	TELOPT_TSPEED         = byte(32)  // terminal speed
	TELOPT_LFLOW          = byte(33)  // remote flow control
	TELOPT_LINEMODE       = byte(34)  // Linemode option
	TELOPT_XDISPLOC       = byte(35)  // X Display Location
	TELOPT_OLD_ENVIRON    = byte(36)  // Old - Environment variables
	TELOPT_AUTHENTICATION = byte(37)  // Authenticate
	TELOPT_ENCRYPT        = byte(38)  // Encryption option
	TELOPT_NEW_ENVIRON    = byte(39)  // New - Environment variables
	TELOPT_TN3270E        = byte(40)  // RFC2355 - TN3270 Enhancements
	TELOPT_CHARSET        = byte(42)  // RFC2066 - Charset
	TELOPT_COMPORT        = byte(44)  // RFC2217 - Com Port Control
	TELOPT_KERMIT         = byte(47)  // RFC2840 - Kermit
	TELOPT_EXOPL          = byte(255) // extended-options-list
)

const ( // sub-option qualifiers
	TELQUAL_IS    = byte(0) // option is...
	TELQUAL_SEND  = byte(1) // send option
	TELQUAL_INFO  = byte(2) // ENVIRON: informational version of IS
	TELQUAL_REPLY = byte(2) // AUTHENTICATION: client version of IS
	TELQUAL_NAME  = byte(3) // AUTHENTICATION: client version of IS
)

const (
	LFLOW_OFF         = byte(0) // disable remote flow control
	LFLOW_ON          = byte(1) // enable remote flow control
	LFLOW_RESTART_ANY = byte(2) // restart output on any char
	LFLOW_RESTART_XON = byte(3) // restart output only on XON
)

const ( // LINEMODE suboptions
	LM_MODE        = byte(1)
	LM_FORWARDMASK = byte(2)
	LM_SLC         = byte(3)
)

const (
	SLC_SYNCH = 1
	SLC_BRK   = 2
	SLC_IP    = 3
	SLC_AO    = 4
	SLC_AYT   = 5
	SLC_EOR   = 6
	SLC_ABORT = 7
	SLC_EOF   = 8
	SLC_SUSP  = 9
	SLC_EC    = 10
	SLC_EL    = 11
	SLC_EW    = 12
	SLC_RP    = 13
	SLC_LNEXT = 14
	SLC_XON   = 15
	SLC_XOFF  = 16
	SLC_FORW1 = 17
	SLC_FORW2 = 18
	SLC_MCL   = 19
	SLC_MCR   = 20
	SLC_MCWL  = 21
	SLC_MCWR  = 22
	SLC_MCBOL = 23
	SLC_MCEOL = 24
	SLC_INSRT = 25
	SLC_OVER  = 26
	SLC_ECR   = 27
	SLC_EWR   = 28
	SLC_EBOL  = 29
	SLC_EEOL  = 30
)
