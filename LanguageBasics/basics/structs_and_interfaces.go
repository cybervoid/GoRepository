package basics

import (
  "fmt"
  "math"
)

type circle struct {
  diameter, circumference, radius, area float64
}

// In my case, I am both setting the struct's value and
// returning the desired result.  This is entirely a
// personal preference, its not necessary to both setting
// the value and returning it from my function.

func (c *circle) getRadius() float64 {
  c.radius = c.diameter / 2
  return c.radius
}

func (c *circle) getArea() float64 {
    if !(c.radius > 0) {
      c.getRadius()
    }
    c.area = math.Pi * c.radius * c.radius
    return c.area
}
func (c *circle) getCircumference() float64 {
  if !(c.radius > 0) {
    c.getRadius()
  }
  c.circumference = 2 * math.Pi * c.radius
  return c.circumference
}

func StructsInterfacesExample() {
  c := new(circle)
  c.diameter = 12
  rad := c.getRadius()
  c.getArea()
  c.getCircumference()
  fmt.Println("radius: ", c.radius)
  fmt.Println("area: ", c.area)
  fmt.Println("circumference: ", c.circumference)
}
