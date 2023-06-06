package main

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example.com/concurrent-golang/model"
)

func main() {

	log.Println("Starting")

	DATABASE_URL := "postgresql://root@localhost:26257?sslmode=disable"

	db, err := gorm.Open(postgres.Open(DATABASE_URL+"&application_name=movies"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Automatically create the "movie_box_offices" table based on the `MovieBoxoffice`
	// model.
	db.AutoMigrate(&model.MovieBoxOffice{})
	db.AutoMigrate(&model.MovieTickets{})

	var ids []uuid.UUID

	var m model.Movie[int] = &model.MovieTickets{MovieID: uuid.Nil, TicketsLeft: 0}

	var id = m.Insert(db, 1000)
	ids = append(ids, id)

	var movieTicket model.MovieTickets
	db.Where("movie_id=?", id).First(&movieTicket) //.MovieID is movie_id field automatically in db

	var mm model.Movie[int] = &movieTicket
	mm.Update(db, 10)

	/* NUMROWS := 10000

	start_time_blocking := time.Now()
	service.GenerateMovies(db, NUMROWS)
	blocking_duration := time.Since(start_time_blocking)
	log.Printf("Blocking time: %v", blocking_duration.Seconds())

	BATCH_SIZE := 100
	BATCHES := NUMROWS / BATCH_SIZE
	log.Printf("Batches %d", BATCHES)
	var wg sync.WaitGroup
	wg.Add(BATCHES)
	start_time_non_blocking := time.Now()
	for batch := 1; batch <= BATCHES; batch++ {

		go service.GenerateMoviesNonBlocking(db, BATCH_SIZE, &wg)

	}
	wg.Wait()
	non_blocking_duration := time.Since(start_time_non_blocking)
	log.Printf("Non Blocking time: %v", non_blocking_duration.Seconds()) */
}
