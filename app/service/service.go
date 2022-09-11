package service

// Services all service object injected here
type Services struct {
	Auth        AuthService
	Merchant    MerchantService
	Outlet      OutletService
	Transaction TransactionService
}
