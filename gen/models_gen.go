// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CommentFilterType struct {
	And             []*CommentFilterType `json:"AND"`
	Or              []*CommentFilterType `json:"OR"`
	ID              *string              `json:"id"`
	IDNe            *string              `json:"id_ne"`
	IDGt            *string              `json:"id_gt"`
	IDLt            *string              `json:"id_lt"`
	IDGte           *string              `json:"id_gte"`
	IDLte           *string              `json:"id_lte"`
	IDIn            []string             `json:"id_in"`
	Text            *string              `json:"text"`
	TextNe          *string              `json:"text_ne"`
	TextGt          *string              `json:"text_gt"`
	TextLt          *string              `json:"text_lt"`
	TextGte         *string              `json:"text_gte"`
	TextLte         *string              `json:"text_lte"`
	TextIn          []string             `json:"text_in"`
	TextLike        *string              `json:"text_like"`
	TextPrefix      *string              `json:"text_prefix"`
	TextSuffix      *string              `json:"text_suffix"`
	Reference       *string              `json:"reference"`
	ReferenceNe     *string              `json:"reference_ne"`
	ReferenceGt     *string              `json:"reference_gt"`
	ReferenceLt     *string              `json:"reference_lt"`
	ReferenceGte    *string              `json:"reference_gte"`
	ReferenceLte    *string              `json:"reference_lte"`
	ReferenceIn     []string             `json:"reference_in"`
	ReferenceLike   *string              `json:"reference_like"`
	ReferencePrefix *string              `json:"reference_prefix"`
	ReferenceSuffix *string              `json:"reference_suffix"`
	ReferenceID     *string              `json:"referenceID"`
	ReferenceIDNe   *string              `json:"referenceID_ne"`
	ReferenceIDGt   *string              `json:"referenceID_gt"`
	ReferenceIDLt   *string              `json:"referenceID_lt"`
	ReferenceIDGte  *string              `json:"referenceID_gte"`
	ReferenceIDLte  *string              `json:"referenceID_lte"`
	ReferenceIDIn   []string             `json:"referenceID_in"`
	UpdatedAt       *time.Time           `json:"updatedAt"`
	UpdatedAtNe     *time.Time           `json:"updatedAt_ne"`
	UpdatedAtGt     *time.Time           `json:"updatedAt_gt"`
	UpdatedAtLt     *time.Time           `json:"updatedAt_lt"`
	UpdatedAtGte    *time.Time           `json:"updatedAt_gte"`
	UpdatedAtLte    *time.Time           `json:"updatedAt_lte"`
	UpdatedAtIn     []*time.Time         `json:"updatedAt_in"`
	CreatedAt       *time.Time           `json:"createdAt"`
	CreatedAtNe     *time.Time           `json:"createdAt_ne"`
	CreatedAtGt     *time.Time           `json:"createdAt_gt"`
	CreatedAtLt     *time.Time           `json:"createdAt_lt"`
	CreatedAtGte    *time.Time           `json:"createdAt_gte"`
	CreatedAtLte    *time.Time           `json:"createdAt_lte"`
	CreatedAtIn     []*time.Time         `json:"createdAt_in"`
	UpdatedBy       *string              `json:"updatedBy"`
	UpdatedByNe     *string              `json:"updatedBy_ne"`
	UpdatedByGt     *string              `json:"updatedBy_gt"`
	UpdatedByLt     *string              `json:"updatedBy_lt"`
	UpdatedByGte    *string              `json:"updatedBy_gte"`
	UpdatedByLte    *string              `json:"updatedBy_lte"`
	UpdatedByIn     []string             `json:"updatedBy_in"`
	CreatedBy       *string              `json:"createdBy"`
	CreatedByNe     *string              `json:"createdBy_ne"`
	CreatedByGt     *string              `json:"createdBy_gt"`
	CreatedByLt     *string              `json:"createdBy_lt"`
	CreatedByGte    *string              `json:"createdBy_gte"`
	CreatedByLte    *string              `json:"createdBy_lte"`
	CreatedByIn     []string             `json:"createdBy_in"`
}

type _Service struct {
	Sdl *string `json:"sdl"`
}

type CommentSortType string

const (
	CommentSortTypeIDAsc           CommentSortType = "ID_ASC"
	CommentSortTypeIDDesc          CommentSortType = "ID_DESC"
	CommentSortTypeTextAsc         CommentSortType = "TEXT_ASC"
	CommentSortTypeTextDesc        CommentSortType = "TEXT_DESC"
	CommentSortTypeReferenceAsc    CommentSortType = "REFERENCE_ASC"
	CommentSortTypeReferenceDesc   CommentSortType = "REFERENCE_DESC"
	CommentSortTypeReferenceIDAsc  CommentSortType = "REFERENCE_ID_ASC"
	CommentSortTypeReferenceIDDesc CommentSortType = "REFERENCE_ID_DESC"
	CommentSortTypeUpdatedAtAsc    CommentSortType = "UPDATED_AT_ASC"
	CommentSortTypeUpdatedAtDesc   CommentSortType = "UPDATED_AT_DESC"
	CommentSortTypeCreatedAtAsc    CommentSortType = "CREATED_AT_ASC"
	CommentSortTypeCreatedAtDesc   CommentSortType = "CREATED_AT_DESC"
	CommentSortTypeUpdatedByAsc    CommentSortType = "UPDATED_BY_ASC"
	CommentSortTypeUpdatedByDesc   CommentSortType = "UPDATED_BY_DESC"
	CommentSortTypeCreatedByAsc    CommentSortType = "CREATED_BY_ASC"
	CommentSortTypeCreatedByDesc   CommentSortType = "CREATED_BY_DESC"
)

var AllCommentSortType = []CommentSortType{
	CommentSortTypeIDAsc,
	CommentSortTypeIDDesc,
	CommentSortTypeTextAsc,
	CommentSortTypeTextDesc,
	CommentSortTypeReferenceAsc,
	CommentSortTypeReferenceDesc,
	CommentSortTypeReferenceIDAsc,
	CommentSortTypeReferenceIDDesc,
	CommentSortTypeUpdatedAtAsc,
	CommentSortTypeUpdatedAtDesc,
	CommentSortTypeCreatedAtAsc,
	CommentSortTypeCreatedAtDesc,
	CommentSortTypeUpdatedByAsc,
	CommentSortTypeUpdatedByDesc,
	CommentSortTypeCreatedByAsc,
	CommentSortTypeCreatedByDesc,
}

func (e CommentSortType) IsValid() bool {
	switch e {
	case CommentSortTypeIDAsc, CommentSortTypeIDDesc, CommentSortTypeTextAsc, CommentSortTypeTextDesc, CommentSortTypeReferenceAsc, CommentSortTypeReferenceDesc, CommentSortTypeReferenceIDAsc, CommentSortTypeReferenceIDDesc, CommentSortTypeUpdatedAtAsc, CommentSortTypeUpdatedAtDesc, CommentSortTypeCreatedAtAsc, CommentSortTypeCreatedAtDesc, CommentSortTypeUpdatedByAsc, CommentSortTypeUpdatedByDesc, CommentSortTypeCreatedByAsc, CommentSortTypeCreatedByDesc:
		return true
	}
	return false
}

func (e CommentSortType) String() string {
	return string(e)
}

func (e *CommentSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommentSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommentSortType", str)
	}
	return nil
}

func (e CommentSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
