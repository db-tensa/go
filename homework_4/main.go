package main


// i hate this code... 
import (
	"fmt"
	"math"
	"time"
)

var standardRate float64 = 1.5
var expressRate float64 = 2.5
var standardPackagingRate float64 = 40.0
var reinforcedPackagingRate float64 = 50.0
var premiumPackagingRate float64 = 75.0

func getNumberInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scanln(&input)
	if input <= 0 {
		fmt.Println("Enter a positive number!")
		return getNumberInput(prompt)
	}
	return input
}

func getIntInput(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scanln(&input)
	if input < 1 || input > 3 {
		fmt.Println("Enter a number from 1 to 3!")
		return getIntInput(prompt)
	}
	return input
}

func getYesNoInput(prompt string) bool {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	if input == "yes" {
		return true
	} else if input == "no" {
		return false
	}
	fmt.Println("Enter 'yes' or 'no'!")
	return getYesNoInput(prompt)
}

func calculateBasePrice(weight, distance float64) float64 {
	return math.Round(weight*distance*standardRate*100) / 100
}

func calculateDeliveryTypePrice(basePrice float64, deliveryType int) float64 {
	if deliveryType == 2 {
		return math.Round(basePrice*0.5*100) / 100
	}
	return 0
}

func calculateDiscount(basePrice float64, clientStatus int) float64 {
	if clientStatus == 2 {
		return math.Round(basePrice*0.15*100) / 100
	}
	return 0
}

func calculateFinalPrice(basePrice, additionalPrice, discount float64) float64 {
	return math.Round((basePrice+additionalPrice-discount)*100) / 100
}

func calculateBaseDeliveryTime(distance float64, deliveryType int) float64 {
	var speed float64
	if deliveryType == 1 {
		speed = 100
	} else {
		speed = 200
	}
	return math.Ceil(distance / speed)
}

func addWeatherDelay(baseTime float64, weatherCondition int) float64 {
	switch weatherCondition {
	case 2:
		return baseTime + 0.5
	case 3:
		return baseTime + 1.0
	default:
		return baseTime
	}
}

func calculateFinalDeliveryTime(baseTime, weatherDelay float64, isWeekend bool) float64 {
	totalTime := weatherDelay
	if isWeekend {
		totalTime += 1.0
	}
	return math.Round(totalTime*100) / 100
}

func calculatePackagingMaterial(length, width, height float64) float64 {
	surfaceArea := 2*(length*width+length*height+width*height) / 10000
	return math.Round(surfaceArea*100) / 100
}

func calculatePackagingCost(materialAmount float64, materialType int) float64 {
	var rate float64
	switch materialType {
	case 1:
		rate = standardPackagingRate
	case 2:
		rate = reinforcedPackagingRate
	case 3:
		rate = premiumPackagingRate
	}
	return math.Round(materialAmount*rate*100) / 100
}

func main() {
	for {
		fmt.Println("=== Parcel Delivery Calculator ===")
		fmt.Println("Select an option:")
		fmt.Println("1. Calculate delivery cost")
		fmt.Println("2. Estimate delivery time")
		fmt.Println("3. Calculate packaging materials")
		fmt.Println("4. Exit")
		choice := getIntInput("> ")

		if choice == 4 {
			fmt.Println("Goodbye!")
			break
		}

		switch choice {
		case 1:
			fmt.Println("--- Calculate Delivery Cost ---")
			weight := getNumberInput("Enter parcel weight (kg): ")
			getNumberInput("Enter parcel length (cm): ")
			getNumberInput("Enter parcel width (cm): ")
			getNumberInput("Enter parcel height (cm): ")
			fmt.Println("Select delivery type:")
			fmt.Println("1. Standard")
			fmt.Println("2. Express")
			deliveryType := getIntInput("> ")
			distance := getNumberInput("Enter delivery distance (km): ")
			fmt.Println("Select client status:")
			fmt.Println("1. Regular")
			fmt.Println("2. Loyal")
			clientStatus := getIntInput("> ")

			basePrice := calculateBasePrice(weight, distance)
			additionalPrice := calculateDeliveryTypePrice(basePrice, deliveryType)
			discount := calculateDiscount(basePrice, clientStatus)
			finalPrice := calculateFinalPrice(basePrice, additionalPrice, discount)

			fmt.Println("Delivery cost calculation results:")
			fmt.Printf("Base cost: %.2f UAH\n", basePrice)
			fmt.Printf("Additional cost (Express): %.2f UAH\n", additionalPrice)
			fmt.Printf("Discount (Loyal client): %.2f UAH\n", discount)
			fmt.Printf("Total cost: %.2f UAH\n", finalPrice)

		case 2:
			fmt.Println("--- Estimate Delivery Time ---")
			distance := getNumberInput("Enter delivery distance (km): ")
			fmt.Println("Select delivery type:")
			fmt.Println("1. Standard")
			fmt.Println("2. Express")
			deliveryType := getIntInput("> ")
			fmt.Println("Select weather conditions:")
			fmt.Println("1. Good")
			fmt.Println("2. Fair")
			fmt.Println("3. Poor")
			weatherCondition := getIntInput("> ")
			isWeekend := getYesNoInput("Is today a weekend? (yes/no): ")

			baseTime := calculateBaseDeliveryTime(distance, deliveryType)
			weatherDelay := addWeatherDelay(baseTime, weatherCondition)
			finalTime := calculateFinalDeliveryTime(baseTime, weatherDelay, isWeekend)
			deliveryDate := time.Now().AddDate(0, 0, int(finalTime))

			fmt.Println("Delivery time estimation results:")
			fmt.Printf("Base delivery time (Standard): %.0f days\n", baseTime)
			fmt.Printf("Delay due to weather conditions: %.1f days\n", weatherDelay-baseTime)
			fmt.Printf("Total estimated delivery time: %.1f days\n", finalTime)
			fmt.Printf("Estimated arrival date: %s\n", deliveryDate.Format("02 January 2006"))

		case 3:
			fmt.Println("--- Calculate Packaging Materials ---")
			length := getNumberInput("Enter parcel length (cm): ")
			width := getNumberInput("Enter parcel width (cm): ")
			height := getNumberInput("Enter parcel height (cm): ")
			fmt.Println("Select packaging material type:")
			fmt.Println("1. Standard cardboard")
			fmt.Println("2. Reinforced cardboard with film")
			fmt.Println("3. Premium packaging")
			materialType := getIntInput("> ")

			materialAmount := calculatePackagingMaterial(length, width, height)
			packagingCost := calculatePackagingCost(materialAmount, materialType)

			fmt.Println("Packaging materials calculation results:")
			fmt.Printf("Required packaging material amount: %.2f m²\n", materialAmount)
			fmt.Printf("Packaging materials cost: %.2f UAH\n", packagingCost)
		}

		if !getYesNoInput("Would you like to return to the main menu? (yes/no): ") {
			fmt.Println("Goodbye!")
			break
		}
	}
}
