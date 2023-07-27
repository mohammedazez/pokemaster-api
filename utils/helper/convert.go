package helper

import (
	"strconv"
	"strings"
)

func StringToInt64(sValue string, defValue int64) int64 {
	value := defValue
	iValue, err := strconv.ParseInt(sValue, 10, 64)
	if err == nil {
		value = iValue
	}
	return value
}

func StringToInt(sValue string, defValue int) int {
	value := StringToInt64(sValue, int64(defValue))
	return int(value)
}

func StringToBool(sValue string, defValue bool) bool {
	value := defValue
	if sValue == "1" || strings.ToLower(sValue) == "true" {
		value = true
	} else if sValue == "0" || strings.ToLower(sValue) == "false" {
		value = false
	}
	return value
}

// func (s *Service) Update(form *domain.Pokemon) (domain.Pokemon, error) {
// 	ctx := context.Background()

// 	getPokemon, err := s.repo.GetPokemon(ctx, form.ID)
// 	if err != nil {
// 		log.Printf("Failed to get pokemon: %v", err)
// 		return getPokemon, err
// 	}

// 	fibNum := getNumber(getPokemon.PokemonName)
// 	form.PokemonName = form.PokemonName + "-" + fibNum

// 	update, err := s.repo.UpdatePokemon(ctx, form)
// 	if err != nil {
// 		log.Printf("Failed to update pokemon: %v", err)
// 		return update, err
// 	}

// 	return domain.Pokemon{}, nil
// }

// func getNumber(name string) string {
// 	var n string
// 	var results string

// 	// mendapatakan index setelah "-"
// 	index := strings.LastIndex(name, "-")
// 	fmt.Println(index)

// 	// jika index -1 yang berarti belum ada nomornya maka masukkan "0"
// 	if index == -1 {
// 		results = "0"

// 		// else jika sudah ada nomornya
// 	} else {

// 		// untuk get last value setelah "-"  by example 610, 987, 1597
// 		for i := index + 1; i < len(name); i++ {
// 			n += string(name[i])
// 		}

// 		if n == "0" {
// 			results = "1"
// 		}

// 		// convert value ke int
// 		num, _ := strconv.Atoi(n)
// 		fmt.Println("num", num)

// 		// untuk mendapatkan semua nomor fibonacci
// 		f := fibonacci()
// 		var fibNumbers []int
// 		var fibNumber int
// 		for i := 0; i < num+2; i++ {
// 			fibNumber = f()
// 			fibNumbers = append(fibNumbers, fibNumber)
// 		}

// 		fmt.Println("fibNumbers", fibNumbers)

// 		// assign data fibonacci ke filteredSlice
// 		filteredSlice := make([]int, 0)
// 		for _, value := range fibNumbers {
// 			// saring value kurang dari num dan lebih dari 0
// 			if value <= num && value >= 0 {
// 				filteredSlice = append(filteredSlice, value)
// 			}
// 		}
// 		fmt.Println("filteredSlice", filteredSlice)

// 		// lalu jumlahkan data filteredSlice dari index terakhir dan index sebelum index terakhir
// 		if len(filteredSlice) >= 2 {
// 			results = strconv.Itoa(filteredSlice[len(filteredSlice)-1] + filteredSlice[len(filteredSlice)-2])
// 		}
// 	}

// 	return results
// }

// func fibonacci() func() int {
// 	first, second := 0, 1
// 	return func() int {
// 		ret := first
// 		first, second = second, first+second
// 		return ret
// 	}
// }
