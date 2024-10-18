package youtube

import (
	"context"
	"errors"
	"fmt"
	"math"
	apperrors "telegraminput/internal/errors"
	"telegraminput/internal/logger"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

type Yt struct {
	yt     *youtube.Service
	logger logger.Logger
}

const (
	caption           = "Взгляните на это видео:\n\n"
	baseURL           = "https://www.youtube.com/watch?v="
	radius            = "1000km"
	maxRetries uint64 = 4
)

func New(ctx context.Context, logger logger.Logger, apiKey string) (*Yt, error) {
	ytClient, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create youtube service: %w", err)
	}

	logger.Debug("youtube service has been started")

	return &Yt{
		yt:     ytClient,
		logger: logger,
	}, nil
}

// Too long random characters query decrease can be a problem
// so we will use 1..3 range.
func mapTriesToQLen(retries uint64) uint {
	var (
		maxLen float64 = 3
		minLen float64 = 1
	)

	return uint(math.Max(minLen, math.Min(maxLen, float64(retries))))
}

func (y *Yt) RandVideo() (RandomVideo, error) {
	operation := func() (string, error) {
		tries := maxRetries

		return func() (string, error) {
			qLen := mapTriesToQLen(tries)
			videoID, err := y.randVideoID(qLen)
			tries--

			if err != nil && !errors.Is(err, apperrors.ErrNotExisted) {
				return "", backoff.Permanent(err)
			}

			return videoID, err
		}()
	}

	backoffPolicy := backoff.WithMaxRetries(
		backoff.NewExponentialBackOff(), maxRetries)

	id, err := backoff.RetryWithData(operation, backoffPolicy)
	if err != nil {
		return RandomVideo{}, fmt.Errorf("failed to get rand video id: %w", err)
	}

	return RandomVideo{Link: baseURL + id, Caption: caption}, nil
}

func (y *Yt) randVideoID(qLen uint) (string, error) {
	var (
		query       string
		order       string
		coordinates string
	)

	var errG errgroup.Group

	errG.Go(func() error {
		var err error
		query, err = randString(qLen)

		return err
	})

	errG.Go(func() error {
		var err error
		order, err = randOrder()

		return err
	})

	errG.Go(func() error {
		var err error
		coordinates, err = randCoordinatesStr()

		return err
	})

	if err := errG.Wait(); err != nil {
		return "", fmt.Errorf("cannot create random data: %w", err)
	}

	return y.videoID(query, order, coordinates, radius)
}

func (y *Yt) videoID(
	query string, order string, coordinates string, radius string,
) (string, error) {
	request := y.yt.Search.List([]string{"id"})

	request.Q(query)
	request.MaxResults(1)
	request.Location(coordinates)
	request.LocationRadius(radius)
	request.Order(order)
	request.Type("video")
	request.RegionCode("ua")

	resp, err := request.Do()
	if err != nil {
		return "", errors.Join(apperrors.ErrExternal, err)
	}

	if len(resp.Items) < 1 {
		return "", apperrors.ErrNotExisted
	}

	id := resp.Items[0].Id.VideoId

	return id, nil
}

