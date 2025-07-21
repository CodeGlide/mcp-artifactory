package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CustomerAddress represents the CustomerAddress schema from the OpenAPI specification
type CustomerAddress struct {
	Locality string `json:"locality,omitempty"`
	Postal_code string `json:"postal_code,omitempty"`
	Address_line_1 string `json:"address_line_1,omitempty"`
	Administrative_district_level_1 string `json:"administrative_district_level_1,omitempty"`
}

// OrderReturnLineItemModifier represents the OrderReturnLineItemModifier schema from the OpenAPI specification
type OrderReturnLineItemModifier struct {
	Total_price_money Money `json:"total_price_money,omitempty"`
	Uid string `json:"uid,omitempty"`
	Base_price_money Money `json:"base_price_money,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"`
	Name string `json:"name,omitempty"`
	Source_modifier_uid string `json:"source_modifier_uid,omitempty"`
}

// OrderLineItemAppliedTax represents the OrderLineItemAppliedTax schema from the OpenAPI specification
type OrderLineItemAppliedTax struct {
	Applied_money Money `json:"applied_money,omitempty"`
	Tax_uid string `json:"tax_uid"`
	Uid string `json:"uid,omitempty"`
}

// OrderRefund represents the OrderRefund schema from the OpenAPI specification
type OrderRefund struct {
	Tender_id string `json:"tender_id"`
	Amount_money Money `json:"amount_money"`
	Reason string `json:"reason"`
	Status string `json:"status"`
	Transaction_id string `json:"transaction_id,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id string `json:"id"`
	Processing_fee_money Money `json:"processing_fee_money,omitempty"`
	Location_id string `json:"location_id"`
}

// OrderPricingOptions represents the OrderPricingOptions schema from the OpenAPI specification
type OrderPricingOptions struct {
	Auto_apply_taxes bool `json:"auto_apply_taxes,omitempty"`
	Auto_apply_discounts bool `json:"auto_apply_discounts,omitempty"`
}

// OrderReward represents the OrderReward schema from the OpenAPI specification
type OrderReward struct {
	Id string `json:"id"` // The reward's unique ID.
	Reward_tier_id string `json:"reward_tier_id"` // The ID of the reward tier that the reward belongs to.
}

// FulfillmentRecipient represents the FulfillmentRecipient schema from the OpenAPI specification
type FulfillmentRecipient struct {
	Display_name string `json:"display_name,omitempty"`
	Email_address string `json:"email_address,omitempty"`
	Phone_number string `json:"phone_number,omitempty"`
	Address Address `json:"address,omitempty"`
	Customer_id string `json:"customer_id,omitempty"`
}

// OrderReturnDiscount represents the OrderReturnDiscount schema from the OpenAPI specification
type OrderReturnDiscount struct {
	Catalog_object_id string `json:"catalog_object_id,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"`
	Source_discount_uid string `json:"source_discount_uid,omitempty"`
	Applied_money Money `json:"applied_money,omitempty"`
	Type string `json:"type,omitempty"`
	Percentage string `json:"percentage,omitempty"`
	Amount_money Money `json:"amount_money,omitempty"`
	Name string `json:"name,omitempty"`
	Scope string `json:"scope,omitempty"`
	Uid string `json:"uid,omitempty"`
}

// OrderServiceCharge represents the OrderServiceCharge schema from the OpenAPI specification
type OrderServiceCharge struct {
	Type string `json:"type,omitempty"`
	Percentage string `json:"percentage,omitempty"` // The percentage of the service charge, as a string representation of a decimal number. A value of `7.25` corresponds to a percentage of 7.25%.
	Taxable bool `json:"taxable,omitempty"`
	Applied_money Money `json:"applied_money,omitempty"`
	Total_money Money `json:"total_money,omitempty"`
	Amount_money Money `json:"amount_money,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"` // The version of the service charge's associated CatalogObject, if any.
	Name string `json:"name,omitempty"` // The name of the service charge.
	Uid string `json:"uid,omitempty"` // The service charge's unique ID.
	Applied_taxes []OrderLineItemAppliedTax `json:"applied_taxes,omitempty"`
	Calculation_phase string `json:"calculation_phase,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"` // The ID of the service charge's associated CatalogObject, if any.
	Total_tax_money Money `json:"total_tax_money,omitempty"`
}

// TenderCashDetails represents the TenderCashDetails schema from the OpenAPI specification
type TenderCashDetails struct {
	Buyer_tendered_money Money `json:"buyer_tendered_money,omitempty"`
	Change_back_money Money `json:"change_back_money,omitempty"`
}

// FulfillmentShipmentDetails represents the FulfillmentShipmentDetails schema from the OpenAPI specification
type FulfillmentShipmentDetails struct {
	In_progress_at string `json:"in_progress_at,omitempty"`
	Recipient FulfillmentRecipient `json:"recipient,omitempty"`
	Packaged_at string `json:"packaged_at,omitempty"`
	Tracking_url string `json:"tracking_url,omitempty"`
	Tracking_number string `json:"tracking_number,omitempty"`
	Canceled_at string `json:"canceled_at,omitempty"`
	Failure_reason string `json:"failure_reason,omitempty"`
	Expected_shipped_at string `json:"expected_shipped_at,omitempty"`
	Shipped_at string `json:"shipped_at,omitempty"`
	Carrier string `json:"carrier,omitempty"`
	Failed_at string `json:"failed_at,omitempty"`
	Shipping_note string `json:"shipping_note,omitempty"`
	Shipping_type string `json:"shipping_type,omitempty"`
	Cancel_reason string `json:"cancel_reason,omitempty"`
	Placed_at string `json:"placed_at,omitempty"`
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Administrative_district_level_2 string `json:"administrative_district_level_2,omitempty"`
	Sublocality string `json:"sublocality,omitempty"`
	Administrative_district_level_3 string `json:"administrative_district_level_3,omitempty"`
	Last_name string `json:"last_name,omitempty"`
	Postal_code string `json:"postal_code,omitempty"`
	Sublocality_2 string `json:"sublocality_2,omitempty"`
	Administrative_district_level_1 string `json:"administrative_district_level_1,omitempty"`
	Country string `json:"country,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Locality string `json:"locality,omitempty"`
	Address_line_2 string `json:"address_line_2,omitempty"`
	Sublocality_3 string `json:"sublocality_3,omitempty"`
	Address_line_1 string `json:"address_line_1,omitempty"`
	Address_line_3 string `json:"address_line_3,omitempty"`
}

