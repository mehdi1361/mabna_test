package main

import (
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"mabna_test/models"
	"os"
	"sync"
	"time"
	"strconv"
	"math/rand"
)

func createRandomTrade(quantity int, instrumentName string) (bool, error) {
	var wg sync.WaitGroup
	db, err := models.Connect()
	if err != nil {
		return false, errors.New("can not connect to database")
	}
	var instrument models.Instrument

	_ = db.First(&instrument, "name=?", instrumentName)
	fmt.Println(instrument.ID)

	for i := 0; i < quantity; i++ {
		wg.Add(1)
		go func(da int) {
			defer wg.Done()
			var data = models.Trade{
				DateN:        time.Now(),
				Open:         int32(rand.Intn(1100 - 1000) + 1000),
				High:         int32(rand.Intn(4100 - 4000) + 4000),
				Low:          int32(rand.Intn(350 - 300) + 300),
				Close:        int32(rand.Intn(450 - 400) + 400),
				InstrumentId: instrument.ID,
			}
			db.Create(&data)
			fmt.Printf("test: %d\n", da)
		}(i)
	}
	wg.Wait()
	return true, errors.New("bad")
}

func listinstrument() {
	db, err := models.Connect()
	if err != nil {
		fmt.Println("can not connect to database")
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
		},
	}

	var lstIntruments []models.Instrument
	_ = db.Find(&lstIntruments)

	for _, row := range lstIntruments {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.ID)},
			{Align: simpletable.AlignRight, Text: row.Name},
		}

		table.Body.Cells = append(table.Body.Cells, r)

	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func listTrade(insName string) {
	db, err := models.Connect()
	if err != nil {
		fmt.Println("can not connect to database")
	}
	var instrument models.Instrument

	_ = db.First(&instrument, "name=?", insName)

	var lstTrade []models.Trade

	_ = db.Find(&lstTrade, "instrument_id = ? ", instrument.ID)

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "DATE"},
			{Align: simpletable.AlignCenter, Text: "OPEN"},
			{Align: simpletable.AlignCenter, Text: "HIGH"},
			{Align: simpletable.AlignCenter, Text: "LOW"},
			{Align: simpletable.AlignCenter, Text: "CLOSE"},
		},
	}

	for _, row := range lstTrade {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", row.DateN.Format("2006-01-01"))},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.Open)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.High)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.Low)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.Close)},
		}

		table.Body.Cells = append(table.Body.Cells, r)

	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func createInstrunment(name string) {
	db, err := models.Connect()
	if err != nil {
		fmt.Println("can not connect to database")
	}
	instrument := models.Instrument{Name: name}
	db.Create(&instrument)
}

func main() {
	args := os.Args
	switch {
	case args[1] == "?" || args[1] == "help":
		fmt.Println("help")

	case len(args) == 3 && args[1] == "show" && args[2] == "instrument":
		listinstrument()

	case len(args) == 4 && args[1] == "show" && args[2] == "trade":
		listTrade(args[3])

	case len(args)==4 && args[1] == "create" && args[2] == "instrument":
		createInstrunment(args[3])

	case len(args)==5 && args[1] == "create" && args[2] == "trade":
		quantity, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println("can't convert to integer")
			os.Exit(1)
		}
		_, _ = createRandomTrade(quantity, args[3])
	default:
		fmt.Println("error can find comand")
	}
}
