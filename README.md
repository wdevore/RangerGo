# RangerGo
**RangerGo** A variation of the [Ranger Dart](https://github.com/wdevore/Ranger-Dart) game engine but written in [Go](https://golang.org/) and [SDL](https://www.libsdl.org/download-2.0.php), and possibly [Vulkan](https://www.khronos.org/vulkan/).

# Current Tasks and Goals
* -- **Done** -- Node Dragging
* -- **Done** -- Filters: transform and translate
* -- **Done** -- Zoom Node
* -- **Done** -- Interpolation
* -- **Done** -- Simple motion animations
* -- **Done** -- Circle, AABB
* **Working** AnchorNode
* Particles
* Animation (tweening)
* Enhance raster fonts to allow transforms
* Sprite Textures
* Box2D physics
* OpenGL Vulkan
* Audio (SFXR 8bit sound)
* Joysticks
* Finish lower case Vector font characters.

## Notes

# Tracking
Some **Nodes**/**Objects** may want to *Track* the properties of another **Node**.

For example, an AABB object may wan't to track **Mesh** changes on a node such that it can *rebuild* its internal min/max properties.