// OrderReturnServiceCharge represents the OrderReturnServiceCharge schema from the OpenAPI specification
type OrderReturnServiceCharge struct {
	Source_service_charge_uid string `json:"source_service_charge_uid,omitempty"`
	Taxable bool `json:"taxable,omitempty"`
	Applied_taxes []OrderLineItemAppliedTax `json:"applied_taxes,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"`
	Percentage string `json:"percentage,omitempty"`
	Total_money Money `json:"total_money,omitempty"`
	Total_tax_money Money `json:"total_tax_money,omitempty"`
	Amount_money Money `json:"amount_money,omitempty"`
	Calculation_phase string `json:"calculation_phase,omitempty"`
	Uid string `json:"uid,omitempty"`
	Applied_money Money `json:"applied_money,omitempty"`
	Name string `json:"name,omitempty"`
}

// OrderLineItem represents the OrderLineItem schema from the OpenAPI specification
type OrderLineItem struct {
	Quantity_unit OrderQuantityUnit `json:"quantity_unit,omitempty"`
	Total_discount_money Money `json:"total_discount_money,omitempty"`
	Total_tax_money Money `json:"total_tax_money,omitempty"`
	Uid string `json:"uid,omitempty"` // The line item's unique ID.
	Item_type string `json:"item_type,omitempty"`
	Gross_sales_money Money `json:"gross_sales_money,omitempty"`
	Pricing_blocklists []OrderLineItemPricingBlocklists `json:"pricing_blocklists,omitempty"`
	Applied_discounts []OrderLineItemAppliedDiscount `json:"applied_discounts,omitempty"`
	Modifiers []OrderLineItemModifier `json:"modifiers,omitempty"`
	Base_price_money Money `json:"base_price_money,omitempty"`
	Note string `json:"note,omitempty"`
	Variation_name string `json:"variation_name,omitempty"`
	Variation_total_price_money Money `json:"variation_total_price_money,omitempty"`
	Name string `json:"name,omitempty"` // The name of the line item.
	Applied_taxes []OrderLineItemAppliedTax `json:"applied_taxes,omitempty"`
	Quantity string `json:"quantity"` // The quantity purchased, formatted as a decimal number.
	Total_money Money `json:"total_money,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"`
}

