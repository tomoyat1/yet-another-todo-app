package postgres

import (
	"github.com/jackc/pgx"

	todo "github.com/tomoyat1/yet-another-todo-app"
	"fmt"
)

type PgTodoItemImpl struct {
	ID int
	todo.Item
}

type PgItemRepositoryImpl struct {
	connPool *pgx.ConnPool
}

func NewPgItemRepositoryImpl(connString string) (*PgItemRepositoryImpl, error) {
	connCfg, err := pgx.ParseConnectionString(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create Item repository: %s", err.Error())
	}
	connPoolCfg := pgx.ConnPoolConfig{
		ConnConfig: connCfg,
		MaxConnections: 3,
	}
	c, err := pgx.NewConnPool(connPoolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Item repository: %s", err.Error())
	}
	return &PgItemRepositoryImpl{
		connPool: c,
	}, nil
}

func (r *PgItemRepositoryImpl) Save(t *todo.Item) (err error) {
	const insertQuery = `
INSERT INTO todos (
  id,
  title,
  details,
  done
) VALUES (
  $1,
  $2,
  $3,
  $4
)
`
	const selectIDQuery = `
SELECT id
FROM todos
WHERE id = $1
`
	const updateQuery = `
UPDATE todos 
SET
  title=$2,
  details=$3,
  done=$4	
WHERE id=$1
`
	tx, err := r.connPool.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	row := tx.QueryRow(selectIDQuery, t.ID)
	var existingId string
	err = row.Scan(&existingId)
	if err == pgx.ErrNoRows {
		_, err = tx.Exec(insertQuery, t.ID, t.Title, t.Details, t.Done)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	} else {
		_, err = tx.Exec(updateQuery, t.ID, t.Title, t.Details, t.Done)
		if err != nil {
			return
		}
	}
	tx.Commit()
	return nil
}

func (r *PgItemRepositoryImpl) List() ([]*todo.Item, error) {
	rows, err := r.connPool.Query("SELECT id,title,details,done FROM todos")
	if err != nil {
		return nil, err
	}
	todos := make([]*todo.Item, 0)
	for rows.Next() {
		t := todo.Item{}
		err := rows.Scan(&t.ID, &t.Title, &t.Details, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &t)
	}
	return todos, nil
}
