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

func (pr *SaintRepository) CreateSaint(saint model.Saint) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO saint" +
		"(saint_name, quote)" +
		"VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(saint.Name, saint.Quote).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *SaintRepository) GetSaintById(id_saint int) (*model.Saint, error) {

	query, err := pr.connection.Prepare("SELECT * FROM saint WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var saint model.Saint

	err = query.QueryRow(id_saint).Scan(
		&saint.ID,
		&saint.Name,
		&saint.Quote,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &saint, nil
}
