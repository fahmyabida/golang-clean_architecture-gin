package middleware

//import "github.com/gin-gonic/gin"
//
//const (
//	ACCESS_TOKEN_KEY = "Access-Token"
//)
//
//type goMiddleware struct {
//	// another stuff , may be needed by middleware
//}
//
//type responseError struct {
//	Message string `json:"message"`
//}
//
//func (m *goMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
//	return func(c gin.Context) error {
//		c.Header("Access-Control-Allow-Origin", "*")
//		return next(c)
//	}
//}
//
//func InitMiddleware() *goMiddleware {
//	return &goMiddleware{}
//}
