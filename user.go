package main

import (
	elastic "gopkg.in/olivere/elastic.v3"

	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TYPE_USER = "user"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// checkUser checks whether user is valid

