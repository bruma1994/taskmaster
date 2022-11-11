package taskmaster

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id   int       `json:"id,omitempty"`
	Text string    `json:"text,omitempty"`
	Tags []string  `json:"tags,omitempty"`
	Due  time.Time `json:"due"`
}

type TaskMaster struct {
	sync.Mutex
	tasks  map[int]Task
	nextId int
}

func (t *TaskMaster) CreateTask(text string, tags []string, due time.Time) int {
	t.Lock()
	defer t.Unlock()

	task := Task{
		Id:   t.nextId,
		Text: text,
		Due:  due,
	}
	copy(task.Tags, tags)

	t.tasks[t.nextId] = task
	t.nextId++

	return task.Id
}

func (t *TaskMaster) GetTask(id int) (Task, error) {
	t.Lock()
	defer t.Unlock()

	if task, ok := t.tasks[id]; ok {
		return task, nil
	} else {
		return Task{}, fmt.Errorf("task with id %d not found", id)
	}
}

func (t *TaskMaster) DeleteTask(id int) error {
	t.Lock()
	defer t.Unlock()

	if _, ok := t.tasks[id]; !ok {
		return fmt.Errorf("task with id %d not found", id)
	}

	delete(t.tasks, id)
	return nil
}

func (t *TaskMaster) ClearStore() error {
	t.Lock()
	defer t.Unlock()

	t.tasks = make(map[int]Task)
	return nil
}

func (t *TaskMaster) GetAllTasks() []Task {
	t.Lock()
	defer t.Unlock()

	allTasks := make([]Task, 0, len(t.tasks))
	for task := range t.tasks {
		allTasks = append(allTasks, task)
	}
}

func (t *TaskMaster) GetTaskByTag(tag string) []Task {
	//TODO implement me
	panic("implement me")
}

func (t *TaskMaster) GetTaskByDueDate(year int, month time.Month, day int) []Task {
	//TODO implement me
	panic("implement me")
}

func New() *TaskMaster {
	ts := &TaskMaster{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0

	return ts
}
