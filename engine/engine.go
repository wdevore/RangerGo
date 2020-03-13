package engine

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

const (
	second = 1000000000
	// fps             = 60.0
	// framePeriod     = 1.0 / fps * 1000.0
	updatePerSecond = 30
	updatePeriod    = float64(second) / float64(updatePerSecond)
)

type engine struct {
	// -----------------------------------------
	// Application properties
	// -----------------------------------------
	// App window size
	world api.IWorld

	// -----------------------------------------
	// SDL properties
	// -----------------------------------------
	window  *sdl.Window
	surface *sdl.Surface
	texture *sdl.Texture

	// -----------------------------------------
	// Graphic properties
	// -----------------------------------------
	pixels     *image.RGBA // Drawing buffer
	bounds     image.Rectangle
	clearColor color.RGBA

	// -----------------------------------------
	// Engine properties
	// -----------------------------------------
	running bool

	// -----------------------------------------
	// Scene graph root node.
	// -----------------------------------------
	// root api.INode // deprecate

	nodeMan api.INodeManager
}

// New constructs a Engine object.
// The Engine runs the main loop.
func New(world api.IWorld) api.IEngine {
	o := new(engine)
	o.world = world
	o.running = false
	o.clearColor = rendering.NewPaletteInt64(rendering.Orange).Color()

	// o.root = nodes.NewNode()
	// o.root.Initialize("Root")

	o.nodeMan = nodes.NewNodeManager(world)
	return o
}

func (e *engine) Configure() {
	var err error

	fmt.Println("Initializing SDL..")
	err = sdl.Init(sdl.INIT_TIMER | sdl.INIT_VIDEO | sdl.INIT_EVENTS)
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating window...")
	e.window, err = sdl.CreateWindow(
		e.world.Title(),
		100, 100,
		int32(e.world.WindowSize().X()), int32(e.world.WindowSize().Y()),
		sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	// Using GetSurface requires using window.UpdateSurface() rather than renderer.Present.
	// v.surface, err = v.window.GetSurface()
	// if err != nil {
	// 	panic(err)
	// }
	// v.renderer, err = sdl.CreateSoftwareRenderer(v.surface)
	// OR create renderer manually

	fmt.Println("Creating renderer...")
	renderer, errR := sdl.CreateRenderer(
		e.window, -1, sdl.RENDERER_SOFTWARE)
	if errR != nil {
		panic(err)
	}

	e.world.SetRenderer(renderer)

	fmt.Println("Creating renderer texture...")
	e.texture, err = renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING,
		int32(e.world.WindowSize().X()), int32(e.world.WindowSize().Y()))
	if err != nil {
		panic(err)
	}

	e.bounds = image.Rect(0, 0, int(e.world.WindowSize().X()), int(e.world.WindowSize().Y()))
	e.pixels = image.NewRGBA(e.bounds)

	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	fmt.Println("Configure complete.")
}

// func (e *engine) Root() api.INode {
// 	return e.root
// }

// Start see api.go for docs
func (e *engine) Start() {
	fmt.Println(("Engine starting..."))

	e.running = true

	// ***************************
	// Debugging only
	// ***************************
	stepEnabled := false
	lag := int64(0)
	nsPerUpdate := int64(math.Round(updatePeriod))
	frameDt := float64(nsPerUpdate) / 1000000.0
	upsCnt := 0
	ups := 0
	fpsCnt := 0
	fps := 0
	previousT := time.Now()
	secondCnt := int64(0)
	renderElapsedCnt := int64(0)
	// ***************************

	// var frameStart time.Time
	// var elapsedTime float64
	// var loopTime float64

	// Get a reference to SDL's internal keyboard state. It is updated
	// during sdl.PollEvent()
	// keyState := sdl.GetKeyboardState()

	sdl.SetEventFilterFunc(e.filterEvent, nil)

	for e.running {
		currentT := time.Now()

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Handle Events
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		sdl.PumpEvents()

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Update
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		elapsedT := currentT.Sub(previousT)
		previousT = currentT
		elapsedNano := elapsedT.Nanoseconds()

		// Note: This update loop is based on:
		// https://gameprogrammingpatterns.com/game-loop.html

		if !stepEnabled {
			lag += elapsedNano
			lagging := true
			for lagging {
				if lag >= nsPerUpdate {
					e.nodeMan.Update(frameDt)
					lag -= nsPerUpdate
					upsCnt++
				} else {
					lagging = false
				}
			}
		}

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Render
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		e.nodeMan.PreVisit()
		// **** Any rendering must occur AFTER this point ****

		// Capture time AFTER pre_visit otherwise if vsync is enabled
		// then time includes the vertical refresh which ~16.667ms
		renderT := time.Now()

		interpolation := float64(lag) / float64(nsPerUpdate)

		moreScenes := e.nodeMan.Visit(interpolation)

		// time.Sleep(time.Millisecond * 1)

		if !moreScenes {
			// fmt.Println("No more scenes!")
			// e.running = false
			// continue
		}

		renderElapsedCnt += (time.Now().Sub(renderT)).Nanoseconds()

		secondCnt += elapsedNano
		if secondCnt >= second {
			avgRender := (float64(renderElapsedCnt) / float64(fpsCnt)) / 1000000.0
			fmt.Printf("fps (%2d), ups (%2d) rend (%2.4f)\n", fps, ups, avgRender)
			fps = fpsCnt
			ups = upsCnt
			upsCnt = 0
			fpsCnt = 0
			secondCnt = 0
			renderElapsedCnt = 0
		}

		// time.Sleep(time.Millisecond * 10)

		// Render scene graph
		// e.root.Draw(e.context)
		// e.renderRawOverlay(elapsedTime, loopTime)
		// e.window.UpdateSurface()

		fpsCnt++

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Present
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		e.nodeMan.PostVisit()
	}
}

