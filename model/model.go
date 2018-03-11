package model

type db interface {
    SelectPeople(p *pgDb) ([]*Person, error)
}

type Model struct {
    db
}

func New(db db) *Model {
    return &Model{
        db: db,
    }
}

func (m *Model) People() ([]*Person, error) {
    return m.SelectPeople()
}

type Person struct {
    Id          int64
    First, Last string
}