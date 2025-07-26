package internal

import (
	"math"
	"raytracer/internal/geometry"
)

type Box struct {
	min      geometry.Vector
	max      geometry.Vector
	material Material
}

func NewBox(min, max geometry.Vector, material Material) Box {
	return Box{min: min, max: max, material: material}
}

func (box Box) hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	// Ray-box intersection using the slab method
	invDir := geometry.Vector{
		X: 1.0 / ray.direction.X,
		Y: 1.0 / ray.direction.Y,
		Z: 1.0 / ray.direction.Z,
	}

	// Calculate t values for each axis
	t1 := (box.min.X - ray.origin.X) * invDir.X
	t2 := (box.max.X - ray.origin.X) * invDir.X
	t3 := (box.min.Y - ray.origin.Y) * invDir.Y
	t4 := (box.max.Y - ray.origin.Y) * invDir.Y
	t5 := (box.min.Z - ray.origin.Z) * invDir.Z
	t6 := (box.max.Z - ray.origin.Z) * invDir.Z

	// Find the min and max t values for each axis
	tMinX := math.Min(t1, t2)
	tMaxX := math.Max(t1, t2)
	tMinY := math.Min(t3, t4)
	tMaxY := math.Max(t3, t4)
	tMinZ := math.Min(t5, t6)
	tMaxZ := math.Max(t5, t6)

	// Find the overall min and max t values
	tNear := math.Max(math.Max(tMinX, tMinY), tMinZ)
	tFar := math.Min(math.Min(tMaxX, tMaxY), tMaxZ)

	// Check if ray intersects the box
	if tNear > tFar || tFar < tMin || tNear > tMax {
		return false, HitRecord{}
	}

	// Determine which t to use (closest valid intersection)
	t := tNear
	isExitingBox := false
	if t < tMin {
		t = tFar
		isExitingBox = true
		if t > tMax {
			return false, HitRecord{}
		}
	}

	// Calculate hit point
	hitPoint := ray.pointAtParameter(t)

	// Calculate normal based on which slab was hit
	normal := box.calculateNormalFromSlab(t, tMinX, tMaxX, tMinY, tMaxY, tMinZ, tMaxZ, isExitingBox)

	hr := HitRecord{
		t:        t,
		p:        hitPoint,
		normal:   normal,
		material: box.material,
	}

	return true, hr
}

func (box Box) calculateNormalFromSlab(t, tMinX, tMaxX, tMinY, tMaxY, tMinZ, tMaxZ float64, isExitingBox bool) geometry.Vector {
	const epsilon = 1e-8

	// Determine which slab was hit by comparing t values
	if math.Abs(t-tMinX) < epsilon {
		return geometry.Vector{X: -1, Y: 0, Z: 0} // Left face
	} else if math.Abs(t-tMaxX) < epsilon {
		return geometry.Vector{X: 1, Y: 0, Z: 0} // Right face
	} else if math.Abs(t-tMinY) < epsilon {
		return geometry.Vector{X: 0, Y: -1, Z: 0} // Bottom face
	} else if math.Abs(t-tMaxY) < epsilon {
		return geometry.Vector{X: 0, Y: 1, Z: 0} // Top face
	} else if math.Abs(t-tMinZ) < epsilon {
		return geometry.Vector{X: 0, Y: 0, Z: -1} // Back face
	} else if math.Abs(t-tMaxZ) < epsilon {
		return geometry.Vector{X: 0, Y: 0, Z: 1} // Front face
	}

	// Fallback: determine which slab t is closest to
	switch t {
	case tMinX:
		return geometry.Vector{X: -1, Y: 0, Z: 0}
	case tMaxX:
		return geometry.Vector{X: 1, Y: 0, Z: 0}
	case tMinY:
		return geometry.Vector{X: 0, Y: -1, Z: 0}
	case tMaxY:
		return geometry.Vector{X: 0, Y: 1, Z: 0}
	case tMinZ:
		return geometry.Vector{X: 0, Y: 0, Z: -1}
	case tMaxZ:
		return geometry.Vector{X: 0, Y: 0, Z: 1}
	default:
		// This shouldn't happen if the logic above is correct
		return geometry.Vector{X: 0, Y: 0, Z: 0}
	}
}
