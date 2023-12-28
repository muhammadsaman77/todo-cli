package src

type TodoService interface {
	GetData() *[]Todo
	AddData(newTodo string)
	EditTodo(index int, value string)
	UpdateStatus(index int)
	Delete(index int)
	DeleteAll()
}

type todoService struct {
}

func NewService() TodoService {
	return &todoService{}
}

func (service *todoService) GetData() *[]Todo {
	return &Data
}

func (service *todoService) AddData(newData string) {
	Data = append(Data, Todo{Task: newData, Status: "uncompleted"})
}

func (service *todoService) EditTodo(index int, value string) {
	Data[index].Task = value
}
func (service *todoService) UpdateStatus(index int) {
	if Data[index].Status == "uncompleted" {
		Data[index].Status = "completed"
	} else {
		Data[index].Status = "uncompleted"
	}
}

func (service *todoService) Delete(index int) {

	Data = append(Data[:index], Data[index+1:]...)
}

func (service *todoService) DeleteAll() {
	Data = []Todo{}
}