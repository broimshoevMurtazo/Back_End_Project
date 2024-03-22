package controlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"murtazo/app/cockie"
	"murtazo/app/forreturn"
	"murtazo/app/structs"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var UserShablon structs.Create
	c.ShouldBindJSON(&UserShablon)
	if UserShablon.Name == "" || UserShablon.Surname == "" || UserShablon.Login == "" || UserShablon.Password == "" {
		c.JSON(404, "error empty field")
	} else {
		client, ctx := forreturn.DBConnection()

		DBConnect := client.Database("ClassRoom").Collection("users")

		var found = DBConnect.FindOne(ctx, bson.M{
			"login" : UserShablon.Login,
		})
		var Loginisexxsist structs.Create
		found.Decode(&Loginisexxsist)
		if Loginisexxsist.Login != "" {
			c.JSON(404, "email is exsist")
		} else {

			id := primitive.NewObjectID().Hex()
			Hashed, _ := forreturn.HashPassword(UserShablon.Password)

			DBConnect.InsertOne(ctx, bson.M{
				"_id":      id,
				"name":     UserShablon.Name,
				"surname":  UserShablon.Surname,
				"login":    UserShablon.Login,
				"password": Hashed,
			})
			c.JSON(200,"succes")
		}
	}
}
func Login(c *gin.Context) {
	var LoginTemp structs.Create
	c.ShouldBindJSON(&LoginTemp)

	if LoginTemp.Login == "" || LoginTemp.Password == "" {
		c.JSON(404, "Error empty fild")
	} else {
		client, ctx := forreturn.DBConnection()

		DBConnect := client.Database("ClassRoom").Collection("users")

		result := DBConnect.FindOne(ctx, bson.M{
			"login": LoginTemp.Login,
		})

		var userdata structs.Create
		result.Decode(&userdata)
		isValidPass := forreturn.CompareHashPasswords(userdata.Password, LoginTemp.Password)
		fmt.Println(isValidPass)

		if isValidPass {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    "Classcockie",
				Value:   userdata.Id,
				Expires: time.Now().Add(60 * time.Minute),
			})
			c.JSON(200, "success")
		} else {
			c.JSON(404, "Wrong login or password")
		}
	}
}

