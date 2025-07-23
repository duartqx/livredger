package mimetypes

type MimeType string

const (
	JSON              MimeType = "application/json"
	HTML              MimeType = "text/html; charset=utf-8"
	HTMX              MimeType = "text/htmx"
	PlainText         MimeType = "text/plain; charset=utf-8"
	XML               MimeType = "application/xml"
	FormURLEncoded    MimeType = "application/x-www-form-urlencoded"
	MultipartFormData MimeType = "multipart/form-data"
)
