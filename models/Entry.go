package models

import "time"

type Entry struct {
  stamp time.Time
  message   string
}

func (e *Entry) getReadableStamp() string {
  return e.stamp.Format(time.RFC3339)[0:19]
}
 
func (e *Entry) ToString() string {
  return e.getReadableStamp() + " " + e.message
}
