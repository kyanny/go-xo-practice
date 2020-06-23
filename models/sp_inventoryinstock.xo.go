// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// InventoryInStock calls the stored procedure 'public.inventory_in_stock(integer) boolean' on db.
func InventoryInStock(db XODB, v0 int) (bool, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public.inventory_in_stock($1)`

	// run query
	var ret bool
	XOLog(sqlstr, v0)
	err = db.QueryRow(sqlstr, v0).Scan(&ret)
	if err != nil {
		return false, err
	}

	return ret, nil
}
