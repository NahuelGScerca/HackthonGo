package customers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Store(ctx context.Context, entidad models.Customers) error
	Get(ctx context.Context, id int) (models.Customers, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Customers, error) {
	query := "SELECT * FROM customers WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Customers{}
	err := row.Scan(&b.ID, &b.LastName, &b.FirstName, &b.ConditionState)
	if err != nil {
		return models.Customers{}, err
	}

	return b, nil
}

func (r *repository) Store(ctx context.Context, entidad models.Customers) error {
	// fmt.Println("ENTIDAD:  ", entidad)
	// return nil
	query := "INSERT INTO customers(id,lastName,firstName,conditionState) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println("error en el prepare")
		fmt.Println(query)

		return err
	}

	res, err := stmt.Exec(&entidad.ID, &entidad.LastName, &entidad.FirstName, &entidad.ConditionState)
	if err != nil {
		return err
	}

	if num, err := res.RowsAffected(); num > 0 && err == nil {
		defer stmt.Close()

		return nil
	}

	return fmt.Errorf("Error insert")
}
