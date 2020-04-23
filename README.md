# RangerGo
**RangerGo** A variation of the [Ranger Dart](https://github.com/wdevore/Ranger-Dart) game engine but written in [Go](https://golang.org/) and [SDL](https://www.libsdl.org/download-2.0.php), and possibly [Vulkan](https://www.khronos.org/vulkan/).

# Current Tasks and Goals
* **Done** Node Dragging
* **Done** Filters: transform and translate
* **Done** Zoom Node
* **Done** Interpolation
* **Done** Simple motion animations
* **Done** Circle, AABB
* **Done** AnchorNode
* **Done** Particles
* **Done** Animation (tweening) -- Using tanema's library
* **Working Draft #1 done** Box2D physics (with space ships)
* **Working** Zones (Circle and Rectangle) combined with Zoom
* QuadTree for view-space culling
* OpenGL
* Audio (SFXR 8bit sound) : May build GUI using imGui, May use: https://github.com/faiface/beep
* Simple Widget GUI framework
  * Buttons
  * Checkboxes
  * Text input
  * Text
  * Dialog
  * Grouping (i.e. Radio buttons)
* Sprite Textures (Vulkan)
* Finish lower case Vector font characters.
* Enhance raster fonts to allow transforms
* Joysticks and Gamepads

## Notes
Zooming. We need a global zoom value between zones such that each zone starts
from the global value. This value could be either "from" or "to".

# Tracking (Optional)
Some **Nodes**/**Objects** may want to *Track* the properties of another **Node**.

For example, an AABB object may wan't to track **Mesh** changes on a node such that it can *rebuild* its internal min/max properties.

## Packages

```
go get github.com/tanema/gween
go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}
go get github.com/ByteArena/box2d
```
