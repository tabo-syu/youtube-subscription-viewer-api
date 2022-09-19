package infrastructures

type YoutubeHandler struct{}

func NewYoutubeHandler(config *Youtube) *YoutubeHandler {
	return &YoutubeHandler{}
}
