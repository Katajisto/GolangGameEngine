package main

type Entity interface {
	Draw()
	GetBasePhys() *BasePhys
}


var entityList []Entity

func AddEntity(entity Entity) {
	entityList = append(entityList, entity)
}

func drawEntities() {
	for _, entity := range entityList {
		entity.Draw()
	}
}

func calcEntityPhys() {
	for _, entity := range entityList {
		entity.GetBasePhys().phys()
	}
}