// OrderReturnLineItem represents the OrderReturnLineItem schema from the OpenAPI specification
type OrderReturnLineItem struct {
	Uid string `json:"uid,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"`
	Name string `json:"name,omitempty"`
	Applied_taxes []OrderLineItemAppliedTax `json:"applied_taxes,omitempty"`
	Applied_discounts []OrderLineItemAppliedDiscount `json:"applied_discounts,omitempty"`
	Note string `json:"note,omitempty"`
	Total_money Money `json:"total_money,omitempty"`
	Base_price_money Money `json:"base_price_money,omitempty"`
	Gross_return_money Money `json:"gross_return_money,omitempty"`
	Source_line_item_uid string `json:"source_line_item_uid,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"`
	Return_modifiers []OrderReturnLineItemModifier `json:"return_modifiers,omitempty"`
	Total_discount_money Money `json:"total_discount_money,omitempty"`
	Quantity_unit OrderQuantityUnit `json:"quantity_unit,omitempty"`
	Quantity string `json:"quantity"`
	Item_type string `json:"item_type,omitempty"`
	Variation_name string `json:"variation_name,omitempty"`
	Total_tax_money Money `json:"total_tax_money,omitempty"`
	Variation_total_price_money Money `json:"variation_total_price_money,omitempty"`
}

// CustomerCreatedWebhookEventContext represents the CustomerCreatedWebhookEventContext schema from the OpenAPI specification
type CustomerCreatedWebhookEventContext struct {
	Merge CustomerCreatedWebhookEventContextMerge `json:"merge"` // Information about the merge operation associated with the event.
}

// OrderReturn represents the OrderReturn schema from the OpenAPI specification
type OrderReturn struct {
	Return_discounts []OrderReturnDiscount `json:"return_discounts,omitempty"`
	Return_line_items []OrderReturnLineItem `json:"return_line_items,omitempty"`
	Return_service_charges []OrderReturnServiceCharge `json:"return_service_charges,omitempty"`
	Return_taxes []OrderReturnTax `json:"return_taxes,omitempty"`
	Rounding_adjustment OrderRoundingAdjustment `json:"rounding_adjustment,omitempty"`
	Source_order_id string `json:"source_order_id,omitempty"`
	Uid string `json:"uid,omitempty"`
	Return_amounts OrderMoneyAmounts `json:"return_amounts,omitempty"`
}

// OrderLineItemAppliedDiscount represents the OrderLineItemAppliedDiscount schema from the OpenAPI specification
type OrderLineItemAppliedDiscount struct {
	Uid string `json:"uid,omitempty"`
	Applied_money Money `json:"applied_money,omitempty"`
	Discount_uid string `json:"discount_uid"`
}

// OrderRoundingAdjustment represents the OrderRoundingAdjustment schema from the OpenAPI specification
type OrderRoundingAdjustment struct {
	Amount_money Money `json:"amount_money,omitempty"`
	Name string `json:"name,omitempty"`
	Uid string `json:"uid,omitempty"`
}

