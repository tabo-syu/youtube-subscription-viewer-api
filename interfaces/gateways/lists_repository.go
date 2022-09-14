package gateways

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ListsRepository struct {
	sql     interfaces.SqlHandler
	youtube interfaces.YoutubeHandler
}

var _ ports.ListsRepository = (*ListsRepository)(nil)

func NewListsRepository(s interfaces.SqlHandler, y interfaces.YoutubeHandler) *ListsRepository {
	return &ListsRepository{s, y}
}

func (r *ListsRepository) Create(list *entities.List) (*entities.List, error) {
	return &entities.List{}, nil
}

func (r *ListsRepository) GetAll() ([]*entities.List, error) {
	return []*entities.List{}, nil
}

func (r *ListsRepository) GetById(id entities.ListId) (*entities.List, error) {
	return &entities.List{}, nil
}

func (r *ListsRepository) Update(list *entities.List) (*entities.List, error) {
	return &entities.List{}, nil
}

func (r *ListsRepository) DeleteById(id entities.ListId) error {
	return nil
}

func (r *ListsRepository) AddChannel(channel *entities.Channel) (*entities.Channel, error) {
	return &entities.Channel{}, nil
}

func (r *ListsRepository) GetAllChannels() ([]*entities.Channel, error) {
	return []*entities.Channel{}, nil
}

func (r *ListsRepository) GetFeed() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}
