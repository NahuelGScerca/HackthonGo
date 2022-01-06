package models

type Invoices struct {
	ID         int
	Datetime   string
	IdCustomer int
	Total      float64
}
