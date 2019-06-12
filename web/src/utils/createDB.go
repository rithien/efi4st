/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package utils

import "github.com/jmoiron/sqlx"

// exec the schema or fail; multi-statement Exec behavior varies between
// database drivers;  pq will exec them all, sqlite3 won't, ymmv
func CreateDB(db *sqlx.DB) {
	db.MustExec(projectSchema)
	db.MustExec(firmwareSchema)
}