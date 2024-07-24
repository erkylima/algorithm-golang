package ports

type Command[I any, T any] interface {
	execute() T
}
