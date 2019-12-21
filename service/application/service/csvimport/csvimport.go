package csvimport

import (
	"bufio"
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/infrastructure/utils"
)

// CsvImportInterface defines the interface for a CsvImport
type CsvImportInterface interface {
	Execute(ctx context.Context, reader *bufio.Reader) ([]wine.Wine, error)
}

// CsvImport implements a service to parse a CSV file and import in in the db.
type CsvImport struct {
	CellarRepository cellar.Repository
	WineRepository   wine.Repository
	BottleRepository bottle.Repository
}

// Execute executes the service.
func (service CsvImport) Execute(ctx context.Context, reader *bufio.Reader) ([]wine.Wine, error) {
	wines, err := service.readLineByLine(ctx, reader)
	if err != nil {
		return *new([]wine.Wine), errors.New("CsvImport Execute - " + err.Error())
	}
	log.Printf("CsvImport Execute - %d data lines read", len(wines))
	return wines, nil
}

func (service CsvImport) readLineByLine(ctx context.Context, reader *bufio.Reader) ([]wine.Wine, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	linesRead := 0
	var wines []wine.Wine
	for scanner.Scan() {
		line := scanner.Text()
		if linesRead == 0 {
			if !validateData(line, referenceHeaders) {
				return wines, errors.New("Invalid data")
			}
		} else {
			wine := service.persistDataLine(ctx, line)
			wines = append(wines, *wine)
		}
		linesRead++
	}
	return wines, nil
}

func validateData(header string, referenceHeaders []string) bool {
	headers := parseLine(header)
	return utils.EqualsStringSlices(referenceHeaders, headers)
}

func (service CsvImport) persistDataLine(ctx context.Context, line string) *wine.Wine {
	data := parseLine(line)
	cellar := buildHardcodedCellar()
	stringCellarID, _ := service.CellarRepository.SaveCellar(ctx, cellar)
	cellarID, _ := strconv.ParseInt(stringCellarID, 10, 64)

	wine := dataToWine(data)
	stringWineID, _ := service.WineRepository.SaveWine(ctx, wine)
	wineID, _ := strconv.ParseInt(stringWineID, 10, 64)

	bottle := dataToBottle(data, cellarID, wineID)
	service.BottleRepository.SaveBottle(ctx, bottle)
	return wine
}

func parseLine(line string) []string {
	cleanedLine := strings.TrimSuffix(line, ",")
	return strings.Split(cleanedLine, ",")
}

func buildHardcodedCellar() *cellar.Cellar {
	return cellar.NewCellar("Moiré", 1)
}

func dataToWine(data []string) *wine.Wine {
	return &wine.Wine{
		Name:        data[0],
		Designation: data[1],
		Growth:      data[2],
		Country:     "France",
		Region:      data[4],
		Color:       wine.StringToColor(data[5]),
		Type:        wine.StringToType(data[6]),
		Producer:    data[8],
	}
}

func dataToBottle(data []string, cellarID int64, wineID int64) *bottle.Bottle {
	return &bottle.Bottle{
		Year:     utils.StringToInt(data[3]),
		Size:     bottle.IntToSize(utils.StringToInt(data[10])),
		Quantity: utils.StringToInt(data[7]),
		CellarID: cellarID,
		WineID:   wineID,
	}
}

var referenceHeaders = []string{
	"Nom",
	"Appellation",
	"Cru",
	"Millésime",
	"Région",
	"Couleur",
	"Type",
	"Stock",
	"Producteur",
	"Origine",
	"Format",
}
