package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/tnvmadhav/ExerciseDiary/internal/db"
	"github.com/tnvmadhav/ExerciseDiary/internal/models"
)

func setHandler(c *gin.Context) {

	var formData []models.Set
	var oneSet models.Set
	var weight, reps int

	_ = c.PostFormMap("sets")

	formMap := c.Request.PostForm
	// log.Println("MAP:", formMap)

	len := len(formMap["name"])
	// log.Println("LEN:", len)
	date := formMap["date"][0]

	for i := 0; i < len; i++ {
		oneSet.Date = date
		oneSet.Name = formMap["name"][i]
		weight, _ = strconv.Atoi(formMap["weight"][i])
		reps, _ = strconv.Atoi(formMap["reps"][i])
		oneSet.Weight = weight
		oneSet.Reps = reps

		formData = append(formData, oneSet)
	}

	db.BulkDeleteSetsByDate(appConfig.DBPath, date)
	db.BulkAddSets(appConfig.DBPath, formData)
	exData.Sets = db.SelectSet(appConfig.DBPath)

	// log.Println("FORM DATA:", formData)

	c.Redirect(http.StatusFound, "/")
}
