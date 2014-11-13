package controllers

import "github.com/revel/revel"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

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


    var arr = []Example {
        {{"a", "b", "c"}},
    }
    fmt.Println(arr)    


	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
	
}

func (c App) TestJson() revel.Result {

	//TODO

	return c.Render(App.Index)
}