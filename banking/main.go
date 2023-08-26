package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type BankAccount struct {
	Name    string
	Balance float64
}

var accounts map[string]BankAccount

func main() {
	accounts = make(map[string]BankAccount)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.POST("/create", func(c *gin.Context) {
		name := c.PostForm("name")
		balance := c.PostForm("balance")
		accounts[name] = BankAccount{Name: name, Balance: parseBalance(balance)}
		c.Redirect(302, "/")
	})

	router.POST("/deposit", func(c *gin.Context) {
		name := c.PostForm("account")
		amount := parseBalance(c.PostForm("amount"))
		if acc, ok := accounts[name]; ok {
			acc.Balance += amount
			accounts[name] = acc
			c.Redirect(302, "/")
		} else {
			c.String(404, "Account not found")
		}
	})

	router.POST("/withdraw", func(c *gin.Context) {
		name := c.PostForm("account")
		amount := parseBalance(c.PostForm("amount"))
		if acc, ok := accounts[name]; ok {
			if acc.Balance >= amount {
				acc.Balance -= amount
				accounts[name] = acc
				c.Redirect(302, "/")
			} else {
				c.String(400, "Insufficient funds")
			}
		} else {
			c.String(404, "Account not found")
		}
	})

	router.GET("/balance/:name", func(c *gin.Context) {
		name := c.Param("name")
		if acc, ok := accounts[name]; ok {
			c.String(200, fmt.Sprintf("Balance for %s: %.2f", acc.Name, acc.Balance))
		} else {
			c.String(404, "Account not found")
		}
	})

	router.Run()
}

func parseBalance(balance string) float64 {
	var result float64
	fmt.Sscanf(balance, "%f", &result)
	return result
}
