package infrastructures

type YoutubeHandler struct{}

func NewYoutubeHandler(config *Config) *YoutubeHandler {
	return &YoutubeHandler{}
}
