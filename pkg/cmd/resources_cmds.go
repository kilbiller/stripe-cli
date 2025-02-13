// This file is generated; DO NOT EDIT.

package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	"github.com/stripe/stripe-cli/pkg/cmd/resource"
)

func addAllResourcesCmds(rootCmd *cobra.Command) {
	// Namespace commands
	_ = resource.NewNamespaceCmd(rootCmd, "")
	nsCheckoutCmd := resource.NewNamespaceCmd(rootCmd, "checkout")
	nsIssuingCmd := resource.NewNamespaceCmd(rootCmd, "issuing")
	nsRadarCmd := resource.NewNamespaceCmd(rootCmd, "radar")
	nsReportingCmd := resource.NewNamespaceCmd(rootCmd, "reporting")
	nsTerminalCmd := resource.NewNamespaceCmd(rootCmd, "terminal")

	// Resource commands
	r3dSecureCmd := resource.NewResourceCmd(rootCmd, "3d_secure")
	rAccountLinksCmd := resource.NewResourceCmd(rootCmd, "account_links")
	rAccountsCmd := resource.NewResourceCmd(rootCmd, "accounts")
	rApplePayDomainsCmd := resource.NewResourceCmd(rootCmd, "apple_pay_domains")
	rApplicationFeesCmd := resource.NewResourceCmd(rootCmd, "application_fees")
	rBalanceCmd := resource.NewResourceCmd(rootCmd, "balance")
	rBalanceTransactionsCmd := resource.NewResourceCmd(rootCmd, "balance_transactions")
	rBankAccountsCmd := resource.NewResourceCmd(rootCmd, "bank_accounts")
	rCapabilitiesCmd := resource.NewResourceCmd(rootCmd, "capabilities")
	rCardsCmd := resource.NewResourceCmd(rootCmd, "cards")
	rChargesCmd := resource.NewResourceCmd(rootCmd, "charges")
	rCountrySpecsCmd := resource.NewResourceCmd(rootCmd, "country_specs")
	rCouponsCmd := resource.NewResourceCmd(rootCmd, "coupons")
	rCreditNotesCmd := resource.NewResourceCmd(rootCmd, "credit_notes")
	rCustomerBalanceTransactionsCmd := resource.NewResourceCmd(rootCmd, "customer_balance_transactions")
	rCustomersCmd := resource.NewResourceCmd(rootCmd, "customers")
	rDisputesCmd := resource.NewResourceCmd(rootCmd, "disputes")
	rEphemeralKeysCmd := resource.NewResourceCmd(rootCmd, "ephemeral_keys")
	rEventsCmd := resource.NewResourceCmd(rootCmd, "events")
	rExchangeRatesCmd := resource.NewResourceCmd(rootCmd, "exchange_rates")
	rExternalAccountsCmd := resource.NewResourceCmd(rootCmd, "external_accounts")
	rFeeRefundsCmd := resource.NewResourceCmd(rootCmd, "fee_refunds")
	rFileLinksCmd := resource.NewResourceCmd(rootCmd, "file_links")
	rFilesCmd := resource.NewResourceCmd(rootCmd, "files")
	rInvoiceitemsCmd := resource.NewResourceCmd(rootCmd, "invoiceitems")
	rInvoicesCmd := resource.NewResourceCmd(rootCmd, "invoices")
	rLineItemsCmd := resource.NewResourceCmd(rootCmd, "line_items")
	rLoginLinksCmd := resource.NewResourceCmd(rootCmd, "login_links")
	rOrderReturnsCmd := resource.NewResourceCmd(rootCmd, "order_returns")
	rOrdersCmd := resource.NewResourceCmd(rootCmd, "orders")
	rPaymentIntentsCmd := resource.NewResourceCmd(rootCmd, "payment_intents")
	rPaymentMethodsCmd := resource.NewResourceCmd(rootCmd, "payment_methods")
	rPaymentSourcesCmd := resource.NewResourceCmd(rootCmd, "payment_sources")
	rPayoutsCmd := resource.NewResourceCmd(rootCmd, "payouts")
	rPersonsCmd := resource.NewResourceCmd(rootCmd, "persons")
	rPlansCmd := resource.NewResourceCmd(rootCmd, "plans")
	rProductsCmd := resource.NewResourceCmd(rootCmd, "products")
	rRefundsCmd := resource.NewResourceCmd(rootCmd, "refunds")
	rReviewsCmd := resource.NewResourceCmd(rootCmd, "reviews")
	rScheduledQueryRunsCmd := resource.NewResourceCmd(rootCmd, "scheduled_query_runs")
	rSetupIntentsCmd := resource.NewResourceCmd(rootCmd, "setup_intents")
	rSkusCmd := resource.NewResourceCmd(rootCmd, "skus")
	rSourcesCmd := resource.NewResourceCmd(rootCmd, "sources")
	rSubscriptionItemsCmd := resource.NewResourceCmd(rootCmd, "subscription_items")
	rSubscriptionSchedulesCmd := resource.NewResourceCmd(rootCmd, "subscription_schedules")
	rSubscriptionsCmd := resource.NewResourceCmd(rootCmd, "subscriptions")
	rTaxIdsCmd := resource.NewResourceCmd(rootCmd, "tax_ids")
	rTaxRatesCmd := resource.NewResourceCmd(rootCmd, "tax_rates")
	rTokensCmd := resource.NewResourceCmd(rootCmd, "tokens")
	rTopupsCmd := resource.NewResourceCmd(rootCmd, "topups")
	rTransferReversalsCmd := resource.NewResourceCmd(rootCmd, "transfer_reversals")
	rTransfersCmd := resource.NewResourceCmd(rootCmd, "transfers")
	rUsageRecordsCmd := resource.NewResourceCmd(rootCmd, "usage_records")
	rWebhookEndpointsCmd := resource.NewResourceCmd(rootCmd, "webhook_endpoints")

	rCheckoutSessionsCmd := resource.NewResourceCmd(nsCheckoutCmd.Cmd, "sessions")

	rIssuingAuthorizationsCmd := resource.NewResourceCmd(nsIssuingCmd.Cmd, "authorizations")
	rIssuingCardholdersCmd := resource.NewResourceCmd(nsIssuingCmd.Cmd, "cardholders")
	rIssuingCardsCmd := resource.NewResourceCmd(nsIssuingCmd.Cmd, "cards")
	rIssuingDisputesCmd := resource.NewResourceCmd(nsIssuingCmd.Cmd, "disputes")
	rIssuingTransactionsCmd := resource.NewResourceCmd(nsIssuingCmd.Cmd, "transactions")

	rRadarEarlyFraudWarningsCmd := resource.NewResourceCmd(nsRadarCmd.Cmd, "early_fraud_warnings")
	rRadarValueListItemsCmd := resource.NewResourceCmd(nsRadarCmd.Cmd, "value_list_items")
	rRadarValueListsCmd := resource.NewResourceCmd(nsRadarCmd.Cmd, "value_lists")

	rReportingReportRunsCmd := resource.NewResourceCmd(nsReportingCmd.Cmd, "report_runs")
	rReportingReportTypesCmd := resource.NewResourceCmd(nsReportingCmd.Cmd, "report_types")

	rTerminalConnectionTokensCmd := resource.NewResourceCmd(nsTerminalCmd.Cmd, "connection_tokens")
	rTerminalLocationsCmd := resource.NewResourceCmd(nsTerminalCmd.Cmd, "locations")
	rTerminalReadersCmd := resource.NewResourceCmd(nsTerminalCmd.Cmd, "readers")

	// Operation commands
	resource.NewOperationCmd(r3dSecureCmd.Cmd, "create", "/v1/3d_secure", http.MethodPost, map[string]string{
		"amount":     "integer",
		"card":       "string",
		"currency":   "string",
		"customer":   "string",
		"return_url": "string",
	}, &Config)
	resource.NewOperationCmd(r3dSecureCmd.Cmd, "retrieve", "/v1/3d_secure/{three_d_secure}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rAccountLinksCmd.Cmd, "create", "/v1/account_links", http.MethodPost, map[string]string{
		"account":     "string",
		"collect":     "string",
		"failure_url": "string",
		"success_url": "string",
		"type":        "string",
	}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "capabilities", "/v1/accounts/{account}/capabilities", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "create", "/v1/accounts", http.MethodPost, map[string]string{
		"account_token":    "string",
		"business_type":    "string",
		"country":          "string",
		"default_currency": "string",
		"email":            "string",
		"external_account": "string",
		"type":             "string",
	}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "delete", "/v1/accounts/{account}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "list", "/v1/accounts", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "reject", "/v1/accounts/{account}/reject", http.MethodPost, map[string]string{
		"reason": "string",
	}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "retrieve", "/v1/account", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rAccountsCmd.Cmd, "update", "/v1/accounts/{account}", http.MethodPost, map[string]string{
		"account_token":    "string",
		"business_type":    "string",
		"default_currency": "string",
		"email":            "string",
		"external_account": "string",
	}, &Config)
	resource.NewOperationCmd(rApplePayDomainsCmd.Cmd, "create", "/v1/apple_pay/domains", http.MethodPost, map[string]string{
		"domain_name": "string",
	}, &Config)
	resource.NewOperationCmd(rApplePayDomainsCmd.Cmd, "delete", "/v1/apple_pay/domains/{domain}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rApplePayDomainsCmd.Cmd, "list", "/v1/apple_pay/domains", http.MethodGet, map[string]string{
		"domain_name":    "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rApplePayDomainsCmd.Cmd, "retrieve", "/v1/apple_pay/domains/{domain}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rApplicationFeesCmd.Cmd, "list", "/v1/application_fees", http.MethodGet, map[string]string{
		"charge":         "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rApplicationFeesCmd.Cmd, "retrieve", "/v1/application_fees/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rBalanceCmd.Cmd, "retrieve", "/v1/balance", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rBalanceTransactionsCmd.Cmd, "list", "/v1/balance_transactions", http.MethodGet, map[string]string{
		"currency":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"payout":         "string",
		"source":         "string",
		"starting_after": "string",
		"type":           "string",
	}, &Config)
	resource.NewOperationCmd(rBalanceTransactionsCmd.Cmd, "retrieve", "/v1/balance_transactions/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rBankAccountsCmd.Cmd, "delete", "/v1/customers/{customer}/sources/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rBankAccountsCmd.Cmd, "update", "/v1/customers/{customer}/sources/{id}", http.MethodPost, map[string]string{
		"account_holder_name": "string",
		"account_holder_type": "string",
		"address_city":        "string",
		"address_country":     "string",
		"address_line1":       "string",
		"address_line2":       "string",
		"address_state":       "string",
		"address_zip":         "string",
		"exp_month":           "string",
		"exp_year":            "string",
		"name":                "string",
	}, &Config)
	resource.NewOperationCmd(rBankAccountsCmd.Cmd, "verify", "/v1/customers/{customer}/sources/{id}/verify", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rCapabilitiesCmd.Cmd, "list", "/v1/accounts/{account}/capabilities", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCapabilitiesCmd.Cmd, "retrieve", "/v1/accounts/{account}/capabilities/{capability}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCapabilitiesCmd.Cmd, "update", "/v1/accounts/{account}/capabilities/{capability}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rCardsCmd.Cmd, "delete", "/v1/customers/{customer}/sources/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rCardsCmd.Cmd, "update", "/v1/customers/{customer}/sources/{id}", http.MethodPost, map[string]string{
		"account_holder_name": "string",
		"account_holder_type": "string",
		"address_city":        "string",
		"address_country":     "string",
		"address_line1":       "string",
		"address_line2":       "string",
		"address_state":       "string",
		"address_zip":         "string",
		"exp_month":           "string",
		"exp_year":            "string",
		"name":                "string",
	}, &Config)
	resource.NewOperationCmd(rChargesCmd.Cmd, "capture", "/v1/charges/{charge}/capture", http.MethodPost, map[string]string{
		"amount":                      "integer",
		"application_fee":             "integer",
		"application_fee_amount":      "integer",
		"receipt_email":               "string",
		"statement_descriptor":        "string",
		"statement_descriptor_suffix": "string",
		"transfer_group":              "string",
	}, &Config)
	resource.NewOperationCmd(rChargesCmd.Cmd, "create", "/v1/charges", http.MethodPost, map[string]string{
		"amount":                      "integer",
		"application_fee":             "integer",
		"application_fee_amount":      "integer",
		"currency":                    "string",
		"customer":                    "string",
		"description":                 "string",
		"on_behalf_of":                "string",
		"receipt_email":               "string",
		"source":                      "string",
		"statement_descriptor":        "string",
		"statement_descriptor_suffix": "string",
		"transfer_group":              "string",
	}, &Config)
	resource.NewOperationCmd(rChargesCmd.Cmd, "list", "/v1/charges", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"payment_intent": "string",
		"starting_after": "string",
		"transfer_group": "string",
	}, &Config)
	resource.NewOperationCmd(rChargesCmd.Cmd, "retrieve", "/v1/charges/{charge}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rChargesCmd.Cmd, "update", "/v1/charges/{charge}", http.MethodPost, map[string]string{
		"customer":       "string",
		"description":    "string",
		"receipt_email":  "string",
		"transfer_group": "string",
	}, &Config)
	resource.NewOperationCmd(rCountrySpecsCmd.Cmd, "list", "/v1/country_specs", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rCountrySpecsCmd.Cmd, "retrieve", "/v1/country_specs/{country}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCouponsCmd.Cmd, "create", "/v1/coupons", http.MethodPost, map[string]string{
		"amount_off":         "integer",
		"currency":           "string",
		"duration":           "string",
		"duration_in_months": "integer",
		"id":                 "string",
		"max_redemptions":    "integer",
		"name":               "string",
		"redeem_by":          "integer",
	}, &Config)
	resource.NewOperationCmd(rCouponsCmd.Cmd, "delete", "/v1/coupons/{coupon}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rCouponsCmd.Cmd, "list", "/v1/coupons", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rCouponsCmd.Cmd, "retrieve", "/v1/coupons/{coupon}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCouponsCmd.Cmd, "update", "/v1/coupons/{coupon}", http.MethodPost, map[string]string{
		"name": "string",
	}, &Config)
	resource.NewOperationCmd(rCreditNotesCmd.Cmd, "create", "/v1/credit_notes", http.MethodPost, map[string]string{
		"amount":        "integer",
		"credit_amount": "integer",
		"invoice":       "string",
		"memo":          "string",
		"reason":        "string",
		"refund":        "string",
		"refund_amount": "integer",
	}, &Config)
	resource.NewOperationCmd(rCreditNotesCmd.Cmd, "list", "/v1/credit_notes", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"invoice":        "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rCreditNotesCmd.Cmd, "retrieve", "/v1/credit_notes/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCreditNotesCmd.Cmd, "update", "/v1/credit_notes/{id}", http.MethodPost, map[string]string{
		"memo": "string",
	}, &Config)
	resource.NewOperationCmd(rCreditNotesCmd.Cmd, "void_credit_note", "/v1/credit_notes/{id}/void", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rCustomerBalanceTransactionsCmd.Cmd, "create", "/v1/customers/{customer}/balance_transactions", http.MethodPost, map[string]string{
		"amount":      "integer",
		"currency":    "string",
		"description": "string",
	}, &Config)
	resource.NewOperationCmd(rCustomerBalanceTransactionsCmd.Cmd, "list", "/v1/customers/{customer}/balance_transactions", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rCustomerBalanceTransactionsCmd.Cmd, "retrieve", "/v1/customers/{customer}/balance_transactions/{transaction}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCustomerBalanceTransactionsCmd.Cmd, "update", "/v1/customers/{customer}/balance_transactions/{transaction}", http.MethodPost, map[string]string{
		"description": "string",
	}, &Config)
	resource.NewOperationCmd(rCustomersCmd.Cmd, "create", "/v1/customers", http.MethodPost, map[string]string{
		"balance":        "integer",
		"coupon":         "string",
		"description":    "string",
		"email":          "string",
		"invoice_prefix": "string",
		"name":           "string",
		"payment_method": "string",
		"phone":          "string",
		"source":         "string",
		"tax_exempt":     "string",
	}, &Config)
	resource.NewOperationCmd(rCustomersCmd.Cmd, "delete", "/v1/customers/{customer}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rCustomersCmd.Cmd, "delete_discount", "/v1/customers/{customer}/discount", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rCustomersCmd.Cmd, "list", "/v1/customers", http.MethodGet, map[string]string{
		"email":          "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rCustomersCmd.Cmd, "retrieve", "/v1/customers/{customer}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rCustomersCmd.Cmd, "update", "/v1/customers/{customer}", http.MethodPost, map[string]string{
		"balance":        "integer",
		"coupon":         "string",
		"default_source": "string",
		"description":    "string",
		"email":          "string",
		"invoice_prefix": "string",
		"name":           "string",
		"phone":          "string",
		"source":         "string",
		"tax_exempt":     "string",
	}, &Config)
	resource.NewOperationCmd(rDisputesCmd.Cmd, "close", "/v1/disputes/{dispute}/close", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rDisputesCmd.Cmd, "list", "/v1/disputes", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rDisputesCmd.Cmd, "retrieve", "/v1/disputes/{dispute}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rDisputesCmd.Cmd, "update", "/v1/disputes/{dispute}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rEphemeralKeysCmd.Cmd, "create", "/v1/ephemeral_keys", http.MethodPost, map[string]string{
		"customer":     "string",
		"issuing_card": "string",
	}, &Config)
	resource.NewOperationCmd(rEphemeralKeysCmd.Cmd, "delete", "/v1/ephemeral_keys/{key}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rEventsCmd.Cmd, "list", "/v1/events", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"type":           "string",
	}, &Config)
	resource.NewOperationCmd(rEventsCmd.Cmd, "retrieve", "/v1/events/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rExchangeRatesCmd.Cmd, "list", "/v1/exchange_rates", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rExchangeRatesCmd.Cmd, "retrieve", "/v1/exchange_rates/{currency}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rExternalAccountsCmd.Cmd, "create", "/v1/accounts/{account}/external_accounts", http.MethodPost, map[string]string{
		"external_account": "string",
	}, &Config)
	resource.NewOperationCmd(rExternalAccountsCmd.Cmd, "delete", "/v1/accounts/{account}/external_accounts/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rExternalAccountsCmd.Cmd, "list", "/v1/accounts/{account}/external_accounts", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rExternalAccountsCmd.Cmd, "retrieve", "/v1/accounts/{account}/external_accounts/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rExternalAccountsCmd.Cmd, "update", "/v1/accounts/{account}/external_accounts/{id}", http.MethodPost, map[string]string{
		"account_holder_name": "string",
		"account_holder_type": "string",
		"address_city":        "string",
		"address_country":     "string",
		"address_line1":       "string",
		"address_line2":       "string",
		"address_state":       "string",
		"address_zip":         "string",
		"exp_month":           "string",
		"exp_year":            "string",
		"name":                "string",
	}, &Config)
	resource.NewOperationCmd(rFeeRefundsCmd.Cmd, "create", "/v1/application_fees/{id}/refunds", http.MethodPost, map[string]string{
		"amount": "integer",
	}, &Config)
	resource.NewOperationCmd(rFeeRefundsCmd.Cmd, "list", "/v1/application_fees/{id}/refunds", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rFeeRefundsCmd.Cmd, "retrieve", "/v1/application_fees/{fee}/refunds/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rFeeRefundsCmd.Cmd, "update", "/v1/application_fees/{fee}/refunds/{id}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rFileLinksCmd.Cmd, "create", "/v1/file_links", http.MethodPost, map[string]string{
		"expires_at": "integer",
		"file":       "string",
	}, &Config)
	resource.NewOperationCmd(rFileLinksCmd.Cmd, "list", "/v1/file_links", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"file":           "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rFileLinksCmd.Cmd, "retrieve", "/v1/file_links/{link}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rFileLinksCmd.Cmd, "update", "/v1/file_links/{link}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rFilesCmd.Cmd, "create", "/v1/files", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rFilesCmd.Cmd, "list", "/v1/files", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"purpose":        "string",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rFilesCmd.Cmd, "retrieve", "/v1/files/{file}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoiceitemsCmd.Cmd, "create", "/v1/invoiceitems", http.MethodPost, map[string]string{
		"amount":              "integer",
		"currency":            "string",
		"customer":            "string",
		"description":         "string",
		"invoice":             "string",
		"quantity":            "integer",
		"subscription":        "string",
		"unit_amount":         "integer",
		"unit_amount_decimal": "string",
	}, &Config)
	resource.NewOperationCmd(rInvoiceitemsCmd.Cmd, "delete", "/v1/invoiceitems/{invoiceitem}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoiceitemsCmd.Cmd, "list", "/v1/invoiceitems", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"invoice":        "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rInvoiceitemsCmd.Cmd, "retrieve", "/v1/invoiceitems/{invoiceitem}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoiceitemsCmd.Cmd, "update", "/v1/invoiceitems/{invoiceitem}", http.MethodPost, map[string]string{
		"amount":              "integer",
		"description":         "string",
		"quantity":            "integer",
		"unit_amount":         "integer",
		"unit_amount_decimal": "string",
	}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "create", "/v1/invoices", http.MethodPost, map[string]string{
		"application_fee_amount": "integer",
		"collection_method":      "string",
		"customer":               "string",
		"days_until_due":         "integer",
		"default_payment_method": "string",
		"default_source":         "string",
		"description":            "string",
		"due_date":               "integer",
		"footer":                 "string",
		"statement_descriptor":   "string",
		"subscription":           "string",
	}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "delete", "/v1/invoices/{invoice}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "finalize_invoice", "/v1/invoices/{invoice}/finalize", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "list", "/v1/invoices", http.MethodGet, map[string]string{
		"collection_method": "string",
		"customer":          "string",
		"ending_before":     "string",
		"limit":             "integer",
		"starting_after":    "string",
		"status":            "string",
		"subscription":      "string",
	}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "mark_uncollectible", "/v1/invoices/{invoice}/mark_uncollectible", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "pay", "/v1/invoices/{invoice}/pay", http.MethodPost, map[string]string{
		"payment_method": "string",
		"source":         "string",
	}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "retrieve", "/v1/invoices/{invoice}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "send_invoice", "/v1/invoices/{invoice}/send", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "upcoming", "/v1/invoices/upcoming", http.MethodGet, map[string]string{
		"coupon":                      "string",
		"customer":                    "string",
		"schedule":                    "string",
		"subscription":                "string",
		"subscription_proration_date": "integer",
		"subscription_start_date":     "integer",
	}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "update", "/v1/invoices/{invoice}", http.MethodPost, map[string]string{
		"application_fee_amount": "integer",
		"collection_method":      "string",
		"days_until_due":         "integer",
		"default_payment_method": "string",
		"default_source":         "string",
		"description":            "string",
		"due_date":               "integer",
		"footer":                 "string",
		"statement_descriptor":   "string",
	}, &Config)
	resource.NewOperationCmd(rInvoicesCmd.Cmd, "void_invoice", "/v1/invoices/{invoice}/void", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rLineItemsCmd.Cmd, "list", "/v1/invoices/{invoice}/lines", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rLoginLinksCmd.Cmd, "create", "/v1/accounts/{account}/login_links", http.MethodPost, map[string]string{
		"redirect_url": "string",
	}, &Config)
	resource.NewOperationCmd(rOrderReturnsCmd.Cmd, "list", "/v1/order_returns", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"order":          "string",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rOrderReturnsCmd.Cmd, "retrieve", "/v1/order_returns/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rOrdersCmd.Cmd, "create", "/v1/orders", http.MethodPost, map[string]string{
		"coupon":   "string",
		"currency": "string",
		"customer": "string",
		"email":    "string",
	}, &Config)
	resource.NewOperationCmd(rOrdersCmd.Cmd, "list", "/v1/orders", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"status":         "string",
	}, &Config)
	resource.NewOperationCmd(rOrdersCmd.Cmd, "pay", "/v1/orders/{id}/pay", http.MethodPost, map[string]string{
		"application_fee": "integer",
		"customer":        "string",
		"email":           "string",
		"source":          "string",
	}, &Config)
	resource.NewOperationCmd(rOrdersCmd.Cmd, "retrieve", "/v1/orders/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rOrdersCmd.Cmd, "return_order", "/v1/orders/{id}/returns", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rOrdersCmd.Cmd, "update", "/v1/orders/{id}", http.MethodPost, map[string]string{
		"coupon":                   "string",
		"selected_shipping_method": "string",
		"status":                   "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "cancel", "/v1/payment_intents/{intent}/cancel", http.MethodPost, map[string]string{
		"cancellation_reason": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "capture", "/v1/payment_intents/{intent}/capture", http.MethodPost, map[string]string{
		"amount_to_capture":           "integer",
		"application_fee_amount":      "integer",
		"statement_descriptor":        "string",
		"statement_descriptor_suffix": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "confirm", "/v1/payment_intents/{intent}/confirm", http.MethodPost, map[string]string{
		"payment_method":     "string",
		"return_url":         "string",
		"setup_future_usage": "string",
		"source":             "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "create", "/v1/payment_intents", http.MethodPost, map[string]string{
		"amount":                      "integer",
		"application_fee_amount":      "integer",
		"capture_method":              "string",
		"confirmation_method":         "string",
		"currency":                    "string",
		"customer":                    "string",
		"description":                 "string",
		"on_behalf_of":                "string",
		"payment_method":              "string",
		"receipt_email":               "string",
		"return_url":                  "string",
		"setup_future_usage":          "string",
		"source":                      "string",
		"statement_descriptor":        "string",
		"statement_descriptor_suffix": "string",
		"transfer_group":              "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "list", "/v1/payment_intents", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "retrieve", "/v1/payment_intents/{intent}", http.MethodGet, map[string]string{
		"client_secret": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentIntentsCmd.Cmd, "update", "/v1/payment_intents/{intent}", http.MethodPost, map[string]string{
		"amount":                      "integer",
		"currency":                    "string",
		"customer":                    "string",
		"description":                 "string",
		"payment_method":              "string",
		"setup_future_usage":          "string",
		"source":                      "string",
		"statement_descriptor":        "string",
		"statement_descriptor_suffix": "string",
		"transfer_group":              "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentMethodsCmd.Cmd, "attach", "/v1/payment_methods/{payment_method}/attach", http.MethodPost, map[string]string{
		"customer": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentMethodsCmd.Cmd, "create", "/v1/payment_methods", http.MethodPost, map[string]string{
		"customer":       "string",
		"payment_method": "string",
		"type":           "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentMethodsCmd.Cmd, "detach", "/v1/payment_methods/{payment_method}/detach", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rPaymentMethodsCmd.Cmd, "list", "/v1/payment_methods", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"type":           "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentMethodsCmd.Cmd, "retrieve", "/v1/payment_methods/{payment_method}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rPaymentMethodsCmd.Cmd, "update", "/v1/payment_methods/{payment_method}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rPaymentSourcesCmd.Cmd, "create", "/v1/customers/{customer}/sources", http.MethodPost, map[string]string{
		"source": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentSourcesCmd.Cmd, "list", "/v1/customers/{customer}/sources", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"object":         "string",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rPaymentSourcesCmd.Cmd, "retrieve", "/v1/customers/{customer}/sources/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rPayoutsCmd.Cmd, "cancel", "/v1/payouts/{payout}/cancel", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rPayoutsCmd.Cmd, "create", "/v1/payouts", http.MethodPost, map[string]string{
		"amount":               "integer",
		"currency":             "string",
		"description":          "string",
		"destination":          "string",
		"method":               "string",
		"source_type":          "string",
		"statement_descriptor": "string",
	}, &Config)
	resource.NewOperationCmd(rPayoutsCmd.Cmd, "list", "/v1/payouts", http.MethodGet, map[string]string{
		"destination":    "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"status":         "string",
	}, &Config)
	resource.NewOperationCmd(rPayoutsCmd.Cmd, "retrieve", "/v1/payouts/{payout}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rPayoutsCmd.Cmd, "update", "/v1/payouts/{payout}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rPersonsCmd.Cmd, "create", "/v1/accounts/{account}/persons", http.MethodPost, map[string]string{
		"email":            "string",
		"first_name":       "string",
		"first_name_kana":  "string",
		"first_name_kanji": "string",
		"gender":           "string",
		"id_number":        "string",
		"last_name":        "string",
		"last_name_kana":   "string",
		"last_name_kanji":  "string",
		"maiden_name":      "string",
		"person_token":     "string",
		"phone":            "string",
		"ssn_last_4":       "string",
	}, &Config)
	resource.NewOperationCmd(rPersonsCmd.Cmd, "delete", "/v1/accounts/{account}/persons/{person}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rPersonsCmd.Cmd, "list", "/v1/accounts/{account}/persons", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rPersonsCmd.Cmd, "retrieve", "/v1/accounts/{account}/persons/{person}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rPersonsCmd.Cmd, "update", "/v1/accounts/{account}/persons/{person}", http.MethodPost, map[string]string{
		"email":            "string",
		"first_name":       "string",
		"first_name_kana":  "string",
		"first_name_kanji": "string",
		"gender":           "string",
		"id_number":        "string",
		"last_name":        "string",
		"last_name_kana":   "string",
		"last_name_kanji":  "string",
		"maiden_name":      "string",
		"person_token":     "string",
		"phone":            "string",
		"ssn_last_4":       "string",
	}, &Config)
	resource.NewOperationCmd(rPlansCmd.Cmd, "create", "/v1/plans", http.MethodPost, map[string]string{
		"aggregate_usage":   "string",
		"amount":            "integer",
		"amount_decimal":    "string",
		"billing_scheme":    "string",
		"currency":          "string",
		"id":                "string",
		"interval":          "string",
		"interval_count":    "integer",
		"nickname":          "string",
		"tiers_mode":        "string",
		"trial_period_days": "integer",
		"usage_type":        "string",
	}, &Config)
	resource.NewOperationCmd(rPlansCmd.Cmd, "delete", "/v1/plans/{plan}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rPlansCmd.Cmd, "list", "/v1/plans", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"product":        "string",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rPlansCmd.Cmd, "retrieve", "/v1/plans/{plan}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rPlansCmd.Cmd, "update", "/v1/plans/{plan}", http.MethodPost, map[string]string{
		"nickname":          "string",
		"product":           "string",
		"trial_period_days": "integer",
	}, &Config)
	resource.NewOperationCmd(rProductsCmd.Cmd, "create", "/v1/products", http.MethodPost, map[string]string{
		"caption":              "string",
		"description":          "string",
		"id":                   "string",
		"name":                 "string",
		"statement_descriptor": "string",
		"type":                 "string",
		"unit_label":           "string",
		"url":                  "string",
	}, &Config)
	resource.NewOperationCmd(rProductsCmd.Cmd, "delete", "/v1/products/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rProductsCmd.Cmd, "list", "/v1/products", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"type":           "string",
		"url":            "string",
	}, &Config)
	resource.NewOperationCmd(rProductsCmd.Cmd, "retrieve", "/v1/products/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rProductsCmd.Cmd, "update", "/v1/products/{id}", http.MethodPost, map[string]string{
		"caption":              "string",
		"description":          "string",
		"name":                 "string",
		"statement_descriptor": "string",
		"unit_label":           "string",
		"url":                  "string",
	}, &Config)
	resource.NewOperationCmd(rRefundsCmd.Cmd, "create", "/v1/refunds", http.MethodPost, map[string]string{
		"amount": "integer",
		"charge": "string",
		"reason": "string",
	}, &Config)
	resource.NewOperationCmd(rRefundsCmd.Cmd, "list", "/v1/refunds", http.MethodGet, map[string]string{
		"charge":         "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rRefundsCmd.Cmd, "retrieve", "/v1/refunds/{refund}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rRefundsCmd.Cmd, "update", "/v1/refunds/{refund}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rReviewsCmd.Cmd, "approve", "/v1/reviews/{review}/approve", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rReviewsCmd.Cmd, "list", "/v1/reviews", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rReviewsCmd.Cmd, "retrieve", "/v1/reviews/{review}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rScheduledQueryRunsCmd.Cmd, "list", "/v1/sigma/scheduled_query_runs", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rScheduledQueryRunsCmd.Cmd, "retrieve", "/v1/sigma/scheduled_query_runs/{scheduled_query_run}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rSetupIntentsCmd.Cmd, "cancel", "/v1/setup_intents/{intent}/cancel", http.MethodPost, map[string]string{
		"cancellation_reason": "string",
	}, &Config)
	resource.NewOperationCmd(rSetupIntentsCmd.Cmd, "confirm", "/v1/setup_intents/{intent}/confirm", http.MethodPost, map[string]string{
		"payment_method": "string",
		"return_url":     "string",
	}, &Config)
	resource.NewOperationCmd(rSetupIntentsCmd.Cmd, "create", "/v1/setup_intents", http.MethodPost, map[string]string{
		"customer":       "string",
		"description":    "string",
		"on_behalf_of":   "string",
		"payment_method": "string",
		"return_url":     "string",
	}, &Config)
	resource.NewOperationCmd(rSetupIntentsCmd.Cmd, "list", "/v1/setup_intents", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"payment_method": "string",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rSetupIntentsCmd.Cmd, "retrieve", "/v1/setup_intents/{intent}", http.MethodGet, map[string]string{
		"client_secret": "string",
	}, &Config)
	resource.NewOperationCmd(rSetupIntentsCmd.Cmd, "update", "/v1/setup_intents/{intent}", http.MethodPost, map[string]string{
		"customer":       "string",
		"description":    "string",
		"payment_method": "string",
	}, &Config)
	resource.NewOperationCmd(rSkusCmd.Cmd, "create", "/v1/skus", http.MethodPost, map[string]string{
		"currency": "string",
		"id":       "string",
		"image":    "string",
		"price":    "integer",
		"product":  "string",
	}, &Config)
	resource.NewOperationCmd(rSkusCmd.Cmd, "delete", "/v1/skus/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rSkusCmd.Cmd, "list", "/v1/skus", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"product":        "string",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rSkusCmd.Cmd, "retrieve", "/v1/skus/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rSkusCmd.Cmd, "update", "/v1/skus/{id}", http.MethodPost, map[string]string{
		"currency": "string",
		"image":    "string",
		"price":    "integer",
		"product":  "string",
	}, &Config)
	resource.NewOperationCmd(rSourcesCmd.Cmd, "create", "/v1/sources", http.MethodPost, map[string]string{
		"amount":               "integer",
		"currency":             "string",
		"customer":             "string",
		"flow":                 "string",
		"original_source":      "string",
		"statement_descriptor": "string",
		"token":                "string",
		"type":                 "string",
		"usage":                "string",
	}, &Config)
	resource.NewOperationCmd(rSourcesCmd.Cmd, "detach", "/v1/customers/{customer}/sources/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rSourcesCmd.Cmd, "retrieve", "/v1/sources/{source}", http.MethodGet, map[string]string{
		"client_secret": "string",
	}, &Config)
	resource.NewOperationCmd(rSourcesCmd.Cmd, "source_transactions", "/v1/sources/{source}/source_transactions", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rSourcesCmd.Cmd, "update", "/v1/sources/{source}", http.MethodPost, map[string]string{
		"amount": "integer",
	}, &Config)
	resource.NewOperationCmd(rSourcesCmd.Cmd, "verify", "/v1/sources/{source}/verify", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionItemsCmd.Cmd, "create", "/v1/subscription_items", http.MethodPost, map[string]string{
		"payment_behavior": "string",
		"plan":             "string",
		"proration_date":   "integer",
		"quantity":         "integer",
		"subscription":     "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionItemsCmd.Cmd, "delete", "/v1/subscription_items/{item}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionItemsCmd.Cmd, "list", "/v1/subscription_items", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"subscription":   "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionItemsCmd.Cmd, "retrieve", "/v1/subscription_items/{item}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionItemsCmd.Cmd, "update", "/v1/subscription_items/{item}", http.MethodPost, map[string]string{
		"payment_behavior": "string",
		"plan":             "string",
		"proration_date":   "integer",
		"quantity":         "integer",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionItemsCmd.Cmd, "usage_record_summaries", "/v1/subscription_items/{subscription_item}/usage_record_summaries", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionSchedulesCmd.Cmd, "cancel", "/v1/subscription_schedules/{schedule}/cancel", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionSchedulesCmd.Cmd, "create", "/v1/subscription_schedules", http.MethodPost, map[string]string{
		"collection_method":      "string",
		"customer":               "string",
		"default_payment_method": "string",
		"default_source":         "string",
		"end_behavior":           "string",
		"from_subscription":      "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionSchedulesCmd.Cmd, "list", "/v1/subscription_schedules", http.MethodGet, map[string]string{
		"customer":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionSchedulesCmd.Cmd, "release", "/v1/subscription_schedules/{schedule}/release", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionSchedulesCmd.Cmd, "retrieve", "/v1/subscription_schedules/{schedule}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionSchedulesCmd.Cmd, "update", "/v1/subscription_schedules/{schedule}", http.MethodPost, map[string]string{
		"collection_method":      "string",
		"default_payment_method": "string",
		"default_source":         "string",
		"end_behavior":           "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionsCmd.Cmd, "cancel", "/v1/subscriptions/{subscription_exposed_id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionsCmd.Cmd, "create", "/v1/subscriptions", http.MethodPost, map[string]string{
		"backdate_start_date":    "integer",
		"billing_cycle_anchor":   "integer",
		"cancel_at":              "integer",
		"collection_method":      "string",
		"coupon":                 "string",
		"customer":               "string",
		"days_until_due":         "integer",
		"default_payment_method": "string",
		"default_source":         "string",
		"payment_behavior":       "string",
		"trial_period_days":      "integer",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionsCmd.Cmd, "delete_discount", "/v1/subscriptions/{subscription_exposed_id}/discount", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionsCmd.Cmd, "list", "/v1/subscriptions", http.MethodGet, map[string]string{
		"collection_method": "string",
		"customer":          "string",
		"ending_before":     "string",
		"limit":             "integer",
		"plan":              "string",
		"starting_after":    "string",
		"status":            "string",
	}, &Config)
	resource.NewOperationCmd(rSubscriptionsCmd.Cmd, "retrieve", "/v1/subscriptions/{subscription_exposed_id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rSubscriptionsCmd.Cmd, "update", "/v1/subscriptions/{subscription_exposed_id}", http.MethodPost, map[string]string{
		"billing_cycle_anchor":   "string",
		"collection_method":      "string",
		"coupon":                 "string",
		"days_until_due":         "integer",
		"default_payment_method": "string",
		"default_source":         "string",
		"payment_behavior":       "string",
		"proration_date":         "integer",
	}, &Config)
	resource.NewOperationCmd(rTaxIdsCmd.Cmd, "create", "/v1/customers/{customer}/tax_ids", http.MethodPost, map[string]string{
		"type":  "string",
		"value": "string",
	}, &Config)
	resource.NewOperationCmd(rTaxIdsCmd.Cmd, "delete", "/v1/customers/{customer}/tax_ids/{id}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rTaxIdsCmd.Cmd, "list", "/v1/customers/{customer}/tax_ids", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rTaxIdsCmd.Cmd, "retrieve", "/v1/customers/{customer}/tax_ids/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTaxRatesCmd.Cmd, "create", "/v1/tax_rates", http.MethodPost, map[string]string{
		"description":  "string",
		"display_name": "string",
		"jurisdiction": "string",
	}, &Config)
	resource.NewOperationCmd(rTaxRatesCmd.Cmd, "list", "/v1/tax_rates", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rTaxRatesCmd.Cmd, "retrieve", "/v1/tax_rates/{tax_rate}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTaxRatesCmd.Cmd, "update", "/v1/tax_rates/{tax_rate}", http.MethodPost, map[string]string{
		"description":  "string",
		"display_name": "string",
		"jurisdiction": "string",
	}, &Config)
	resource.NewOperationCmd(rTokensCmd.Cmd, "create", "/v1/tokens", http.MethodPost, map[string]string{
		"customer": "string",
	}, &Config)
	resource.NewOperationCmd(rTokensCmd.Cmd, "retrieve", "/v1/tokens/{token}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTopupsCmd.Cmd, "cancel", "/v1/topups/{topup}/cancel", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rTopupsCmd.Cmd, "create", "/v1/topups", http.MethodPost, map[string]string{
		"amount":               "integer",
		"currency":             "string",
		"description":          "string",
		"source":               "string",
		"statement_descriptor": "string",
		"transfer_group":       "string",
	}, &Config)
	resource.NewOperationCmd(rTopupsCmd.Cmd, "list", "/v1/topups", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"status":         "string",
	}, &Config)
	resource.NewOperationCmd(rTopupsCmd.Cmd, "retrieve", "/v1/topups/{topup}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTopupsCmd.Cmd, "update", "/v1/topups/{topup}", http.MethodPost, map[string]string{
		"description": "string",
	}, &Config)
	resource.NewOperationCmd(rTransferReversalsCmd.Cmd, "create", "/v1/transfers/{id}/reversals", http.MethodPost, map[string]string{
		"amount":      "integer",
		"description": "string",
	}, &Config)
	resource.NewOperationCmd(rTransferReversalsCmd.Cmd, "list", "/v1/transfers/{id}/reversals", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rTransferReversalsCmd.Cmd, "retrieve", "/v1/transfers/{transfer}/reversals/{id}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTransferReversalsCmd.Cmd, "update", "/v1/transfers/{transfer}/reversals/{id}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rTransfersCmd.Cmd, "create", "/v1/transfers", http.MethodPost, map[string]string{
		"amount":             "integer",
		"currency":           "string",
		"description":        "string",
		"destination":        "string",
		"source_transaction": "string",
		"source_type":        "string",
		"transfer_group":     "string",
	}, &Config)
	resource.NewOperationCmd(rTransfersCmd.Cmd, "list", "/v1/transfers", http.MethodGet, map[string]string{
		"destination":    "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"transfer_group": "string",
	}, &Config)
	resource.NewOperationCmd(rTransfersCmd.Cmd, "retrieve", "/v1/transfers/{transfer}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTransfersCmd.Cmd, "update", "/v1/transfers/{transfer}", http.MethodPost, map[string]string{
		"description": "string",
	}, &Config)
	resource.NewOperationCmd(rUsageRecordsCmd.Cmd, "create", "/v1/subscription_items/{subscription_item}/usage_records", http.MethodPost, map[string]string{
		"action":    "string",
		"quantity":  "integer",
		"timestamp": "integer",
	}, &Config)
	resource.NewOperationCmd(rWebhookEndpointsCmd.Cmd, "create", "/v1/webhook_endpoints", http.MethodPost, map[string]string{
		"api_version": "string",
		"url":         "string",
	}, &Config)
	resource.NewOperationCmd(rWebhookEndpointsCmd.Cmd, "delete", "/v1/webhook_endpoints/{webhook_endpoint}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rWebhookEndpointsCmd.Cmd, "list", "/v1/webhook_endpoints", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rWebhookEndpointsCmd.Cmd, "retrieve", "/v1/webhook_endpoints/{webhook_endpoint}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rWebhookEndpointsCmd.Cmd, "update", "/v1/webhook_endpoints/{webhook_endpoint}", http.MethodPost, map[string]string{
		"url": "string",
	}, &Config)
	resource.NewOperationCmd(rCheckoutSessionsCmd.Cmd, "create", "/v1/checkout/sessions", http.MethodPost, map[string]string{
		"billing_address_collection": "string",
		"cancel_url":                 "string",
		"client_reference_id":        "string",
		"customer":                   "string",
		"customer_email":             "string",
		"locale":                     "string",
		"mode":                       "string",
		"submit_type":                "string",
		"success_url":                "string",
	}, &Config)
	resource.NewOperationCmd(rCheckoutSessionsCmd.Cmd, "retrieve", "/v1/checkout/sessions/{session}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingAuthorizationsCmd.Cmd, "approve", "/v1/issuing/authorizations/{authorization}/approve", http.MethodPost, map[string]string{
		"held_amount": "integer",
	}, &Config)
	resource.NewOperationCmd(rIssuingAuthorizationsCmd.Cmd, "decline", "/v1/issuing/authorizations/{authorization}/decline", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingAuthorizationsCmd.Cmd, "list", "/v1/issuing/authorizations", http.MethodGet, map[string]string{
		"card":           "string",
		"cardholder":     "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"status":         "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingAuthorizationsCmd.Cmd, "retrieve", "/v1/issuing/authorizations/{authorization}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingAuthorizationsCmd.Cmd, "update", "/v1/issuing/authorizations/{authorization}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingCardholdersCmd.Cmd, "create", "/v1/issuing/cardholders", http.MethodPost, map[string]string{
		"email":        "string",
		"name":         "string",
		"phone_number": "string",
		"status":       "string",
		"type":         "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingCardholdersCmd.Cmd, "list", "/v1/issuing/cardholders", http.MethodGet, map[string]string{
		"email":          "string",
		"ending_before":  "string",
		"limit":          "integer",
		"phone_number":   "string",
		"starting_after": "string",
		"status":         "string",
		"type":           "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingCardholdersCmd.Cmd, "retrieve", "/v1/issuing/cardholders/{cardholder}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingCardholdersCmd.Cmd, "update", "/v1/issuing/cardholders/{cardholder}", http.MethodPost, map[string]string{
		"email":        "string",
		"phone_number": "string",
		"status":       "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingCardsCmd.Cmd, "create", "/v1/issuing/cards", http.MethodPost, map[string]string{
		"cardholder":         "string",
		"currency":           "string",
		"replacement_for":    "string",
		"replacement_reason": "string",
		"status":             "string",
		"type":               "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingCardsCmd.Cmd, "details", "/v1/issuing/cards/{card}/details", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingCardsCmd.Cmd, "list", "/v1/issuing/cards", http.MethodGet, map[string]string{
		"cardholder":     "string",
		"ending_before":  "string",
		"exp_month":      "integer",
		"exp_year":       "integer",
		"last4":          "string",
		"limit":          "integer",
		"name":           "string",
		"source":         "string",
		"starting_after": "string",
		"status":         "string",
		"type":           "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingCardsCmd.Cmd, "retrieve", "/v1/issuing/cards/{card}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingCardsCmd.Cmd, "update", "/v1/issuing/cards/{card}", http.MethodPost, map[string]string{
		"cardholder": "string",
		"status":     "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingDisputesCmd.Cmd, "create", "/v1/issuing/disputes", http.MethodPost, map[string]string{
		"amount":               "integer",
		"disputed_transaction": "string",
		"reason":               "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingDisputesCmd.Cmd, "list", "/v1/issuing/disputes", http.MethodGet, map[string]string{
		"disputed_transaction": "string",
		"ending_before":        "string",
		"limit":                "integer",
		"starting_after":       "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingDisputesCmd.Cmd, "retrieve", "/v1/issuing/disputes/{dispute}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingDisputesCmd.Cmd, "update", "/v1/issuing/disputes/{dispute}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingTransactionsCmd.Cmd, "list", "/v1/issuing/transactions", http.MethodGet, map[string]string{
		"card":           "string",
		"cardholder":     "string",
		"dispute":        "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rIssuingTransactionsCmd.Cmd, "retrieve", "/v1/issuing/transactions/{transaction}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rIssuingTransactionsCmd.Cmd, "update", "/v1/issuing/transactions/{transaction}", http.MethodPost, map[string]string{}, &Config)
	resource.NewOperationCmd(rRadarEarlyFraudWarningsCmd.Cmd, "list", "/v1/radar/early_fraud_warnings", http.MethodGet, map[string]string{
		"charge":         "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rRadarEarlyFraudWarningsCmd.Cmd, "retrieve", "/v1/radar/early_fraud_warnings/{early_fraud_warning}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rRadarValueListItemsCmd.Cmd, "create", "/v1/radar/value_list_items", http.MethodPost, map[string]string{
		"value":      "string",
		"value_list": "string",
	}, &Config)
	resource.NewOperationCmd(rRadarValueListItemsCmd.Cmd, "delete", "/v1/radar/value_list_items/{item}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rRadarValueListItemsCmd.Cmd, "list", "/v1/radar/value_list_items", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
		"value":          "string",
		"value_list":     "string",
	}, &Config)
	resource.NewOperationCmd(rRadarValueListItemsCmd.Cmd, "retrieve", "/v1/radar/value_list_items/{item}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rRadarValueListsCmd.Cmd, "create", "/v1/radar/value_lists", http.MethodPost, map[string]string{
		"alias":     "string",
		"item_type": "string",
		"name":      "string",
	}, &Config)
	resource.NewOperationCmd(rRadarValueListsCmd.Cmd, "delete", "/v1/radar/value_lists/{value_list}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rRadarValueListsCmd.Cmd, "list", "/v1/radar/value_lists", http.MethodGet, map[string]string{
		"alias":          "string",
		"contains":       "string",
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rRadarValueListsCmd.Cmd, "retrieve", "/v1/radar/value_lists/{value_list}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rRadarValueListsCmd.Cmd, "update", "/v1/radar/value_lists/{value_list}", http.MethodPost, map[string]string{
		"alias": "string",
		"name":  "string",
	}, &Config)
	resource.NewOperationCmd(rReportingReportRunsCmd.Cmd, "create", "/v1/reporting/report_runs", http.MethodPost, map[string]string{
		"report_type": "string",
	}, &Config)
	resource.NewOperationCmd(rReportingReportRunsCmd.Cmd, "list", "/v1/reporting/report_runs", http.MethodGet, map[string]string{
		"ending_before":  "string",
		"limit":          "integer",
		"starting_after": "string",
	}, &Config)
	resource.NewOperationCmd(rReportingReportRunsCmd.Cmd, "retrieve", "/v1/reporting/report_runs/{report_run}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rReportingReportTypesCmd.Cmd, "list", "/v1/reporting/report_types", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rReportingReportTypesCmd.Cmd, "retrieve", "/v1/reporting/report_types/{report_type}", http.MethodGet, map[string]string{}, &Config)
	resource.NewOperationCmd(rTerminalConnectionTokensCmd.Cmd, "create", "/v1/terminal/connection_tokens", http.MethodPost, map[string]string{
		"location":         "string",
		"operator_account": "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalLocationsCmd.Cmd, "create", "/v1/terminal/locations", http.MethodPost, map[string]string{
		"display_name":     "string",
		"operator_account": "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalLocationsCmd.Cmd, "delete", "/v1/terminal/locations/{location}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rTerminalLocationsCmd.Cmd, "list", "/v1/terminal/locations", http.MethodGet, map[string]string{
		"ending_before":    "string",
		"limit":            "integer",
		"operator_account": "string",
		"starting_after":   "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalLocationsCmd.Cmd, "retrieve", "/v1/terminal/locations/{location}", http.MethodGet, map[string]string{
		"operator_account": "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalLocationsCmd.Cmd, "update", "/v1/terminal/locations/{location}", http.MethodPost, map[string]string{
		"display_name":     "string",
		"operator_account": "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalReadersCmd.Cmd, "create", "/v1/terminal/readers", http.MethodPost, map[string]string{
		"label":             "string",
		"location":          "string",
		"operator_account":  "string",
		"registration_code": "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalReadersCmd.Cmd, "delete", "/v1/terminal/readers/{reader}", http.MethodDelete, map[string]string{}, &Config)
	resource.NewOperationCmd(rTerminalReadersCmd.Cmd, "list", "/v1/terminal/readers", http.MethodGet, map[string]string{
		"ending_before":    "string",
		"limit":            "integer",
		"location":         "string",
		"operator_account": "string",
		"starting_after":   "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalReadersCmd.Cmd, "retrieve", "/v1/terminal/readers/{reader}", http.MethodGet, map[string]string{
		"operator_account": "string",
	}, &Config)
	resource.NewOperationCmd(rTerminalReadersCmd.Cmd, "update", "/v1/terminal/readers/{reader}", http.MethodPost, map[string]string{
		"label":            "string",
		"operator_account": "string",
	}, &Config)
}
