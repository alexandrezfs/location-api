package controllers

import "github.com/revel/revel"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "os"
import "encoding/json"
import "io/ioutil"

type App struct {
	*revel.Controller
}

type Example struct {
    text  []string
 }

func (c App) Index() revel.Result {

	db, err := sql.Open("mysql", "root:kazuki69@/golang")

    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
    if err != nil {
    	fmt.Printf("%s", err.Error());
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    // Prepare statement for reading data
    stmtOut, err := db.Prepare("SELECT squareNumber FROM squareNum WHERE number = ?")
    if err != nil {
    	fmt.Printf("%s", err.Error());
    }
    defer stmtOut.Close()

    // Insert square numbers for 0-24 in the database
    for i := 0; i < 25; i++ {
        _, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
        if err != nil {
    		fmt.Printf("%s", err.Error());
        }
    }

    var squareNum int // we "scan" the result in here

    // Query the square-number of 13
    err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
    if err != nil {
    	fmt.Printf("%s", err.Error());
    }
    fmt.Printf("The square number of 13 is: %d", squareNum)

    // Query another number.. 1 maybe?
    err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
    if err != nil {
    	fmt.Printf("%s", err.Error());
    }
    fmt.Printf("The square number of 1 is: %d", squareNum)

	greeting := "test json"


    type jsonobject struct {
    Object ObjectType
}
 
type ObjectType struct {
    Buffer_size int
    Databases   []DatabasesType
}
 
type DatabasesType struct {
    Host   string
    User   string
    Pass   string
    Type   string
    Name   string
    Tables []TablesType
}
 
type TablesType struct {
    Name     string
    Statment string
    Regex    string
    Types    []TypesType
}
 
type TypesType struct {
    Id    string
    Value string
}
 
// Main function
// I realize this function is much too simple I am simply at a loss to
 
func main() {
    file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
 
    //m := new(Dispatch)
    //var m interface{}
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)
} 


	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 10).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
	
}

func (c App) TestJson() revel.Result {


    type jsonobject struct {
        Object ObjectType
    }

    type ObjectType struct {
        Names   []NamesType
    }

    type NamesType struct {
        Firstname   string
        Lastname   string
    }

	//TODO

	return c.Render(App.Index)
}