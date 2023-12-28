package src

type Todo struct {
	Task   string
	Status string
}

var Data = []Todo{{Task: "Buat Laporan", Status: "uncompleted"}, {Task: "Buat PPT", Status: "completed"}}