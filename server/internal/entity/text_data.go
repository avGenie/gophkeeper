package entity

// TextDataObjects Array of TextData objects
type TextDataObjects []TextData

// TextData Contains text data as data content
type TextData struct {
	Name     string
	Text     string
	Metadata Metadata
}
