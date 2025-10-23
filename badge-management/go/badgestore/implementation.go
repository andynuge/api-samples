// This mock store simulates a database client for badges.
// It logs creation and returns a Badge object with a generated ID.

package badgestore

import (
	"log"
	"time"
)

// BadgeStore defines the interface for badge storage operations.
type BadgeStore interface {
	Create(serialNumber string, version string) (Badge, error)
}

// Badge represents a badge entity.
type Badge struct {
	// ID is the identifier of the badge in the Database
	ID int `json:"id"`

	// SerialNumber of the badge that is unique per badge
	SerialNumber string `json:"serialNumber"`

	// Version of the badge firmware
	Version string `json:"version"`
}

type badgeStoreImpl struct{}

// Create creates a new badge in the store.
// For this example, it simply logs the creation and returns a Badge with a mock ID.
// In a real implementation, this would interact with a database.
func (b *badgeStoreImpl) Create(serialNumber string, version string) (Badge, error) {
	log.Printf("Created badge with SerialNumber: %s, Version: %s", serialNumber, version)
	return Badge{
		ID:           int(time.Now().UnixMilli()),
		SerialNumber: serialNumber,
		Version:      version,
	}, nil
}

// Ensure badgeStoreImpl implements BadgeStore interface
var _ BadgeStore = (*badgeStoreImpl)(nil)

func New() *badgeStoreImpl {
	return &badgeStoreImpl{}
}
