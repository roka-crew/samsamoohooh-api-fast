package handler

import "github.com/roka-crew/internal/topic/service"

type TopicHandler struct {
	topicService service.TopicService
}

func NewTopicHandler(topicService service.TopicService) *TopicHandler {
	topicHandler := &TopicHandler{
		topicService: topicService,
	}

	return topicHandler
}
