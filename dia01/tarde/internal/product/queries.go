package product

const (
	GetAll  = "SELECT p.id, p.name, p.price, p.quantity, p.type FROM products AS p"
	GetById = "SELECT p.id, p.name, p.price, p.quantity, p.type FROM products AS p WHERE p.id = ?"
	Create  = "INSERT INTO products (name, price, quantity, type) VALUES (?, ?, ?, ?)"
	Delete  = "DELETE FROM products WHERE id=?"
)
