package voxel

const (
	T_UNKNOWN VoxelType = iota
	T_GAME
	T_PLANE
	T_WORLD
	T_REGION
	T_LOCALE
	T_PROVINCE
	T_BUILDING
	T_ROOM
	T_NPC
	T_CHARACTER
	T_ITEM
	T_CONTAINER
)

type VoxelType int
