package workers

import (
	"context"
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/services"
)

type TokenCleanupWorker struct {
	tokenService *services.FCMTokenService
	interval     time.Duration
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewTokenCleanupWorker(tokenService *services.FCMTokenService, interval time.Duration) *TokenCleanupWorker {
	ctx, cancel := context.WithCancel(context.Background())
	return &TokenCleanupWorker{
		tokenService: tokenService,
		interval:     interval,
		ctx:          ctx,
		cancel:       cancel,
	}
}

func (w *TokenCleanupWorker) Start() {
	go func() {
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				w.tokenService.CleanupTokens()
			case <-w.ctx.Done():
				return
			}
		}
	}()
}

func (w *TokenCleanupWorker) Stop() {
	w.cancel()
}
