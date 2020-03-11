package main

import (
	"database/sql"
	"fmt"
)

type Customer struct {
	ID                int
	FirstName         string
	LastName          string
	Email             string
	State             string
	CustomerStatus    string
	LoanStatus        string
	BirthDate         string
	SSN               string
	FundingDate       string
	DueDateAdjusted   string
}

func paydayPaidOffCustomer(db sql.DB) {
	sqlStatement := `SELECT customers.id,
	   loan.gov_law_state_cd,
     customers.email,
	   customers.status_cd,
	   loan.status_cd,
	   people.birth_date,
	   people.ssn,
	   loan.funding_date,
	   loan.due_date_adjusted,
	   people.first_name,
	   people.last_name
FROM
  customers
JOIN loans loan
ON loan.customer_id = customers.id
AND loan.loan_type_cd = 'payday'
JOIN people
ON people.id = customers.person_id
WHERE customers.status_cd = 'active'
AND NOT EXISTS (SELECT 1 FROM
				loans l2 WHERE l2.customer_id = loan.customer_id
				AND l2.status_cd IN ('applied', 'approved', 'issued', 'issued_pmt_proc', 'in_default', 'in_default_pmt_proc', 'on_hold', 'withdrawn'))
AND loan.gov_law_state_cd in ('ND')
AND loan.base_loan_id is NULL
AND EXTRACT(DAY FROM people.birth_date)::Int % 2 = 0
ORDER BY customers.id DESC LIMIT 1;`
	var customer Customer
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&customer.ID, &customer.Email, &customer.State, &customer.BirthDate,
		&customer.FirstName, &customer.LastName, &customer.SSN, &customer.CustomerStatus,
		&customer.LoanStatus, &customer.FundingDate, &customer.DueDateAdjusted)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(customer)
	default:
		panic(err)
	}
}
