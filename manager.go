package main

type Manager struct {
	results     Results
	CurrentView int
}

func NewManager() Manager {
	r := NewResults()
	return Manager{
		results:     r,
		CurrentView: 0,
	}
}

type Items map[string]any

type Results struct {
	items Items
}

func NewResults() Results {
	return Results{}
}

func (r Results) GetItems() Items {
	return r.items
}

func (r *Results) SetItem(key string, value any) {
	_, ok := r.items[key]
	if ok {
		panic("Item key already exists!")
	}

	r.items[key] = value
}

func (r Results) GetItemValue(key string) any {
	return r.items[key]
}

func (r Results) ContainsKey(key string) bool {
	_, ok := r.items[key]
	return ok
}
