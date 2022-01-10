package customers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
	"github.com/NahuelGScerca/HackthonGo/utils"
)

type Service interface {
	ExportData(ctx context.Context) error
	Get(ctx context.Context, id int) (models.Customers, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (serv *service) Get(ctx context.Context, id int) (models.Customers, error) {
	customer, err := serv.repository.Get(ctx, id)
	if err != nil {
		return models.Customers{}, err
	}
	return customer, nil
}

func (serv *service) ExportData(ctx context.Context) error {

	data, err := utils.ReadData("customers")

	if err != nil {
		fmt.Println("ERROR")
	}

	for _, sale := range data {
		parseData := strings.Split(sale, "#$%#")
		idParsed, _ := strconv.Atoi(parseData[0])
		customerToSend := models.Customers{
			ID:             idParsed,
			LastName:       parseData[1],
			FirstName:      parseData[2],
			ConditionState: parseData[3],
		}

		err := serv.repository.Store(context.Background(), customerToSend)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	return nil

}
