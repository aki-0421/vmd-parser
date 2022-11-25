package vmd

const VMD_MAGIC = "Vocaloid Motion Data 0002"

type VMD struct {
	Header
	Motions []Motion
	Morphs  []Morph
	Cameras []Camera
}

type Header struct {
	Coordinate  string
	Name        string
	MotionCount int
	MorphCount  int
	CameraCount int
}

type Motion struct {
	BoneName      string
	FrameNum      int
	Position      []float32
	Rotation      []float32
	Interpolation []int8
}

type Morph struct {
	MorphName string
	FrameNum  int
	Weight    float32
}

type Camera struct {
	FrameNum      int
	Distance      float32
	Position      []float32
	Rotation      []float32
	Interpolation []int8
	Fov           int
	Perspective   int8
}
