package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func InitDB(db *sql.DB) {
    checkTypeQuery := `SELECT 1 FROM pg_type WHERE typname = 'organization_type';`
    var exists int
    err := db.QueryRow(checkTypeQuery).Scan(&exists)
    
    if err == sql.ErrNoRows {
        _, err = db.Exec(`CREATE TYPE organization_type AS ENUM ('IE', 'LLC', 'JSC');`)
        if err != nil {
            log.Fatalf("Unable to create type organization_type: %v\n", err)
        }
    } else if err != nil {
        log.Fatalf("Error checking for type organization_type: %v\n", err)
    }


	queries := []string{
        `CREATE TABLE IF NOT EXISTS employee (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50) UNIQUE NOT NULL,
            first_name VARCHAR(50),
            last_name VARCHAR(50),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
        `CREATE TABLE IF NOT EXISTS organization (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            description TEXT,
            type organization_type,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
        `CREATE TABLE IF NOT EXISTS organization_responsible (
            id SERIAL PRIMARY KEY,
            organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
            user_id INT REFERENCES employee(id) ON DELETE CASCADE
        );`,
    }

    for _, query := range queries {
        _, err := db.Exec(query)
        if err != nil {
            log.Fatalf("Unable to execute the query: %v\nQuery: %s\n", err, query)
        }
    }
}