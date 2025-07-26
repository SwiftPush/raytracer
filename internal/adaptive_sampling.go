package internal

import (
	"image/color"
	"math"
	"math/rand"
	"raytracer/internal/geometry"
)

type AdaptiveSamplingConfig struct {
	minSamples        int
	maxSamples        int
	varianceThreshold float64
	batchSize         int
}

func newAdaptiveSamplingConfig(maxSamples int) AdaptiveSamplingConfig {
	return AdaptiveSamplingConfig{
		minSamples:        16,
		maxSamples:        maxSamples,
		varianceThreshold: 0.0001,
		batchSize:         8,
	}
}

type colorAccumulator struct {
	colorSum     geometry.Vector
	colorSumSq   geometry.Vector
	sampleCount  int
}

func newColorAccumulator() colorAccumulator {
	return colorAccumulator{
		colorSum:    geometry.Vector{X: 0, Y: 0, Z: 0},
		colorSumSq:  geometry.Vector{X: 0, Y: 0, Z: 0},
		sampleCount: 0,
	}
}

func (acc *colorAccumulator) addSamples(scene Scene, rnd *rand.Rand, io ImageOptions, i, j, numSamples int) {
	for s := 0; s < numSamples; s++ {
		sampleColor := acc.takeSample(scene, rnd, io, i, j)
		acc.colorSum = acc.colorSum.Add(sampleColor)
		acc.colorSumSq = acc.colorSumSq.Add(geometry.Vector{
			X: sampleColor.X * sampleColor.X,
			Y: sampleColor.Y * sampleColor.Y,
			Z: sampleColor.Z * sampleColor.Z,
		})
		acc.sampleCount++
	}
}

func (acc *colorAccumulator) takeSample(scene Scene, rnd *rand.Rand, io ImageOptions, i, j int) geometry.Vector {
	u := (float64(i) + rnd.Float64()) / float64(io.nX)
	v := (float64(j) + rnd.Float64()) / float64(io.nY)
	ray := scene.camera.getRay(rnd, u, v)
	return colour(ray, scene.objects, 0, rnd)
}

func (acc *colorAccumulator) hasConverged(threshold float64) bool {
	if acc.sampleCount == 0 {
		return false
	}
	
	mean := acc.colorSum.DivideScalar(float64(acc.sampleCount))
	meanSquared := acc.colorSumSq.DivideScalar(float64(acc.sampleCount))
	
	variance := geometry.Vector{
		X: meanSquared.X - mean.X*mean.X,
		Y: meanSquared.Y - mean.Y*mean.Y,
		Z: meanSquared.Z - mean.Z*mean.Z,
	}
	
	maxVariance := math.Max(math.Max(variance.X, variance.Y), variance.Z)
	return maxVariance < threshold
}

func (acc *colorAccumulator) finalColor() color.RGBA {
	finalColor := acc.colorSum.DivideScalar(float64(acc.sampleCount))
	
	return color.RGBA{
		R: uint8(math.Min(255.99*finalColor.X, 255)),
		G: uint8(math.Min(255.99*finalColor.Y, 255)),
		B: uint8(math.Min(255.99*finalColor.Z, 255)),
		A: 255,
	}
}

func renderPixelAdaptive(scene Scene, rnd *rand.Rand, io ImageOptions, i, j int) color.RGBA {
	config := newAdaptiveSamplingConfig(io.nS)
	accumulator := newColorAccumulator()
	
	// Take initial minimum samples
	accumulator.addSamples(scene, rnd, io, i, j, config.minSamples)
	
	// Adaptive sampling loop - continue until converged or max samples reached
	for accumulator.sampleCount < config.maxSamples {
		if accumulator.hasConverged(config.varianceThreshold) {
			break
		}
		
		batchSize := min(config.batchSize, config.maxSamples-accumulator.sampleCount)
		accumulator.addSamples(scene, rnd, io, i, j, batchSize)
	}
	
	return accumulator.finalColor()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}