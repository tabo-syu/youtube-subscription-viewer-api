package gateways

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ChannelsRepository struct {
	sql interfaces.SqlHandler
}

var _ ports.ChannelsRepository = (*ChannelsRepository)(nil)

func NewChannelsRepository(s interfaces.SqlHandler) *ChannelsRepository {
	return &ChannelsRepository{s}
}

func (r *ChannelsRepository) GetFeed() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}

func (r *ChannelsRepository) BulkSave(channels []*entities.Channel) error {
	tx, err := r.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO channels (id, name, thumbnail, url) VALUES ($1, $2, $3, $4)
		ON CONFLICT (id)
		DO UPDATE SET name = $2, thumbnail = $3, url = $4
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, channel := range channels {
		_, err := stmt.Exec(channel.Id, channel.Name, channel.Thumbnail, channel.Url)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
