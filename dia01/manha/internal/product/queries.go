package product

const (
	GetAll  = "SELECT p.id, p.name, p.price, p.quantity, p.type FROM products AS p"
	GetById = "SELECT p.id, p.name, p.price, p.quantity, p.type FROM products AS p WHERE p.id = ?"
)
