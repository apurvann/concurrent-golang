package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"gorm.io/gorm"

	"example.com/concurrent-golang/repo"
)

func GenerateMovies(db *gorm.DB, NUMROWS int) {

	//log.Println("Generating Random movies")
	if err := crdbgorm.ExecuteTx(context.Background(), db, nil,
		func(tx *gorm.DB) error {
			return repo.InsertMovies(db, NUMROWS)
		},
	); err != nil {
		fmt.Println(err)
	}

	//var movieNames []string
	//movieNames = append(movieNames, )
	// To handle potential transaction retry errors, we wrap the call
	// to `deleteAccounts` in `crdbgorm.ExecuteTx`
	//if err := crdbgorm.ExecuteTx(context.Background(), db, nil,
	//	func(tx *gorm.DB) error {
	//		return repo.DeleteMovies(db, movieNames)
	//	},
	//); err != nil {
	//	fmt.Println(err)
	//}

}

func GenerateMoviesNonBlocking(db *gorm.DB, NUMROWS int, wg *sync.WaitGroup) {

	//log.Println("Generating Random movies NonBlocking")
	if err := crdbgorm.ExecuteTx(context.Background(), db, nil,
		func(tx *gorm.DB) error {
			return repo.InsertMovies(db, NUMROWS)
		},
	); err != nil {
		fmt.Println(err)
	}
	defer wg.Done()
}
