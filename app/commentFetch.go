package app

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
)

func Fetch(apiKey, receivedVideoID string) []string {
	var commentSlice []string

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
		return nil
	}

	var nextPageToken string
	for {
		commentsResponse, err := youtubeService.CommentThreads.List([]string{"snippet"}).
			VideoId(receivedVideoID).
			TextFormat("plainText").
			MaxResults(100). // Adjust the number of comments per page as needed.
			PageToken(nextPageToken).
			Do()
		if err != nil {
			log.Fatalf("Error retrieving comments: %v", err)
		}

		for _, comment := range commentsResponse.Items {
			commentText := comment.Snippet.TopLevelComment.Snippet.TextDisplay
			commentSlice = append(commentSlice, commentText)
		}

		// Check if there are more pages of comments.
		nextPageToken = commentsResponse.NextPageToken
		if nextPageToken == "" {
			break // No more pages of comments
		}
	}
	return commentSlice
}

func FetchInstagram(apiKey, receivedID string) {

}
