package seeder

import (
	"awesome-content/config"
	"awesome-content/internal/adapters/driven/repository/model/mysql"
	"awesome-content/internal/database"
	"fmt"

	"github.com/jaswdr/faker/v2"
)

type ContentSeeder struct {
	Title string
	Rate  int
	Text  string
}

func (s *ContentSeeder) TableName() string {
	return "contents"
}

func SeedContent(appConfig *config.AppConfig) {
	fake := faker.New()
	seeder := ContentSeeder{}

	seeder.Title = fake.Lorem().Word()
	seeder.Text = fake.Lorem().Sentence(10)
	seeder.Rate = fake.RandomNumber(1)

	content := mysql.ContentModel{
		Title: seeder.Title,
		Rate:  seeder.Rate,
		Text:  seeder.Text,
	}

	db, err := database.InitDbClient(appConfig)
	if err != nil {
		panic("Could not open connection to the database")
	}

	db.Create(&content)

	fmt.Printf("Record created: %+v", seeder)
}
