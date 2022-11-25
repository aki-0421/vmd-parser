package vmd

import "errors"

func (c *client) parse() (*VMD, error) {
	err := c.parseHeader()
	if err != nil {
		return nil, err
	}

	err = c.parseMotion()
	if err != nil {
		return nil, err
	}

	err = c.parseMorph()
	if err != nil {
		return nil, err
	}

	err = c.parseCamera()
	if err != nil {
		return nil, err
	}

	v := c.VMD

	return v, nil
}

func (c *client) parseHeader() error {
	m, err := c.GetChar(30)
	if err != nil {
		return err
	}
	if m != VMD_MAGIC {
		return errors.New("file is not vmd")
	}

	name, err := c.GetUnicodeChar(20)
	if err != nil {
		return err
	}

	c.Header.Coordinate = "left"
	c.Header.Name = name

	return nil
}

func (c *client) parseMotion() error {
	// motion length
	l, err := c.GetInt()
	if err != nil {
		return err
	}

	// parse motions
	ms := make([]Motion, 0, l)
	for i := 0; i < l; i++ {
		// BoneName
		n, err := c.GetUnicodeChar(15)
		if err != nil {
			return err
		}

		// FrameNum
		f, err := c.GetInt()
		if err != nil {
			return err
		}

		// Position
		p, err := c.GetFloat32Array(3)
		if err != nil {
			return err
		}

		// Rotation
		r, err := c.GetFloat32Array(4)
		if err != nil {
			return err
		}

		// Interpolation
		ip, err := c.GetInt8Array(64)
		if err != nil {
			return err
		}

		m := Motion{
			BoneName:      n,
			FrameNum:      f,
			Position:      p,
			Rotation:      r,
			Interpolation: ip,
		}

		ms = append(ms, m)
	}

	c.Header.MotionCount = l
	c.Motions = ms

	return nil
}

func (c *client) parseMorph() error {
	// morph length
	l, err := c.GetInt()
	if err != nil {
		return err
	}

	// parse morphs
	ms := make([]Morph, 0, l)
	for i := 0; i < l; i++ {
		n, err := c.GetUnicodeChar(15)
		if err != nil {
			return err
		}

		f, err := c.GetInt()
		if err != nil {
			return err
		}

		w, err := c.GetFloat32()
		if err != nil {
			return err
		}

		m := Morph{
			MorphName: n,
			FrameNum:  f,
			Weight:    w,
		}

		ms = append(ms, m)
	}

	c.Header.MorphCount = l
	c.Morphs = ms

	return nil
}

func (c *client) parseCamera() error {
	// camera length
	l, err := c.GetInt()
	if err != nil {
		return err
	}

	// parse camera
	cs := make([]Camera, 0, l)
	for i := 0; i < l; i++ {
		// FrameNum
		fn, err := c.GetInt()
		if err != nil {
			return err
		}

		// Distance
		d, err := c.GetFloat32()
		if err != nil {
			return err
		}

		// Position
		pos, err := c.GetFloat32Array(3)
		if err != nil {
			return err
		}

		// Rotation
		r, err := c.GetFloat32Array(3)
		if err != nil {
			return err
		}

		// Interpolation
		i, err := c.GetInt8Array(24)
		if err != nil {
			return err
		}

		// Fov
		f, err := c.GetInt()
		if err != nil {
			return err
		}

		// Perspective
		p, err := c.GetInt8()
		if err != nil {
			return err
		}

		c := Camera{
			FrameNum:      fn,
			Distance:      d,
			Position:      pos,
			Rotation:      r,
			Interpolation: i,
			Fov:           f,
			Perspective:   p,
		}

		cs = append(cs, c)
	}

	c.Header.CameraCount = l
	c.Cameras = cs

	return nil
}
