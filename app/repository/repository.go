package repository

// Repositories all repo object injected here
type Repositories struct {
	Merchant    MerchantRepository
	Outlet      OutletRepository
	Transaction TransactionRepository
	User        UserRepository
}
