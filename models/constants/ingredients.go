package constants

type Ingredient struct {
	Name  IngredientType `json:"name"`
	Price int            `json:"price"`
}

var availableIngredients = map[IngredientType]Ingredient{
	Potato: {
		Name:  Potato,
		Price: 1,
	},
	Fish: {
		Name:  Fish,
		Price: 20,
	},
	Chicken: {
		Name:  Chicken,
		Price: 15,
	},
	Beef: {
		Name:  Beef,
		Price: 25,
	},
	Tomato: {
		Name:  Tomato,
		Price: 3,
	},
	Onion: {
		Name:  Onion,
		Price: 1,
	},
	Bread: {
		Name:  Bread,
		Price: 10,
	},
	Egg: {
		Name:  Egg,
		Price: 5,
	},
}

func AvailableIngredients() map[IngredientType]Ingredient {
	return availableIngredients
}