// OrderReturnTax represents the OrderReturnTax schema from the OpenAPI specification
type OrderReturnTax struct {
	Applied_money Money `json:"applied_money,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Catalog_version int `json:"catalog_version,omitempty"`
	Source_tax_uid string `json:"source_tax_uid,omitempty"`
	Uid string `json:"uid,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"`
	Percentage string `json:"percentage,omitempty"`
	Scope string `json:"scope,omitempty"`
}

// WebhookSubscription represents the WebhookSubscription schema from the OpenAPI specification
type WebhookSubscription struct {
	Api_version string `json:"api_version,omitempty"`
	Enabled bool `json:"enabled"`
	Event_types []string `json:"event_types"`
	Id string `json:"id"`
	Name string `json:"name"`
	Notification_url string `json:"notification_url"`
	Signature_key string `json:"signature_key"`
}

// OrderCreatedWebhookEvent represents the OrderCreatedWebhookEvent schema from the OpenAPI specification
type OrderCreatedWebhookEvent struct {
	Type string `json:"type"` // The type of the event.
	Created_at string `json:"created_at"`
	Data OrderCreatedWebhookData `json:"data"` // The data included in the event.
	Event_id string `json:"event_id"` // The ID of the event.
	Merchant_id string `json:"merchant_id"` // The ID of the merchant that created the order.
}

// CustomerPreferences represents the CustomerPreferences schema from the OpenAPI specification
type CustomerPreferences struct {
	Email_unsubscribed bool `json:"email_unsubscribed,omitempty"` // Indicates whether the customer has unsubscribed from marketing campaign emails. A value of true means that the customer chose to opt out of email marketing from the current Square seller or from all Square sellers. This value is read-only from the Customers API.
}

// TenderCardDetails represents the TenderCardDetails schema from the OpenAPI specification
type TenderCardDetails struct {
	Card Card `json:"card,omitempty"`
	Entry_method string `json:"entry_method,omitempty"`
	Status string `json:"status,omitempty"`
}

// OrderCreatedWebhookOrder represents the OrderCreatedWebhookOrder schema from the OpenAPI specification
type OrderCreatedWebhookOrder struct {
	Created_at string `json:"created_at"` // The timestamp when the order was created, in RFC 3339 format.
	Location_id string `json:"location_id"` // The ID of the location where the order was created.
	Order_id string `json:"order_id"` // The ID of the order.
	State string `json:"state"` // The state of the order.
	Version int `json:"version"` // The version of the order.
}

// CustomerCreatedWebhookObject represents the CustomerCreatedWebhookObject schema from the OpenAPI specification
type CustomerCreatedWebhookObject struct {
	Customer Customer `json:"customer"`
	Event_context CustomerCreatedWebhookEventContext `json:"event_context,omitempty"` // Information about the change that triggered the event. This field is returned only if the customer is created by a merge operation.
}

// CustomerCreatedWebhookEvent represents the CustomerCreatedWebhookEvent schema from the OpenAPI specification
type CustomerCreatedWebhookEvent struct {
	Event_id string `json:"event_id"` // The unique ID of the event, which is used for [idempotency support](https://developer.squareup.com/docs/webhooks/step4manage#webhooks-best-practices)
	Merchant_id string `json:"merchant_id"` // The ID of the seller associated with the event.
	Type string `json:"type"` // The type of event. For this object, the value is customer.created.
	Created_at string `json:"created_at"` // Read only The timestamp of when the event was created, in RFC 3339 format. Examples for January 25th, 2020 6:25:34pm Pacific Standard Time: UTC: 2020-01-26T02:25:34Z Pacific Standard Time with UTC offset: 2020-01-25T18:25:34-08:00
	Data CustomerCreatedWebhookData `json:"data"` // The data associated with the event.
}

// OrderLineItemModifier represents the OrderLineItemModifier schema from the OpenAPI specification
type OrderLineItemModifier struct {
	Total_price_money Money `json:"total_price_money,omitempty"`
	Uid string `json:"uid,omitempty"` // The modifier's unique ID.
	Base_price_money Money `json:"base_price_money,omitempty"`
	Catalog_object_id string `json:"catalog_object_id,omitempty"` // The ID of the item modifier applied to the line item.
	Catalog_version int `json:"catalog_version,omitempty"` // The version of the item modifier applied to the line item.
	Name string `json:"name,omitempty"` // The name of the item modifier.
}

// WebhookSubscriptionResult represents the WebhookSubscriptionResult schema from the OpenAPI specification
type WebhookSubscriptionResult struct {
	Subscription WebhookSubscription `json:"subscription"`
}

// Fulfillment represents the Fulfillment schema from the OpenAPI specification
type Fulfillment struct {
	Line_iten_application string `json:"line_iten_application,omitempty"`
	Pickup_details FulfillmentPickupDetails `json:"pickup_details,omitempty"`
	State string `json:"state,omitempty"` // The current state of this fulfillment.
	Type string `json:"type,omitempty"` // The type of fulfillment.
	Uid string `json:"uid,omitempty"` // The fulfillment's unique ID.
	Entries []FulfillmentEntry `json:"entries,omitempty"`
}

// Card represents the Card schema from the OpenAPI specification
type Card struct {
	Billing_address Address `json:"billing_address,omitempty"`
	Reference_id string `json:"reference_id,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Card_brand string `json:"card_brand,omitempty"`
	Customer_id string `json:"customer_id,omitempty"`
	Prepaid_type string `json:"prepaid_type,omitempty"`
	Bin string `json:"bin,omitempty"`
	Exp_year int `json:"exp_year,omitempty"`
	Card_co_brand string `json:"card_co_brand,omitempty"`
	Version string `json:"version,omitempty"`
	Card_type string `json:"card_type,omitempty"`
	Merchant_id string `json:"merchant_id,omitempty"`
	Exp_month int `json:"exp_month,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last_4 string `json:"last_4,omitempty"`
	Cardholder_name string `json:"cardholder_name,omitempty"`
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Note string `json:"note,omitempty"` // A custom note associated with the customer profile.
	Reference_id string `json:"reference_id,omitempty"` // An optional second ID used to associate the customer profile with an entity in another system.
	Preferences CustomerPreferences `json:"preferences"` // Represents general customer preferences.
	Address CustomerAddress `json:"address"` // The physical address associated with the customer profile.
	Company_name string `json:"company_name,omitempty"` // A business name associated with the customer profile.
	Segment_ids []string `json:"segment_ids,omitempty"` // The IDs of segments the customer belongs to.
	Nickname string `json:"nickname,omitempty"` // A nickname for the customer profile.
	Id string `json:"id"` // A unique Square-assigned ID for the customer profile. If you need this ID for an API request, use the ID returned when you created the customer profile or call the SearchCustomers or ListCustomers endpoint.
	Tax_ids CustomerTaxIds `json:"tax_ids,omitempty"` // The tax ID associated with the customer profile. This field is present only for customers of sellers in EU countries or the United Kingdom. For more information, see Customer tax IDs.
	Updated_at string `json:"updated_at"` // Read only The timestamp when the customer profile was last updated, in RFC 3339 format. Examples for January 25th, 2020 6:25:34pm Pacific Standard Time: UTC: 2020-01-26T02:25:34Z Pacific Standard Time with UTC offset: 2020-01-25T18:25:34-08:00
	Given_name string `json:"given_name"` // The given name (that is, the first name) associated with the customer profile.
	Phone_number string `json:"phone_number"` // The phone number associated with the customer profile. A phone number can contain 9–16 digits, with an optional + prefix.
	Email_address string `json:"email_address"` // The email address associated with the customer profile.
	Version float64 `json:"version"` // The Square-assigned version number of the customer profile. The version number is incremented each time an update is committed to the customer profile, except for changes to customer segment membership and cards on file.
	Birthday string `json:"birthday,omitempty"` // The birthday associated with the customer profile, in RFC 3339 format. The year is optional. The timezone and time are not allowed. For example, 0000-09-21T00:00:00-00:00 represents a birthday on September 21 and 1998-09-21T00:00:00-00:00 represents a birthday on September 21, 1998.
	Created_at string `json:"created_at"` // Read only The timestamp when the customer profile was created, in RFC 3339 format. Examples for January 25th, 2020 6:25:34pm Pacific Standard Time: UTC: 2020-01-26T02:25:34Z Pacific Standard Time with UTC offset: 2020-01-25T18:25:34-08:00
	Family_name string `json:"family_name"` // The family name (that is, the last name) associated with the customer profile.
	Group_ids []string `json:"group_ids,omitempty"` // The IDs of customer groups the customer belongs to.
	Creation_source string `json:"creation_source"` // The method used to create the customer profile.
}

// CustomerCreatedWebhookEventContextMerge represents the CustomerCreatedWebhookEventContextMerge schema from the OpenAPI specification
type CustomerCreatedWebhookEventContextMerge struct {
	From_customer_ids []string `json:"from_customer_ids"` // The IDs of the existing customers that were merged and then deleted.
	To_customer_id string `json:"to_customer_id"` // The ID of the new customer created by the merge.
}

// Money represents the Money schema from the OpenAPI specification
type Money struct {
	Amount int `json:"amount,omitempty"` // The amount of money, in the smallest denomination of the currency indicated by currency.
	Currency string `json:"currency,omitempty"` // The currency code.
}

// FulfillmentEntry represents the FulfillmentEntry schema from the OpenAPI specification
type FulfillmentEntry struct {
	Line_item_uid string `json:"line_item_uid"`
	Quantity int `json:"quantity"`
	Uid string `json:"uid,omitempty"`
}

// CustomerCreatedWebhookData represents the CustomerCreatedWebhookData schema from the OpenAPI specification
type CustomerCreatedWebhookData struct {
	Id string `json:"id"` // The ID of the new customer.
	Object CustomerCreatedWebhookObject `json:"object"` // An object that contains the new customer.
	Type string `json:"type"` // The type of object affected by the event. For this event, the value is "customer".
}

// Tender represents the Tender schema from the OpenAPI specification
type Tender struct {
	Tip_money Money `json:"tip_money,omitempty"`
	Type string `json:"type,omitempty"`
	Amount_money Money `json:"amount_money,omitempty"`
	Location_id string `json:"location_id,omitempty"` // The ID of the tender's associated location.
	Customer_id string `json:"customer_id,omitempty"`
	Processing_fee_money Money `json:"processing_fee_money,omitempty"`
	Transaction_id string `json:"transaction_id,omitempty"` // The ID of the tender's associated transaction.
	Card_details TenderCardDetails `json:"card_details,omitempty"`
	Cash_details TenderCashDetails `json:"cash_details,omitempty"`
	Id string `json:"id,omitempty"` // The tender's unique ID.
	Note string `json:"note,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Payment_id string `json:"payment_id,omitempty"`
}

