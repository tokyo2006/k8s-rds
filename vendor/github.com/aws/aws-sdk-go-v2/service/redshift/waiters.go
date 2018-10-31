// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilClusterAvailable uses the Amazon Redshift API operation
// DescribeClusters to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Redshift) WaitUntilClusterAvailable(input *DescribeClustersInput) error {
	return c.WaitUntilClusterAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterAvailableWithContext is an extended version of WaitUntilClusterAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Redshift) WaitUntilClusterAvailableWithContext(ctx aws.Context, input *DescribeClustersInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilClusterAvailable",
		MaxAttempts: 30,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Clusters[].ClusterStatus",
				Expected: "available",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Clusters[].ClusterStatus",
				Expected: "deleting",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ClusterNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeClustersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeClustersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilClusterDeleted uses the Amazon Redshift API operation
// DescribeClusters to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Redshift) WaitUntilClusterDeleted(input *DescribeClustersInput) error {
	return c.WaitUntilClusterDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterDeletedWithContext is an extended version of WaitUntilClusterDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Redshift) WaitUntilClusterDeletedWithContext(ctx aws.Context, input *DescribeClustersInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilClusterDeleted",
		MaxAttempts: 30,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ClusterNotFound",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Clusters[].ClusterStatus",
				Expected: "creating",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Clusters[].ClusterStatus",
				Expected: "modifying",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeClustersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeClustersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilClusterRestored uses the Amazon Redshift API operation
// DescribeClusters to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Redshift) WaitUntilClusterRestored(input *DescribeClustersInput) error {
	return c.WaitUntilClusterRestoredWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterRestoredWithContext is an extended version of WaitUntilClusterRestored.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Redshift) WaitUntilClusterRestoredWithContext(ctx aws.Context, input *DescribeClustersInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilClusterRestored",
		MaxAttempts: 30,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Clusters[].RestoreStatus.Status",
				Expected: "completed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Clusters[].ClusterStatus",
				Expected: "deleting",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeClustersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeClustersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilSnapshotAvailable uses the Amazon Redshift API operation
// DescribeClusterSnapshots to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Redshift) WaitUntilSnapshotAvailable(input *DescribeClusterSnapshotsInput) error {
	return c.WaitUntilSnapshotAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilSnapshotAvailableWithContext is an extended version of WaitUntilSnapshotAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Redshift) WaitUntilSnapshotAvailableWithContext(ctx aws.Context, input *DescribeClusterSnapshotsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilSnapshotAvailable",
		MaxAttempts: 20,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "Snapshots[].Status",
				Expected: "available",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Snapshots[].Status",
				Expected: "failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "Snapshots[].Status",
				Expected: "deleted",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeClusterSnapshotsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeClusterSnapshotsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}