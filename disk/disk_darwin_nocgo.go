//go:build (darwin && !cgo) || ios
// +build darwin,!cgo ios

package disk

import (
	"context"

	"github.com/somma-inc/gopsutil/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
