package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	dogFunction, err := Animal(dog)
	catFunction, err := Animal(cat)
	hamsterFunction, err := Animal(hamster)
	tarantulaFunction, err := Animal(tarantula)

	if err != nil {
		fmt.Println("Animal inserido é inválido")
		return
	}

	total := 0.0
	total += dogFunction(25)
	total += catFunction(18)
	total += hamsterFunction(3)
	total += tarantulaFunction(18)

	fmt.Printf("No total, precisaremos de %.2f KG de\nalimentos para todos os animai", total/1000)
}

var animalsMapFoodWight = map[string]float64{
	dog:       10000.00,
	cat:       5000.00,
	hamster:   250.00,
	tarantula: 150.00,
}

func Animal(animal string) (func(animalQuantity float64) float64, error) {
	var animalsMapFunc = map[string]func(float64) float64{
		dog:       dogFunc,
		cat:       catFunc,
		hamster:   hamsterFunc,
		tarantula: tarangulaFunc,
	}

	animalFunction, ok := animalsMapFunc[animal]

	if !ok {
		return nil, fmt.Errorf("Não foi encontrado um animal do tipo %s", animal)
	}

	return animalFunction, nil
}

func dogFunc(animalQuantity float64) float64 {
	return animalQuantity * animalsMapFoodWight[dog]
}
func catFunc(animalQuantity float64) float64 {
	return animalQuantity * animalsMapFoodWight[cat]
}
func hamsterFunc(animalQuantity float64) float64 {
	return animalQuantity * animalsMapFoodWight[hamster]
}
func tarangulaFunc(animalQuantity float64) float64 {
	return animalQuantity * animalsMapFoodWight[tarantula]
}
