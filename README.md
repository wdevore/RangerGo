# RangerGo
**RangerGo** is a variation of the [Ranger Dart](https://github.com/wdevore/Ranger-Dart) game engine but written in [Go](https://golang.org/) and [SDL](https://www.libsdl.org/download-2.0.php)

# **Update**
This version of Ranger has served its purpose and that was to refine the engine core while using a simple rendering backend (aka SDL). Now that that has been completed work has shifted over to [Ranger-Go-IGE](https://github.com/wdevore/Ranger-Go-IGE) whos goal is to switch out SDL in favor of OpenGL.

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
* **done** Box2D physics (with space ship)
* **Done** Zones combined with Zoom

## Notes

# Tracking (Optional)
Some **Nodes**/**Objects** may want to *Track* the properties of another **Node**.

For example, an AABB object may wan't to track **Mesh** changes on a node such that it can *rebuild* its internal min/max properties.

## Packages

```
go get github.com/tanema/gween
go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}
go get github.com/ByteArena/box2d
```
