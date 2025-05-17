package sample

type Entity struct {
	ID     int64  `db:"ID"`
	Name   string `db:"Name"`
	Active bool   `db:"Active"`
}
