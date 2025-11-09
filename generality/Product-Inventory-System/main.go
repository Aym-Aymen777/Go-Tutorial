package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

func LoadProducts() ([]Product, error) {
	data, err := os.ReadFile("products.json")
	if err != nil || len(data) == 0 {
		if os.IsNotExist(err) {
			return []Product{}, nil
		}
		return nil, err
	}
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func AddProduct() ([]Product, error) {
	newProducts := []Product{}
	products, err := LoadProducts()
	if err != nil {
		fmt.Println("❌ Error loading products :",err)
	}
    for{
	var choice string
	fmt.Print("Do you want to add a new product? (y/n): ")
	fmt.Scan(&choice)
	if choice == "n" || choice == "N" || choice == "no" || choice == "No" || choice == "NO" {
		break
	}else if choice == "y" || choice == "Y" || choice == "yes" || choice == "Yes" || choice == "YES" {
	var newProduct Product
	fmt.Printf("Enter the product ID: ")
	fmt.Scan(&newProduct.ID)
	for i := 0; i < len(products); i++ {
		if products[i].ID == newProduct.ID {
			fmt.Println("Product ID already exists. Please enter a unique ID.")
			fmt.Scan(&newProduct.ID)
			i = -1
		}
	}
	fmt.Print("Enter the product name: ")
	fmt.Scan(&newProduct.Name)
	fmt.Print("Enter the product price: ")
	fmt.Scan(&newProduct.Price)
	fmt.Print("Enter the product quantity: ")
	fmt.Scan(&newProduct.Quantity)
	newProduct.Total = newProduct.Price * float64(newProduct.Quantity)

	newProducts = append(newProducts, newProduct)
	}else{
		fmt.Println("Invalid choice. Please enter 'y' or 'n'.")
	}
	}



	products = append(products, newProducts...)

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println("❌ Error marshalling products :",err)
	}

	err = os.WriteFile("products.json", data, 0644)
	if err != nil {
		fmt.Println("❌ Error saving products :",err)
	}

	fmt.Println("Product added successfully ✅")

	return newProducts, nil
}


func TotalInventoryValue([]Product) float64 {
	products,err :=LoadProducts()
	if err != nil {
		fmt.Println("❌ Error loading products :",err)
	}
	total := 0.0
	for i :=0 ;i<len(products);i++{
		total += products[i].Total
	}
	return total
}

func DisplayProducts(){
	products,err :=LoadProducts()
	if err != nil {
		fmt.Println("❌ Error loading products :",err)
	}
	fmt.Println("===== Product Inventory =====")
	for i :=0 ;i<len(products);i++{
		fmt.Println("ID : ",products[i].ID," | Name : ",products[i].Name," | Price : ",products[i].Price," | Quantity : ",products[i].Quantity," | Total : ",products[i].Total)	
	}
	fmt.Println("Total Inventory Value : ",TotalInventoryValue(products))
}

func SearchForProduct( name string){
	products,err :=LoadProducts()
	if err != nil {
		fmt.Println("❌ Error loading products :",err)
	}
	for i :=0 ;i<len(products);i++{
		if products[i].Name == name {
			fmt.Println("Product Found ✅ =====")
			fmt.Println("ID : ",products[i].ID," | Name : ",products[i].Name," | Price : ",products[i].Price," | Quantity : ",products[i].Quantity," | Total : ",products[i].Total)	
		    return
		}
	}
	fmt.Println("Product Not Found ❌")

}

func main() {
	fmt.Println("===== Product Inventory System =====")
	for {
		fmt.Println("1. Add Product ➕")
		fmt.Println("2. Display Products 📦")
		fmt.Println("3. Search For Product 🔎")
		fmt.Println("4. Exit 🚪")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			AddProduct()
		case 2:
			DisplayProducts()
		case 3:
			var name string
			fmt.Print("Enter the product name: ")
			fmt.Scan(&name)
			SearchForProduct(name)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
	//TODO: improve the search function by using map instead of slices because it's more faster ; and add the delete/update modules
	
}
