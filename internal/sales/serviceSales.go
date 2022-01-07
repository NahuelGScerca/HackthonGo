package sales

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
	"github.com/NahuelGScerca/HackthonGo/utils"
)

type Service interface {
	Get(ctx context.Context, id int) (models.Sales, error)
	ExportData(ctx context.Context) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (serv *service) Get(ctx context.Context, id int) (models.Sales, error) {
	sale, err := serv.repository.Get(ctx, id)
	if err != nil {
		return models.Sales{}, err
	}
	return sale, nil
}

func (serv *service) ExportData(ctx context.Context) error {

	data, err := utils.ReadData("sales")

	if err != nil {
		fmt.Println("ERROR")
	}

	for _, sale := range data {
		parseData := strings.Split(sale, "#$%#")

		idParsed, _ := strconv.Atoi(parseData[0])
		idProductParsed, _ := strconv.Atoi(parseData[1])
		idInvoiceParsed, _ := strconv.Atoi(parseData[2])
		idQuantityParsed, _ := strconv.ParseFloat(parseData[2], 64)

		salesToSend := models.Sales{
			ID:        idParsed,
			IdProduct: idProductParsed,
			IdInvoice: idInvoiceParsed,
			Quantity:  idQuantityParsed,
		}

		err := serv.repository.Store(context.Background(), salesToSend)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	return nil

}
