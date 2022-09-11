package repository

// Repositories all repo object injected here
type Repositories struct {
	Merchant    MerchantRepository
	Transaction TransactionRepository
	User        UserRepository
}
