// Определение типа данных верхнего уровня
package main

import (
	"errors"
	"time"
)

var ErrNoRecords = errors.New("no records found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
