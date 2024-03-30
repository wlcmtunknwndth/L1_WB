package main

type Storage interface {
	Save(d *Data)
	Get(name string)
}

func main() {
	storages := []Storage{NewSqlDbAdapter(&SqlDb{}), NewCacherAdapter(&Cacher{})}

	for i, _ := range storages {
		storages[i].Save(&Data{})
		storages[i].Get("idk")
	}
}
