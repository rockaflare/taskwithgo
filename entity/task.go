package entity

type Task struct {
	ID      int    `json:"id"`
	Pegawai string `json:"pegawai"`
	Detail  string `json:"detail"`
	IsDone  string `json:"isdone"`
}
