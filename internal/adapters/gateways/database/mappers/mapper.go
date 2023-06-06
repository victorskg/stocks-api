package mappers

type EntityToDomainMapper[T any, P any] interface {
	FromEntityToDomain(entity T) *P
}

type DomainToEntityMapper[T any, P any] interface {
	FromDomainToEntity(domain P) *T
}

type Mapper[T any, P any] interface {
	FromEntityToDomain(entity T) *P
	FromDomainToEntity(domain P) *T
}
