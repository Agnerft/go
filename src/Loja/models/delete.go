package models

import "Loja/db"

func Delete(int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `DELETE FROM todos WHERE id=$1`

	res, err := conn.Exec(sql, todo.ID)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
