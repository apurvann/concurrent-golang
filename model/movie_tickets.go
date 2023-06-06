package model

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MovieTickets maps to movie_tickets table
type MovieTickets struct {
	MovieID     uuid.UUID
	TicketsLeft int
}

func (mt MovieTickets) Insert(db *gorm.DB, initialTickets int) uuid.UUID {
	mt.MovieID = uuid.New()
	mt.TicketsLeft = initialTickets
	log.Println(mt)

	if err := db.Create(&mt).Error; err != nil {
		log.Fatal(err)
	}

	return mt.MovieID

	//log.Println(mt.TicketsLeft)

}
func (mt MovieTickets) Update(db *gorm.DB, ticketsBought int) {
	var ticketsLeft = mt.TicketsLeft - ticketsBought
	log.Println(mt)
	if err := db.Model(&mt).
		Where("movie_id=?", mt.MovieID).
		Update("tickets_left", ticketsLeft).Error; err != nil {
		log.Fatal(err)
	}

}
func (mt MovieTickets) Delete(db *gorm.DB, id uuid.UUID) {

}
