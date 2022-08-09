package entity

import "time"

type Task struct {
	ID       int       `json:"id"`
	Pegawai  string    `json:"pegawai"`
	Detail   string    `json:"detail"`
	Deadline time.Time `json:"deadline"`
	IsDone   bool      `json:"isdone"`
}
