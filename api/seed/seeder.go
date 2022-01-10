package seed

import (
	"log"

	"github.com/dkantikorn/go-fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Sarawutt Bureekeaw",
		Email:    "sarawutt.b@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Peter Parker",
		Email:    "pparker@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Robert Downey Jr",
		Email:    "robert.d@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Golang",
		Content: "Golang content",
	},
	models.Post{
		Title:   "Java",
		Content: "Java Content",
	},
	models.Post{
		Title:   "Java Spring Boot",
		Content: "Java Spring Boot Content",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
