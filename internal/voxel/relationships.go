package voxel

const (
	R_UNKNOWN Relationship = iota
	R_ROOT
	R_IS
	R_CARRIES
	R_CONTAINS
	R_HOLDS
	R_WEARS
)

type Relationship int
