// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// GroupConcat calls the stored procedure 'public.group_concat(text) text' on db.
func GroupConcat(db XODB, v0 string) (string, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public.group_concat($1)`

	// run query
	var ret string
	XOLog(sqlstr, v0)
	err = db.QueryRow(sqlstr, v0).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}

// GroupConcat calls the stored procedure 'public._group_concat(text, text) text' on db.
func GroupConcat(db XODB, v0 string, v1 string) (string, error) {
	var err error

	// sql query
	const sqlstr = `SELECT public._group_concat($1, $2)`

	// run query
	var ret string
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}
