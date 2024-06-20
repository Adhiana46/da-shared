package httpserver

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	Constants "github.com/adhiana46/da-shared/constants"

	Errors "github.com/adhiana46/da-shared/errors"
	PkgMiddleware "github.com/adhiana46/da-shared/pkg/middlewares/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) initValidator() {
	// Validator
	validatorInstance := validator.New()
	// register function to get tag name from json tags.
	validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("name"), ",", 2)[0]
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})
	s.engine.Validator = NewEchoValidator(validatorInstance)
}

func (s *Server) initErrorHandler() {
	s.engine.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			statusCode int
			message    string
			errorsData any
		)
		stackTraces := []string{
			err.Error(),
		}
		stackTraces = append(stackTraces, strings.Split(strings.ReplaceAll(fmt.Sprintf("%+v", err), "\t", "  "), "\n")...)

		// Handle Errors
		switch {
		case errors.As(err, &validator.ValidationErrors{}):
			errValidation := validator.ValidationErrors{}
			errors.As(err, &errValidation)

			message = Constants.MsgErrorValidation
			statusCode = 400
			errorsData = func(validationErrs validator.ValidationErrors) map[string][]string {
				errorFields := map[string][]string{}
				for _, e := range validationErrs {
					errorFields[e.Field()] = append(errorFields[e.Field()], e.Tag())
				}

				return errorFields
			}(errValidation)
		case errors.As(err, new(Errors.InternalError)):
			errInternal := Errors.NewInternalError()
			errors.As(err, &errInternal)

			message = errInternal.Error()
			statusCode = errInternal.HttpStatusCode()
		default:
			message = Constants.MsgErrorInternalServer
			statusCode = http.StatusInternalServerError

			// Log unexpected error
			slog.Error("Unexpected: ", slog.String("error", err.Error()))
			// TODO: log error perhaps using Sentry.io or another logging service
		}

		details := map[string]any{
			"method":        c.Request().Method,
			"endpoint":      c.Request().URL.String(),
			"client_ip":     c.RealIP(),
			"user_agent":    c.Request().UserAgent(),
			"payload":       c.Get("request_body"),
			"error_message": err.Error(),
			"stack_trace":   stackTraces,
		}
		slog.Error("Error processing HTTP request", slog.Any("details", details))

		err = c.JSON(statusCode, map[string]any{
			"status":       false,
			"message":      message,
			"errors":       errorsData,
			"stack_traces": stackTraces,
		})
		if err != nil {
			slog.Error("Error writing response json", slog.String("error", err.Error()))
		}
	}
}

func (s *Server) initMiddlewares() {
	// CORS Middleware
	s.engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}))

	s.engine.Use(PkgMiddleware.ContextRequestBodyMiddleware())
	s.engine.Use(PkgMiddleware.ContextMetadataMiddleware())
	s.engine.Use(PkgMiddleware.RequestLoggingMiddleware())
}
