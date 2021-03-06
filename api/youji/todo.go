package youji

import (
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateOperatorsFromParams generates a set of MongoDB update operators for the
// todo given the set of update parameters.
func (t *Todo) UpdateOperatorsFromParams(in *UpdateTodoParams) bson.M {
	currentDate := make(map[string]bool)
	set := make(map[string]interface{})
	unset := make(map[string]string)

	if in.Done && t.CompletedAt == nil {
		currentDate["completedat"] = true
	} else if !in.Done && t.CompletedAt != nil {
		unset["completedat"] = ""
	}

	if text := in.GetText(); text != "" {
		set["text"] = text
	}

	if len(currentDate) > 0 || len(set) > 0 || len(unset) > 0 {
		currentDate["updatedat"] = true
	}

	update := bson.M{}

	if len(currentDate) > 0 {
		update["$currentDate"] = currentDate
	}
	if len(set) > 0 {
		update["$set"] = set
	}
	if len(unset) > 0 {
		update["$unset"] = unset
	}

	return update
}
