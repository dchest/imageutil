package imageutil

import (
	"os"
	"image/png"
	"image"
	"testing"
)

func testImage() (f *os.File, size int64, img image.Image) {
	f, err := os.Open("testdata/1.png")
	if err != nil {
		panic(err)
	}
	img, err = png.Decode(f)
	if err != nil {
		panic(err)
	}
	size = int64(img.Bounds().Max.X * img.Bounds().Max.Y * 4)
	return
}

func BenchmarkNearestNeighbor(b *testing.B) {
	b.StopTimer()
	f, size, img := testImage()
	defer f.Close()
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ResizeNearestNeighbor(img, w*2, h*2)
		b.SetBytes(size)
	}
}

func BenchmarkBilinear(b *testing.B) {
	b.StopTimer()
	f, size, img := testImage()
	nimg, ok := img.(*image.NRGBA)
	if !ok {
		panic("image not NRGBA")
	}
	defer f.Close()
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ResizeBilinear(nimg, w*2, h*2)
		b.SetBytes(size)
	}
}