// OrderCreatedWebhookData represents the OrderCreatedWebhookData schema from the OpenAPI specification
type OrderCreatedWebhookData struct {
	Id string `json:"id"` // The ID of the order.
	Object OrderCreatedWebhookObject `json:"object"` // The object included in the event.
	Type string `json:"type"` // The type of the event.
}

// OrderQuantityUnit represents the OrderQuantityUnit schema from the OpenAPI specification
type OrderQuantityUnit struct {
	Catalog_object_id string `json:"catalog_object_id,omitempty"` // The ID of the CatalogObject representing the unit used to measure the line item quantity.
	Catalog_version int `json:"catalog_version,omitempty"` // The version of the CatalogObject representing the unit used to measure the line item quantity.
	Measurement_unit map[string]interface{} `json:"measurement_unit,omitempty"`
	Precision int `json:"precision,omitempty"` // The precision of the quantity, in number of decimal places. For example, a precision of 2 indicates that the quantity can be split into hundredths.
}

// FulfillmentDeliveryDetails represents the FulfillmentDeliveryDetails schema from the OpenAPI specification
type FulfillmentDeliveryDetails struct {
	Note string `json:"note,omitempty"`
	Courier_pickup_window_duration string `json:"courier_pickup_window_duration,omitempty"`
	Courier_provider_name string `json:"courier_provider_name,omitempty"`
	Deliver_at string `json:"deliver_at,omitempty"`
	Rejected_at string `json:"rejected_at,omitempty"`
	Deliver_window_duration string `json:"deliver_window_duration,omitempty"`
	In_progress_at string `json:"in_progress_at,omitempty"`
	Placed_at string `json:"placed_at,omitempty"`
	Cancel_reason string `json:"cancel_reason,omitempty"`
	Completed_at string `json:"completed_at,omitempty"`
	Delivered_at string `json:"delivered_at,omitempty"`
	External_delivery_id string `json:"external_delivery_id,omitempty"`
	Courier_pickup_at string `json:"courier_pickup_at,omitempty"`
	Courier_support_phone_number string `json:"courier_support_phone_number,omitempty"`
	Square_delivery_id string `json:"square_delivery_id,omitempty"`
	Managed_delivery bool `json:"managed_delivery,omitempty"`
	Schedule_type string `json:"schedule_type,omitempty"`
	Ready_at string `json:"ready_at,omitempty"`
	Is_no_contact_delivery bool `json:"is_no_contact_delivery,omitempty"`
	Prep_time_duration string `json:"prep_time_duration,omitempty"`
	Recipient FulfillmentRecipient `json:"recipient,omitempty"`
	Dropoff_notes string `json:"dropoff_notes,omitempty"`
	Canceled_at string `json:"canceled_at,omitempty"`
}

