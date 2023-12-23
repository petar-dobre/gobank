package models

type Server interface {
	Run ()
	Routes ()
	AccountHandler
}
