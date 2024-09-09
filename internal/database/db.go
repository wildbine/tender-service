package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func InitDB(db *sql.DB) {
    _, err := db.Exec(`
    CREATE TYPE IF NOT EXISTS organization_type AS ENUM (
        'IE',
        'LLC',
        'JSC'
    );

    CREATE TABLE IF NOT EXISTS employee (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS organization (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        description TEXT,
        type organization_type,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS organization_responsible (
        id SERIAL PRIMARY KEY,
        organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
        user_id INT REFERENCES employee(id) ON DELETE CASCADE
    );
    `)
    if err != nil {
        log.Fatalf("Unable to execute the queries: %v\n", err)
    }

}