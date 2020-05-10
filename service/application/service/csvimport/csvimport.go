package csvimport

import (
	"bufio"
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/lapostoj/winemanager/service/application/service/createbottle"
	"github.com/lapostoj/winemanager/service/application/service/createcellar"
	"github.com/lapostoj/winemanager/service/application/service/createwine"
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
	CreateCellar createcellar.CreateCellarService
	CreateBottle createbottle.CreateBottleService
	CreateWine   createwine.CreateWineService
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
	var cellarID int64

	for scanner.Scan() {
		line := scanner.Text()
		if linesRead == 0 {
			if !validateHeaders(line, referenceHeaders) {
				return wines, errors.New("Invalid data")
			}
			cellar := buildHardcodedCellar()
			stringCellarID, _ := service.CreateCellar.Execute(ctx, cellar)
			log.Printf("StringCellarId: %s", stringCellarID)
			cellarID, _ = strconv.ParseInt(stringCellarID, 10, 64)
			log.Printf("CellarID: %q", cellarID)
		} else {
			wine := service.persistDataLine(ctx, line, cellarID)
			wines = append(wines, *wine)
		}
		linesRead++
	}
	return wines, nil
}

func (service CsvImport) persistDataLine(ctx context.Context, line string, cellarID int64) *wine.Wine {
	data := parseLine(line)

	wine := dataToWine(data)
	stringWineID, _ := service.CreateWine.Execute(ctx, wine)
	wineID, _ := strconv.ParseInt(stringWineID, 10, 64)
	log.Printf("wineID: %q", wineID)

	bottle := dataToBottle(data, cellarID, wineID)
	service.CreateBottle.Execute(ctx, bottle)
	return wine
}

func validateHeaders(line string, referenceHeaders []string) bool {
	headers := parseLine(line)
	return utils.EqualsStringSlices(referenceHeaders, headers)
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
