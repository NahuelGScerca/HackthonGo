package invoices

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
	Get(ctx context.Context, id int) (models.Invoices, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (serv *service) Get(ctx context.Context, id int) (models.Invoices, error) {
	invoice, err := serv.repository.Get(ctx, id)
	if err != nil {
		return models.Invoices{}, err
	}
	return invoice, nil
}

func (serv *service) ExportData(ctx context.Context) error {

	data, err := utils.ReadData("invoices")

	if err != nil {
		fmt.Println("ERROR")
	}

	for _, sale := range data {
		parseData := strings.Split(sale, "#$%#")
		idParsed, _ := strconv.Atoi(parseData[0])
		idCustomerParsed, _ := strconv.Atoi(parseData[2])
		total, _ := strconv.ParseFloat(parseData[3], 64)

		invoiceToSend := models.Invoices{
			ID:         idParsed,
			Datetime:   parseData[1],
			IdCustomer: idCustomerParsed,
			Total:      total,
		}

		err := serv.repository.Store(context.Background(), invoiceToSend)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	return nil

}
