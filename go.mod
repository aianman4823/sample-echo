module github.com/aianman4823/sample-echo

go 1.14

replace local.packages/handler => ./handler

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/labstack/echo/v4 v4.1.15
	local.packages/handler v0.0.0-00010101000000-000000000000 // indirect
)
