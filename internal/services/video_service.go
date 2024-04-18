package services

import (
	"github.com/kzar1n/alura-challanges-backend-1/internal/infra/database"
	"github.com/kzar1n/alura-challanges-backend-1/internal/models"
	"github.com/kzar1n/alura-challanges-backend-1/internal/repository"
)

func GetAll() ([]models.Video, error) {
	videoRepository := repository.NewVideoRepository(database.DB)
	videos, err := videoRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func GetByID(id string) (models.Video, error) {
	videoRepository := repository.NewVideoRepository(database.DB)
	video, err := videoRepository.GetByID(id)

	if err != nil {
		return models.Video{}, err
	}

	return video, nil
}

func Create(video models.Video) error {
	videoRepository := repository.NewVideoRepository(database.DB)
	err := videoRepository.Create(video)

	if err != nil {
		return err
	}

	return nil
}

func Update(id int, video models.Video) (models.Video, error) {
	video.ID = id
	videoRepository := repository.NewVideoRepository(database.DB)
	err := videoRepository.Update(id, video)

	if err != nil {
		return video, err
	}

	return video, nil
}

func Delete(id string) error {
	videoRepository := repository.NewVideoRepository(database.DB)
	err := videoRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
