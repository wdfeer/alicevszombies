package internal

type Entity = uint32
type World struct {
	nextID Entity
	// TODO: add components like positions, etc
}

func NewWorld() World {
	return World{
		nextID: 1,
	}
}

func (world *World) NewEntity() Entity {
	id := world.nextID
	world.nextID++
	return id
}

func (world *World) Update() {

}

func (world World) Render() {

}
