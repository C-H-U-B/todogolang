package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// TodoList stores tasks
type TodoList []TodoItem

// Save TodoList
func (t *TodoList) Save(path string) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0644)
}

// LoadTodoList returns ToDoList
func LoadTodoList(path string) (TodoList, error) {

	var list TodoList
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)

	if err != nil {
		return list, err
	}

	return list, nil

}

// TodoItem stores information about a task
type TodoItem struct {
	// Add field description
	Description string

	// Add field due date
	Date time.Time
}

// Add item to TodoList
func (t *TodoList) Add(item TodoItem) {
	// Append item to `t` (list)

	*t = append(*t, item)
	// ...
}


func (t *TodoList) Delete(index int){
	*t = append((*t)[:index], (*t)[index+1:]...)
}

func main() {
	// todo, err := LoadTodoList("./todo.json")
	// if err != nil {
	// 	todo = TodoList{}
	// }

	// todo.Add(TodoItem{
	// 	Description: "test",
	// 	Date:        time.Now().Add(time.Hour * 24),
	// })

	// todo.Add(TodoItem{
	// 	Description: "test",
	// 	Date:        time.Now().Add(time.Hour * 24),
	// })

	// err = todo.Save("./todo.json")
	// if err != nil {
	// 	panic(err)
	// }
	Run()
}
