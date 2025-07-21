package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PaymentListResponse represents the PaymentListResponse schema from the OpenAPI specification
type PaymentListResponse struct {
	PagesCount float64 `json:"pagesCount,omitempty"`
	Total float64 `json:"total,omitempty"`
	Data []Payment `json:"data,omitempty"`
	Limit float64 `json:"limit,omitempty"`
	Page float64 `json:"page,omitempty"`
}

// ApproximatePrice represents the ApproximatePrice schema from the OpenAPI specification
type ApproximatePrice struct {
	Currency_from string `json:"currency_from,omitempty"`
	Currency_to string `json:"currency_to,omitempty"`
	Estimated_amount float64 `json:"estimated_amount,omitempty"`
	Amount_from float64 `json:"amount_from,omitempty"`
}

// MinimumPayment represents the MinimumPayment schema from the OpenAPI specification
type MinimumPayment struct {
	Currency_from string `json:"currency_from,omitempty"`
	Currency_to string `json:"currency_to,omitempty"`
	Min_amount float64 `json:"min_amount,omitempty"`
}

// Payment represents the Payment schema from the OpenAPI specification
type Payment struct {
	Order_description string `json:"order_description,omitempty"`
	Outcome_currency string `json:"outcome_currency,omitempty"`
	Pay_address string `json:"pay_address,omitempty"`
	Pay_amount float64 `json:"pay_amount,omitempty"`
	Pay_currency string `json:"pay_currency,omitempty"`
	Order_id string `json:"order_id,omitempty"`
	Actually_paid float64 `json:"actually_paid,omitempty"`
	Payment_id float64 `json:"payment_id,omitempty"`
	Purchase_id string `json:"purchase_id,omitempty"`
	Outcome_amount float64 `json:"outcome_amount,omitempty"`
	Payment_status string `json:"payment_status,omitempty"`
	Price_amount float64 `json:"price_amount,omitempty"`
	Price_currency string `json:"price_currency,omitempty"`
}
