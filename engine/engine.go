package engine

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

const (
	second          = 1000000000
	framesPerSec    = 60.0
	framePeriod     = 1.0 / framesPerSec * 1000.0
	updatePerSecond = 30
	updatePeriod    = float64(second) / float64(updatePerSecond)

	// The "present" call is impacted by this flag, by an order of magnitude.
	enabledVSync = true

	displayStatsVisibility = true

	renderMaxCnt = int64(50)
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
	// texture *sdl.Texture

	// -----------------------------------------
	// Graphic properties
	// -----------------------------------------
	// pixels     *image.RGBA // Drawing buffer
	// bounds     image.Rectangle
	clearColor color.RGBA

	// -----------------------------------------
	// Engine properties
	// -----------------------------------------
	running bool

	// -----------------------------------------
	// Scene graph is a node manager
	// -----------------------------------------
	sceneGraph api.INodeManager

	// -----------------------------------------
	// Debug
	// -----------------------------------------
	stepEnabled bool
	statsColor  api.IPalette
}

// New constructs a Engine object.
// The Engine runs the main loop.
func New(world api.IWorld) api.IEngine {
	o := new(engine)

	o.world = world
	o.running = false
	o.clearColor = rendering.NewPaletteInt64(rendering.Orange).Color()
	o.stepEnabled = false

	o.sceneGraph = nodes.NewNodeManager(world)

	o.statsColor = rendering.NewPaletteInt64(rendering.Orange)

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
	flags := uint32(sdl.RENDERER_ACCELERATED)
	if enabledVSync {
		flags |= sdl.RENDERER_PRESENTVSYNC
	}
	renderer, errR := sdl.CreateRenderer(
		e.window, -1, flags)
	if errR != nil {
		panic(err)
	}

	e.world.SetRenderer(renderer)

	// Setup a callback called during PumpEvents
	sdl.SetEventFilterFunc(e.filterEvent, nil)

	// fmt.Println("Creating renderer texture...")
	// e.texture, err = renderer.CreateTexture(
	// 	sdl.PIXELFORMAT_ABGR8888,
	// 	sdl.TEXTUREACCESS_STREAMING,
	// 	int32(e.world.WindowSize().X()), int32(e.world.WindowSize().Y()))
	// if err != nil {
	// 	panic(err)
	// }

	// e.bounds = image.Rect(0, 0, int(e.world.WindowSize().X()), int(e.world.WindowSize().Y()))
	// e.pixels = image.NewRGBA(e.bounds)

	// renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	fmt.Println("Configure complete.")
}

// Start see api.go for docs
func (e *engine) Start() {
	fmt.Println(("Engine starting..."))

	e.running = true

	// ***************************
	// Debugging only
	// ***************************
	lag := int64(0)
	nsPerUpdate := int64(math.Round(updatePeriod))
	frameDt := float64(nsPerUpdate) / 1000000.0
	upsCnt := 0
	ups := 0
	fpsCnt := 0
	fps := 0
	previousT := time.Now()
	secondCnt := int64(0)
	renderElapsedTime := int64(0)

	// presentElapsedCnt := int64(0)
	renderCnt := int64(0)
	avgRender := 0.0
	// ***************************

	for e.running {
		currentT := time.Now()

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Handle Events
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// In order for filterFunc to trigger we need to repeatedly call
		// PumpEvents.
		sdl.PumpEvents()

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Update
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		elapsedNano := (currentT.Sub(previousT)).Nanoseconds()

		// Note: This update loop is based on:
		// https://gameprogrammingpatterns.com/game-loop.html

		if !e.stepEnabled {
			lag += elapsedNano
			lagging := true
			for lagging {
				if lag >= nsPerUpdate {
					e.sceneGraph.Update(frameDt)
					lag -= nsPerUpdate
					upsCnt++
				} else {
					lagging = false
				}
			}
		}

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Render Scenegraph by visiting the nodes
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		renderT := time.Now()

		e.sceneGraph.PreVisit()
		// **** Any rendering must occur AFTER this point ****

		// Calc interpolation for nodes that need it.
		interpolation := float64(lag) / float64(nsPerUpdate)

		// Once the last scene has exited the stage we stop running.
		moreScenes := e.sceneGraph.Visit(interpolation)

		if !moreScenes {
			e.running = false
			continue
		}

		if renderCnt >= renderMaxCnt {
			avgRender = float64(renderElapsedTime) / float64(renderMaxCnt) / 1000.0
			renderCnt = 0
			renderElapsedTime = 0
		} else {
			renderElapsedTime += (time.Now().Sub(renderT)).Microseconds()
			renderCnt++
		}

		secondCnt += elapsedNano
		if secondCnt >= second {
			// fmt.Printf("fps (%2d), ups (%2d), rend (%2.4f)\n", fps, ups, avgRender)
			// fmt.Printf("secCnt %d, fpsCnt %d, presC %d\n", secondCnt, fpsCnt, presentElapsedCnt)

			fps = fpsCnt
			ups = upsCnt
			upsCnt = 0
			fpsCnt = 0
			secondCnt = 0
		}

		if displayStatsVisibility {
			e.drawStats(fps, ups, avgRender)
		}

		// time.Sleep(time.Millisecond * 1)

		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// Finish rendering
		// ~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--~--
		// presentT := time.Now()
		// SDL present's elapsed time is different if vsync is on or off.
		e.sceneGraph.PostVisit()
		// presentElapsedCnt = (time.Now().Sub(presentT)).Nanoseconds()

		fpsCnt++
		previousT = currentT
	}
}

func (e *engine) End() {
	fmt.Println("Engine shutting down...")

	// fmt.Println("Disposing texture...")
	// e.texture.Destroy()
	fmt.Println("Disposing renderer...")
	e.world.Renderer().Destroy()
	fmt.Println("Disposing window...")
	e.window.Destroy()
	fmt.Println("Quitting SDL...")
	sdl.Quit()

	fmt.Println("Done. Goodbye.")
}

func (e *engine) EnableStepping(enable bool) {
	e.stepEnabled = enable
}

func (e *engine) PushStart(node api.INode) {
	e.sceneGraph.PushNode(node)
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
		e.sceneGraph.RouteEvents(event)

		// fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
		return false // We handled it. Don't allow it to be added to the queue.
	case *sdl.MouseButtonEvent:
		event.SetType(api.IOTypeMouseButton)
		event.SetWhich(t.Which)
		event.SetClicks(t.Clicks)
		event.SetButton(t.Button)
		event.SetMousePosition(t.X, t.Y)
		e.sceneGraph.RouteEvents(event)
		return false
		// fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	case *sdl.MouseWheelEvent:
		event.SetType(api.IOTypeMouseWheel)
		event.SetWhich(t.Which)
		event.SetMouseRelMovement(t.X, t.Y)
		event.SetDirection(t.Direction)
		e.sceneGraph.RouteEvents(event)
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
		e.sceneGraph.RouteEvents(event)
		// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
		// 	t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
		return false
	}

	// True means we didn't handled it. Allow it to be queued.
	return true
}

func (e *engine) drawStats(fps, ups int, avgRend float64) {
	world := e.world
	text := fmt.Sprintf("%2d, %2d, %2.4f", fps, ups, avgRend)

	x := 15.0
	y := world.WindowSize().Y() - 15.0

	world.Context().SetDrawColor(e.statsColor)
	world.Context().DrawText(x, y, text, 1, 1, false)
}
