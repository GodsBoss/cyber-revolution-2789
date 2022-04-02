package scale

type ByInteger struct {
	UnscaledHeight int
	UnscaledWidth  int

	HorizontalMargin int
	VerticalMargin   int

	scale int
}

func (byInteger *ByInteger) Scale() int {
	return byInteger.scale
}

func (byInteger *ByInteger) Recalculate(availableWidth, availableHeight int) {
	wf := (availableWidth - byInteger.HorizontalMargin) / byInteger.UnscaledWidth
	hf := (availableHeight - byInteger.VerticalMargin) / byInteger.UnscaledHeight

	f := wf
	if hf < f {
		f = hf
	}
	if f < 1 {
		f = 1
	}

	byInteger.scale = f
}

func (byInteger *ByInteger) RealSize() (realWidth, realHeight int) {
	return byInteger.UnscaledWidth * byInteger.scale, byInteger.UnscaledHeight * byInteger.scale
}
