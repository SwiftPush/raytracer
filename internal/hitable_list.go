package internal

type HitableList struct {
	hitables []Hitable
}

func (h *HitableList) hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	tempRecord := HitRecord{}
	hitAnything := false
	closestSoFar := tMax
	for _, hitable := range h.hitables {
		if hit, hitRecord := hitable.hit(ray, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = hitRecord.t
			tempRecord = hitRecord
		}
	}
	return hitAnything, tempRecord
}
