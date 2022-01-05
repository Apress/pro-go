package services

type lifecycle int

const (
    Transient lifecycle = iota
    Singleton
    Scoped
)
