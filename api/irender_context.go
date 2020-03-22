package api

const (
	// FILLED polygon
	FILLED = 0
	// OUTLINED polygon
	OUTLINED = 1
	// FILLOUTLINED both fill and outlined
	FILLOUTLINED = 2

	// CLOSED indicates a polygon should be rendered closed
	CLOSED = 0
	// OPEN indicates a polygon should be rendered open
	OPEN = 1
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

	// TransformPoints transforms a line/rectangle-corners using the current context.
	TransformPoints(p1, p2, out1, out2 IPoint)

	// TransformArray transforms a array of vertices using the current context
	// into output bucket.
	TransformArray(vertices, bucket []IPoint)

	TransformMesh(mesh IMesh)

	TransformPolygon(poly IPolygon)

	// ,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.
	// Draw functions render directly to the device.
	// ,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.
	SetDrawColor(color IPalette)

	DrawPoint(x, y int32)

	DrawLine(x1, y1, x2, y2 int32)
	DrawLineUsing(p1, p2 IPoint)

	DrawRectangle(rect IRectangle)

	DrawFilledRectangle(rect IRectangle)

	DrawCheckerBoard(size int)

	DrawText(x, y float64, text string, scale int, fill int, invert bool)

	// ,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.,--.
	// Render functions render based on transformed vertices
	// The Render functions use the Draw functions above.
	// ,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.,__.

	RenderLine(x1, y1, x2, y2 float64)

	RenderLines(mesh IMesh)

	RenderPolygon(poly IPolygon, style int)

	// Render an axis aligned rectangle. Rotating any of the vertices
	// will cause strange rendering behaviours
	RenderAARectangle(min, max IPoint, fill int)

	RenderCheckerBoard(mesh IMesh, oddColor IPalette, evenColor IPalette)
}
