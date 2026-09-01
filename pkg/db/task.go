package db

import (
	"database/sql"
	"time"

	"go_final_project/pkg/scheduler"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := "INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)"
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func GetTasks(limit int, search string) ([]*Task, error) {
	var tasks []*Task

	rows, err := findTasks(limit, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	var task Task
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"
	err := DB.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	return &task, err
}

func UpdateTask(task *Task) error {
	query := "UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?"
	resp, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func DeleteTask(id string) error {
	query := "DELETE FROM scheduler WHERE id = ?"
	resp, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

const searchDateFormat = "02.01.2006"

func findTasks(limit int, search string) (*sql.Rows, error) {
	baseQuery := "SELECT id, date, title, comment, repeat FROM scheduler"

	if search == "" {
		query := baseQuery + " ORDER BY date LIMIT ?"
		return DB.Query(query, limit)
	}

	if date, err := time.Parse(searchDateFormat, search); err == nil {
		query := baseQuery + " WHERE date = ? ORDER BY date LIMIT ?"
		return DB.Query(query, date.Format(scheduler.DateFormat), limit)
	}

	query := baseQuery + " WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?"
	like := "%" + search + "%"
	return DB.Query(query, like, like, limit)
}
