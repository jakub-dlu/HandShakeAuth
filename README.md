# Hand-Shake Authentication
Hand Shake Authentication (HSAuth) in GO-Lang.
Advanced authentication system using two pairs of keys. The server and the user only exchange public keys and calculate the 'shared secret'. If they are the same, it means that the authentication process was successful. The security phenomenon of this method is that we do not store the password on the server.

# Libraries used in the project
List of libraries used:

# GIN
Gin was used for routing and, above all, communication with the client. Gin is a high-performance web framework for Go that is designed to create web applications and APIs with ease.

# GORM & SQLITE3
Gorm was used to communicate with the database. GORM is a popular ORM (Object-Relational Mapping) library for Golang that simplifies database interactions by allowing developers to work with Go structs instead of raw SQL queries. SQLite3 is a lightweight, disk-based database that doesn’t require a separate server process and allows accessing the database using a nonstandard variant of the SQL query language.

# UUID
Used for unique user id.
