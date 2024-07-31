package entity

// CardObjects Array of CardData objects
type CardObjects []CardData

// CardData Contains bank card data as data content
type CardData struct {
	Name            ObjectName
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	Code            int
	Cardholder      string
	Metadata        Metadata
}
