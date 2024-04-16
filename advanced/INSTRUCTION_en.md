# Let's Draw Art Using the Game Engine Ebitengine

[Japanese version](INSTRUCTION_ja.md)

## Final Image

![bubbles](images/goal.png)

## Explanation of the Final Image

We will draw countless ripples spreading in a circle inside a circle.

Drawing rules:

- Place small circles representing ripples at random positions inside the outer circle
  - Assume an outer circle (a circle with a radius of half of screenWidth or screenHeight)
  - The distance from the center of the small circle is determined by log10(1-1000 random number)/3 * radius of the outer circle
  - The angle of the small circle is determined by a random number from 0 - 2π
- With each drawing, the ripples are enlarged and the color is made lighter
- If the maximum number of ripples is exceeded, delete from the old ones
- With each drawing, a new ripple is placed

## Install Ebitengine

Follow the installation instructions on the official site.
https://ebitengine.org/en/documents/install.html?os=linux

Let's proceed until "Hello, World!" is displayed.

## Let's Draw a Circle

Basic vector graphics are provided by the vector package, so let's start by drawing a circle.
https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2/vector

<details>
  <summary>Hint</summary>

You can draw a circle using the `Draw` function with `vector.StrokeCircle`.
</details>

## Let's Change the Circle

Let's change the size (radius r) and color of the circle.
To manage the changes, define a `Bubble struct` with radius r and color c as fields.

<details>
  <summary>Hint</summary>

```go
var (
	Magnification    float32 = 1.03 
	ColorAttenuation float32 = 0.98
)

type Bubble struct {
	R float32
	C uint8 // Gray scale color
}
```

To change the values, modify them in the Update function.

Also, the speed of change depends on the TPS (tick per second) of ebitengine, so adjust it as needed with ebiten.SetTPS(n).

</details>

## Let's Randomly Change the Position of the Circle

Let's also manage the position (x, y) of the circle with the Bubble struct.
You can easily generate random numbers using the `math/rand/v2` package.

## Let's Place a Large Number of Circles

Now that we can place and change one circle, let's make it possible to place a large number of them.
As a start, let's aim to place 100 circles.

To accurately manage the number of circles, we can confirm the count by displaying it using `ebitenutil.DebugPrint()`.

Circles are added and deleted repeatedly, but Go's slice is not suitable for continuous addition and deletion, so it would be good to use container/list.
https://pkg.go.dev/container/list

If you want to use it like a queue, you can use it as follows:

- Add with list.PushFront() (enqueue operation)
- Retrieve and delete with list.Remove(list.Back()) (dequeue operation)

To inspect all elements, follow the documentation.

```go
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}
```

## Let's Make the Circles Align with the Outer Circle

It's nice as it is, but let's give some regularity to the generation of circles.
Let's make a large number of small circles appear along a large outer circle.
To randomly place along the outer circle, it would be good to use trigonometric functions.
Also, let's make the distance from the center of the outer circle random.

<details>
  <summary>Hint</summary>

To determine the distance d from the center of the outer circle with radius R randomly.

```go
d := rand.Float32() * R
```

To place along a circle using trigonometric functions, it is good to decide the angle theta (0-2π) randomly, so

```go
theta := rand.Float32() * math.Pi * 2
```

The image is to place a point at (0, d) and rotate it by angle theta.
The rotation operation can be performed with the following calculation when the origin is (0,0).

```
x2 = x * cos(θ) - y * sin(θ)
y2 = x * sin(θ) + y * cos(θ)
```

</details>

## (Next step) Let's Play Around

- Change the parameters
- Change the color
- Try drawing with something other than a circle
- Accept user input and cause changes
- etc...
