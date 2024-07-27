package entity

// CardData Contains bank card data as data content
type CardData struct {
	Name            string
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	Code            int
	Cardholder      string
	Metadata        Metadata
}
