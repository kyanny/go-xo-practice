// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Store represents a row from 'public.store'.
type Store struct {
	StoreID        int       `json:"store_id"`         // store_id
	ManagerStaffID int16     `json:"manager_staff_id"` // manager_staff_id
	AddressID      int16     `json:"address_id"`       // address_id
	LastUpdate     time.Time `json:"last_update"`      // last_update

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Store exists in the database.
func (s *Store) Exists() bool {
	return s._exists
}

// Deleted provides information if the Store has been deleted from the database.
func (s *Store) Deleted() bool {
	return s._deleted
}

// Insert inserts the Store to the database.
func (s *Store) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.store (` +
		`store_id, manager_staff_id, address_id, last_update` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`

	// run query
	XOLog(sqlstr, s.StoreID, s.ManagerStaffID, s.AddressID, s.LastUpdate)
	err = db.QueryRow(sqlstr, s.StoreID, s.ManagerStaffID, s.AddressID, s.LastUpdate).Scan(&s.StoreID)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the Store in the database.
func (s *Store) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.store SET (` +
		`manager_staff_id, address_id, last_update` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE store_id = $4`

	// run query
	XOLog(sqlstr, s.ManagerStaffID, s.AddressID, s.LastUpdate, s.StoreID)
	_, err = db.Exec(sqlstr, s.ManagerStaffID, s.AddressID, s.LastUpdate, s.StoreID)
	return err
}

// Save saves the Store to the database.
func (s *Store) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for Store.
//
// NOTE: PostgreSQL 9.5+ only
func (s *Store) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.store (` +
		`store_id, manager_staff_id, address_id, last_update` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (store_id) DO UPDATE SET (` +
		`store_id, manager_staff_id, address_id, last_update` +
		`) = (` +
		`EXCLUDED.store_id, EXCLUDED.manager_staff_id, EXCLUDED.address_id, EXCLUDED.last_update` +
		`)`

	// run query
	XOLog(sqlstr, s.StoreID, s.ManagerStaffID, s.AddressID, s.LastUpdate)
	_, err = db.Exec(sqlstr, s.StoreID, s.ManagerStaffID, s.AddressID, s.LastUpdate)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the Store from the database.
func (s *Store) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.store WHERE store_id = $1`

	// run query
	XOLog(sqlstr, s.StoreID)
	_, err = db.Exec(sqlstr, s.StoreID)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

// Address returns the Address associated with the Store's AddressID (address_id).
//
// Generated from foreign key 'store_address_id_fkey'.
func (s *Store) Address(db XODB) (*Address, error) {
	return AddressByAddressID(db, int(s.AddressID))
}

// Staff returns the Staff associated with the Store's ManagerStaffID (manager_staff_id).
//
// Generated from foreign key 'store_manager_staff_id_fkey'.
func (s *Store) Staff(db XODB) (*Staff, error) {
	return StaffByStaffID(db, int(s.ManagerStaffID))
}

// StoreByManagerStaffID retrieves a row from 'public.store' as a Store.
//
// Generated from index 'idx_unq_manager_staff_id'.
func StoreByManagerStaffID(db XODB, managerStaffID int16) (*Store, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`store_id, manager_staff_id, address_id, last_update ` +
		`FROM public.store ` +
		`WHERE manager_staff_id = $1`

	// run query
	XOLog(sqlstr, managerStaffID)
	s := Store{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, managerStaffID).Scan(&s.StoreID, &s.ManagerStaffID, &s.AddressID, &s.LastUpdate)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// StoreByStoreID retrieves a row from 'public.store' as a Store.
//
// Generated from index 'store_pkey'.
func StoreByStoreID(db XODB, storeID int) (*Store, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`store_id, manager_staff_id, address_id, last_update ` +
		`FROM public.store ` +
		`WHERE store_id = $1`

	// run query
	XOLog(sqlstr, storeID)
	s := Store{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, storeID).Scan(&s.StoreID, &s.ManagerStaffID, &s.AddressID, &s.LastUpdate)
	if err != nil {
		return nil, err
	}

	return &s, nil
}