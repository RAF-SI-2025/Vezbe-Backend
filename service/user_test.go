package service

import (
	"testing"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func resetUsers() {
	users = nil
}

func TestCreateUser(t *testing.T) {
	resetUsers()

	user, err := CreateUser("John", "Doe", "john@example.com", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.FirstName != "John" || user.LastName != "Doe" || user.Email != "john@example.com" {
		t.Errorf("user fields mismatch: %+v", user)
	}
	if user.ID == (uuid.UUID{}) {
		t.Error("expected non-zero UUID")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("secret")); err != nil {
		t.Error("password was not hashed correctly")
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user in store, got %d", len(users))
	}
}

func TestGetAllUsers(t *testing.T) {
	resetUsers()

	if got := GetAllUsers(); len(got) != 0 {
		t.Errorf("expected empty slice, got %d users", len(got))
	}

	if _, err := CreateUser("Alice", "Smith", "alice@example.com", "pass"); err != nil {
		t.Fatal(err)
	}
	if _, err := CreateUser("Bob", "Jones", "bob@example.com", "pass"); err != nil {
		t.Fatal(err)
	}

	if got := GetAllUsers(); len(got) != 2 {
		t.Errorf("expected 2 users, got %d", len(got))
	}
}

func TestGetUserByID(t *testing.T) {
	resetUsers()

	created, _ := CreateUser("Jane", "Doe", "jane@example.com", "pass")

	got, err := GetUserByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected user %v, got %v", created.ID, got.ID)
	}

	_, err = GetUserByID(uuid.New())
	if err == nil {
		t.Error("expected error for non-existent ID")
	}
}

func TestUpdateUser(t *testing.T) {
	resetUsers()

	created, _ := CreateUser("Old", "Name", "old@example.com", "pass")

	updated, err := UpdateUser(created.ID, "New", "Name", "new@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.FirstName != "New" || updated.Email != "new@example.com" {
		t.Errorf("update not applied: %+v", updated)
	}

	_, err = UpdateUser(uuid.New(), "X", "Y", "z@example.com")
	if err == nil {
		t.Error("expected error for non-existent ID")
	}
}

func TestDeleteUser(t *testing.T) {
	resetUsers()

	created, _ := CreateUser("Del", "Me", "del@example.com", "pass")

	if err := DeleteUser(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("expected 0 users after delete, got %d", len(users))
	}

	if err := DeleteUser(uuid.New()); err == nil {
		t.Error("expected error for non-existent ID")
	}
}

func TestChangePassword(t *testing.T) {
	resetUsers()

	created, _ := CreateUser("Pass", "User", "pass@example.com", "oldpass")

	if err := ChangePassword(created.ID, "oldpass", "newpass"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Verify new password is stored
	if err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte("newpass")); err != nil {
		t.Error("new password hash does not match")
	}

	// Wrong old password
	if err := ChangePassword(created.ID, "wrongpass", "newpass"); err == nil {
		t.Error("expected error for wrong old password")
	}

	// Non-existent user
	if err := ChangePassword(uuid.New(), "oldpass", "newpass"); err == nil {
		t.Error("expected error for non-existent ID")
	}
}

func TestCheckPassword(t *testing.T) {
	resetUsers()

	user, _ := CreateUser("Check", "Pass", "check@example.com", "mypassword")

	if err := CheckPassword(user, "mypassword"); err != nil {
		t.Errorf("expected no error for correct password: %v", err)
	}

	if err := CheckPassword(user, "wrongpassword"); err == nil {
		t.Error("expected error for incorrect password")
	}
}
