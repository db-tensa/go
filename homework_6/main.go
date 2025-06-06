package main

import (
	"fmt"
	"strings"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    string
	Stock       int
}

type Customer struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}

type CartItem struct {
	Product  Product
	Quantity int
}

type Cart struct {
	CustomerID int
	Items      []CartItem
	Discount   float64
}

type Order struct {
	ID         int
	CustomerID int
	Items      []CartItem
	Total      float64
	Status     string
	CreatedAt  time.Time
}

type Store struct {
	Products  []Product
	Customers []Customer
	Carts     []Cart
	Orders    []Order
}

func (s *Store) AddProduct(name, description, category string, price float64, stock int) {
	s.Products = append(s.Products, Product{
		ID:          len(s.Products) + 1,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Stock:       stock,
	})
	fmt.Printf("Product %s added\n", name)
}

func (s *Store) ListProducts() {
	fmt.Println("--- All products ---")
	for _, p := range s.Products {
		fmt.Printf("ID: %d | %s | Category: %s | Price: %.2f | Stock: %d\n",
			p.ID, p.Name, p.Category, p.Price, p.Stock)
	}
}

func (s *Store) AddCustomer(name, email, phone, address string) {
	s.Customers = append(s.Customers, Customer{
		ID:      len(s.Customers) + 1,
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
	})
	fmt.Printf("Client %s registered\n", name)
}

func (s *Store) AddToCart(customerID, productID, quantity int) {
	var cart *Cart
	for i, c := range s.Carts {
		if c.CustomerID == customerID {
			cart = &s.Carts[i]
			break
		}
	}
	if cart == nil {
		s.Carts = append(s.Carts, Cart{CustomerID: customerID})
		cart = &s.Carts[len(s.Carts)-1]
	}

	for _, p := range s.Products {
		if p.ID == productID {
			if p.Stock >= quantity {
				cart.Items = append(cart.Items, CartItem{Product: p, Quantity: quantity})
				fmt.Printf("Product %s (%d) added to cart\n", p.Name, quantity)
			} else {
				fmt.Println("Not enough stock")
			}
			return
		}
	}
	fmt.Println("Product not found")
}

func (s *Store) ViewCart(customerID int) {
	for _, c := range s.Carts {
		if c.CustomerID == customerID {
			fmt.Println("--- Your cart ---")
			for _, customer := range s.Customers {
				if customer.ID == customerID {
					fmt.Printf("Client: %s\n", customer.Name)
					break
				}
			}
			total := 0.0
			for i, item := range c.Items {
				itemTotal := item.Product.Price * float64(item.Quantity)
				fmt.Printf("%d. %s x%d - %.2f\n", i+1, item.Product.Name, item.Quantity, itemTotal)
				total += itemTotal
			}
			if c.Discount > 0 {
				fmt.Printf("Discount: %.0f%%\n", c.Discount*100)
				total *= (1 - c.Discount)
			}
			fmt.Printf("Total: %.2f\n", total)
			return
		}
	}
	fmt.Println("Cart empty or client not found")
}

func main() {
	store := Store{}
	var choice int

	for {
		fmt.Println("\n=== TechStore Online ===")
		fmt.Println("Main menu:")
		fmt.Println("1. Manage products")
		fmt.Println("2. Manage clients")
		fmt.Println("3. Shopping cart")
		fmt.Println("4. Orders")
		fmt.Println("5. Store stats")
		fmt.Println("6. Exit")
		fmt.Print("> ")
		fmt.Scan(&choice)

		if choice == 6 {
			break
		}

		switch choice {
		case 1:
			for {
				fmt.Println("\n--- Product menu ---")
				fmt.Println("1. Add product")
				fmt.Println("2. List all products")
				fmt.Println("3. Find product by ID")
				fmt.Println("4. Search by category")
				fmt.Println("5. Update product")
				fmt.Println("6. Back to main menu")
				fmt.Print("> ")
				fmt.Scan(&choice)

				if choice == 6 {
					break
				}

				switch choice {
				case 1:
					var name, description, category string
					var price float64
					var stock int
					fmt.Println("\n--- Add product ---")
					fmt.Print("Enter product name: ")
					fmt.Scan(&name)
					fmt.Print("Enter description: ")
					fmt.Scan(&description)
					fmt.Print("Enter price: ")
					fmt.Scan(&price)
					fmt.Print("Enter category: ")
					fmt.Scan(&category)
					fmt.Print("Enter stock quantity: ")
					fmt.Scan(&stock)
					store.AddProduct(name, description, category, price, stock)
				case 2:
					store.ListProducts()
				case 3:
					// Find product by ID (in main, no separate function)
					var id int
					fmt.Print("Enter product ID: ")
					fmt.Scan(&id)
					for _, p := range store.Products {
						if p.ID == id {
							fmt.Printf("Found: %s | Price: %.2f | Stock: %d\n",
								p.Name, p.Price, p.Stock)
							break
						}
					}
				case 4:
					var category string
					fmt.Print("Enter category: ")
					fmt.Scan(&category)
					fmt.Println("--- Products in", category, "---")
					for _, p := range store.Products {
						if strings.EqualFold(p.Category, category) {
							fmt.Printf("ID: %d | %s | Price: %.2f\n", p.ID, p.Name, p.Price)
						}
					}
				case 5:
					fmt.Println("Product update not implemented")
				}
			}
		case 2:
			fmt.Println("Client management not fully implemented")
			// Add customer directly in main
			var name, email, phone, address string
			fmt.Print("Enter client name: ")
			fmt.Scan(&name)
			fmt.Print("Enter email: ")
			fmt.Scan(&email)
			fmt.Print("Enter phone: ")
			fmt.Scan(&phone)
			fmt.Print("Enter address: ")
			fmt.Scan(&address)
			store.AddCustomer(name, email, phone, address)
		case 3:
			var customerID int
			fmt.Print("Enter client ID: ")
			fmt.Scan(&customerID)
			for {
				fmt.Println("\n--- Cart menu ---")
				fmt.Println("1. Add product to cart")
				fmt.Println("2. Remove product from cart")
				fmt.Println("3. View cart")
				fmt.Println("4. Apply discount")
				fmt.Println("5. Place order")
				fmt.Println("6. Back to main menu")
				fmt.Print("> ")
				fmt.Scan(&choice)

				if choice == 6 {
					break
				}

				switch choice {
				case 1:
					var productID, quantity int
					fmt.Print("Enter product ID: ")
					fmt.Scan(&productID)
					fmt.Print("Enter quantity: ")
					fmt.Scan(&quantity)
					store.AddToCart(customerID, productID, quantity)
				case 3:
					store.ViewCart(customerID)
				default:
					fmt.Println("Function not implemented")
				}
			}
		case 4:
			fmt.Println("Order system not implemented")
		case 5:
			fmt.Println("Store stats not implemented")
		}
	}
}
