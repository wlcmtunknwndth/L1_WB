package main

type SqlDbAdapter struct {
	db *SqlDb
}

func NewSqlDbAdapter(db *SqlDb) *SqlDbAdapter {
	return &SqlDbAdapter{db: db}
}

func (s *SqlDbAdapter) Save(d *Data) {
	s.db.SaveData(d)
}

func (s *SqlDbAdapter) Get(name string) {
	s.db.GetData(name)
}
