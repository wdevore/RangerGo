package api

const (
	// FILLED polygon
	FILLED = 0
	// OUTLINED polygon
	OUTLINED = 1
	// FILLOUTLINED both fill and outlined
	FILLOUTLINED = 2
)

// IRenderContext represents visual rendering context
type IRenderContext interface {
	// Initialize render context
	Initialize()

	// Apply transform to current context transform
	Apply(IAffineTransform)

	// Pre draw
	Pre()

	// Save pushes the current state onto the stack
	Save()

	// Restore pops the current state
	Restore()

	// Post
	Post()

	// TransformPoint transforms an IPoint using the current context.
	TransformPoint(p, out IPoint)

	// TransformLine transforms a line/rectangle-corners using the current context.
	TransformLine(p1, p2, out1, out2 IPoint)

	// TransformArray transforms a array of vertices using the current context
	// into output bucket.
	TransformArray(vertices, bucket []IPoint)

	TransformMesh(mesh IMesh)

	TransformPolygon(poly IPolygon)

	// ,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.
	// Draw functions render directly to the device.
	// ,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.

	DrawPoint(x, y int32)

	DrawLine(x1, y1, x2, y2 int32)

	DrawRectangle(rect IRectangle)

	DrawFilledRectangle(rect IRectangle)

	DrawCheckerBoard(size int)

	// ,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.
	// Render functions render based on transformed vertices
	// The Render functions use the Draw functions above.
	// ,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.

	RenderLine(x1, y1, x2, y2 float64)

	RenderMesh(mesh IMesh)

	// Render an axis aligned rectangle. Rotating any of the vertices
	// will cause strange rendering behaviours
	RenderAARectangle(min, max IPoint, fill int)
}
