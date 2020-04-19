package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	user struct {
		ID           int    `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		Email        string `json:"email,omitempty"`
		Password     string `json:"password,omitempty"`
		MobileNumber string `json:"mobile_number,omitempty"`
	}

	cabs struct {
		ID           int    `json:"id,omitempty"`
		CabNumber    string `json:"cab_no,omitempty"`
		DriverName   string `json:"driver_name,omitempty"`
		DriverNumber string `json:"driver_number,omitempty"`
		Status       string `json:"status,omitempty"`
	}

	bookings struct {
		ID       int       `json:"id,omitempty"`
		UserID   int       `json:"user_id,omitempty"`
		CabID    int       `json:"cab_id,omitempty"`
		FromLoc  string    `json:"from_loc,omitempty"`
		ToLoc    string    `json:"to_loc,omitempty"`
		BookedAt time.Time `json:"booked_at,omitempty"`
		Status   string    `json:"status,omitempty"`
	}
)

var db *gorm.DB

func init() {

	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/booking")
	{
		v1.POST("/createUser", createUser)
		v1.GET("/cabs", getNearByCabs)
		v1.POST("/book", bookCab)
		v1.GET("/getBookings/:user_id", getPastBookings)
	}
	return router
}

// createUser create new user
func createUser(c *gin.Context) {

	body := c.Request.Body
	decoder := json.NewDecoder(body)
	var param user
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		fmt.Println(decodeErr.Error())
		return
	}

	db.Save(&param)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "mobile_number": param.MobileNumber, "email": param.Email, "name": param.Name, "user_Id": param.ID})
}

// getNearbyCabs get cabs near by user location
func getNearByCabs(c *gin.Context) {

	var cabDetails []cabs
	var cabDetailsRet []cabs

	findLocation := c.Query("location")

	if findLocation == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Please provide current location"})
		return
	}
	db.Where("current_loc = ?", findLocation).Find(&cabDetails)

	for _, item := range cabDetails {
		if item.Status == "available" {
			cabDetailsRet = append(cabDetailsRet, cabs{ID: item.ID, CabNumber: item.CabNumber, DriverName: item.DriverName, DriverNumber: item.DriverNumber})
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": cabDetailsRet})

}

// bookCab from point A to point B
func bookCab(c *gin.Context) {

	body := c.Request.Body
	decoder := json.NewDecoder(body)
	var param bookings
	var cabDetails cabs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		fmt.Println(decodeErr.Error())
		return
	}

	if param.CabID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "message": "Please provide valid cab Id!"})
		return
	}

	db.Where("id = ?", param.CabID).Find(&cabDetails)

	if cabDetails.DriverName == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "message": "sdasdasPlease provide valid cab Id!"})
		return
	}

	booking := bookings{UserID: param.UserID, CabID: param.CabID, BookedAt: time.Now(), FromLoc: param.FromLoc, ToLoc: param.ToLoc, Status: "booked"}
	db.Save(&booking)

	db.Model(&cabDetails).Update("status", "booked")

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Cab booked successfully!", "driver_number": cabDetails.DriverNumber, "driver_name": cabDetails.DriverName, "cab_number": cabDetails.CabNumber, "bookingId": booking.ID})

}

// getPastBookings get all past booking for user
func getPastBookings(c *gin.Context) {

	var bookingDetails []bookings
	var bookingReturn []bookings

	userID := c.Param("user_id")

	if userID == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Please provide user ID"})
		return
	}

	db.Where("user_id = ?", userID).Find(&bookingDetails)

	for _, item := range bookingDetails {
		bookingReturn = append(bookingReturn, bookings{ID: item.ID, FromLoc: item.FromLoc, ToLoc: item.ToLoc, BookedAt: item.BookedAt})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": bookingReturn})

}
