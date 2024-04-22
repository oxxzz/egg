package egg

import "strconv"

func (c *Context) PathValueInt(key string) int {
	// 32-bit signed integer
	v, _ := strconv.Atoi(c.params[key])
	return v
}

func (c *Context) PathValueInt64(key string) int64 {
	// 64-bit signed integer
	v, _ := strconv.ParseInt(c.params[key], 10, 64)
	return v
}

func (c *Context) PathValueUint(key string) uint {
	// 32-bit unsigned integer
	v, _ := strconv.ParseUint(c.params[key], 10, 0)
	return uint(v)
}

func (c *Context) PathValueUint64(key string) uint64 {
	// 64-bit unsigned integer
	v, _ := strconv.ParseUint(c.params[key], 10, 64)
	return v
}

func (c *Context) PathValueFloat32(key string) float32 {
	// 32-bit float
	v, _ := strconv.ParseFloat(c.params[key], 32)
	return float32(v)
}

func (c *Context) PathValueFloat64(key string) float64 {
	// 64-bit float
	v, _ := strconv.ParseFloat(c.params[key], 64)
	return v
}

func (c *Context) PathValueBool(key string) bool {
	// bool
	v, _ := strconv.ParseBool(c.params[key])
	return v
}
