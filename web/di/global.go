package di

var globalDI DependencyInjecter

func SetGlobalDI(di DepencyInjecter) {
	globalDI = di
}
