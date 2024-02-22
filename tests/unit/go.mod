module github.com/ibrahim-shaaban-hussein/structure/tests/unit

go 1.18

require (
	github.com/gin-gonic/gin v1.9.1 // Add this line for Gin framework
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers v0.0.0-20240222095042-bec9cdd758b9
	go.mongodb.org/mongo-driver v1.14.0
)

replace (
  // Replace indirect dependencies to specific versions if needed, e.g.:
  github.com/stretchr/testify => github.com/stretchr/testify v1.8.3
)
