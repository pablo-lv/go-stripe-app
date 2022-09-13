package main

import (
	"net/http"
	"stripeapp/internal/models"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["publishable_key"] = app.config.stripe.key

	if err := app.renderTemplate(w, r, "terminal", &templateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
	}

	cardHolder := r.Form.Get("cardholder-name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["paymentIntent"] = paymentIntent
	data["paymentMethod"] = paymentMethod
	data["paymentAmount"] = paymentAmount
	data["paymentCurrency"] = paymentCurrency

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) ChargeOne(w http.ResponseWriter, r *http.Request) {
	widget := models.Widget{
		ID:             1,
		Name:           "custom",
		Description:    "A very nice widget",
		InventoryLevel: 10,
		Price:          1000,
	}

	data := make(map[string]interface{})

	data["widget"] = widget

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}
