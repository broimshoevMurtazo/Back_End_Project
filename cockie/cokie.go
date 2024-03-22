package cockie

import (
	"murtazo/app/forreturn"
	"murtazo/app/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ForGin(c *gin.Context) bool {
	var proverka bool = false
	var CookieData, CookieError = c.Request.Cookie("Classcockie")
	if CookieError != nil {

		c.JSON(404, "Error Not Cookie found")
	}	
	cannect,ctx := forreturn.DBConnection()
	var createDB1 = cannect.Database("ClassRoom").Collection("users")
	var singlerezult = createDB1.FindOne(ctx, bson.M{
		"_id": CookieData.Value,
	})
	var datafromdb structs.Create
	singlerezult.Decode(&datafromdb)
	if datafromdb.Id !=""{
		proverka = true
	}
	return proverka
}