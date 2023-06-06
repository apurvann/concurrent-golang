package repo

import (
	"encoding/json"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/xyproto/randomstring"
	"gorm.io/gorm"

	model "example.com/concurrent-golang/model"
)

func InsertMovies(db *gorm.DB, numRows int) error {
	log.Printf("Creating %d new accounts...", numRows)
	for i := 0; i < numRows; i++ {
		randomID := uuid.New()
		randomMovieName := randomstring.HumanFriendlyEnglishString(rand.Intn(20))
		randomCollectionDollars := rand.Float32()
		//randomTicketsLeft := rand.Intn(100)
		if err := db.Create(
			&model.MovieBoxOffice{ID: randomID, Name: randomMovieName, Collection: randomCollectionDollars}).Error; err != nil {
			return err
		}
	}
	log.Println("Movies created.")
	return nil
}

// Get movie object
func getMovie(db *gorm.DB, movieName string) {
	var movieToFind model.MovieBoxOffice
	db.Where("Name = ?", movieName).Find(&movieToFind)
	b, err := json.Marshal(movieToFind)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}

// Delete specific movies
func DeleteMovies(db *gorm.DB, movieNames []string) error {

	err := db.Where("Name IN ?", movieNames).Delete(model.MovieBoxOffice{}).Error
	if err != nil {
		return err
	}
	log.Println("Movies deleted.")
	return nil
}
