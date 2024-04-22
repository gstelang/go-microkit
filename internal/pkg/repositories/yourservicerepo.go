// Define the operations you want to perform.
// Example if you need something from a db, specify the interface....

// UserRepository defines the repository interface for user-related operations.

// type User struct {
// 	ID   int
// 	Name string
// }

// type NewUserServiceRepo interface {
// 	GetUserByID(id int) (*User, error)
// 	CreateUser(user *User) error
// }

// ----------------------------------------
// You can use CQRS to separate read and write.

// ReadService defines the interface for read operations.
// type ReadServiceRepo interface {
// 	GetUserByID(id int) (*User, error)
// }

// // WriteService defines the interface for write operations.
// type WriteServiceRepo interface {
// 	CreateUser(user *User) error
// }
