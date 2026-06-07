package server

import (
	"emperror.dev/errors"
)

var (
	ErrIsRunning            = errors.New("Serveur démarré")
	ErrSuspended            = errors.New("Serveur suspendu")
	ErrServerIsInstalling   = errors.New("Installation en cours")
	ErrServerIsTransferring = errors.New("server is currently being transferred")
	ErrServerIsRestoring    = errors.New("server is currently being restored")
)

type crashTooFrequent struct{}

func (e *crashTooFrequent) Error() string {
	return "[FrameMC] Le serveur a crashé trop de fois, vérifier les erreurs"
}

func IsTooFrequentCrashError(err error) bool {
	_, ok := err.(*crashTooFrequent)

	return ok
}

type serverDoesNotExist struct{}

func (e *serverDoesNotExist) Error() string {
	return "server does not exist on remote system"
}

func IsServerDoesNotExistError(err error) bool {
	_, ok := err.(*serverDoesNotExist)

	return ok
}
