package execute

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"distri/internal/proto"
)

var (
	ErrToManyRetries = errors.New("too many retries")
)

type Worker interface {
	ExecuteAction(ctx context.Context, name string, data []byte) (any, error)
}

type ServerConfig struct {
	ActionTimeout time.Duration
	RetriesAmount byte
}

type Server struct {
	proto.UnimplementedActionExecutorServiceServer
	cfg    ServerConfig
	worker Worker
}

func NewServer(cfg ServerConfig, worker Worker) *Server {
	return &Server{
		cfg:    cfg,
		worker: worker,
	}
}

func (s *Server) Execute(ctx context.Context, req *proto.ExecuteActionRequest) (*proto.ExecuteActionResponse, error) {
	ctx, cancel := context.WithTimeoutCause(ctx, s.cfg.ActionTimeout, fmt.Errorf("execution of action: %s timeouted", req.ActionName))
	defer cancel()

	retried := 0
	for range s.cfg.RetriesAmount {
		if retried >= int(s.cfg.RetriesAmount) {
			return nil, fmt.Errorf("%w: count: %d", ErrToManyRetries, retried)
		}

		res, err := s.worker.ExecuteAction(ctx, req.ActionName, req.Payload)
		if err != nil {
			retried++

			continue
		}

		payload, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("marshal response: %w", err)
		}

		return &proto.ExecuteActionResponse{Result: payload}, nil
	}

	return nil, nil
}
