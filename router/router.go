package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mabna_test/models"
	"time"
	//	"github.com/gin-gonic/gin/binding"
)

func PingHandler(c *gin.Context) {
	db, err := models.Connect()
	if err != nil {
		c.JSON(200, gin.H{"message": fmt.Sprintf("error in connect to database %s", err)})
		return
	}

	var instruments []models.Instrument
	_ = db.Preload("Trades").Find(&instruments)

	var instrument models.Instrument

	_ = db.First(&instrument, "id=?", 1)
	fmt.Println(instrument.ID)

	var data = models.Trade{
		DateN: time.Now(),
		Open:  1,
		High:  1,
		Low:   1,
		Close: 1,
		InstrumentId: instrument.ID,
	}
	db.Create(&data)

	c.JSON(200, gin.H{
		"message": instruments,
	})

}
