package common

import (
	"github.com/pkg/errors"
)

var (
	OwnerNotFound        = errors.New("owner not found")
	OwnedByOtherIdentity = errors.New("owned by other identity")
)
