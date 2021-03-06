// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Payment represents a row from 'public.payment'.
type Payment struct {
	PaymentID   int       `json:"payment_id"`   // payment_id
	CustomerID  int16     `json:"customer_id"`  // customer_id
	StaffID     int16     `json:"staff_id"`     // staff_id
	RentalID    int       `json:"rental_id"`    // rental_id
	Amount      float64   `json:"amount"`       // amount
	PaymentDate time.Time `json:"payment_date"` // payment_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Payment exists in the database.
func (p *Payment) Exists() bool {
	return p._exists
}

// Deleted provides information if the Payment has been deleted from the database.
func (p *Payment) Deleted() bool {
	return p._deleted
}

// Insert inserts the Payment to the database.
func (p *Payment) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.payment (` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`

	// run query
	XOLog(sqlstr, p.PaymentID, p.CustomerID, p.StaffID, p.RentalID, p.Amount, p.PaymentDate)
	err = db.QueryRow(sqlstr, p.PaymentID, p.CustomerID, p.StaffID, p.RentalID, p.Amount, p.PaymentDate).Scan(&p.PaymentID)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Payment in the database.
func (p *Payment) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.payment SET (` +
		`customer_id, staff_id, rental_id, amount, payment_date` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE payment_id = $6`

	// run query
	XOLog(sqlstr, p.CustomerID, p.StaffID, p.RentalID, p.Amount, p.PaymentDate, p.PaymentID)
	_, err = db.Exec(sqlstr, p.CustomerID, p.StaffID, p.RentalID, p.Amount, p.PaymentDate, p.PaymentID)
	return err
}

// Save saves the Payment to the database.
func (p *Payment) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Payment.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Payment) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.payment (` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (payment_id) DO UPDATE SET (` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date` +
		`) = (` +
		`EXCLUDED.payment_id, EXCLUDED.customer_id, EXCLUDED.staff_id, EXCLUDED.rental_id, EXCLUDED.amount, EXCLUDED.payment_date` +
		`)`

	// run query
	XOLog(sqlstr, p.PaymentID, p.CustomerID, p.StaffID, p.RentalID, p.Amount, p.PaymentDate)
	_, err = db.Exec(sqlstr, p.PaymentID, p.CustomerID, p.StaffID, p.RentalID, p.Amount, p.PaymentDate)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Payment from the database.
func (p *Payment) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.payment WHERE payment_id = $1`

	// run query
	XOLog(sqlstr, p.PaymentID)
	_, err = db.Exec(sqlstr, p.PaymentID)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// Customer returns the Customer associated with the Payment's CustomerID (customer_id).
//
// Generated from foreign key 'payment_customer_id_fkey'.
func (p *Payment) Customer(db XODB) (*Customer, error) {
	return CustomerByCustomerID(db, int(p.CustomerID))
}

// Rental returns the Rental associated with the Payment's RentalID (rental_id).
//
// Generated from foreign key 'payment_rental_id_fkey'.
func (p *Payment) Rental(db XODB) (*Rental, error) {
	return RentalByRentalID(db, p.RentalID)
}

// Staff returns the Staff associated with the Payment's StaffID (staff_id).
//
// Generated from foreign key 'payment_staff_id_fkey'.
func (p *Payment) Staff(db XODB) (*Staff, error) {
	return StaffByStaffID(db, int(p.StaffID))
}

// PaymentsByCustomerID retrieves a row from 'public.payment' as a Payment.
//
// Generated from index 'idx_fk_customer_id'.
func PaymentsByCustomerID(db XODB, customerID int16) ([]*Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date ` +
		`FROM public.payment ` +
		`WHERE customer_id = $1`

	// run query
	XOLog(sqlstr, customerID)
	q, err := db.Query(sqlstr, customerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Payment{}
	for q.Next() {
		p := Payment{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.PaymentID, &p.CustomerID, &p.StaffID, &p.RentalID, &p.Amount, &p.PaymentDate)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

// PaymentsByRentalID retrieves a row from 'public.payment' as a Payment.
//
// Generated from index 'idx_fk_rental_id'.
func PaymentsByRentalID(db XODB, rentalID int) ([]*Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date ` +
		`FROM public.payment ` +
		`WHERE rental_id = $1`

	// run query
	XOLog(sqlstr, rentalID)
	q, err := db.Query(sqlstr, rentalID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Payment{}
	for q.Next() {
		p := Payment{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.PaymentID, &p.CustomerID, &p.StaffID, &p.RentalID, &p.Amount, &p.PaymentDate)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

// PaymentsByStaffID retrieves a row from 'public.payment' as a Payment.
//
// Generated from index 'idx_fk_staff_id'.
func PaymentsByStaffID(db XODB, staffID int16) ([]*Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date ` +
		`FROM public.payment ` +
		`WHERE staff_id = $1`

	// run query
	XOLog(sqlstr, staffID)
	q, err := db.Query(sqlstr, staffID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Payment{}
	for q.Next() {
		p := Payment{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.PaymentID, &p.CustomerID, &p.StaffID, &p.RentalID, &p.Amount, &p.PaymentDate)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

// PaymentByPaymentID retrieves a row from 'public.payment' as a Payment.
//
// Generated from index 'payment_pkey'.
func PaymentByPaymentID(db XODB, paymentID int) (*Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`payment_id, customer_id, staff_id, rental_id, amount, payment_date ` +
		`FROM public.payment ` +
		`WHERE payment_id = $1`

	// run query
	XOLog(sqlstr, paymentID)
	p := Payment{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, paymentID).Scan(&p.PaymentID, &p.CustomerID, &p.StaffID, &p.RentalID, &p.Amount, &p.PaymentDate)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
