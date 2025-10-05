package main

/*

Create a new type: vehicle.
	○ The underlying type is a struct.
	○ The fields:
		■ doors
		■ color
● Create two new types: truck & sedan.
	○ The underlying type of each of these new types is a struct.
	○ Embed the “vehicle” type in both truck & sedan
	○ Give truck the field “fourWheel” which will be set to bool.
	○ Give sedan the field “luxury” which will be set to bool. solution
● Using the vehicle, truck, and sedan structs:
	○ using a composite literal, create a value of type truck and assign values to the
	fields;
	○ using a composite literal, create a value of type sedan and assign values to the
	fields.
● Print out each of these values.
● Print out a single field from each of these values.
*/
import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: "4",
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: "2",
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(t.doors, t.color, t.fourWheel)
	fmt.Println(s.doors, s.color, s.luxury)
}
