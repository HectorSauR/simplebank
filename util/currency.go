package util

// Constants for all supported currencies.
const (
	USD = "USD"
	EUR = "EUR"
	MXN = "MXN"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, MXN:
		return true
	default:
		return false
	}
}
