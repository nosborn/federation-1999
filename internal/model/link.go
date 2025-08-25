package model

// const (
// 	LINK_ADDRESS_SIZE   =  16
// 	LINK_HOSTNAME_SIZE  =  64
// )

// const (
// 	LINK_FORMAT = "%u %s %s %d\n"
// 	LINK_TOKENS = 4
// 	LINK_SOCKET = "%s/.fedtpd.socket" /* add homeDir() */
// )

// type ConnectInfo struct {
//    addr [LINK_ADDRESS_SIZE]byte,
//    name [LINK_HOSTNAME_SIZE]byte,
// }

const (
	DLE byte = 16
)

const (
	LeAck byte = 1 + iota
	LeSpy
	LeTrace
)
