package service

import (
	"context"
	pb "forum-service/forum-protos/genprotos"
	st "forum-service/storage"
	"regexp"
)

type TagService struct {
	storage st.Storage
	pb.UnimplementedTagServiceServer
}

func NewTagService(storage *st.Storage) *TagService {
	return &TagService{storage: *storage}
}

func (s *TagService) GetPopular(ctx context.Context, req *pb.Pagination) (*pb.TagPopularRes, error) {
	return s.storage.TagS.GetPopular(req)
}

func ValidateTags(tags string) (bool, []string) {
	re := regexp.MustCompile(`^(#\w+(\s*,?\s*#\w+)*)$`)

	if re.MatchString(tags) {
		splitTags := regexp.MustCompile(`[,\s]+`).Split(tags, -1)
		uniqueTagsMap := make(map[string]bool)
		var result []string

		for _, tag := range splitTags {
			if !uniqueTagsMap[tag] {
				uniqueTagsMap[tag] = true
				result = append(result, tag)
			}
		}

		return true, result
	}
	return false, nil
}
