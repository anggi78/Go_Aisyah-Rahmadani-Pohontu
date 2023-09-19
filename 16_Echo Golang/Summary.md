# Echo Golang

Perpustakaan pihak ketiga golang:
- echo
- Go Kit
- Logrus
- gorm.io
- cobra

Echo adalah sebuah framework web yang ringan dan berkinerja tinggi yang digunakan untuk membangun aplikasi web dengan Go. Framework ini dirancang untuk kecepatan dan kegunaan yang maksimal, dan biasanya digunakan untuk membuat aplikasi web RESTful.

Keuntungan menggunakan echo:
1. Router yang dioptimalkan
2. Rendering data
3. Data binding
4. Dapat diskalakan
5. Middleware

### Basic Routing & Controller
import (

	"net/http"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)

	e.Start(":8000")
}

func UserController(e echo.Context) error {

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user" : user,
		"search" : search,
		"short" : short,
	})
}

### Data Rendering
import (

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	
	Id int `json: "id" form: "id"`
	Age int `json: "age" form: "age"`
	Email string `json: "email" form: "email"`
	Name string `json: "name" form: "name"`
}

func main() {

	e := echo.New()

	//routing
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)

	e.Start(":8000")
}

func UserController(e echo.Context) error {

	user := User{Id: id, Age: age, Email: "anggikiyowo@gmail.com", Name: "Anggi Cantik"}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user" : user,
		"search" : search,
		"short" : short,
	})
}

### Retrieve Data
type User struct {

	Id int `json: "id" form: "id"`
	Age int `json: "age" form: "age"`
	Email string `json: "email" form: "email"`
	Name string `json: "name" form: "name"`
}

func main() {
	
	e := echo.New()

	//routing
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)

	e.Start(":8000")
}

func UserController(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	user := User{Id: id, Age: age, Email: "anggikiyowo@gmail.com", Name: "Anggi Cantik"}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user" : user,
		"search" : search,
		"short" : short,
	})
}

### Binding Data
func RegisterController(c echo.Context) error {
	
	// email := c.FormValue("email")
	// name := c.FormValue("name")

	user := User{}

	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user" : user,
	})
}