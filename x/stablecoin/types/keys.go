package types

const (
	// ModuleName defines the module name
	ModuleName = "stablecoin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stablecoin"
)

var (
	ParamsKey = []byte("p_stablecoin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
