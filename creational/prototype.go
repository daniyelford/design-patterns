package main

import "fmt"

// Prototype interface
type Prototype interface {
	Clone() Prototype
	Show()
}

// User struct
type User struct {
	Name        string
	Role        string
	Permissions []string
}

func (u *User) Clone() Prototype {
	// ایجاد کپی عمیق از Slice
	newPermissions := make([]string, len(u.Permissions))
	copy(newPermissions, u.Permissions)

	return &User{
		Name:        u.Name,
		Role:        u.Role,
		Permissions: newPermissions,
	}
}

func (u *User) Show() {
	fmt.Printf("👤 %s (%s)\n", u.Name, u.Role)
	fmt.Printf("🔑 Permissions: %v\n\n", u.Permissions)
}

func main() {
	admin := &User{"Admin", "Administrator", []string{"create", "edit", "delete"}}
	admin.Show()

	user1 := admin.Clone().(*User)
	user1.Name = "Daniyal"
	user1.Role = "Editor"
	user1.Permissions = append(user1.Permissions, "publish")

	user1.Show()
	admin.Show() // admin هنوز بدون تغییره
}
