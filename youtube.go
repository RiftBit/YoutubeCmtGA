package main

import (
	"github.com/parnurzeal/gorequest"
	"regexp"
	"encoding/json"
	"time"
	"math/rand"
)

var CommentsDB map[string]ParsedComment


type ParsedComment struct {
	Id string
	AuthorChannelURL string
	AuthorChannelID string
	AuthorDisplayName string
	TextDisplay string
	AuthorProfileImageURL string
}


type GetCommentsParams struct {
	Url string
	Pattern *regexp.Regexp
	PageToken string
}


type ThreadObj struct {
	Etag  string `json:"etag"`
	NextPageToken  string `json:"nextPageToken"`
	Items []struct {
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Kind    string `json:"kind"`
		Snippet struct {
				CanReply        bool `json:"canReply"`
				IsPublic        bool `json:"isPublic"`
				TopLevelComment struct {
							Etag    string `json:"etag"`
							ID      string `json:"id"`
							Kind    string `json:"kind"`
							Snippet struct {
									AuthorChannelID struct {
												Value string `json:"value"`
											} `json:"authorChannelId"`
									AuthorChannelURL      string `json:"authorChannelUrl"`
									AuthorDisplayName     string `json:"authorDisplayName"`
									AuthorProfileImageURL string `json:"authorProfileImageUrl"`
									CanRate               bool   `json:"canRate"`
									LikeCount             int    `json:"likeCount"`
									PublishedAt           string `json:"publishedAt"`
									TextDisplay           string `json:"textDisplay"`
									UpdatedAt             string `json:"updatedAt"`
									VideoID               string `json:"videoId"`
									ViewerRating          string `json:"viewerRating"`
								} `json:"snippet"`
						} `json:"topLevelComment"`
				TotalReplyCount int    `json:"totalReplyCount"`
				VideoID         string `json:"videoId"`
			} `json:"snippet"`
	} `json:"items"`
	Kind     string `json:"kind"`
	PageInfo struct {
		      ResultsPerPage int `json:"resultsPerPage"`
		      TotalResults   int `json:"totalResults"`
	      } `json:"pageInfo"`
}

func GetComments(params GetCommentsParams) {
	var finalUrl string
	if params.PageToken != "" {
		finalUrl = "https://www.googleapis.com/youtube/v3/commentThreads?part=snippet&maxResults=100&textFormat=plainText&fields=items,nextPageToken&pageToken=" + params.PageToken + "&videoId=" + params.Url + "&key="+ Configs.Youtube.ApiKey
	} else {
		finalUrl = "https://www.googleapis.com/youtube/v3/commentThreads?part=snippet&maxResults=100&textFormat=plainText&fields=items,nextPageToken&videoId=" + params.Url + "&key="+ Configs.Youtube.ApiKey
	}
	_, body, errs := gorequest.New().Get(finalUrl).End()
	if len(errs) > 0 {
		panic(errs)
	}

	m := &ThreadObj{}
	_ = json.Unmarshal([]byte(body), &m)

	if len(m.Items) > 0 {

		for _, item := range m.Items {
			//Проверяем комментарий на совпадение по нашему регулярному выражению
			if params.Pattern.MatchString(item.Snippet.TopLevelComment.Snippet.TextDisplay) {
				comment := ParsedComment{}
				comment.AuthorDisplayName = item.Snippet.TopLevelComment.Snippet.AuthorDisplayName
				comment.TextDisplay = item.Snippet.TopLevelComment.Snippet.TextDisplay
				comment.AuthorProfileImageURL = item.Snippet.TopLevelComment.Snippet.AuthorProfileImageURL
				comment.AuthorChannelURL = item.Snippet.TopLevelComment.Snippet.AuthorChannelURL
				comment.AuthorChannelID = item.Snippet.TopLevelComment.Snippet.AuthorChannelID.Value
				comment.Id = item.Snippet.TopLevelComment.ID
				//Добавляя по ключу (ID пользователя) мы убираем дубли от комментаторов, дабы уравнять шансы у всех участников
				CommentsDB[comment.AuthorChannelID] = comment
			}
		}

		//Если есть другие страницы, то рекурсивно начинаем обрабатывать и их
		if m.NextPageToken != "" {
			GetComments(GetCommentsParams{Url: params.Url, Pattern: params.Pattern, PageToken: m.NextPageToken})
		}
	}
	//fmt.Println(len(CommentsDB))
}

func GetRandComment() ParsedComment {
	//Превращаем Map в Slice для того, чтобы потом рандомно выбрать из него комментарий
	sliceOfComments := make([]ParsedComment, 0, len(CommentsDB))
	for _, comment := range CommentsDB {
		sliceOfComments = append(sliceOfComments, comment)
	}

	//генерируем случайный номер комментария
	rand.Seed(time.Now().Unix())
	CommentNumer := rand.Int() % len(sliceOfComments)

	result := ParsedComment{}
	result.Id = sliceOfComments[CommentNumer].Id
	result.AuthorChannelURL = sliceOfComments[CommentNumer].AuthorChannelURL
	result.AuthorChannelID = sliceOfComments[CommentNumer].AuthorChannelID
	result.AuthorDisplayName = sliceOfComments[CommentNumer].AuthorDisplayName
	result.TextDisplay = sliceOfComments[CommentNumer].TextDisplay
	result.AuthorProfileImageURL = sliceOfComments[CommentNumer].AuthorProfileImageURL

	return result
}