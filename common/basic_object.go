package common

type BasicObject struct {
	id        int64
	transform Matrix
	inverseTransform Matrix
	inverseTransposeTransform Matrix
	material  Material
}

func (bo *BasicObject) transformPointToObject(p Vector) Vector {
	return bo.inverseTransform.Transform(p)
}

func (bo *BasicObject) transformRayToObject(r Ray) Ray {
	return r.Transform(bo.inverseTransform)
}

func (bo *BasicObject) transformVectorToWorld(v Vector) Vector {
	return bo.inverseTransposeTransform.Transform(v)
}

func (bo *BasicObject) Id() int64 {
	return bo.id
}

func (bo *BasicObject) SetTransform(m Matrix) {
	bo.transform = m
	bo.inverseTransform = m.Inverse()
	bo.inverseTransposeTransform = m.Inverse().Transpose()
}

func (bo *BasicObject) SetMaterial(m Material) {
	bo.material = m
}

func (bo *BasicObject) Material() Material {
	return bo.material
}
