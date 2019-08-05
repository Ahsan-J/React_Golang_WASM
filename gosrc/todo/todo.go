package todo

import (
	"react_go_wasm/gosrc/helper"
	"syscall/js"
)

var todoItems []Item

// AddItem adds Todo Item
func addItem(i []js.Value) {
	createdAt := helper.GetCurrentISOTime()
	updatedAt := helper.GetCurrentISOTime()
	task := js.ValueOf(i[0].String()).String()
	todoItems = append(todoItems, Item{task, createdAt, updatedAt, 0})
	js.Global().Set("todo_items", helper.GetStructVal(todoItems))
}

func editItem(i []js.Value) {

}

func deleteItem(i []js.Value) {

}

// RegisterCallbacks regsters functions in golang and creates it with JS
func RegisterCallbacks() {
	js.Global().Set("addItem", js.NewCallback(addItem))
	js.Global().Set("editItem", js.NewCallback(editItem))
	js.Global().Set("deleteItem", js.NewCallback(deleteItem))
	js.Global().Set("todo_items", helper.GetStructVal(todoItems))
}
