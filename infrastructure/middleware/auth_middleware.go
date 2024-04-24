package middleware

// import (
//     "net/http"
//     "github.com/filmons/go-ingnerd-backend/authentication"
// )

// // AuthMiddleware is a middleware for authentication
// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Check authentication logic here
//         authenticated := authentication.CheckAuthentication(r)
//         if !authenticated {
//             http.Error(w, "Unauthorized", http.StatusUnauthorized)
//             return
//         }
//         // Call the next handler
//         next.ServeHTTP(w, r)
//     })
// }
