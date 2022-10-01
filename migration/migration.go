package migration

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/infrastructures"
)

func Migrate(sql *infrastructures.SqlHandler) error {
	_, err := sql.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id            varchar(24)                            PRIMARY KEY,
			name          text                                   NOT NULL,
			thumbnail     text                                   NOT NULL,
			access_token  text                                   NOT NULL,
			refresh_token text                                   NOT NULL,
			expiry        timestamp with time zone               NOT NULL,
			created_at    timestamp with time zone DEFAULT NOW() NOT NULL,
			updated_at    timestamp with time zone DEFAULT NOW() NOT NULL
		);

		CREATE TABLE IF NOT EXISTS lists (
			id         uuid                     DEFAULT gen_random_uuid() PRIMARY KEY,
			title      varchar(100)                                       NOT NULL,
			user_id    varchar(24)                                        NOT NULL,
			created_at timestamp with time zone DEFAULT NOW()             NOT NULL,
			updated_at timestamp with time zone DEFAULT NOW()             NOT NULL,

			FOREIGN KEY (user_id) REFERENCES users (id)
		);
		
		CREATE TABLE IF NOT EXISTS channels (
			id           varchar(24) PRIMARY KEY,
			name         text        NOT NULL,
			thumbnail    text        NOT NULL,
			url          text        NOT NULL
		);

		CREATE TABLE IF NOT EXISTS channels_belonging_list (
			id         serial      PRIMARY KEY,
			list_id    uuid        NOT NULL,
			channel_id varchar(24) NOT NULL,
			
			FOREIGN KEY (list_id)    REFERENCES lists    (id),
			FOREIGN KEY (channel_id) REFERENCES channels (id)
		);

		CREATE TABLE IF NOT EXISTS user_subscribe_channels (
			id         serial      PRIMARY KEY,
			user_id    varchar(24) NOT NULL,
			channel_id varchar(24) NOT NULL,
			
			FOREIGN KEY (user_id)    REFERENCES users    (id),
			FOREIGN KEY (channel_id) REFERENCES channels (id)
		);

		CREATE TABLE IF NOT EXISTS videos (
			id           varchar(11)              PRIMARY KEY,
			title        varchar(100)             NOT NULL,
			thumbnail    text                     NOT NULL,
			url          text                     NOT NULL,
			published_at timestamp with time zone NOT NULL,
			channel_id   varchar(24)              NOT NULL,
			
			FOREIGN KEY (channel_id) REFERENCES channels (id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
