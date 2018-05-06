package basics

import (
  "fmt"
  "math"
)
//define types:
type geometry interface {
  getArea() float64
  getPerim() float64
}
type rectangle struct {
  width, height, perimeter, area float64
}
type circle struct {
  diameter, radius, circumference, area float64
}
// CIRCLE
// interface defined funcs - circle
func (c *circle) getArea() float64 {
    if !(c.radius > 0) {
      c.getRadius()
    }
    c.area = math.Pi * c.radius * c.radius
    return c.area
}
func (c *circle) getPerim() float64 {
  if !(c.radius > 0) {
    c.getRadius()
  }
  c.circumference = 2 * math.Pi * c.radius
  return c.circumference
}
//other interface defined funcs for circle
func (c *circle) getRadius() float64 {
  c.radius = c.diameter / 2
  return c.radius
}

// SQUARE
func (r *rectangle) getArea() float64 {
  return r.width * r.height
}

func (r *rectangle) getPerim() float64 {
  return (2 * r.width) + (2 * r.height)
}



// In my case, I am both setting the struct's value and
// returning the desired result.  This is entirely a
// personal preference, its not necessary to both setting
// the value and returning it from my function.



func StructsInterfacesExample() {
  c := new(circle)
  c.diameter = 12

  runGeometry(c)
  r := new(rectangle)
  r.height = 3
  r.width = 4
  runGeometry(r)
  //circleSpecific(c)
}

func circleSpecific(c circle) {
  c.getArea()
  c.getPerim()
  fmt.Println("radius: ", c.radius)
  fmt.Println("area: ", c.area)
  fmt.Println("circumference: ", c.circumference)
}

func runGeometry(g geometry) {
  fmt.Printf("Running geometry on type: %T \n", g)
  fmt.Println("area: ",g.getArea())
  fmt.Println("perimeter: ",g.getPerim())
}
