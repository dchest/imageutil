**Warning: this package is a work in progress; in a very primitive state right now.**

Package imageutil
=====================

	import "github.com/dchest/imageutil"

Package imageutil implements functions for the manipulation of images.


Functions
---------

### func ResizeBilinear

	func ResizeBilinear(img *image.NRGBA, newWidth, newHeight int) *image.NRGBA
	
ResizeBilinear returns a new NRGBA image with the given width and height
created by resizing the given NRGBA image using the bilinear interpolation.

### func ResizeNearestNeighbor

	func ResizeNearestNeighbor(img image.Image, newWidth, newHeight int) *image.NRGBA
	
ResizeNearestNeighbor returns a new NRGBA image with the given width and
height created by resizing the given image using the nearest neighbor
algorithm.