func (e *engine) End() {
	fmt.Println("Engine shutting down...")
	fmt.Println("Disposing texture...")
	e.texture.Destroy()
	fmt.Println("Disposing renderer...")
	e.world.Renderer().Destroy()
	fmt.Println("Disposing window...")
	e.window.Destroy()
	fmt.Println("Quitting SDL...")
	sdl.Quit()

	fmt.Println("Done. Goodbye.")
}

// DisplaySize see api.go for docs
func (e *engine) DisplaySize() (w, h int) {
	return int(e.world.WindowSize().X()), int(e.world.WindowSize().Y())
}

// SetClearColor see api.go for docs
func (e *engine) SetClearColor(color color.RGBA) {
	e.clearColor = color
}

// ==============================================================
// Internals
// ==============================================================
var event = nodes.NewEvent()

// filterEvent returns false if it handled the event. Returning false
// prevents the event from being added to the queue.
func (e *engine) filterEvent(ev sdl.Event, userdata interface{}) bool {
	switch t := ev.(type) {
	case *sdl.QuitEvent:
		e.running = false
		return false // We handled it. Don't allow it to be added to the queue.
	case *sdl.MouseMotionEvent:
		event.SetType(api.IOTypeMouseMotion)
		event.SetState(t.State)
		event.SetWhich(t.Which)
		event.SetMousePosition(t.X, t.Y)
		event.SetMouseRelMovement(t.XRel, t.YRel)
		e.nodeMan.RouteEvents(event)

		// fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
		return false // We handled it. Don't allow it to be added to the queue.
	case *sdl.MouseButtonEvent:
		event.SetType(api.IOTypeMouseButton)
		event.SetWhich(t.Which)
		event.SetClicks(t.Clicks)
		event.SetButton(t.Button)
		event.SetMousePosition(t.X, t.Y)
		e.nodeMan.RouteEvents(event)
		return false
		// fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	case *sdl.MouseWheelEvent:
		event.SetType(api.IOTypeMouseWheel)
		event.SetWhich(t.Which)
		event.SetMouseRelMovement(t.X, t.Y)
		event.SetDirection(t.Direction)
		e.nodeMan.RouteEvents(event)
		return false
		// fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y)
	case *sdl.KeyboardEvent:
		if t.State == sdl.PRESSED {
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_ESCAPE:
				e.running = false
				return false
			}
		}
		event.SetType(api.IOTypeKeyboard)
		event.SetState(uint32(t.State))
		event.SetRepeat(t.Repeat)
		event.SetKeyScan(uint32(t.Keysym.Scancode))
		event.SetKeyCode(uint32(t.Keysym.Sym))
		e.nodeMan.RouteEvents(event)
		// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
		// 	t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
		return false
	}

	// True means we didn't handled it. Allow it to be queued.
	return true
}

// deprecated
func (e *engine) clearDisplay() {
	// for y := 0; y < int(e.height); y++ {
	// 	for x := 0; x < int(e.width); x++ {
	// 		e.pixels.SetRGBA(x, y, e.clearColor)
	// 	}
	// }
	c := e.clearColor
	renderer := e.world.Renderer()

	renderer.SetDrawColor(c.R, c.G, c.B, c.A)
	renderer.Clear()
	// e.renderer.Present()
}
