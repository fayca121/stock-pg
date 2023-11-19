package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/fayca121/stock-pg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetStock(c *gin.Context) {
	var stock models.Stock

	reqParamId := c.Param("id")
	stockId, err := strconv.Atoi(reqParamId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest,
			fmt.Errorf("Unable to convert the string into int.  %v", err))
		return
	}

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlStatement, stockId)
	err = row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNoContent, gin.H{"error": "No records found"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
			return
		}
	}

	c.JSON(http.StatusOK, stock)
}

func GetAllStock(c *gin.Context) {
	var stocks []models.Stock

	sqlStatement := `SELECT * FROM stocks`

	db := c.MustGet("db").(*sql.DB)

	rows, err := db.Query(sqlStatement)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
			return
		}

		stocks = append(stocks, stock)
	}

	c.JSON(http.StatusOK, stocks)
}

func CreateStock(c *gin.Context) {
	var stock models.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)

	var response response = response{
		ID:      id,
		Message: "Stock created successfully",
	}

	c.JSON(http.StatusCreated, response)
}

func UpdateStock(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	db := c.MustGet("db").(*sql.DB)
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var stock models.Stock

	if err := c.ShouldBindJSON(&stock); err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	c.JSON(http.StatusCreated, stock)
}

func DeleteStock(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	db := c.MustGet("db").(*sql.DB)

	sqlStatement := "DELETE FROM stocks where stockid=$1"

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    id,
	})

}
