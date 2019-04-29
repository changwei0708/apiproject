package video

import (
	"apiproject/model/video"
)

func FindVideoList() []video.VideoModel {
	var videoModelList []video.VideoModel
	videoModelList = append(videoModelList, video.VideoModel{"id1", "title1"})
	videoModelList = append(videoModelList, video.VideoModel{"id2", "title2"})

	return videoModelList
}
