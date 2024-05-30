package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"campusCard/model"
	"campusCard/cache"
	"time"
	"context"

	"github.com/google/uuid"
)

type LibraryController struct {
}

var Ctx = context.Background()

func (l LibraryController)ReserveSeat(c *gin.Context) {
	var reservation model.Reservation
	if err := c.BindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation.ID = uuid.New().String()
	reservation.Status = "reserved"
	cache.Rdb.Set(Ctx, reservation.ID, reservation, 0)
	c.JSON(http.StatusOK, gin.H{"message": "Seat reserved", "reservation": reservation})
}

func (l LibraryController)QuickReserveSeat(c *gin.Context) {
	// Implement logic to quickly reserve a seat
}

func (l LibraryController)ScanReserveSeat(c *gin.Context) {
	// Implement logic to reserve a seat by scanning a QR code
}

func (l LibraryController)CheckIn(c *gin.Context) {
	var checkIn model.CheckIn
	if err := c.BindJSON(&checkIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	checkIn.ID = uuid.New().String()
	checkIn.CheckInTime = time.Now().Format(time.RFC3339)
	cache.Rdb.Set(Ctx, checkIn.ID, checkIn, 0)
	c.JSON(http.StatusOK, gin.H{"message": "Checked in", "check_in": checkIn})
}

func (l LibraryController)TemporaryLeave(c *gin.Context) {
	// Implement logic for temporary leave
}

func (l LibraryController)CheckOut(c *gin.Context) {
	// Implement logic for check out
}
