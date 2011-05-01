// Package imageutil implements functions for the manipulation of images.
package imageutil

import "image"

// ResizeNearestNeighbor returns a new NRGBA image with the given width and
// height created by resizing the given image using the nearest neighbor
// algorithm.
func ResizeNearestNeighbor(img image.Image, newWidth, newHeight int) *image.NRGBA {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	nimg := image.NewNRGBA(newWidth, newHeight)

	xr := (w<<16)/newWidth + 1
	yr := (h<<16)/newHeight + 1

	for yo := 0; yo < newHeight; yo++ {
		y2 := (yo * yr) >> 16
		for xo := 0; xo < newWidth; xo++ {
			x2 := (xo * xr) >> 16
			nimg.Set(xo, yo, img.At(x2, y2))
			//Much faster, but requires some image type.
			//nimg.Pix[offset] = img.Pix[y2*w+x2]
			//offset++
		}
	}
	return nimg
}

// ResizeBilinear returns a new NRGBA image with the given width and height
// created by resizing the given NRGBA image using the bilinear interpolation.
func ResizeBilinear(img *image.NRGBA, newWidth, newHeight int) *image.NRGBA {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	xr := float32(w-1) / float32(newWidth)
	yr := float32(h-1) / float32(newHeight)

	nimg := image.NewNRGBA(newWidth, newHeight)
	offset := 0

	for yo := 0; yo < newHeight; yo++ {
		y := int(yr * float32(yo))
		dy := yr*float32(yo) - float32(y)
		ody := 1.0 - dy
		for xo := 0; xo < newWidth; xo++ {
			x := int(xr * float32(xo))
			dx := xr*float32(xo) - float32(x)
			odx := 1.0 - dx
			i := y*w + x
			a := img.Pix[i]
			b := img.Pix[i+1]
			c := img.Pix[i+w]
			d := img.Pix[i+w+1]

			alpha := float32(a.A)*odx*ody + float32(b.A)*dx*ody +
				float32(c.A)*dy*odx + float32(d.A)*dx*dy

			red := float32(a.R)*float32(a.A)*odx*ody + float32(b.R)*float32(b.A)*dx*ody +
				float32(c.R)*float32(c.A)*dy*odx + float32(d.R)*float32(d.A)*dx*dy

			green := float32(a.G)*float32(a.A)*odx*ody + float32(b.G)*float32(b.A)*dx*ody +
				float32(c.G)*float32(c.A)*dy*odx + float32(d.G)*float32(d.A)*dx*dy

			blue := float32(a.B)*float32(a.A)*odx*ody + float32(b.B)*float32(b.A)*dx*ody +
				float32(c.B)*float32(c.A)*dy*odx + float32(d.B)*float32(d.A)*dx*dy

			aavg := (float32(a.A) + float32(b.A) + float32(c.A) + float32(d.A)) / 4.0
			if aavg > 0 {
				red /= aavg
				green /= aavg
				blue /= aavg
				if red > 255 {
					red = 255
				}
				if green > 255 {
					green = 255
				}
				if blue > 255 {
					blue = 255
				}
			} else {
				red = 0
				green = 0
				blue = 0
			}
			nimg.Pix[offset] = image.NRGBAColor{
				uint8(red),
				uint8(green),
				uint8(blue),
				uint8(alpha),
			}
			offset++
		}
	}
	return nimg
}
