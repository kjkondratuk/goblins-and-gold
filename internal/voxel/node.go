package voxel

type Voxel interface {
	Key() Relationship
	Name() string
	Type() VoxelType

	Children() map[Relationship]Voxel
	Parents() map[Relationship]Voxel

	SetChild(Relationship, Voxel)
	SetChildren(map[Relationship]Voxel)
	RemoveChildrenByKey(...Relationship)
	RemoveChildrenByName(...string)
	RemoveChildrenByType(...VoxelType)

	SetParent(Relationship, Voxel)
	SetParents(map[Relationship]Voxel)
	RemoveParentsByKey(...Relationship)
	RemoveParentsByName(...string)
	RemoveParentsByType(...VoxelType)
}

type voxel struct {
	_key      Relationship
	_name     string
	_type     VoxelType
	_children map[Relationship]Voxel
	_parents  map[Relationship]Voxel
}

func NewVoxel(nKey Relationship, nName string, nType VoxelType) Voxel {
	return &voxel{
		_key:      nKey,
		_name:     nName,
		_type:     nType,
		_parents:  make(map[Relationship]Voxel),
		_children: make(map[Relationship]Voxel),
	}
}

func (v *voxel) Key() Relationship {
	return v._key
}

func (v *voxel) Name() string {
	return v._name
}

func (v *voxel) Type() VoxelType {
	return v._type
}

func (v *voxel) Children() map[Relationship]Voxel {
	return v._children
}

func (v *voxel) Parents() map[Relationship]Voxel {
	return v._parents
}

func (v *voxel) SetChild(key Relationship, value Voxel) {
	v._children[key] = value
}

func (v *voxel) SetChildren(toSet map[Relationship]Voxel) {
	for k, val := range toSet {
		v.SetChild(k, val)
	}
}

func (v *voxel) RemoveChildrenByKey(keys ...Relationship) {
	for _, k := range keys {
		if _, ok := v._children[k]; ok {
			delete(v._children, k)
		}
	}
}

func (v *voxel) RemoveChildrenByName(names ...string) {
	for i, child := range v._children {
		for _, name := range names {
			if child.Name() == name {
				delete(v._children, i)
			}
		}
	}
}

func (v *voxel) RemoveChildrenByType(types ...VoxelType) {
	for i, child := range v._children {
		for _, t := range types {
			if child.Type() == t {
				delete(v._children, i)
			}
		}
	}
}

func (v *voxel) SetParent(key Relationship, val Voxel) {
	v._parents[key] = val
}

func (v *voxel) SetParents(toSet map[Relationship]Voxel) {
	for k, val := range toSet {
		v.SetParent(k, val)
	}
}

func (v *voxel) RemoveParentsByKey(keys ...Relationship) {
	for _, k := range keys {
		if _, ok := v._parents[k]; ok {
			delete(v._parents, k)
		}
	}
}

func (v *voxel) RemoveParentsByName(names ...string) {
	for i, parent := range v._parents {
		for _, name := range names {
			if parent.Name() == name {
				delete(v._parents, i)
			}
		}
	}
}

func (v *voxel) RemoveParentsByType(types ...VoxelType) {
	for i, parent := range v._parents {
		for _, t := range types {
			if parent.Type() == t {
				delete(v._parents, i)
			}
		}
	}
}
