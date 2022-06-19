package imbue

// WithGroupedOption is an option that changes the behavior of a call to
// WithXGrouped().
type WithGroupedOption interface {
	applyWithGroupedOptionToContainer(*Container)
	applyWithGroupedOptionToContext(*Context)
}