// CreateWebhookSubscription represents the CreateWebhookSubscription schema from the OpenAPI specification
type CreateWebhookSubscription struct {
	Idompotency_key string `json:"idompotency_key,omitempty"`
	Subscription WebhookSubscription `json:"subscription,omitempty"`
}

// FulfillmentPickupDetails represents the FulfillmentPickupDetails schema from the OpenAPI specification
type FulfillmentPickupDetails struct {
	Rejected_at string `json:"rejected_at,omitempty"`
	Pickup_window_duration string `json:"pickup_window_duration,omitempty"`
	Recipient FulfillmentRecipient `json:"recipient,omitempty"`
	Picked_up_at string `json:"picked_up_at,omitempty"`
	Note string `json:"note,omitempty"`
	Prep_time_duration string `json:"prep_time_duration,omitempty"`
	Placed_at string `json:"placed_at,omitempty"`
	Accepted_at string `json:"accepted_at,omitempty"`
	Expired_at string `json:"expired_at,omitempty"`
	Schedule_type string `json:"schedule_type,omitempty"`
	Canceled_at string `json:"canceled_at,omitempty"`
	Pickup_at string `json:"pickup_at,omitempty"`
	Ready_at string `json:"ready_at,omitempty"`
	Is_curbside_pickup bool `json:"is_curbside_pickup,omitempty"`
	Expires_at string `json:"expires_at,omitempty"`
	Cancel_reason string `json:"cancel_reason,omitempty"`
	Curbside_pickup_details map[string]interface{} `json:"curbside_pickup_details,omitempty"`
	Auto_complete_duration string `json:"auto_complete_duration,omitempty"`
}

