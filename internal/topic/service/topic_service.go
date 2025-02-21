package service

import "github.com/roka-crew/internal/topic/store"

type TopicService struct {
	topicStore store.TopicStore
}

func NewTopicService(topicStore store.TopicStore) *TopicService {
	return &TopicService{
		topicStore: topicStore,
	}
}
