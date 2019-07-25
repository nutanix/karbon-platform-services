package io_utils

import (
	"fmt"
	"strings"
)

// ToSQLList returns a list of strings aas is used in SQL `WHERE` clause.
// for example: ('1','2','3','4')
func ToSQLList(input []string) string {
	ids := make([]string, 0, len(input))
	for _, id := range input {
		ids = append(ids, fmt.Sprintf("'%s'", id))
	}
	return fmt.Sprintf("%s", strings.Join(ids, ","))
}
