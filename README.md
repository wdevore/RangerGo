# RangerGo
**RangerGo** A variation of the [Ranger Dart](https://github.com/wdevore/Ranger-Dart) game engine but written in [Go](https://golang.org/) and [SDL](https://www.libsdl.org/download-2.0.php), and possibly [Vulkan](https://www.khronos.org/vulkan/).

# Current Tasks and Goals
* Node Dragging
* Zoom Node
* Particles
* Animation
* Interpolation
* Filters: transform and translate
* Enhance raster fonts to allow transforms
* More custom nodes: Circle, AnchorNode, AABB
* Sprite Textures
* Box2D physics
* OpenGL Vulkan
* Audio (SFXR 8bit sound)
* Joysticks
* Complete lower case Vector font characters.

# Examples
The examples are designed to show how to use **Ranger**. Ultimately you will build your own Nodes as you build your game.

All examples can run simply by using:

```> go run .```

at the *CLI* within the specific directory.

## Basic
Basic is as simple as it gets. It starts the engine, prints the display size and waits for 3 seconds.

## Minimal
Minimal takes the next step and creates the two basic **Nodes** that make up a minimalistic shell: a Boot and Scene nodes.

The Boot **Node** exist solely to demonstrate transitioning from one Scene to another. A Boot node would typically be used for things that need to happen before any other node runs, perhaps a connection to the network.

The Scene node shows an example of drawing a background--checkboard in this example--and a *hand coded* line. Typically you would create an actually **Node**, for example, there is a custom **LineNode** you would normally use.

## GameLayer
GameLayer emphasizes a child *Layer* Node as the background instead of a Scene Node. Games should always use the Scene Node for transitions and *Layer* type Nodes for backgrounds and game content.

## Raster text
Raster text is a simple example of using the raster font. The font is a stripped down port from the [font8x8](https://crates.io/crates/font8x8/0.2.3) [Rust](https://www.rust-lang.org/) crate. At the time of this writing enhancements to allow transformations on the font are planned.

## Vector text
Vector text is a simple example of the using a handcrafted Vector font. The lower case characters are not complete at the time of this writing, however, it is quite usable. **TODO** complete Vector font lower case characters.

## Space mappings 1
There are several main coordinates spaces: World, Device, View and Node.

* **World**-space is infinite in size
* **Device**-space is finite in size and limited by the physical structure of the device, for example, a cell phone or mouse coordinates (aka **mouse-space**).
* **View**-space is a virtual space artificially defined by the developer, similar in concept to virtual consoles like [Pico8](https://www.lexaloffle.com/pico-8.php). It is still infinite in size but virtual bounds are mapped to **Device** space.
* **Node**-space is the coordinate space relative to a **Node**, typically seen as local-space.

Space mappings #1 example shows how to map mouse-space to view-space using the *MapDeviceToView* method.

Every **Node** can register for events, either Timming or IO. This example shows how a node registers itself for IO events via Enter/Exit lifecycle events.

## Space mappings 2
Space mapping #2 adds to #1 by using the provided custom/RectangleNode. This node uses the *MapDeviceToNode* and *PointInside* methods, and serves as an example for mapping techniques.

```
    nodes.MapDeviceToNode(r.world, mx, my, r, r.localPosition)
	r.pointInside = r.polygon.PointInside(r.localPosition)
```

This allows the node's *Draw* to change the node's draw color according to the *pointInside* boolean.

```
	if r.pointInside {
		context.SetDrawColor(r.insideColor)
	} else {
		context.SetDrawColor(r.color)
	}
	context.RenderPolygon(r.polygon, api.CLOSED)
```

### Example Image
<img src="docs/RangerGo_space-mappings-2.png" alt="alt text" width="500" >

