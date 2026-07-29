package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error){
	//bcrypt package to convert password to hashsecure password, but first we convert to byte slice the strign password. We input 14(how secure) as secure hashpassword
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	return  string(bytes),err
}


func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))

	//it returns true if password is valid and invalid if it return error
	return err == nil
}