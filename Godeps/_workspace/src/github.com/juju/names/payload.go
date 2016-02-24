// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/juju/utils"
)

const (
	// PayloadTagKind is used as the prefix for the string
	// representation of payload tags.
	PayloadTagKind = "payload"
)

// IsValidPayload returns whether id is a valid Juju ID for
// a charm payload. The ID must be a valid UUID.
func IsValidPayload(id string) bool {
	return utils.IsValidUUIDString(id)
}

// PayloadTag represents a charm payload.
type PayloadTag struct {
	id string
}

// NewPayloadTag returns the tag for a charm's payload with the given id.
func NewPayloadTag(id string) PayloadTag {
	return PayloadTag{
		id: id,
	}
}

// ParsePayloadTag parses a payload tag string.
// So ParsePayloadTag(tag.String()) === tag.
func ParsePayloadTag(tag string) (PayloadTag, error) {
	t, err := ParseTag(tag)
	if err != nil {
		return PayloadTag{}, err
	}
	pt, ok := t.(PayloadTag)
	if !ok {
		return PayloadTag{}, invalidTagError(tag, PayloadTagKind)
	}
	return pt, nil
}

// Kind implements Tag.
func (t PayloadTag) Kind() string {
	return PayloadTagKind
}

// Id implements Tag.Id. It always returns the same ID with which
// it was created. So NewPayloadTag(x).Id() == x for all valid x.
func (t PayloadTag) Id() string {
	return t.id
}

// String implements Tag.
func (t PayloadTag) String() string {
	return tagString(t)
}