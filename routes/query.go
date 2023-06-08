package routes

const (
	ALL_TASKS   string = "SELECT * FROM list.task"
	SET_POST    string = "INSERT INTO list.task (title, task) VALUES ($1, $2)"
	DELETE_TASK string = "DELETE FROM list.task WHERE id = $1"
	UPDATE_TASK string = "UPDATE list.task SET title = $1, task = $2 WHERE id = $3"
)
