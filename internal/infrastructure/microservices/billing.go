package dcollect

import (
	"fmt"
	"graduatework/internal/model"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

// ReadBilling is a method to get billing data from simulator
func (m *MicroServiceStr) ReadBilling() (outputData model.BillingData) {

	body, err := ioutil.ReadFile("./simulator/billing.data")
	if err != nil {
		log.Print("error reading the billing data", err)
	}

	var num uint8
	for i := len(body) - 1; i >= 0; i-- {
		a := body[i]
		b := string(a)

		x, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println("error to assert string to int: ", err)
		}

		degr := float64(x)

		if degr == 0 {
			continue
		} else {
			degr = float64(i)
		}
		num = num + uint8(math.Pow(2, degr))

	}

	if num&1 == 1 {
		outputData.CheckoutPage = true
	}
	if num&2 == 2 {
		outputData.FraudControl = true
	}
	if num&4 == 4 {
		outputData.Recurring = true
	}
	if num&8 == 8 {
		outputData.Payout = true

	}
	if num&16 == 16 {
		outputData.Purchase = true
	}
	if num&32 == 32 {
		outputData.CreateCustomer = true
	}
	return outputData
}
