package types

const (
	// ModuleName defines the module name
	ModuleName = "loan"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_loan"
)

var (
	ParamsKey = []byte("p_loan")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LoanKey      = "Loan/value/"
	LoanCountKey = "Loan/count/"
)
