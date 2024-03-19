package database

func (database *Database) InitStatements() {
	// AUTH
	// register
	database.prepareStatement("register", "INSERT INTO users (name, hash) VALUES ($1, $2) RETURNING id")
	// user existance
	database.prepareStatement("userExists", "SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)")
	// hash and user id
	database.prepareStatement("getUserAndHash", "SELECT id, hash FROM users WHERE name = $1")

	// GAMES
	// full list
	database.prepareStatement("getGames", "SELECT id, title FROM games")
}
