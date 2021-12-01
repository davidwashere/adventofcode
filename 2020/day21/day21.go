package day21

import (
	"aoc2020/util"
	"fmt"
	"sort"
	"strings"
)

type foods struct {
	// mapIngredientsToAllergens map[string][]string
	// mapAllergensToIngredients map[string][]string
	IngredientCounts map[string]int
	All              []food
	AllIngredients   []string
}

type food struct {
	ingredients []string
	allergens   []string
}

func parsefile(inputfile string) foods {
	data, _ := util.ReadFileToStringSlice(inputfile)

	foods := foods{}
	foods.IngredientCounts = map[string]int{}

	for _, line := range data {
		lineS := strings.Split(line, " (contains ")
		ingredients := util.ParseStrs(lineS[0])
		allergens := util.ParseStrs(lineS[1])
		food := food{
			ingredients: ingredients,
			allergens:   allergens,
		}

		for _, ingredient := range ingredients {
			if val, ok := foods.IngredientCounts[ingredient]; ok {
				foods.IngredientCounts[ingredient] = val + 1
			} else {
				foods.IngredientCounts[ingredient] = 1
			}

		}

		foods.All = append(foods.All, food)
	}

	return foods
}

func part1(inputfile string) int {
	foods := parsefile(inputfile)

	allToIngs := map[string][]string{}
	for _, food := range foods.All {
		for _, allergen := range food.allergens {
			if _, ok := allToIngs[allergen]; !ok {
				allToIngs[allergen] = append([]string{}, food.ingredients...)
			} else {
				repeats := []string{}
				for _, ingredient := range food.ingredients {
					if util.IsStringIn(allToIngs[allergen], ingredient) {
						repeats = append(repeats, ingredient)
					}
				}

				// Remove anything that isn't a repeat
				allToIngs[allergen] = repeats
				if len(repeats) == 1 {
					// remove solo ingredent from every other allergen
					ing := repeats[0]
					for a := range allToIngs {
						if a == allergen {
							// skip the alergen the must have it
							continue
						}

						allToIngs[a] = util.RemoveStringFromSlice(allToIngs[a], ing)
					}
				}
			}
		}
	}

	allSolo := false
	for !allSolo {
		allSolo = true
		for a, ings := range allToIngs {
			if len(ings) > 1 {
				allSolo = false
				continue
			}

			// solo
			ing := ings[0]
			for k := range allToIngs {
				if k == a {
					continue
				}
				allToIngs[k] = util.RemoveStringFromSlice(allToIngs[k], ing)
			}
		}
	}

	ingsWithAllergy := []string{}
	for k, v := range allToIngs {
		fmt.Println(k, v)
		ingsWithAllergy = append(ingsWithAllergy, v[0])
	}

	result := 0
	for ing, count := range foods.IngredientCounts {
		if !util.IsStringIn(ingsWithAllergy, ing) {
			result += count
		}
	}

	return result
}

func part2(inputfile string) string {
	foods := parsefile(inputfile)

	allToIngs := map[string][]string{}
	for _, food := range foods.All {
		for _, allergen := range food.allergens {
			if _, ok := allToIngs[allergen]; !ok {
				allToIngs[allergen] = append([]string{}, food.ingredients...)
			} else {
				repeats := []string{}
				for _, ingredient := range food.ingredients {
					if util.IsStringIn(allToIngs[allergen], ingredient) {
						repeats = append(repeats, ingredient)
					}
				}

				// Remove anything that isn't a repeat
				allToIngs[allergen] = repeats
				if len(repeats) == 1 {
					// remove solo ingredent from every other allergen
					ing := repeats[0]
					for a := range allToIngs {
						if a == allergen {
							// skip the alergen the must have it
							continue
						}

						allToIngs[a] = util.RemoveStringFromSlice(allToIngs[a], ing)
					}
				}
			}
		}
	}

	allSolo := false
	for !allSolo {
		allSolo = true
		for a, ings := range allToIngs {
			if len(ings) > 1 {
				allSolo = false
				continue
			}

			// solo
			ing := ings[0]
			for k := range allToIngs {
				if k == a {
					continue
				}
				allToIngs[k] = util.RemoveStringFromSlice(allToIngs[k], ing)
			}
		}
	}

	allergens := []string{}
	ingsWithAllergy := []string{}
	for k, v := range allToIngs {
		allergens = append(allergens, k)
		fmt.Println(k, v)
		ingsWithAllergy = append(ingsWithAllergy, v[0])
	}

	sort.Strings(allergens)
	result := ""
	for i, allergen := range allergens {
		result += allToIngs[allergen][0]
		if i < len(allergens)-1 {
			result += ","
		}
	}

	return result
}
