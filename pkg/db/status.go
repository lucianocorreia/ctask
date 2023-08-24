package db

type status int

const (
	todo status = iota
	inProgress
	done
)

// String returns the string representation of the status
func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

// Next returns the next status
func (s status) Next() int {
	if s == done {
		return int(todo)
	}

	return int(s + 1)
}

// Prev returns the previous status
func (s status) Prev() int {
	if s == todo {
		return int(done)
	}

	return int(s - 1)
}

// Int returns the int representation of the status
func (s status) Int() int {
	return int(s)
}
