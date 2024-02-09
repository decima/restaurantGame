package main

import (
	"encoding/json"
	"fmt"
	. "restaurantAPI/models"
	"restaurantAPI/models/builders"
	"restaurantAPI/models/utils"
)

func main() {

	inputs := []string{
		"fried(chopped(potato),chopped(fish))",
		"Frozen(Tomato,Chopped(Pepper),Roasted(Chopped(Sausage),Cheese),Grilled(Sliced(Bread)),Egg)",
		"Dished(Sliced(Bread),Ham,Sliced(Bread))",
		"Fried(Chopped(Potato)),Fried(Chopped(Fish))",
	}

	for _, input := range inputs {
		fmt.Println("=============================")
		fmt.Println(input)
		test := monParser(input)

		j, _ := json.Marshal(test)
		fmt.Println(string(j))
		fmt.Println("=============================")
		//break

	}

}

func monParser(expr string) Ingredient {

	childrenResponses := [][]Ingredient{}
	level := 0
	responses := []Ingredient{}
	currentToken := []rune{}
	funcName := ""
	increaseArgPosition := false
	argPosition := 0
	//quantity := 1
	for _, char := range expr {
		switch char {
		/*
			// handle quantities if needed
			case '[':
				//do nothing, just to not process the [ as part of the string
			case ']':
				quantity, _ = strconv.Atoi(string(currentToken))
				currentToken = []rune{}
		*/
		case '(':
			level++
			if level > 1 {
				currentToken = append(currentToken, char)
				continue
			}
			funcName = string(currentToken)
			currentToken = []rune{}

		case ')':
			level--

			if level > 0 {
				currentToken = append(currentToken, char)
				continue
			}

			currentToken, childrenResponses = processCurrentToken(currentToken, childrenResponses, argPosition)
			responses = append(responses, builders.IngredientsBuild(childrenResponses[argPosition]).Transform(utils.TransformType(funcName)).Build())
		case ',':

			if level > 1 {
				currentToken = append(currentToken, char)
				continue
			}
			if level == 0 {
				increaseArgPosition = true
			}
			//level is 1, so we are in argument list.
			//processing previous argument
			currentToken, childrenResponses = processCurrentToken(currentToken, childrenResponses, argPosition)

			if increaseArgPosition {
				argPosition++
			}

		default:
			currentToken = append(currentToken, char)

		}

	}

	if len(currentToken) > 0 {
		responses = append(responses, builders.RawIngredientBuild(utils.IngredientType(currentToken)).Build())
	}
	switch len(responses) {
	case 0:
		return RawIngredient{Name: expr}
	case 1:
		return responses[0]
	default:
		return CombinedIngredients{Ingredients: responses}
	}
}

func processCurrentToken(currentToken []rune, childrenResponses [][]Ingredient, argPosition int) ([]rune, [][]Ingredient) {
	processedArgument := monParser(string(currentToken))
	currentToken = []rune{}
	if len(childrenResponses) < argPosition+1 {
		childrenResponses = append(childrenResponses, []Ingredient{})
	}
	childrenResponses[argPosition] = append(childrenResponses[argPosition], processedArgument)
	return currentToken, childrenResponses
}
