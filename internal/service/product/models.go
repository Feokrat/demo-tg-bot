package product

var allProducts = []Product{
	{Title: "Milk"},
	{Title: "Water"},
	{Title: "Eggs"},
	{Title: "Meat"},
}

type Product struct {
	Title string
}
