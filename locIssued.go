package main

import (
	"database/sql"
	"fmt"
)

func locIssuedCustomer(db sql.DB) {
	sqlStatement := `SELECT customers.id as customer_id, 
       customers.email,
       customers.status_cd
FROM customers
JOIN customers_extra
ON customers_extra.customer_id = customers.id
join loans
ON loans.customer_id = customers.id  AND loans.loan_type_cd = 'oec'
WHERE customers_extra.upswing_flg = false
AND customers.status_cd = 'active'
AND loans.status_cd = 'issued'
--AND loans.funding_date = '2020-02-27'
--AND loans.gov_law_state_cd = 'LA'
LIMIT 1;`
	var customer Customer
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&customer.ID, &customer.Email, &customer.CustomerStatus)
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