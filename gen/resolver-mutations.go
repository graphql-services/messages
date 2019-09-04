package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
	"github.com/novacloudcz/graphql-orm/resolvers"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateComment(ctx context.Context, input map[string]interface{}) (item *Comment, err error) {
	return r.Handlers.CreateComment(ctx, r.GeneratedResolver, input)
}
func CreateCommentHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Comment, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Comment{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Comment",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CommentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["text"]; ok && (item.Text != changes.Text) && (item.Text == nil || changes.Text == nil || *item.Text != *changes.Text) {
		item.Text = changes.Text
		event.AddNewValue("text", changes.Text)
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) && (item.Reference == nil || changes.Reference == nil || *item.Reference != *changes.Reference) {
		item.Reference = changes.Reference
		event.AddNewValue("reference", changes.Reference)
	}

	if _, ok := input["referenceID"]; ok && (item.ReferenceID != changes.ReferenceID) && (item.ReferenceID == nil || changes.ReferenceID == nil || *item.ReferenceID != *changes.ReferenceID) {
		item.ReferenceID = changes.ReferenceID
		event.AddNewValue("referenceID", changes.ReferenceID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateComment(ctx context.Context, id string, input map[string]interface{}) (item *Comment, err error) {
	return r.Handlers.UpdateComment(ctx, r.GeneratedResolver, id, input)
}
func UpdateCommentHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Comment, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Comment{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Comment",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CommentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["text"]; ok && (item.Text != changes.Text) && (item.Text == nil || changes.Text == nil || *item.Text != *changes.Text) {
		event.AddOldValue("text", item.Text)
		event.AddNewValue("text", changes.Text)
		item.Text = changes.Text
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) && (item.Reference == nil || changes.Reference == nil || *item.Reference != *changes.Reference) {
		event.AddOldValue("reference", item.Reference)
		event.AddNewValue("reference", changes.Reference)
		item.Reference = changes.Reference
	}

	if _, ok := input["referenceID"]; ok && (item.ReferenceID != changes.ReferenceID) && (item.ReferenceID == nil || changes.ReferenceID == nil || *item.ReferenceID != *changes.ReferenceID) {
		event.AddOldValue("referenceID", item.ReferenceID)
		event.AddNewValue("referenceID", changes.ReferenceID)
		item.ReferenceID = changes.ReferenceID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteComment(ctx context.Context, id string) (item *Comment, err error) {
	return r.Handlers.DeleteComment(ctx, r.GeneratedResolver, id)
}
func DeleteCommentHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Comment, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Comment{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Comment",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, "comments.id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllComments(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllComments(ctx, r.GeneratedResolver)
}
func DeleteAllCommentsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Comment{}).Error
	return err == nil, err
}
