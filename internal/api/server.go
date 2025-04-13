package api

import (
	"net/http"

	"github.com/EnricRG/resteduler-backend/api/mapper"
	"github.com/EnricRG/resteduler-backend/logging"
	"github.com/EnricRG/resteduler-backend/service"
	"github.com/labstack/echo/v4"
	log "github.com/rs/zerolog"
)

func (s *server) Api() []ApiEntry {
	return []ApiEntry{
		{"GET", "/api/v1/users/:uid/schedules", s.getUserSchedules},
	}
}

type server struct {
	port     uint16
	services service.Services
	logger   log.Logger
}

func NewServer(port uint16, services service.Services) *server {
	return &server{
		port,
		services,
		logging.GetLogger(),
	}
}

func (s *server) getUserSchedules(ctx echo.Context) (err error) {
	s.logUrl(ctx)

	var userId uint
	if err = echo.PathParamsBinder(ctx).Uint("uid", &userId).BindError(); err != nil {
		s.logger.Error().Err(err).Msg("Could not bind user identifier from url")
		return ctx.String(http.StatusBadRequest, "Invalid user identifier")
	}

	userSchedules, err := s.services.Shcedules.GetUserSchedules(userId)
	if err != nil {
		s.logger.Error().Err(err).Msgf("Unexpected error while retrieving schedules for user '%d'", userId)
		return s.internalServerError(ctx)
	}

	viewSchedules := mapper.NewScheduleMapper().ToViews(userSchedules)
	return ctx.JSON(http.StatusOK, viewSchedules)
}

// logUrl logs request URI.
func (s *server) logUrl(ctx echo.Context) {
	s.logger.Info().Str("url", ctx.Request().URL.String())
}

// internalServerError returns generic internal server error into context.
func (s *server) internalServerError(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "Internal Server Error")
}
