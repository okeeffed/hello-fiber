package main

import (
	"github.com/gofiber/fiber"
	"log"
	"strings"
)

// Person field names should start with an uppercase letter
type Person struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Age uint8 `json:"age" xml:"age" form:"age" query:"age"`
}

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello, World 👋!")
	})
	
	app.Post("/first", func(c *fiber.Ctx) {
		// Get raw body from POST request:
		body := c.Body() // user=john
		c.Send(body)
	})

	app.Post("/person", func(c *fiber.Ctx) {
		// Create new person p
		p := new(Person)

		// Bind data to p or log error
		if err := c.BodyParser(p); err != nil {
				log.Println(err)
				c.Status(500).Send("Failed")
				return
		}

		log.Println(p.Name) 
		log.Println(p.Age) 

		c.Send("Success")
	})

	app.Post("/json", func(c *fiber.Ctx) {
		// Create new person p
		p := new(Person)
	
		// Bind data to p or log error
		if err := c.BodyParser(p); err != nil {
			log.Println(err)
			c.Status(500).Send("Failed")
			return
		}
	
		// Create data struct:
		data := Person{
			Name: strings.ToUpper(p.Name),
			Age:  p.Age + 10,
		}
	
		if err := c.JSON(data); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

  app.Listen(3000)
}