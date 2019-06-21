package service

import (
	"github.com/BrenQ/Mutant/models"
	"github.com/BrenQ/Mutant/utils"
)


type DnaService struct {
	*models.Dna
}

/**
@method Inicializa los valores del struct dna
*/
func (d *DnaService) Init() {
	d.N = 0
	d.Size = 0
	d.Pattern = 0
	d.Sequence = make([]rune, 0)
}


/**
@method Almacena los valores necesarios para operar
*/
func (d *DnaService) Register(data []string) bool {

	if len(data) == 0 {
		return false
	}
	d.N = len(data)
	for _, value := range data {
		if len(value) != d.N {
			return false
		}
		for _, letter := range value {

			if ! utils.ValidateLetter(letter) {
				return false
			}
			d.Sequence = append(d.Sequence, letter)
		}
	}
	d.Size = len(d.Sequence)
	return true
}

/**
@method  Verifica que un humano es mutando
*/
func (d *DnaService) IsMutant(data [] string) bool {

	const Sequences = 2 // Cantidad de secuencias requeridas para comprobar que es un mutante

	d.Init()

	registered := d.Register(data)
	// Recorro los elementos para verificar si se registra un patron
	if ! registered {
		return false
	}

	for index, letter := range d.Sequence {
		if d.Pattern > Sequences {
			d.Mutant = true
			return true
		}
		iterate(d, index, letter)
	}

	d.Mutant =  false
	return false
}

/**
@method Funcion auxiliar que recorre un array de funciones.
		Este array de funciones contiene la logica que define
		con que valores comparar segun la ubicacion
*/
func iterate(d* DnaService, index int, letter rune) {

	functions := []func(index int, n int) int{
		func(index int, n int) int { return index + 1 },
		func(index int, n int) int { return index + n },
		func(index int, n int) int { return index - n - 1 },
		func(index int, n int) int { return index - n + 1 },
		func(index int, n int) int { return index + n - 1 },
		func(index int, n int) int { return index + n + 1 },
	}

	for _, function := range functions {
		if check(d, function, index, letter, 0) {
			d.Pattern++
		}
	}
}
/**
@method Funcion para validar si se detecta algun patron
*/

func check(d* DnaService, function func(index int, n int) int, index int, letter rune, pattern int) bool {
	result := false
	newIndex := function(index, d.N)

	if pattern == 3 {
		result = true
	}

	if (newIndex > 0 && newIndex < d.Size) && (d.Sequence[newIndex] == letter && pattern < 4) {
		pattern++
		return check(d, function, newIndex, letter, pattern)
	}
	d.Iterations++
	return result
}