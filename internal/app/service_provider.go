package app

import (
	"context"
	"example.com/m/internal/api/note"
	"example.com/m/internal/config"
	"example.com/m/internal/repo"
	"example.com/m/internal/repo/repoUser"
	"example.com/m/internal/service"
	noteService "example.com/m/internal/service/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
	noteRepository repo.NoteRepo

	noteService service.Service

	noteImpl *note.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		s.pgPool = pool
	}

	// В данном контексте не указан s.dbClient, поэтому предположим, что вы хотели использовать s.pgPool
	return s.pgPool
}

func (s *serviceProvider) NoteRepository(ctx context.Context) repo.NoteRepo {
	if s.noteRepository == nil {
		s.noteRepository = repoUser.NewRepository(s.DBClient(ctx))
	}
	return s.noteRepository
}

func (s *serviceProvider) NoteService(ctx context.Context) service.Service {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.NoteRepository(ctx))
	}

	return s.noteService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteService(ctx))
	}

	return s.noteImpl
}