// Order represents the Order schema from the OpenAPI specification
type Order struct {
	Returns []OrderReturn `json:"returns,omitempty"`
	Closed_at string `json:"closed_at,omitempty"` // The time when the order was closed, in RFC 3339 format.
	Tenders []Tender `json:"tenders,omitempty"`
	Line_items []OrderLineItem `json:"line_items,omitempty"` // The line items included in the order.
	Pricing_options OrderPricingOptions `json:"pricing_options,omitempty"`
	Total_service_charge_money Money `json:"total_service_charge_money,omitempty"`
	Rounding_adjustment OrderRoundingAdjustment `json:"rounding_adjustment,omitempty"`
	Location_id string `json:"location_id"` // The ID of the order's associated location.
	Version int `json:"version,omitempty"` // The version number, which is incremented each time an update is committed to the order.
	Created_at string `json:"created_at,omitempty"` // The time when the order was created, in RFC 3339 format.
	Reference_id string `json:"reference_id,omitempty"` // An optional ID you can associate with the order for your own purposes (such as to associate the order with an entity ID in your own database).
	Total_discount_money Money `json:"total_discount_money,omitempty"`
	Total_money Money `json:"total_money,omitempty"`
	Return_amounts OrderMoneyAmounts `json:"return_amounts,omitempty"`
	Total_tax_money Money `json:"total_tax_money,omitempty"`
	Total_tip_money Money `json:"total_tip_money,omitempty"`
	Fulfillments []Fulfillment `json:"fulfillments,omitempty"`
	Net_amounts OrderMoneyAmounts `json:"net_amounts,omitempty"`
	Id string `json:"id,omitempty"` // The order's unique ID.
	Service_charges []OrderServiceCharge `json:"service_charges,omitempty"`
	Rewards []OrderReward `json:"rewards,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The time when the order was last updated, in RFC 3339 format.
	Customer_id string `json:"customer_id,omitempty"` // The ID of the customer associated with the order.
	Net_amount_due_money Money `json:"net_amount_due_money,omitempty"`
	State string `json:"state,omitempty"` // The current state of the order.
	Ticket_name string `json:"ticket_name,omitempty"` // The name of the order's associated ticket.
	Refunds []OrderRefund `json:"refunds,omitempty"`
	Source map[string]interface{} `json:"source,omitempty"`
}

// OrderCreatedWebhookObject represents the OrderCreatedWebhookObject schema from the OpenAPI specification
type OrderCreatedWebhookObject struct {
	Order_created OrderCreatedWebhookOrder `json:"order_created"` // The order included in the event.
}

// CustomerTaxIds represents the CustomerTaxIds schema from the OpenAPI specification
type CustomerTaxIds struct {
	Eu_vat string `json:"eu_vat,omitempty"` // The EU VAT identification number for the customer. For example, IE3426675K. The ID can contain alphanumeric characters only.
}

// MeasurementUnit represents the MeasurementUnit schema from the OpenAPI specification
type MeasurementUnit struct {
	Type string `json:"type,omitempty"` // The type of the unit.
	Volume_unit string `json:"volume_unit,omitempty"` // Represents a standard volume unit.
	Weight_unit string `json:"weight_unit,omitempty"` // Represents a standard weight unit.
	Area_unit string `json:"area_unit,omitempty"` // Represents a standard area unit.
	Custom_unit map[string]interface{} `json:"custom_unit,omitempty"` // If set, the unit is a custom unit. The name of the custom unit is set in the name field of this object. The precision field is not used for custom units.
	Generic_unit string `json:"generic_unit,omitempty"` // Represents a standard unit of a non-specific quantity.
	Length_unit string `json:"length_unit,omitempty"` // Represents a standard length unit.
	Time_unit string `json:"time_unit,omitempty"` // Represents a standard time unit.
}

// OrderLineItemPricingBlocklists represents the OrderLineItemPricingBlocklists schema from the OpenAPI specification
type OrderLineItemPricingBlocklists struct {
	Blocked_taxes []map[string]interface{} `json:"blocked_taxes,omitempty"`
	Blocked_discounts []map[string]interface{} `json:"blocked_discounts,omitempty"`
}

// OrderMoneyAmounts represents the OrderMoneyAmounts schema from the OpenAPI specification
type OrderMoneyAmounts struct {
	Tip_money Money `json:"tip_money,omitempty"`
	Total_money Money `json:"total_money,omitempty"`
	Discount_money Money `json:"discount_money,omitempty"`
	Service_charge_money Money `json:"service_charge_money,omitempty"`
	Tax_money Money `json:"tax_money,omitempty"`
}
