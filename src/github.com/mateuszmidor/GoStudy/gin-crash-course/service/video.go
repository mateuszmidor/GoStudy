package service

import (
	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"github.com/mateuszmidor/GoStudy/gin-crash-course/repo"
)

type VideoService interface {
	Save(entity.Video) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() ([]entity.Video, error)
}

type videoService struct {
	videos repo.VideoRepo
}

func New(repo repo.VideoRepo) VideoService {
	return &videoService{videos: repo}
}

func (s *videoService) Save(v entity.Video) error {
	return s.videos.Save(v)
}

func (s *videoService) Update(v entity.Video) error {
	return s.videos.Update(v)
}

func (s *videoService) Delete(v entity.Video) error {
	return s.videos.Delete(v)
}

func (s *videoService) FindAll() ([]entity.Video, error) {
	return s.videos.FindAll()
}
