package mocks

type Database struct{}

func (db *Database) Save(data string) error {
	// Simulate saving data
	return nil
}
