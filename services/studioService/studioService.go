package studioService

import "FinalProject/models"

type StudioService interface {
	AddStudio(studio *models.AddStudio) (*models.Studio, error)
	GetStudio() (*[]models.ResultStudioById, error)
	GetStudioById(id int) (*models.ResultStudioById, error)
	UpdateStudio(id int, studio *models.Studio) (int, error)
	DeleteStudio(id int) (int, error)
}