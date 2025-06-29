package comandos

import "regexp"

var sqliteFalhouInserirRow = regexp.MustCompile("failed to get next row\nerror code = 1: Error fetching next row: SQLite failure: `(.*?)`")
