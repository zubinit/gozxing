package gozxing

type Writer interface {
	/**
	 * Encode a barcode using the default settings.
	 *
	 * params: contents The contents to encode in the barcode
	 * params: format The barcode format to generate
	 * params: width The preferred width in pixels
	 * params: height The preferred height in pixels
	 * return: {@link BitMatrix} representing encoded barcode image
	 * throws WriterException if contents cannot be encoded legally in a format
	 */
	EncodeWithoutHint(contents string, format BarcodeFormat, width, height int) (*BitMatrix, error)

	/**
	 * params: contents The contents to encode in the barcode
	 * params: format The barcode format to generate
	 * params: width The preferred width in pixels
	 * params: height The preferred height in pixels
	 * params: hints Additional parameters to supply to the encoder
	 * return: {@link BitMatrix} representing encoded barcode image
	 * throws WriterException if contents cannot be encoded legally in a format
	 */
	Encode(contents string, format BarcodeFormat, width, height int, hints map[EncodeHintType]interface{}) (*BitMatrix, error)
}
