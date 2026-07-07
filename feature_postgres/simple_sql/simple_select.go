package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT * FROM tasks 
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next(){
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)

		if err != nil {
			return nil, err
		}

		fmt.Println(
			task.ID, 
			task.Title, 
			task.Description, 
			task.Completed, 
			task.CreatedAt, 
			task.CompletedAt,
		)
	}	
		return tasks, nil
}