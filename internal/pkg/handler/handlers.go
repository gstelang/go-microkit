// package handlers

// injects the service and returns the http.handler

// func CreateYourServiceHandler(ys *YourService, logger *log.Logger) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 				// Use the YourService to perform user-related operations
// 				// Example:
// 				// user, err := YourService.GetUserByID(1)
// 				// if err != nil {
// 				//   // Handle error
// 				// }
// 		}
// 	)
// }