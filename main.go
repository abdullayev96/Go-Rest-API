package main

// import (
// 	"net/http"

// 	"time"

// 	"github.com/gin-gonic/gin"
// )



//import "github.com/gin-gonic/gin"


//  Birinchi holat 
//import "fmt"

// func main() {
// 	fmt.Println("Task Manager API")
// }


//  Ikkinchi  holat 
// func main() {
// 	router := gin.Default()
// 	router.GET("/ping", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	router.Run() // Listen and serve on 0.0.0.0:8080
// }

//    Uchinchi holat 



// func main() {
// 	// Router yaratish
// 	r := gin.Default()

// 	// Oddiy GET endpoint
// 	r.GET("/hello", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Salom, Gin API ishlayapti!",
// 		})
// 	})

// 	// Serverni ishga tushirish
// 	r.Run(":8080")
// }


// type Task struct {
// 	ID          string    `json:"id"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	DueDate     time.Time `json:"due_date"`
// 	Status      string    `json:"status"`
// }

// // Mock data for tasks
// var tasks = []Task{
// 	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
// 	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
// 	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
// }


// func Logger() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         // Middleware logic before request
//         c.Next()
//         // Middleware logic after request
//     }
// }




// func main() {
// 	router := gin.Default()

// 	router.GET("/tasks", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
// 	})

// 	router.GET("/tasks/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")

// 		for _, task := range tasks {
// 			if task.ID == id {
// 				ctx.JSON(http.StatusOK, task)
// 				return
// 			}
// 		}

// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// 	})

	

	// router.PUT("/tasks/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")

	// 	var updatedTask Task

	// 	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	for i, task := range tasks {
	// 		if task.ID == id {
	// 			if updatedTask.Title != "" {
	// 				tasks[i].Title = updatedTask.Title
	// 			}
	// 			if updatedTask.Description != "" {
	// 				tasks[i].Description = updatedTask.Description
	// 			}
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
	// 			return
	// 		}
	// 	}

	// 	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	// })

	// router.DELETE("/tasks/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")

	// 	for i, val := range tasks {
	// 		if val.ID == id {
	// 			tasks = append(tasks[:i], tasks[i+1:]...)
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
	// 			return
	// 		}
	// 	}

	// 	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	// })

	// router.POST("/tasks", func(ctx *gin.Context) {
	// 	var newTask Task

	// 	if err := ctx.ShouldBindJSON(&newTask); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	tasks = append(tasks, newTask)
	// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
	// })

// 	router.Run()
// }

//     Author, Category, Book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	docs "Task_manager_API_GO/docs" // shu yerga sizning module nomingiz
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

)



type Category struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"` // 1 category → ko‘p book
}


type Author struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"` // 1 author → ko‘p book
}


type Book struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
	AuthorID   uint   `json:"author_id"`

	Category Category `json:"category"`
	Author   Author   `json:"author"`
}



// ==== DATABASE ====
var db *gorm.DB
var err error

func initDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Databasega ulanib bo‘lmadi!")
	}

	// Jadval yaratish
	db.AutoMigrate(&Category{}, &Author{}, &Book{})
}


// ==== HANDLERLAR ====

//  Categoriya 


// createCategory godoc
// @Summary      Yangi Category yaratish
// @Description  Bazaga yangi Category qo‘shadi
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      Category  true  "Category qo‘shish modeli"
// @Success      201  {object}  Category
// @Failure      400  {object}  map[string]string
// @Router       /categories [post]
func createCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&category)
	c.JSON(http.StatusCreated, category)
}



// getCategories godoc
// @Summary      Barcha kategoriyalarni olish
// @Description  Bazadagi barcha kategoriyalarni chiqaradi
// @Tags         categories
// @Produce      json
// @Success      200  {array}  Category
// @Router       /categories [get]
func getCategories(c *gin.Context) {
	var categories []Category
	db.Preload("Books").Find(&categories)
	c.JSON(http.StatusOK, categories)
}



// getCategories godoc
// @Summary      Barcha kategoriyalarni olish
// @Description  Bazadagi barcha kategoriyalarni chiqaradi
// @Tags         categories
// @Produce      json
// @Success      200  {array}  Category
// @Router       /categories [get]
func getCategoryByID(c *gin.Context) {
    id := c.Param("id") // URL dan id olish

    var category Category
    // GORM bilan bazadan qidirish
    if err := db.Preload("Books").First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, category)
}



// createAuthor godoc
// @Summary      Yangi Author yaratish
// @Description  Bazaga yangi Author qo‘shadi
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        author  body      Author  true  "Author qo‘shish modeli"
// @Success      201  {object}  Author
// @Failure      400  {object}  map[string]string
// @Router       /authors [post]
func createAuthor(c *gin.Context) {
	var author Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&author)
	c.JSON(http.StatusCreated, author)
}


// getAuthors godoc
// @Summary      Barcha Authorlarni olish
// @Description  Bazadagi barcha Authorlarni chiqaradi
// @Tags         authors
// @Produce      json
// @Success      200  {array}  Author
// @Failure      404  {object}  map[string]string
// @Router       /authors [get]
func getAuthors(c *gin.Context) {
	var authors []Author
	db.Preload("Books").Find(&authors)
	c.JSON(http.StatusOK, authors)
}





//  Books

// createBook godoc
// @Summary      Yangi Book yaratish
// @Description  Bazaga yangi Book qo‘shadi (category_id va author_id kerak bo‘ladi)
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      Book  true  "Book qo‘shish modeli"
// @Success      201  {object}  Book
// @Failure      400  {object}  map[string]string
// @Router       /books [post]
func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}


// getBooks godoc
// @Summary      Barcha Booklarni olish
// @Description  Bazadagi barcha Booklarni Category va Author bilan birga chiqaradi
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array}   Book
// @Router       /books [get]
func getBooks(c *gin.Context) {
	var books []Book
	db.Preload("Category").Preload("Author").Find(&books)
	c.JSON(http.StatusOK, books)
}




// ==== MAIN ====
func main() {
	initDB()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Category
	r.POST("/categories", createCategory)
	r.GET("/categories", getCategories)
	r.GET("/categories/:id", getCategoryByID)

	// Author
	r.POST("/authors", createAuthor)
	r.GET("/authors", getAuthors)

	// Book
	r.POST("/books", createBook)
	r.GET("/books", getBooks)

	r.Run(":8080")
}
