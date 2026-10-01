package auth

//* Private Method
func extractSession() string {
	return "Loggedin"
}

//* public Method
func Get_session() string {
	return extractSession()
}
