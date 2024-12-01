package report

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email string
}

type Product struct {
	Count    int
	Name     string
	Planilla int
}

type OrderAlegra struct {
	Total       int
	Status      string
	Tip         int
	Delivery    int
	Electronic  bool
	PaymentType string
	User        User
	Products    []Product
}

type OrdersArray struct {
	Orders []OrderAlegra
}

type ResponseReport struct {
	Total           int
	TotalElectronic int
	TotalCash       int
	TotalBank       int
	TotalDelivery   int
	TotalTip        int
	TotalCobrado    int
}

type ProductReport struct {
	Count int
}

type finalReport struct {
	GeneralReport ResponseReport
	UserReport    map[string]ResponseReport
	ProductReport map[string]int
}

func CalculateReport(c *fiber.Ctx) error {
	fmt.Println("calculating report...")
	//fmt.Printf("%+v", string(c.Body())[:])
	fmt.Println()
	fmt.Println()
	fmt.Println("...starting playing with response")

	var jsonResponse OrdersArray
	err1 := json.Unmarshal(c.Body(), &jsonResponse)

	if err1 != nil {
		fmt.Println("error in body:", err1)
	}

	if len(jsonResponse.Orders) == 0 {
		fmt.Println("Sorry No data found...")
		return c.SendString("No data found...")
	} else {
		fmt.Printf("%+v", jsonResponse.Orders[0])
	}

	fmt.Println()

	//var reportResponse ResponseReport
	TotalReponse := 0
	TotalElectronic := 0
	TotalBank := 0
	TotalTip := 0
	TotalDelivery := 0
	TotalCobrado := 0

	usersMap := make(map[string]ResponseReport)
	productMap := make(map[string]int)

	for _, v := range jsonResponse.Orders {

		TotalReponse = TotalReponse + v.Total

		userEmail := v.User.Email

		userOrder, exists := usersMap[userEmail]

		if !exists {
			userOrder = ResponseReport{Total: 0, TotalElectronic: 0, TotalCash: 0, TotalBank: 0, TotalDelivery: 0, TotalTip: 0, TotalCobrado: 0}
		}

		userOrder.Total += v.Total

		if v.Electronic {
			TotalElectronic = TotalElectronic + v.Total
			userOrder.TotalElectronic += v.Total
		}

		if v.PaymentType != "efectivo" && v.Status == "pagada" {
			TotalBank = TotalBank + v.Total + v.Tip + v.Delivery
			userOrder.TotalBank = userOrder.TotalBank + v.Total + v.Tip + v.Delivery
		}

		if v.Status == "pagada" {
			TotalCobrado = TotalCobrado + v.Total
			TotalTip = TotalTip + v.Tip
			TotalDelivery = TotalDelivery + v.Delivery
			userOrder.TotalCobrado += v.Total
			userOrder.TotalTip += v.Tip
			userOrder.TotalDelivery += v.Delivery
		}

		userOrder.TotalCash = userOrder.Total - userOrder.TotalBank

		usersMap[userEmail] = userOrder

		//Analizing products total

		for _, product := range v.Products {
			productName := product.Name

			value, exists := productMap[productName]

			if !exists {
				productMap[productName] = product.Count
			}

			productMap[productName] = value + product.Count
		}

	}

	reportResponse := ResponseReport{
		Total:           TotalReponse,
		TotalElectronic: TotalElectronic,
		TotalCash:       TotalReponse - TotalBank,
		TotalBank:       TotalBank,
		TotalDelivery:   TotalDelivery,
		TotalTip:        TotalTip,
		TotalCobrado:    TotalCobrado,
	}

	finalReport := finalReport{
		GeneralReport: reportResponse,
		UserReport:    usersMap,
		ProductReport: productMap,
	}

	fmt.Printf("%+v", reportResponse)
	fmt.Println()
	fmt.Println("Printing map..")
	fmt.Println()
	fmt.Printf("%+v", usersMap)

	fmt.Println()
	fmt.Println()

	// var jsonBodyparser OrdersArray
	// if err3 := c.BodyParser(&jsonBodyparser); err3 != nil {
	// 	fmt.Println("Error in body parser:", err3)
	// }

	// fmt.Printf("%+v", jsonBodyparser)
	// fmt.Println()

	fmt.Println("Ending....")
	return c.Status(fiber.StatusCreated).JSON(finalReport)
}
