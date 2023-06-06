package routes

const ALL_TASKS string = "SELECT * FROM list.task"
const SET_POST string = "INSERT INTO list.task (title, task) VALUES ($1, $2)"
const DELETE_TASK string = "DELETE FROM list.task WHERE id = $1"
const UPDATE_TASK string = "UPDATE list.task SET title = $1, task = $2 WHERE id = $3"
