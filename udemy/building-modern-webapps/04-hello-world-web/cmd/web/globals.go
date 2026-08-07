package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/gerdreiss/go-course/pkg/config"
)

var app config.AppConfig
var session *scs.SessionManager

const addr = ":8080"
