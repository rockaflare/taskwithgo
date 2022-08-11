package entity

type Task struct {
	ID       uint   `json:"id"`
	Pegawai  string `json:"pegawai"`
	Detail   string `json:"detail"`
	Deadline string `json:"deadline"`
	IsDone   bool   `json:"isdone"`
}
