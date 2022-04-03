package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsCurrencySupported returns true if the currency is supported
func IsSupportedCurrenncy(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
