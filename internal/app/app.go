package app

import (
	"context"

	"github.com/go-chi/jwtauth/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	auth_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/auth"
	dashboard_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/dashboard"
	group_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/group"
	jobcontrol_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/jobs_control"
	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
	prediction_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/prediction"
	user_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/user"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
	group_repository "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/repository/group"
	jobcontrol_repository "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/repository/job_control"
	match_repository "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/repository/match"
	prediction_repository "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/repository/prediction"
	user_repository "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/repository/user"
)

type App struct {
	Config config.Config

	DB *postgres.Client

	HashService hashService
	JWTService  jwtService

	FootballAPI footballAPIClient
}

type hashService interface {
	Hash(password string) (string, error)
	Compare(hash string, password string) bool
}

type jwtService interface {
	Generate(user entity.User) (string, error)
	Auth() *jwtauth.JWTAuth
}

type footballAPIClient interface {
	FetchWorldCupMatches(ctx context.Context) ([]match_usecase.ExternalMatchOutput, error)
}

func New(
	ctx context.Context,
	cfg config.Config,
	db *postgres.Client,
	hashService hashService,
	jwtService jwtService,
	footballAPI footballAPIClient,
) *App {
	return &App{
		Config:      cfg,
		DB:          db,
		HashService: hashService,
		JWTService:  jwtService,
		FootballAPI: footballAPI,
	}
}
func (a *App) NewAuthUseCase(dbtx postgres.DBTX) auth_usecase.AuthUseCaseInterface {
	return auth_usecase.NewAuthUseCase(
		user_repository.NewUserRepository(dbtx),
		a.HashService,
		a.JWTService,
	)
}

func (a *App) NewUserUseCase(dbtx postgres.DBTX) user_usecase.UserUseCaseInterface {
	return user_usecase.NewUserUseCase(
		user_repository.NewUserRepository(dbtx),
	)
}

func (a *App) NewGroupUseCase(dbtx postgres.DBTX) group_usecase.GroupUseCaseInterface {
	return group_usecase.NewGroupUseCase(
		group_repository.NewGroupRepository(dbtx),
	)
}

func (a *App) NewMatchUseCase(dbtx postgres.DBTX) match_usecase.MatchUseCaseInterface {
	return match_usecase.NewMatchUseCase(
		match_repository.NewMatchRepository(dbtx),
		a.FootballAPI,
	)
}

func (a *App) NewPredictionUseCase(dbtx postgres.DBTX) prediction_usecase.PredictionUseCaseInterface {
	return prediction_usecase.NewPredictionUseCase(
		prediction_repository.NewPredictionRepository(dbtx),
		match_repository.NewMatchRepository(dbtx),
		group_repository.NewGroupRepository(dbtx),
	)
}

func (a *App) NewJobControlUseCase(dbtx postgres.DBTX) jobcontrol_usecase.JobControlUsecaseInterface {
	return jobcontrol_usecase.NewJobControlUsecase(
		jobcontrol_repository.NewJobControlRepository(dbtx),
		a.NewMatchUseCase(dbtx),
		a.NewPredictionUseCase(dbtx),
	)
}

func (a *App) NewDashboardUseCase(dbtx postgres.DBTX) dashboard_usecase.DashboardUseCaseInterface {
	return dashboard_usecase.NewDashboardUseCase(
		group_repository.NewGroupRepository(dbtx),
		match_repository.NewMatchRepository(dbtx),
		prediction_repository.NewPredictionRepository(dbtx),
	)
}
