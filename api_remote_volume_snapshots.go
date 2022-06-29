
/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type RemoteVolumeSnapshotsApiService service
/*
RemoteVolumeSnapshotsApiService Eradicate a remote volume snapshot
Eradicates a remote volume snapshot that has been destroyed and is pending eradication. Eradicated remote volume snapshots cannot be recovered. Remote volume snapshots are destroyed through the &#x60;PATCH&#x60; method. The &#x60;names&#x60; parameter represents the name of the volume snapshot. The &#x60;on&#x60; parameter represents the name of the offload target. The &#x60;names&#x60; and &#x60;on&#x60; parameters are required and must be used together.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsDeleteOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "On" (optional.String) -  Performs the operation on the target name specified. For example, &#x60;targetName01&#x60;.
     * @param "ReplicationSnapshot" (optional.Bool) -  If set to &#x60;true&#x60;, allow destruction/eradication of snapshots in use by replication. If set to &#x60;false&#x60;, allow destruction/eradication of snapshots not in use by replication. If not specified, defaults to &#x60;false&#x60;.

*/

type RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsDeleteOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Names optional.Interface
    On optional.String
    ReplicationSnapshot optional.Bool
}

func (a *RemoteVolumeSnapshotsApiService) Api28RemoteVolumeSnapshotsDelete(ctx context.Context, localVarOptionals *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/remote-volume-snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.On.IsSet() {
		localVarQueryParams.Add("on", parameterToString(localVarOptionals.On.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReplicationSnapshot.IsSet() {
		localVarQueryParams.Add("replication_snapshot", parameterToString(localVarOptionals.ReplicationSnapshot.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
RemoteVolumeSnapshotsApiService List remote volume snapshots
Displays a list of remote volume snapshots.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsGetOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Destroyed" (optional.Bool) -  If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds.
     * @param "Filter" (optional.String) -  Narrows down the results to only the response objects that satisfy the filter criteria.
     * @param "Ids" (optional.Interface of []string) -  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together.
     * @param "Limit" (optional.Int32) -  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "Offset" (optional.Int32) -  The starting position based on the results of the query in relation to the full set of response objects returned.
     * @param "On" (optional.Interface of []string) -  Performs the operation on the target name specified. Enter multiple target names in comma-separated format. For example, &#x60;targetName01,targetName02&#x60;.
     * @param "Sort" (optional.Interface of []string) -  Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.
     * @param "SourceIds" (optional.Interface of []string) -  Performs the operation on the source ID specified. Enter multiple source IDs in comma-separated format.
     * @param "SourceNames" (optional.Interface of []string) -  Performs the operation on the source name specified. Enter multiple source names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "TotalItemCount" (optional.Bool) -  If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;.
@return InlineResponse200132
*/

type RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsGetOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Destroyed optional.Bool
    Filter optional.String
    Ids optional.Interface
    Limit optional.Int32
    Names optional.Interface
    Offset optional.Int32
    On optional.Interface
    Sort optional.Interface
    SourceIds optional.Interface
    SourceNames optional.Interface
    TotalItemCount optional.Bool
}

func (a *RemoteVolumeSnapshotsApiService) Api28RemoteVolumeSnapshotsGet(ctx context.Context, localVarOptionals *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsGetOpts) (InlineResponse200132, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200132
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/remote-volume-snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Destroyed.IsSet() {
		localVarQueryParams.Add("destroyed", parameterToString(localVarOptionals.Destroyed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ids.IsSet() {
		localVarQueryParams.Add("ids", parameterToString(localVarOptionals.Ids.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.On.IsSet() {
		localVarQueryParams.Add("on", parameterToString(localVarOptionals.On.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SourceIds.IsSet() {
		localVarQueryParams.Add("source_ids", parameterToString(localVarOptionals.SourceIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SourceNames.IsSet() {
		localVarQueryParams.Add("source_names", parameterToString(localVarOptionals.SourceNames.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.TotalItemCount.IsSet() {
		localVarQueryParams.Add("total_item_count", parameterToString(localVarOptionals.TotalItemCount.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse200132
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RemoteVolumeSnapshotsApiService Manage a remote volume snapshot
Destroy or recover a remote volume snapshot from the offload target. The &#x60;on&#x60; parameter represents the name of the offload target. The &#x60;names&#x60; parameter and the &#x60;on&#x60; parameter are required and must be used together.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param optional nil or *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsPatchOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "On" (optional.String) -  Performs the operation on the target name specified. For example, &#x60;targetName01&#x60;.
     * @param "ReplicationSnapshot" (optional.Bool) -  If set to &#x60;true&#x60;, allow destruction/eradication of snapshots in use by replication. If set to &#x60;false&#x60;, allow destruction/eradication of snapshots not in use by replication. If not specified, defaults to &#x60;false&#x60;.
@return InlineResponse200133
*/

type RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsPatchOpts struct {
    Authorization optional.String
    XRequestID optional.String
    Names optional.Interface
    On optional.String
    ReplicationSnapshot optional.Bool
}

func (a *RemoteVolumeSnapshotsApiService) Api28RemoteVolumeSnapshotsPatch(ctx context.Context, body interface{}, localVarOptionals *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsPatchOpts) (InlineResponse200133, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200133
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/remote-volume-snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.On.IsSet() {
		localVarQueryParams.Add("on", parameterToString(localVarOptionals.On.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReplicationSnapshot.IsSet() {
		localVarQueryParams.Add("replication_snapshot", parameterToString(localVarOptionals.ReplicationSnapshot.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse200133
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RemoteVolumeSnapshotsApiService Create a volume snapshot on a connected remote target or offload target
Creates a volume snapshot on the specified connected remote target or offload target.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsPostOpts - Optional Parameters:
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "SourceIds" (optional.Interface of []string) -  Performs the operation on the source ID specified. Enter multiple source IDs in comma-separated format.
     * @param "SourceNames" (optional.Interface of []string) -  Performs the operation on the source name specified. Enter multiple source names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "On" (optional.String) -  Performs the operation on the target name specified. For example, &#x60;targetName01&#x60;.
@return InlineResponse200133
*/

type RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsPostOpts struct {
    Authorization optional.String
    XRequestID optional.String
    SourceIds optional.Interface
    SourceNames optional.Interface
    On optional.String
}

func (a *RemoteVolumeSnapshotsApiService) Api28RemoteVolumeSnapshotsPost(ctx context.Context, localVarOptionals *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsPostOpts) (InlineResponse200133, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200133
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/remote-volume-snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SourceIds.IsSet() {
		localVarQueryParams.Add("source_ids", parameterToString(localVarOptionals.SourceIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SourceNames.IsSet() {
		localVarQueryParams.Add("source_names", parameterToString(localVarOptionals.SourceNames.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.On.IsSet() {
		localVarQueryParams.Add("on", parameterToString(localVarOptionals.On.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse200133
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RemoteVolumeSnapshotsApiService List remote volume snapshots with transfer statistics
Returns a list of remote volume snapshots and their transfer statistics.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsTransferGetOpts - Optional Parameters:
     * @param "Names" (optional.Interface of []string) -  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "Authorization" (optional.String) -  Access token (in JWT format) required to use any API endpoint (except &#x60;/oauth2&#x60;, &#x60;/login&#x60;, and &#x60;/logout&#x60;)
     * @param "XRequestID" (optional.String) -  Supplied by client during request or generated by server.
     * @param "Destroyed" (optional.Bool) -  If set to &#x60;true&#x60;, lists only destroyed objects that are in the eradication pending state. If set to &#x60;false&#x60;, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds.
     * @param "Filter" (optional.String) -  Narrows down the results to only the response objects that satisfy the filter criteria.
     * @param "Ids" (optional.Interface of []string) -  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The &#x60;ids&#x60; and &#x60;names&#x60; parameters cannot be provided together.
     * @param "Limit" (optional.Int32) -  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set &#x60;limit&#x3D;0&#x60;. The total number of resources is returned as a &#x60;total_item_count&#x60; value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.
     * @param "Offset" (optional.Int32) -  The starting position based on the results of the query in relation to the full set of response objects returned.
     * @param "On" (optional.Interface of []string) -  Performs the operation on the target name specified. Enter multiple target names in comma-separated format. For example, &#x60;targetName01,targetName02&#x60;.
     * @param "Sort" (optional.Interface of []string) -  Returns the response objects in the order specified. Set &#x60;sort&#x60; to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (&#x60;-&#x60;) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.
     * @param "SourceIds" (optional.Interface of []string) -  Performs the operation on the source ID specified. Enter multiple source IDs in comma-separated format.
     * @param "SourceNames" (optional.Interface of []string) -  Performs the operation on the source name specified. Enter multiple source names in comma-separated format. For example, &#x60;name01,name02&#x60;.
     * @param "TotalItemCount" (optional.Bool) -  If set to &#x60;true&#x60;, the &#x60;total_item_count&#x60; matching the specified query parameters is calculated and returned in the response. If set to &#x60;false&#x60;, the &#x60;total_item_count&#x60; is &#x60;null&#x60; in the response. This may speed up queries where the &#x60;total_item_count&#x60; is large. If not specified, defaults to &#x60;false&#x60;.
     * @param "TotalOnly" (optional.Bool) -  If set to &#x60;true&#x60;, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If &#x60;total_only&#x3D;true&#x60;, the &#x60;items&#x60; list will be empty.
@return InlineResponse200134
*/

type RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsTransferGetOpts struct {
    Names optional.Interface
    Authorization optional.String
    XRequestID optional.String
    Destroyed optional.Bool
    Filter optional.String
    Ids optional.Interface
    Limit optional.Int32
    Offset optional.Int32
    On optional.Interface
    Sort optional.Interface
    SourceIds optional.Interface
    SourceNames optional.Interface
    TotalItemCount optional.Bool
    TotalOnly optional.Bool
}

func (a *RemoteVolumeSnapshotsApiService) Api28RemoteVolumeSnapshotsTransferGet(ctx context.Context, localVarOptionals *RemoteVolumeSnapshotsApiApi28RemoteVolumeSnapshotsTransferGetOpts) (InlineResponse200134, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200134
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.8/remote-volume-snapshots/transfer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Destroyed.IsSet() {
		localVarQueryParams.Add("destroyed", parameterToString(localVarOptionals.Destroyed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ids.IsSet() {
		localVarQueryParams.Add("ids", parameterToString(localVarOptionals.Ids.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.On.IsSet() {
		localVarQueryParams.Add("on", parameterToString(localVarOptionals.On.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SourceIds.IsSet() {
		localVarQueryParams.Add("source_ids", parameterToString(localVarOptionals.SourceIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SourceNames.IsSet() {
		localVarQueryParams.Add("source_names", parameterToString(localVarOptionals.SourceNames.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.TotalItemCount.IsSet() {
		localVarQueryParams.Add("total_item_count", parameterToString(localVarOptionals.TotalItemCount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TotalOnly.IsSet() {
		localVarQueryParams.Add("total_only", parameterToString(localVarOptionals.TotalOnly.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse200134
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
