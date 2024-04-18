package repository

import (
	"database/sql"
	"log"

	"github.com/kzar1n/alura-challanges-backend-1/internal/models"
)

type VideoRepository struct {
	DB *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{
		DB: db,
	}
}

func (v *VideoRepository) GetAll() ([]models.Video, error) {
	rows, err := v.DB.Query("SELECT * FROM videos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var videos []models.Video

	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.URL)
		if err != nil {
			log.Fatal(err)
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (v *VideoRepository) GetByID(id string) (models.Video, error) {
	var video models.Video
	err := v.DB.QueryRow("SELECT * FROM videos WHERE id = ?", id).Scan(&video.ID, &video.Title, &video.Description, &video.URL)
	if err != nil {
		return models.Video{}, err
	}

	return video, nil
}

func (v *VideoRepository) Create(video models.Video) error {
	_, err := v.DB.Exec("INSERT INTO videos (id, title, description, url) VALUES (?, ?, ?, ?)", video.ID, video.Title, video.Description, video.URL)
	if err != nil {
		return err
	}

	return nil
}

func (v *VideoRepository) Update(id int, video models.Video) error {
	_, err := v.DB.Exec("UPDATE videos SET title = ?, description = ?, url = ? WHERE id = ?", video.Title, video.Description, video.URL, video.ID)
	if err != nil {
		return err
	}

	return nil
}

func (v *VideoRepository) Delete(id string) error {
	_, err := v.DB.Exec("DELETE FROM videos WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
