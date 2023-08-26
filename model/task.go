package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OwnerModel struct {
	CreatedBy *string `bson:"createdBy" json:"createdBy"`
}

type HistoryModel struct {
	CreatedAt *primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	UpdatedAt *primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`
}

// Concrete Model

// Task Models.

type Task struct {
	OwnerModel
	HistoryModel
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description *string            `bson:"description" json:"description,omitempty"`
	Actions     []*Action          `json:"actions,omitempty"`
}

func (t *Task) ToGQLModel() TaskData {
	var actions []*ActionData
	if len(t.Actions) != 0 {
		actions = make([]*ActionData, len(t.Actions))
		for i := 0; i < len(t.Actions); i++ {
			actionGQLResult := t.Actions[i].ToGQLModel()
			actions[i] = &actionGQLResult
		}
	}

	return TaskData{
		ID:          t.ID.Hex(),
		Title:       t.Title,
		Description: t.Description,
		Actions:     actions,
		CreatedBy:   t.CreatedBy,
		CreatedAt:   convertTimestampToInt(t.CreatedAt),
		UpdateAt:    convertTimestampToInt(t.UpdatedAt),
	}
}

type Action struct {
	OwnerModel
	HistoryModel
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Task        primitive.ObjectID `bson:"task"`
	Title       string             `bson:"title" json:"title"`
	Description *string            `bson:"description" json:"description"`
}

func (a *Action) ToGQLModel() ActionData {
	return ActionData{
		ID:          a.ID.Hex(),
		Title:       a.Title,
		Description: a.Description,
		Task:        a.Task.Hex(),
		CreatedBy:   a.CreatedBy,
		CreatedAt:   convertTimestampToInt(a.CreatedAt),
		UpdateAt:    convertTimestampToInt(a.UpdatedAt),
	}
}

func convertTimestampToInt(timeObj *primitive.Timestamp) *int {
	var result *int

	if timeObj == nil {
		return nil
	}

	if !timeObj.IsZero() {
		timeInt := int(timeObj.T)
		result = &timeInt
	}
	return result
}