func Addstudent(c *gin.Context) {
	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST == true {
		var usertempt structs.Student
		c.ShouldBindJSON(&usertempt)

		if usertempt.StudentName == "" || usertempt.StudentSurname == "" || usertempt.StudentEmail == "" || usertempt.StudentPhone == 0 || usertempt.StudentLogin == "" || usertempt.StudentPassword == "" {
			c.JSON(404, "error")
		} else {
			client, ctx := forreturn.DBConnection()

			var createDB = client.Database("ClassRoom").Collection("student")

			var found = createDB.FindOne(ctx, bson.M{
				"login" : usertempt.StudentEmail,
			})
			var Loginisexxsist structs.Student
			found.Decode(&Loginisexxsist)
			if Loginisexxsist.StudentEmail != "" {
				c.JSON(404, "email is exsist")
			} else {
				Hashed, _ := forreturn.HashPassword(usertempt.StudentPassword)
				ID := primitive.NewObjectID().Hex()
				var insertrezult, inserterror = createDB.InsertOne(ctx, bson.M{
					"_id":             ID,
					"studentname":     usertempt.StudentName,
					"studentsurname":  usertempt.StudentSurname,
					"studentemail":    usertempt.StudentEmail,
					"studentphone":    usertempt.StudentPhone,
					"studentlogin":    usertempt.StudentLogin,
					"studentpassword": Hashed,
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}

		}
	}

}

func DeleteStudent(c *gin.Context) {

	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST {
		var usertempt structs.DeleteStudent
		c.ShouldBindJSON(&usertempt)

		connect, ctx := forreturn.DBConnection()

		var createDB = connect.Database("ClassRoom").Collection("student")
		var deletrezult, deleteerror = createDB.DeleteOne(ctx, bson.M{
			"_id": usertempt.Student_id,
		})
		if deleteerror != nil {
			fmt.Printf("deleteerror: %v\n", deleteerror)
		}
		if deletrezult.DeletedCount == 1 {
			c.JSON(200, "succes")
			fmt.Printf("deletrezult: %v\n", deletrezult)
		} else {
			c.JSON(404, "error")
		}

	}
}

func Addproject(c *gin.Context) {
	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST == true {
		var usertempt structs.Forproject
		c.ShouldBindJSON(&usertempt)

		if usertempt.Project_name == "" || usertempt.Discription == "" {
			c.JSON(404, "error1")
		} else {
			client, ctx := forreturn.DBConnection()
			var createDB = client.Database("ClassRoom").Collection("project")
			var found = createDB.FindOne(ctx, usertempt)
			var projectisexsist structs.Forproject
			found.Decode(&projectisexsist)
			if projectisexsist.Project_name == usertempt.Discription {
				c.JSON(404, "project is exsist")
			} else {
				Id := primitive.NewObjectID().Hex()

				var insertrezult, inserterror = createDB.InsertOne(ctx, bson.M{
					"_id":          Id,
					"project_name": usertempt.Project_name,
					"discription":  usertempt.Discription,
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}

		}
	}

}

func DeleteProject(c *gin.Context) {
	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST == true {
		var usertempt structs.DeleteProject
		c.ShouldBindJSON(&usertempt)
		client, ctx := forreturn.DBConnection()
		var createDB = client.Database("ClassRoom").Collection("project")
		deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
			"_id": usertempt.Project_id,
		})
		if deleteerror != nil {
			fmt.Printf("deleteerror: %v\n", deleteerror)
		}
		if deletrezult.DeletedCount == 1 {
			c.JSON(200, "succes")
			fmt.Printf("deletrezult: %v\n", deletrezult)
		} else {
			c.JSON(404, "error")
		}

		client, ctx = forreturn.DBConnection()
		var createDB1 = client.Database("ClassRoom").Collection("JoinStudent")
		deletrezult, deleteerror = createDB1.DeleteMany(ctx, bson.M{

			"project_id": usertempt.Project_id,
		})
		if deleteerror != nil {
			fmt.Printf("deleteerror: %v\n", deleteerror)
		}
		if deletrezult.DeletedCount == 1 {
			c.JSON(200, "succes")
			fmt.Printf("deletrezult: %v\n", deletrezult)
		} else {
			c.JSON(404, "error1")
		}
	}

}
func JoinStudent(c *gin.Context) {
	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST {
		var usertempt structs.JoinStudent
		c.ShouldBindJSON(&usertempt)

		if usertempt.Student_email == "" || usertempt.Project_id == "" {
			c.JSON(404, "error2")
		} else {
			client, ctx := forreturn.DBConnection()

			ConnectProject := client.Database("ClassRoom").Collection("project")
			ConnectStudent := client.Database("ClassRoom").Collection("student")

			result := ConnectProject.FindOne(ctx, bson.M{
				"_id": usertempt.Project_id,
			})
			student := ConnectStudent.FindOne(ctx, bson.M{
				"studentemail": usertempt.Student_email,
			})

			var DBTemp structs.Forproject
			result.Decode(&DBTemp)

			var StudentTemp structs.Student
			student.Decode(&StudentTemp)

			if DBTemp.Id == "" || StudentTemp.StudentEmail == "" {
				c.JSON(404, "problem")
			} else {
				joinstudent, ctx := forreturn.DBConnection()
				JoinOne := joinstudent.Database("ClassRoom").Collection("JoinStudent")

				// result, _ :=
				JoinOne.InsertOne(ctx, bson.M{
					"_id":           primitive.NewObjectID().Hex(),
					"project_id":    DBTemp.Id,
					"student_email": StudentTemp.StudentEmail,
					"owner_id":      primitive.NewObjectID().Hex(),
					"student_phone": StudentTemp.StudentPhone,
				})
				c.JSON(200,"succes")

			}
		}

	}

}

func JoinesList(c *gin.Context) {
	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST  {

		var Forlist = []structs.Forproject{}

		connect,ctx := forreturn.DBConnection()

		var createDB = connect.Database("ClassRoom").Collection("project")

		var singlerezult1, singerror = createDB.Find(ctx, bson.M{})
		if singerror != nil {
			fmt.Printf("singerror: %v\n", singerror)
		}

		for singlerezult1.Next(ctx) {
			var datafromdb structs.Forproject
			fmt.Printf("datafromdb: %v\n", datafromdb)
			singlerezult1.Decode(&datafromdb)

			Forlist = append(Forlist, datafromdb)
		}
		c.JSON(200, Forlist)
	}

}

func Projectlist(c *gin.Context) {
	var iSEXSIST = cockie.ForGin(c)
	if iSEXSIST {
		var Forlist structs.OneProjectList

		var shablon structs.Forproject
		c.ShouldBindJSON(&shablon)

		connect, ctx := forreturn.DBConnection()

		var createDB = connect.Database("ClassRoom").Collection("project")

		var singlerezult1 = createDB.FindOne(ctx, bson.M{
			"_id": shablon.Id,
		})
		singlerezult1.Decode(&Forlist.Project)
		fmt.Printf("singlerezult1: %v\n", singlerezult1)

		var createDB1 = connect.Database("ClassRoom").Collection("JoinStudent")

		var ManyRezult, _ = createDB1.Find(ctx, bson.M{
			"project_id": shablon.Id,
		})
		fmt.Printf("ManyRezult: %v\n", ManyRezult)
		for ManyRezult.Next(ctx) {
			var datafromdb structs.JoinStudent
			ManyRezult.Decode(&datafromdb)

			Forlist.Join_students = append(Forlist.Join_students, datafromdb)
		}
		c.JSON(200, Forlist)
	}

}

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.43.45:5500")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	}

	c.Next()
}
