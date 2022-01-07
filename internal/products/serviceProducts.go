package products

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
	"github.com/NahuelGScerca/HackthonGo/utils"
)

type Service interface {
	Get(ctx context.Context, id int) (models.Products, error)
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

func (serv *service) Get(ctx context.Context, id int) (models.Products, error) {
	product, err := serv.repository.Get(ctx, id)
	if err != nil {
		return models.Products{}, err
	}
	return product, nil
}

func (serv *service) ExportData(ctx context.Context) error {

	data, err := utils.ReadData("products")

	if err != nil {
		fmt.Println("ERROR")
	}

	for _, sale := range data {
		parseData := strings.Split(sale, "#$%#")

		idParsed, _ := strconv.Atoi(parseData[0])
		priceParsed, _ := strconv.ParseFloat(parseData[2], 64)

		productsToSend := models.Products{
			ID:          idParsed,
			Description: parseData[1],
			Price:       priceParsed,
		}

		err := serv.repository.Store(context.Background(), productsToSend)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	return nil

}
