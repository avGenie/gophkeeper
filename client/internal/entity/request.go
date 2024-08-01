package entity

type DataRequestType int

const (
	DataRequestLoginPassword DataRequestType = iota
	DataRequestText
	DataRequestCard
	DataRequestInvalid
)
