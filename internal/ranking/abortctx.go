package ranking

import "context"

func abortRankContext() error {
	ctx := liveRankContext()
	if ctx == nil {
		return nil
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func liveRankContext() context.Context {
	return context.Background()
}
