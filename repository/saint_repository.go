package repository

import (
	"database/sql"
	"fmt"
	"saints-api/model"
)

type SaintRepository struct {
	connection *sql.DB
}

func NewSaintRepository(connection *sql.DB) SaintRepository {
	return SaintRepository{
		connection: connection,
	}
}

func (pr *SaintRepository) GetSaints() ([]model.Saint, error) {

	query := "SELECT id, saint_name, quote FROM saint"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Saint{}, err
	}

	var saintList []model.Saint
	var saintObj model.Saint

	for rows.Next() {
		err = rows.Scan(
			&saintObj.ID,
			&saintObj.Name,
			&saintObj.Quote)

		if err != nil {
			fmt.Println(err)
			return []model.Saint{}, err
		}

		saintList = append(saintList, saintObj)
	}

	rows.Close()

	return saintList, nil
}
