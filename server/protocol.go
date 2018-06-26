package server

import (
	"github.com/ethereum/go-ethereum/p2p/enr"
)

type ResponseStatus uint
type MessageType uint

const (
	REGISTER MessageType = iota
	REGISTER_RESPONSE
	UNREGISTER
	DISCOVER
	DISCOVER_RESPONSE

	OK                  ResponseStatus = 0
	E_INVALID_NAMESPACE ResponseStatus = 100
	E_INVALID_ENR       ResponseStatus = 101
	E_INVALID_TTL       ResponseStatus = 102
	E_INVALID_COOKIE    ResponseStatus = 103
	E_INVALID_CONTENT   ResponseStatus = 104
	E_NOT_AUTHORIZED    ResponseStatus = 200
	E_INTERNAL_ERROR    ResponseStatus = 300
)

type Register struct {
	Topic  string
	Record enr.Record
	TTL    uint
}

type RegisterResponse struct {
	Status  ResponseStatus
	Message string
}

type Discover struct {
	Limit uint
	Topic string
}

type DiscoverResponse struct {
	Status  ResponseStatus
	Message string
	Records []enr.Record
}

type Unregister struct {
	Topic  string
	Record enr.Record
}
