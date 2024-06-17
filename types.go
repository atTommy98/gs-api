package main

type UnitPostJSON struct {
	Value string
}

type PackageSizeJSON struct {
	Value string
}

type SPackage struct {
	Id       int
	Size     int
	Required int
}
