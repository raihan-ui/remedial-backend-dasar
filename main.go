package main

import (
	"fmt"
)

// Student struct represents a student
type Student struct {
	ID        string
	FirstName string
	LastName  string
}

// Course struct represents a course
type Course struct {
	ID   string
	Name string
}

// Grade struct represents a grade for a course
type Grade struct {
	CourseID string
	Grade    int
}

func main() {
	var students []Student
	var courses []Course
	var grades []Grade

	for {
		fmt.Println("Pilih operasi:")
		fmt.Println("1. Input Data Mahasiswa")
		fmt.Println("2. Input Detail Mata Kuliah")
		fmt.Println("3. Lihat Detail Mahasiswa")
		fmt.Println("4. Lihat Detail Mata Kuliah")
		fmt.Println("5. Lihat Nilai")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var student Student
			fmt.Print("ID Mahasiswa: ")
			fmt.Scanln(&student.ID)
			fmt.Print("Nama Depan Mahasiswa: ")
			fmt.Scanln(&student.FirstName)
			fmt.Print("Nama Belakang Mahasiswa: ")
			fmt.Scanln(&student.LastName)
			students = append(students, student)
			fmt.Println("Data Mahasiswa berhasil disimpan.")
		case 2:
			var course Course
			fmt.Print("ID Mata Kuliah: ")
			fmt.Scanln(&course.ID)
			fmt.Print("Nama Mata Kuliah: ")
			fmt.Scanln(&course.Name)
			courses = append(courses, course)
			fmt.Println("Data Mata Kuliah berhasil disimpan.")
		case 3:
			fmt.Println("Detail Mahasiswa:")
			for _, student := range students {
				fmt.Printf("ID: %s, Nama: %s %s\n", student.ID, student.FirstName, student.LastName)
			}
		case 4:
			fmt.Println("Detail Mata Kuliah:")
			for _, course := range courses {
				fmt.Printf("ID: %s, Nama: %s\n", course.ID, course.Name)
			}
		case 5:
			fmt.Println("Nilai Mata Kuliah:")
			for _, grade := range grades {
				for _, student := range students {
					for _, course := range courses {
						if grade.CourseID == course.ID && grade.CourseID == student.ID {
							fmt.Printf("Mahasiswa: %s %s, Mata Kuliah: %s, Nilai: %d\n", student.FirstName, student.LastName, course.Name, grade.Grade)
						}
					}
				}
			}
		case 0:
			fmt.Println("Terima kasih telah menggunakan program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}
