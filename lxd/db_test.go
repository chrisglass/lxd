package main

import (
    "fmt"
    "testing"
    "database/sql"
)

func Test_deleting_a_container_cascades_on_profiles(t *testing.T){
    var db *sql.DB
    var err error
    var count int

    db, err = initializeDbObject(":memory:")
    if err != nil {
        t.Error(err)
    }

    // Insert a container and a related profile (for example)
    statements := `
    INSERT INTO containers (name, architecture, type) VALUES ('thename', 1, 1);
    INSERT INTO profiles (name) VALUES ('theprofile');
    INSERT INTO containers_profiles (container_id, profile_id) VALUES (1, 1);`

    _, err = db.Exec(statements)
    if err != nil {
        t.Error(err)
    }

    // Drop the container we just created.
    statements = `DELETE FROM containers WHERE name = 'thename';`

    _, err = db.Exec(statements)
    if err != nil {
        t.Error(fmt.Sprintf("Error deleting container! %s", err))
    }

    // Make sure there is 0 container_profiles entries left.
    statements = `SELECT count(*) FROM containers_profiles;`
    err = db.QueryRow(statements).Scan(&count)

    if count != 0 {
        t.Error(fmt.Sprintf("Deleting a container didn't delete the profile association! There are %d left", count))
    }
}
