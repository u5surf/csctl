package resource

// strptr is intended to be used to provide a pointer to a
// static string for easily building test resources
func strptr(s string) *string {
	return &s
}

// int32ptr is intended to be used to provide a pointer to
// an int32 for easily building test resources
func int32ptr(i int32) *int32 {
	return &i
}
