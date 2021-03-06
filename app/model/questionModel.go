package model

import (
	"intelliq/app/enums"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

//Question question model
type Question struct {
	QuestionID       bson.ObjectId        `json:"quesId" bson:"_id,omitempty"`
	Title            string               `json:"title" bson:"title"`
	TitleHTML        string               `json:"titleHtml" bson:"titleHtml"`
	Std              int                  `json:"std" bson:"std"`
	Subject          string               `json:"subject" bson:"subject"`
	Topic            string               `json:"topic" bson:"topic"`
	Difficulty       enums.QuesDifficulty `json:"difficulty" bson:"difficulty"`
	Length           enums.QuesLength     `json:"length" bson:"length"`
	Marks            int                  `json:"marks" bson:"marks"`
	Status           enums.QuestionStatus `json:"status" bson:"status"`
	Tags             []string             `json:"tags" bson:"tags"`
	Category         string               `json:"category" bson:"category"`
	ImageURL         string               `json:"imageUrl" bson:"imageUrl"`
	Owner            Contributor          `json:"owner" bson:"owner"`
	Reviewer         Contributor          `json:"reviewer" bson:"reviewer"`
	School           School               `json:"school" bson:"school"`
	GroupCode        string               `json:"groupCode" bson:"groupCode"`
	CreateDate       time.Time            `json:"createDate" bson:"createDate"`
	LastModifiedDate time.Time            `json:"lastModifiedDate" bson:"lastModifiedDate"`
	RejectDesc       string               `json:"rejectDesc" bson:"rejectDesc"`
	OriginID         *bson.ObjectId       `json:"originId" bson:"originId,omitempty"`
}

type Contributor struct {
	UserID   bson.ObjectId `json:"userId" bson:"_id,omitempty"`
	UserName string        `json:"userName" bson:"userName"`
}

//Questions question array
type Questions []Question

//FormatTopicTags formats topic and tags to lowercase
func (question *Question) FormatTopicTags() {
	question.Topic = strings.Title(strings.ToLower(strings.TrimSpace(question.Topic)))
	question.Category = strings.Title(strings.ToLower(strings.TrimSpace(question.Category)))
	var tags []string
	for _, tag := range question.Tags {
		tag := strings.TrimSpace(tag)
		if len(tag) > 0 {
			tags = append(tags, strings.Title(strings.ToLower(tag)))
		}
	}
	question.Tags = tags
}
