package ratelimit

import (
	"errors"
	"fmt"
	"net"
)

type RateLimiter struct {
	policy             Policy
	buckets            map[RequestKey]*TokenBucket
	api                Api
	tokenBucketFactory func() *TokenBucket
}

func NewRateLimiter(policy Policy, api Api, tokenBucketFactory func() *TokenBucket) *RateLimiter {
	return &RateLimiter{
		policy:             policy,
		buckets:            make(map[RequestKey]*TokenBucket),
		api:                api,
		tokenBucketFactory: tokenBucketFactory,
	}
}

func (rl *RateLimiter) Process(r Request) (ApiResult, error) {
	key := rl.policy.getRequestKey(r)
	bucket, ok := rl.buckets[key]

	if !ok {
		// no bucket for this request yet, adding
		bucket = rl.tokenBucketFactory()
		rl.buckets[key] = bucket
		err := bucket.Start()
		if err != nil {
			return "", err
		}
	}

	// Bucket exists so it must have been Started already
	if !bucket.Started {
		panic("bucket exists but isn't started")
	}

	if bucket.IsOverflown() {
		return "", errors.New("too many requests")
	}

	return rl.api.execute(r)
}

func (rl *RateLimiter) StopAll() error {
	for key, bucket := range rl.buckets {
		err := bucket.Stop()
		if err != nil {
			return fmt.Errorf("error stopping bucket for %s: %w", key, err)
		}
	}

	return nil
}

type RequestKey string

type Request struct {
	UserId    string
	IpAddress net.IPAddr
}

type Policy interface {
	getRequestKey(request Request) RequestKey
}

type UserIdPolicy struct {
}

func (p UserIdPolicy) getRequestKey(request Request) RequestKey {
	return RequestKey(request.UserId)
}

type IpPolicy struct {
}

func (p IpPolicy) getRequestKey(request Request) RequestKey {
	return RequestKey(request.IpAddress.String())
}

type ApiResult string

type Api interface {
	execute(r Request) (ApiResult, error)
}

type BasicApi struct {
}

func (api BasicApi) execute(r Request) (ApiResult, error) {
	result := ApiResult(fmt.Sprintf("executed request for %s", r.UserId))
	return result, nil
}
