package input

type RestaurantCreation struct {
	Name     string  `json:"name"`
	Email    *string `json:"email,omitempty"`
	CookName *string `json:"cook_name,omitempty"`
}

type RestaurantJoin struct {
	Token    string  `json:"token"`
	CookName *string `json:"cook_name,omitempty"`
}
