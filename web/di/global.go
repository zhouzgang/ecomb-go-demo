package di

var globalDI DependencyInjecter

func SetGlobalDI(di DependencyInjecter) {
	globalDI = di
}

func GlobalDI() *Container {
	return globalDI.(*Container)
}