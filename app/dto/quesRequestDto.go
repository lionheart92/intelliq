package dto

import (
	"intelliq/app/enums"

	"github.com/globalsign/mgo/bson"
)

//QuesRequestDto request params
type QuesRequestDto struct {
	GroupCode string               `json:"groupCode"`
	SchoolID  bson.ObjectId        `json:"schoolId"`
	UserID    bson.ObjectId        `json:"userId"` // reviewer || teacher
	Std       uint16               `json:"std"`
	Subject   string               `json:"subject"`
	Status    enums.QuestionStatus `json:"status"` // quest status
	Page      int                  `json:"page"`   // for pagination
	Limit     int                  `json:"limit"`  // max records per request
}
