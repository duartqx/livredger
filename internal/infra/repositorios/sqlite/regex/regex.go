package regex

import "regexp"

var SqliteFalhouInserirRow = regexp.MustCompile("failed to get next row\nerror code = 1: Error fetching next row: SQLite failure: `(.*?)`")
var SqliteFalhouCommitar = regexp.MustCompile("failed to execute query COMMIT\nerror code = .: Error executing statement: SQLite failure: `(.*?)`")
