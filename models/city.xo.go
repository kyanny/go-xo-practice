// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// City represents a row from 'public.city'.
type City struct {
	CityID     int       `json:"city_id"`     // city_id
	City       string    `json:"city"`        // city
	CountryID  int16     `json:"country_id"`  // country_id
	LastUpdate time.Time `json:"last_update"` // last_update

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the City exists in the database.
func (c *City) Exists() bool {
	return c._exists
}

// Deleted provides information if the City has been deleted from the database.
func (c *City) Deleted() bool {
	return c._deleted
}

// Insert inserts the City to the database.
func (c *City) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.city (` +
		`city_id, city, country_id, last_update` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`

	// run query
	XOLog(sqlstr, c.CityID, c.City, c.CountryID, c.LastUpdate)
	err = db.QueryRow(sqlstr, c.CityID, c.City, c.CountryID, c.LastUpdate).Scan(&c.CityID)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the City in the database.
func (c *City) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.city SET (` +
		`city, country_id, last_update` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE city_id = $4`

	// run query
	XOLog(sqlstr, c.City, c.CountryID, c.LastUpdate, c.CityID)
	_, err = db.Exec(sqlstr, c.City, c.CountryID, c.LastUpdate, c.CityID)
	return err
}

// Save saves the City to the database.
func (c *City) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for City.
//
// NOTE: PostgreSQL 9.5+ only
func (c *City) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.city (` +
		`city_id, city, country_id, last_update` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (city_id) DO UPDATE SET (` +
		`city_id, city, country_id, last_update` +
		`) = (` +
		`EXCLUDED.city_id, EXCLUDED.city, EXCLUDED.country_id, EXCLUDED.last_update` +
		`)`

	// run query
	XOLog(sqlstr, c.CityID, c.City, c.CountryID, c.LastUpdate)
	_, err = db.Exec(sqlstr, c.CityID, c.City, c.CountryID, c.LastUpdate)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the City from the database.
func (c *City) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.city WHERE city_id = $1`

	// run query
	XOLog(sqlstr, c.CityID)
	_, err = db.Exec(sqlstr, c.CityID)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// Country returns the Country associated with the City's CountryID (country_id).
//
// Generated from foreign key 'fk_city'.
func (c *City) Country(db XODB) (*Country, error) {
	return CountryByCountryID(db, int(c.CountryID))
}

// CityByCityID retrieves a row from 'public.city' as a City.
//
// Generated from index 'city_pkey'.
func CityByCityID(db XODB, cityID int) (*City, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`city_id, city, country_id, last_update ` +
		`FROM public.city ` +
		`WHERE city_id = $1`

	// run query
	XOLog(sqlstr, cityID)
	c := City{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, cityID).Scan(&c.CityID, &c.City, &c.CountryID, &c.LastUpdate)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// CitiesByCountryID retrieves a row from 'public.city' as a City.
//
// Generated from index 'idx_fk_country_id'.
func CitiesByCountryID(db XODB, countryID int16) ([]*City, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`city_id, city, country_id, last_update ` +
		`FROM public.city ` +
		`WHERE country_id = $1`

	// run query
	XOLog(sqlstr, countryID)
	q, err := db.Query(sqlstr, countryID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*City{}
	for q.Next() {
		c := City{
			_exists: true,
		}

		// scan
		err = q.Scan(&c.CityID, &c.City, &c.CountryID, &c.LastUpdate)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}