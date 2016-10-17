package webpage

type renderFormat string

const (
	RenderFormatPNG  renderFormat = "PNG"
	RenderFormatGIF               = "GIF"
	RenderFormatJPEG              = "JPEG"
	RenderFormatPDF               = "PDF"
	RenderFormatBMP               = "BMP"
	RenderFormatPPM               = "PPM"
)

type RenderOptions struct {
	Format  renderFormat
	Quality int
}
