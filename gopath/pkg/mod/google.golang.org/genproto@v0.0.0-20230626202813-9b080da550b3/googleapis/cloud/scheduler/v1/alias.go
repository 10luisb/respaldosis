// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package scheduler aliases all exported identifiers in package
// "cloud.google.com/go/scheduler/apiv1/schedulerpb".
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package scheduler

import (
	src "cloud.google.com/go/scheduler/apiv1/schedulerpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/scheduler/apiv1/schedulerpb
const (
	HttpMethod_DELETE                  = src.HttpMethod_DELETE
	HttpMethod_GET                     = src.HttpMethod_GET
	HttpMethod_HEAD                    = src.HttpMethod_HEAD
	HttpMethod_HTTP_METHOD_UNSPECIFIED = src.HttpMethod_HTTP_METHOD_UNSPECIFIED
	HttpMethod_OPTIONS                 = src.HttpMethod_OPTIONS
	HttpMethod_PATCH                   = src.HttpMethod_PATCH
	HttpMethod_POST                    = src.HttpMethod_POST
	HttpMethod_PUT                     = src.HttpMethod_PUT
	Job_DISABLED                       = src.Job_DISABLED
	Job_ENABLED                        = src.Job_ENABLED
	Job_PAUSED                         = src.Job_PAUSED
	Job_STATE_UNSPECIFIED              = src.Job_STATE_UNSPECIFIED
	Job_UPDATE_FAILED                  = src.Job_UPDATE_FAILED
)

// Deprecated: Please use vars in: cloud.google.com/go/scheduler/apiv1/schedulerpb
var (
	File_google_cloud_scheduler_v1_cloudscheduler_proto = src.File_google_cloud_scheduler_v1_cloudscheduler_proto
	File_google_cloud_scheduler_v1_job_proto            = src.File_google_cloud_scheduler_v1_job_proto
	File_google_cloud_scheduler_v1_target_proto         = src.File_google_cloud_scheduler_v1_target_proto
	HttpMethod_name                                     = src.HttpMethod_name
	HttpMethod_value                                    = src.HttpMethod_value
	Job_State_name                                      = src.Job_State_name
	Job_State_value                                     = src.Job_State_value
)

// App Engine target. The job will be pushed to a job handler by means of an
// HTTP request via an
// [http_method][google.cloud.scheduler.v1.AppEngineHttpTarget.http_method]
// such as HTTP POST, HTTP GET, etc. The job is acknowledged by means of an
// HTTP response code in the range [200 - 299]. Error 503 is considered an App
// Engine system error instead of an application error. Requests returning
// error 503 will be retried regardless of retry configuration and not counted
// against retry counts. Any other response code, or a failure to receive a
// response before the deadline, constitutes a failed attempt.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type AppEngineHttpTarget = src.AppEngineHttpTarget

// App Engine Routing. For more information about services, versions, and
// instances see [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type AppEngineRouting = src.AppEngineRouting

// CloudSchedulerClient is the client API for CloudScheduler service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type CloudSchedulerClient = src.CloudSchedulerClient

// CloudSchedulerServer is the server API for CloudScheduler service.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type CloudSchedulerServer = src.CloudSchedulerServer

// Request message for
// [CreateJob][google.cloud.scheduler.v1.CloudScheduler.CreateJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type CreateJobRequest = src.CreateJobRequest

// Request message for deleting a job using
// [DeleteJob][google.cloud.scheduler.v1.CloudScheduler.DeleteJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type DeleteJobRequest = src.DeleteJobRequest

// Request message for
// [GetJob][google.cloud.scheduler.v1.CloudScheduler.GetJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type GetJobRequest = src.GetJobRequest

// The HTTP method used to execute the job.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type HttpMethod = src.HttpMethod

// Http target. The job will be pushed to the job handler by means of an HTTP
// request via an
// [http_method][google.cloud.scheduler.v1.HttpTarget.http_method] such as HTTP
// POST, HTTP GET, etc. The job is acknowledged by means of an HTTP response
// code in the range [200 - 299]. A failure to receive a response constitutes a
// failed execution. For a redirected request, the response returned by the
// redirected request is considered.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type HttpTarget = src.HttpTarget
type HttpTarget_OauthToken = src.HttpTarget_OauthToken
type HttpTarget_OidcToken = src.HttpTarget_OidcToken

// Configuration for a job. The maximum allowed size for a job is 100KB.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type Job = src.Job
type Job_AppEngineHttpTarget = src.Job_AppEngineHttpTarget
type Job_HttpTarget = src.Job_HttpTarget
type Job_PubsubTarget = src.Job_PubsubTarget

// State of the job.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type Job_State = src.Job_State

// Request message for listing jobs using
// [ListJobs][google.cloud.scheduler.v1.CloudScheduler.ListJobs].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type ListJobsRequest = src.ListJobsRequest

// Response message for listing jobs using
// [ListJobs][google.cloud.scheduler.v1.CloudScheduler.ListJobs].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type ListJobsResponse = src.ListJobsResponse

// Contains information needed for generating an [OAuth
// token](https://developers.google.com/identity/protocols/OAuth2). This type
// of authorization should generally only be used when calling Google APIs
// hosted on *.googleapis.com.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type OAuthToken = src.OAuthToken

// Contains information needed for generating an [OpenID Connect
// token](https://developers.google.com/identity/protocols/OpenIDConnect). This
// type of authorization can be used for many scenarios, including calling
// Cloud Run, or endpoints where you intend to validate the token yourself.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type OidcToken = src.OidcToken

// Request message for
// [PauseJob][google.cloud.scheduler.v1.CloudScheduler.PauseJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type PauseJobRequest = src.PauseJobRequest

// Pub/Sub target. The job will be delivered by publishing a message to the
// given Pub/Sub topic.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type PubsubTarget = src.PubsubTarget

// Request message for
// [ResumeJob][google.cloud.scheduler.v1.CloudScheduler.ResumeJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type ResumeJobRequest = src.ResumeJobRequest

// Settings that determine the retry behavior. By default, if a job does not
// complete successfully (meaning that an acknowledgement is not received from
// the handler, then it will be retried with exponential backoff according to
// the settings in [RetryConfig][google.cloud.scheduler.v1.RetryConfig].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type RetryConfig = src.RetryConfig

// Request message for forcing a job to run now using
// [RunJob][google.cloud.scheduler.v1.CloudScheduler.RunJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type RunJobRequest = src.RunJobRequest

// UnimplementedCloudSchedulerServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type UnimplementedCloudSchedulerServer = src.UnimplementedCloudSchedulerServer

// Request message for
// [UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob].
//
// Deprecated: Please use types in: cloud.google.com/go/scheduler/apiv1/schedulerpb
type UpdateJobRequest = src.UpdateJobRequest

// Deprecated: Please use funcs in: cloud.google.com/go/scheduler/apiv1/schedulerpb
func NewCloudSchedulerClient(cc grpc.ClientConnInterface) CloudSchedulerClient {
	return src.NewCloudSchedulerClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/scheduler/apiv1/schedulerpb
func RegisterCloudSchedulerServer(s *grpc.Server, srv CloudSchedulerServer) {
	src.RegisterCloudSchedulerServer(s, srv)
}