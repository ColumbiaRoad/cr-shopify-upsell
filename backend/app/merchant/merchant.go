package merchant

import "context"

// Storage is the interface to interact with the database, the methods are implemented in db/
type Storage interface {
	PersistUser(ctx context.Context, merchant string) (int64, error)
	CheckMerchantByShop(ctx context.Context, shopURL string) (bool, error)
}

type merchant struct {
	storage Storage
}

// New is a new User
func New(storage Storage) Merchants {
	return &merchant{
		storage: storage,
	}
}

// Merchants represents the interface to "Users"
type Merchants interface {
}
