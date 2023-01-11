package multi

import (
	"github.com/zubinit/gozxing"
)

// MultipleBarcodeReader Implementation of this interface attempt to read several barcodes from one image.
//
// @see com.google.zxing.Reader
type MultipleBarcodeReader interface {
	DecodeMultipleWithoutHint(image *gozxing.BinaryBitmap) ([]*gozxing.Result, error)

	DecodeMultiple(image *gozxing.BinaryBitmap, hints map[gozxing.DecodeHintType]interface{}) ([]*gozxing.Result, error)
}
