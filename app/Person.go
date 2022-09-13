package app

type Person struct {
	ID        int64  //`json:"id"`
	email     string //`json:"name"`
	phone     string //`json:"created_at"`
	firstName string //`json:"updated_at"`
	lastName  string //`json:"updated_at"`
}

// AuthorRepository represent the author's repository contract
//type AuthorRepository interface {
//	GetByID(ctx context.Context, id int64) (Author, error)
//}
