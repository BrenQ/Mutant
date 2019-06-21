package utils

/**
Funcion auxiliar para verificar si una letra es valida
*/

func ValidateLetter(letter rune) bool {
	letters := []rune{'A', 'T', 'C', 'G'}

	for _, value := range letters {
		if value == letter {
			return true
		}
	}
	return false
}
