package user

import (
  "time"
)

type Person struct {
  ID            uint64
  FirstName     string
  LastName      string
  DateOfBirth   time.Time
}
