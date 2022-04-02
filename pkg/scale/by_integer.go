package scale

type ByInteger struct {
	UnscaledHeight int
	UnscaledWidth  int

	HorizontalMargin int
	VerticalMargin   int
}

func (byInteger ByInteger) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scale float64) {
	if byInteger.UnscaledWidth <= 0 || byInteger.UnscaledHeight <= 0 {
		return 0, 0, 1
	}

	wf, hf := (availableWidth-byInteger.HorizontalMargin)/byInteger.UnscaledWidth, (availableHeight-byInteger.VerticalMargin)/byInteger.UnscaledHeight

	f := wf
	if hf < f {
		f = hf
	}
	if f < 1 {
		f = 1
	}

	return byInteger.UnscaledWidth * f, byInteger.UnscaledHeight * f, float64(f)
}
