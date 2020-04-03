# Physics examples
The physic examples are based on ByteArena's [Go port](https://github.com/ByteArena/box2d) of [Box2D](https://box2d.org/).

Some of the examples are built following [iforce2d](https://www.iforce2d.net/)'s tutorials.

All examples can run simply by using:

```> go run .```

at the *CLI* within the specific directory.

-----------------------------------------------------------------
## Basic
Basic is as simple as it gets.

First a circle node is create for the visual aspect of the physic engine:

```Go
g.circleNode = custom.NewCircleNode("Orange Circle", world, g)
g.circleNode.SetScale(5.0)
g.circleNode.SetPosition(0.0, -100.0)
```

Then we coerse the node to **CircleNode** concrete type so we can call the circle's specific behaviours:

```Go
gr := g.circleNode.(*custom.CircleNode)
gr.Configure(12, 1.0)
gr.SetColor(rendering.NewPaletteInt64(rendering.Orange))
```

Next we start building the physics world by using Box2D. That amounts to creating **Gravity**, a **World**, a definition object for creating a **Body**, and finally a **Shape** to attach to a **Fixture**:

```Go
// Define the gravity vector.
// Ranger's coordinate space is defined as:
// .--------> +X
// |
// |
// |
// v +Y
// Thus gravity is specified as positive for downward motion.
g.b2Gravity = box2d.MakeB2Vec2(0.0, 9.8)

// Construct a world object, which will hold and simulate the rigid bodies.
g.b2World = box2d.MakeB2World(g.b2Gravity)

// A body def used to create body
bd := box2d.MakeB2BodyDef()
bd.Type = box2d.B2BodyType.B2_dynamicBody
bd.Position.Set(g.circleNode.Position().X(), g.circleNode.Position().Y())

// An instance of a body to contain the Fixtures
g.b2CircleBody = g.b2World.CreateBody(&bd)

// Every Fixture has a shape
circleShape := box2d.MakeB2CircleShape()
circleShape.M_p.Set(0.0, 0.0) // Relative to body position
circleShape.M_radius = 5

fd := box2d.MakeB2FixtureDef()
fd.Shape = &circleShape
fd.Density = 10.0
g.b2CircleBody.CreateFixtureFromDef(&fd) // attach Fixture to body
```

Now we need to *Step* the physics **World** and update the visual based on the *Step*. Note: we pass the *fractional seconds* instead of milliseconds because Box2D works by seconds not milliseconds:

```Go
// Instruct the world to perform a single step of simulation.
// It is generally best to keep the time step and iterations fixed.
g.b2World.Step(secPerUpdate, api.VelocityIterations, api.PositionIterations)

pos := g.b2CircleBody.GetPosition()
if g.b2CircleBody.IsActive() {
   g.circleNode.SetPosition(pos.X, pos.Y)
}
```

-----------------------------------------------------------------

## Basic Ground

*Basic Ground* takes *Basic* to the next level by adding a **Static** body to represent the ground *and* adding a rotation component to the circle. I purposefully placed the **Ground** a *little-off to the side* so that the circle *glances* off edge which will cause a **Rotation** to be applied.

In order to see the rotation a White line was added to a custom circle.

Just like *Basic*, *Visual* nodes must be created in order to play the part of what you see--remember Box2D has **No** visual component:

```Go
g.circleNode = NewCircleNode("Orange Circle", world, g)
gr := g.circleNode.(*CircleNode)
gr.Configure(6, 1.0)
gr.SetColor(rendering.NewPaletteInt64(rendering.Orange))
gr.SetScale(3.0)
gr.SetPosition(100.0, -100.0)

g.groundLineNode = custom.NewLineNode("Ground", world, g)
gln := g.groundLineNode.(*custom.LineNode)
gln.SetColor(rendering.NewPaletteInt64(rendering.White))
gln.SetPoints(-1.0, 0.0, 1.0, 0.0) // Set by unit coordinates
gln.SetPosition(76.0+50.0, 0.0)
gln.SetScale(25.0)
```

You can see above I was lazy by using the *Concrete* type to call all methods including the **INode** methods. 

As in *Basic* we now create the *Physics* components, but this time we create *two* bodies **and** *two* shapes:

```Go
// Every Fixture has a shape
circleShape := box2d.MakeB2CircleShape()
circleShape.M_p.Set(0.0, 0.0) // Relative to body position
circleShape.M_radius = g.circleNode.Scale()

fd := box2d.MakeB2FixtureDef()
fd.Shape = &circleShape
fd.Density = 1.0
g.b2CircleBody.CreateFixtureFromDef(&fd) // attach Fixture to body

// -------------------------------------------
// The Ground = body + fixture + shape
bDef.Type = box2d.B2BodyType.B2_staticBody
bDef.Position.Set(g.groundLineNode.Position().X(), g.groundLineNode.Position().Y())

g.b2GroundBody = g.b2World.CreateBody(&bDef)

groundShape := box2d.MakeB2EdgeShape()
groundShape.Set(box2d.MakeB2Vec2(-g.groundLineNode.Scale(), 0.0), box2d.MakeB2Vec2(g.groundLineNode.Scale(), 0.0))

fDef := box2d.MakeB2FixtureDef()
fDef.Shape = &groundShape
fDef.Density = 1.0
g.b2GroundBody.CreateFixtureFromDef(&fDef) // attach Fixture to body
```

Next we add the physics *Step* code along with updating the visuals based on the *Step*:

```Go
g.b2World.Step(secPerUpdate, api.VelocityIterations, api.PositionIterations)

if g.b2CircleBody.IsActive() {
   pos := g.b2CircleBody.GetPosition()
   g.circleNode.SetPosition(pos.X, pos.Y)

   rot := g.b2CircleBody.GetAngle()  // <-- New rotational code
   g.circleNode.SetRotation(rot)
}

if g.b2GroundBody.IsActive() {
   pos := g.b2GroundBody.GetPosition()
   g.groundLineNode.SetPosition(pos.X, pos.Y)
}
```
