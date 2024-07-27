package entity

type DataName string

type DataRequestType int

const (
	DataRequestLoginPassword DataRequestType = iota
	DataRequestText
	DataRequestCard
	DataRequestInvalid
)

type DataRequest struct {
	Type DataRequestType
	Name DataName
}
