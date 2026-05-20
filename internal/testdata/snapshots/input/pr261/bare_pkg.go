package pr261

import "sort"

// Legitimate use so goimports doesn't strip the import.
var _ = sort.Slice

// Bare PkgName reference that triggered the panic.
var _ = sort
