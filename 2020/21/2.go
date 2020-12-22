package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type StringSet map[string]struct{}

func (ss StringSet) add(s string) {
	ss[s] = struct{}{}
}
func (ss StringSet) remove(s string) {
	delete(ss, s)
}
func (ss StringSet) contains(s string) bool {
	_, ok := ss[s]
	return ok
}
func (ss StringSet) intersect(o StringSet) {
	keys := make([]string, 0)
	for key := range ss {
		keys = append(keys, key)
	}
	for _, a := range keys {
		_, exists := o[a]
		if !exists {
			ss.remove(a)
		}
	}
}
func (ss StringSet) union(o StringSet) {
	for a := range o {
		ss.add(a)
	}
}
func (ss StringSet) copy() StringSet {
	c := make(StringSet)
	for k, v := range ss {
		c[k] = v
	}
	return c
}
func (ss StringSet) toArray() []string {
	a := make([]string, 0)
	for s := range ss {
		a = append(a, s)
	}
	return a
}

type Recipe struct {
	ingredients StringSet
	allergens   StringSet
}

func (r Recipe) copy() Recipe {
	ingredients := make(StringSet)
	for k, v := range r.ingredients {
		ingredients[k] = v
	}
	allergens := make(StringSet)
	for k, v := range r.allergens {
		allergens[k] = v
	}
	return Recipe{ingredients, allergens}
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/21/input.txt")
	r := regexp.MustCompile(`^(.+) \(contains (.+)\)$`)
	recipes := make([]Recipe, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		matches := r.FindStringSubmatch(line)
		ingredients := make(StringSet)
		for _, ingredient := range strings.Fields(matches[1]) {
			ingredients.add(ingredient)
		}

		allergens := make(StringSet)
		for _, allergen := range strings.Split(matches[2], ", ") {
			allergens.add(allergen)
		}
		recipes = append(recipes, Recipe{ingredients, allergens})
	}

	allergenCandidates := make(map[string]StringSet)
	for _, recipe := range recipes {
		for allergen := range recipe.allergens {
			_, exists := allergenCandidates[allergen]
			if !exists {
				allergenCandidates[allergen] = recipe.ingredients.copy()
			}
			allergenCandidates[allergen].intersect(recipe.ingredients)
		}
	}

	allergenIngredients := make(StringSet)
	for _, v := range allergenCandidates {
		allergenIngredients.union(v)
	}

	count := 0
	for _, recipe := range recipes {
		for ingredient := range recipe.ingredients {
			if !allergenIngredients.contains(ingredient) {
				count++
			}
		}
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", count)

	allAllergens := make(StringSet)
	for _, recipe := range recipes {
		for allergen := range recipe.allergens {
			allAllergens.add(allergen)
		}
	}

	finalMap := make(map[string]string)
	numAllergens := len(allAllergens)
	for len(finalMap) != numAllergens {
		for allergen := range allAllergens {
			ingredients := allergenCandidates[allergen]
			if len(ingredients) == 1 {
				matchingIngredient := ingredients.toArray()[0]
				finalMap[allergen] = matchingIngredient
				for a, candidates := range allergenCandidates {
					if a == allergen {
						continue
					}
					candidates.remove(matchingIngredient)
				}
				break
			}
		}
	}

	keys := make([]string, 0)
	for k := range finalMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	sortedIngredients := make([]string, len(keys))
	for i, key := range keys {
		sortedIngredients[i] = finalMap[key]
	}

	fmt.Printf("Part 2: %v\n", strings.Join(sortedIngredients, ","))
}
