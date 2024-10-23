package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/szaborics/vacations-api/database"
	"github.com/szaborics/vacations-api/models"
)



func main() {
	database.MongodbConnect()

}
