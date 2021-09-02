//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"time"
)

type blockBlobClient struct {
	con *connection
}

// CommitBlockList - The Commit Block List operation writes a blob by specifying the list of block IDs that make up the blob. In order to be written as
// part of a blob, a block must have been successfully written to the
// server in a prior Put Block operation. You can call Put Block List to update a blob by uploading only those blocks that have changed, then committing
// the new and existing blocks together. You can do
// this by specifying whether to commit a block from the committed block list or from the uncommitted block list, or to commit the most recently uploaded
// version of the block, whichever list it may
// belong to.
// If the operation fails it returns the *StorageError error type.
func (client *blockBlobClient) CommitBlockList(ctx context.Context, blocks BlockLookupList, blockBlobCommitBlockListOptions *BlockBlobCommitBlockListOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (BlockBlobCommitBlockListResponse, error) {
	req, err := client.commitBlockListCreateRequest(ctx, blocks, blockBlobCommitBlockListOptions, blobHTTPHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return BlockBlobCommitBlockListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlockBlobCommitBlockListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return BlockBlobCommitBlockListResponse{}, client.commitBlockListHandleError(resp)
	}
	return client.commitBlockListHandleResponse(resp)
}

// commitBlockListCreateRequest creates the CommitBlockList request.
func (client *blockBlobClient) commitBlockListCreateRequest(ctx context.Context, blocks BlockLookupList, blockBlobCommitBlockListOptions *BlockBlobCommitBlockListOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "blocklist")
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*blockBlobCommitBlockListOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobHTTPHeaders.BlobCacheControl)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobHTTPHeaders.BlobContentType)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobHTTPHeaders.BlobContentEncoding)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobHTTPHeaders.BlobContentLanguage)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(blobHTTPHeaders.BlobContentMD5))
	}
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.TransactionalContentMD5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(blockBlobCommitBlockListOptions.TransactionalContentMD5))
	}
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.TransactionalContentCRC64 != nil {
		req.Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(blockBlobCommitBlockListOptions.TransactionalContentCRC64))
	}
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.Metadata != nil {
		for k, v := range blockBlobCommitBlockListOptions.Metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobHTTPHeaders.BlobContentDisposition)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.Tier != nil {
		req.Header.Set("x-ms-access-tier", string(*blockBlobCommitBlockListOptions.Tier))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *blockBlobCommitBlockListOptions.RequestID)
	}
	if blockBlobCommitBlockListOptions != nil && blockBlobCommitBlockListOptions.BlobTagsString != nil {
		req.Header.Set("x-ms-tags", *blockBlobCommitBlockListOptions.BlobTagsString)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(blocks)
}

// commitBlockListHandleResponse handles the CommitBlockList response.
func (client *blockBlobClient) commitBlockListHandleResponse(resp *azcore.Response) (BlockBlobCommitBlockListResponse, error) {
	result := BlockBlobCommitBlockListResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockBlobCommitBlockListResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobCommitBlockListResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobCommitBlockListResponse{}, err
		}
		result.XMSContentCRC64 = xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-version-id"); val != "" {
		result.VersionID = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockBlobCommitBlockListResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return BlockBlobCommitBlockListResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// commitBlockListHandleError handles the CommitBlockList error response.
func (client *blockBlobClient) commitBlockListHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetBlockList - The Get Block List operation retrieves the list of blocks that have been uploaded as part of a block blob
// If the operation fails it returns the *StorageError error type.
func (client *blockBlobClient) GetBlockList(ctx context.Context, listType BlockListType, blockBlobGetBlockListOptions *BlockBlobGetBlockListOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (BlockListResponse, error) {
	req, err := client.getBlockListCreateRequest(ctx, listType, blockBlobGetBlockListOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return BlockListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlockListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BlockListResponse{}, client.getBlockListHandleError(resp)
	}
	return client.getBlockListHandleResponse(resp)
}

// getBlockListCreateRequest creates the GetBlockList request.
func (client *blockBlobClient) getBlockListCreateRequest(ctx context.Context, listType BlockListType, blockBlobGetBlockListOptions *BlockBlobGetBlockListOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "blocklist")
	if blockBlobGetBlockListOptions != nil && blockBlobGetBlockListOptions.Snapshot != nil {
		reqQP.Set("snapshot", *blockBlobGetBlockListOptions.Snapshot)
	}
	reqQP.Set("blocklisttype", string(listType))
	if blockBlobGetBlockListOptions != nil && blockBlobGetBlockListOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*blockBlobGetBlockListOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if blockBlobGetBlockListOptions != nil && blockBlobGetBlockListOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *blockBlobGetBlockListOptions.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getBlockListHandleResponse handles the GetBlockList response.
func (client *blockBlobClient) getBlockListHandleResponse(resp *azcore.Response) (BlockListResponse, error) {
	var val *BlockList
	if err := resp.UnmarshalAsXML(&val); err != nil {
		return BlockListResponse{}, err
	}
	result := BlockListResponse{RawResponse: resp.Response, BlockList: val}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockListResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-blob-content-length"); val != "" {
		blobContentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return BlockListResponse{}, err
		}
		result.BlobContentLength = &blobContentLength
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockListResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// getBlockListHandleError handles the GetBlockList error response.
func (client *blockBlobClient) getBlockListHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// StageBlock - The Stage Block operation creates a new block to be committed as part of a blob
// If the operation fails it returns the *StorageError error type.
func (client *blockBlobClient) StageBlock(ctx context.Context, blockID string, contentLength int64, body azcore.ReadSeekCloser, blockBlobStageBlockOptions *BlockBlobStageBlockOptions, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo) (BlockBlobStageBlockResponse, error) {
	req, err := client.stageBlockCreateRequest(ctx, blockID, contentLength, body, blockBlobStageBlockOptions, leaseAccessConditions, cpkInfo, cpkScopeInfo)
	if err != nil {
		return BlockBlobStageBlockResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlockBlobStageBlockResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return BlockBlobStageBlockResponse{}, client.stageBlockHandleError(resp)
	}
	return client.stageBlockHandleResponse(resp)
}

// stageBlockCreateRequest creates the StageBlock request.
func (client *blockBlobClient) stageBlockCreateRequest(ctx context.Context, blockID string, contentLength int64, body azcore.ReadSeekCloser, blockBlobStageBlockOptions *BlockBlobStageBlockOptions, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "block")
	reqQP.Set("blockid", blockID)
	if blockBlobStageBlockOptions != nil && blockBlobStageBlockOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*blockBlobStageBlockOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blockBlobStageBlockOptions != nil && blockBlobStageBlockOptions.TransactionalContentMD5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(blockBlobStageBlockOptions.TransactionalContentMD5))
	}
	if blockBlobStageBlockOptions != nil && blockBlobStageBlockOptions.TransactionalContentCRC64 != nil {
		req.Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(blockBlobStageBlockOptions.TransactionalContentCRC64))
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if blockBlobStageBlockOptions != nil && blockBlobStageBlockOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *blockBlobStageBlockOptions.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.SetBody(body, "application/octet-stream")
}

// stageBlockHandleResponse handles the StageBlock response.
func (client *blockBlobClient) stageBlockHandleResponse(resp *azcore.Response) (BlockBlobStageBlockResponse, error) {
	result := BlockBlobStageBlockResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobStageBlockResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockBlobStageBlockResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobStageBlockResponse{}, err
		}
		result.XMSContentCRC64 = xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return BlockBlobStageBlockResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// stageBlockHandleError handles the StageBlock error response.
func (client *blockBlobClient) stageBlockHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// StageBlockFromURL - The Stage Block operation creates a new block to be committed as part of a blob where the contents are read from a URL.
// If the operation fails it returns the *StorageError error type.
func (client *blockBlobClient) StageBlockFromURL(ctx context.Context, blockID string, contentLength int64, sourceURL string, blockBlobStageBlockFromURLOptions *BlockBlobStageBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (BlockBlobStageBlockFromURLResponse, error) {
	req, err := client.stageBlockFromURLCreateRequest(ctx, blockID, contentLength, sourceURL, blockBlobStageBlockFromURLOptions, cpkInfo, cpkScopeInfo, leaseAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return BlockBlobStageBlockFromURLResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlockBlobStageBlockFromURLResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return BlockBlobStageBlockFromURLResponse{}, client.stageBlockFromURLHandleError(resp)
	}
	return client.stageBlockFromURLHandleResponse(resp)
}

// stageBlockFromURLCreateRequest creates the StageBlockFromURL request.
func (client *blockBlobClient) stageBlockFromURLCreateRequest(ctx context.Context, blockID string, contentLength int64, sourceURL string, blockBlobStageBlockFromURLOptions *BlockBlobStageBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "block")
	reqQP.Set("blockid", blockID)
	if blockBlobStageBlockFromURLOptions != nil && blockBlobStageBlockFromURLOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*blockBlobStageBlockFromURLOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("x-ms-copy-source", sourceURL)
	if blockBlobStageBlockFromURLOptions != nil && blockBlobStageBlockFromURLOptions.SourceRange != nil {
		req.Header.Set("x-ms-source-range", *blockBlobStageBlockFromURLOptions.SourceRange)
	}
	if blockBlobStageBlockFromURLOptions != nil && blockBlobStageBlockFromURLOptions.SourceContentMD5 != nil {
		req.Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(blockBlobStageBlockFromURLOptions.SourceContentMD5))
	}
	if blockBlobStageBlockFromURLOptions != nil && blockBlobStageBlockFromURLOptions.SourceContentcrc64 != nil {
		req.Header.Set("x-ms-source-content-crc64", base64.StdEncoding.EncodeToString(blockBlobStageBlockFromURLOptions.SourceContentcrc64))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", sourceModifiedAccessConditions.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", sourceModifiedAccessConditions.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", *sourceModifiedAccessConditions.SourceIfMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", *sourceModifiedAccessConditions.SourceIfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if blockBlobStageBlockFromURLOptions != nil && blockBlobStageBlockFromURLOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *blockBlobStageBlockFromURLOptions.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// stageBlockFromURLHandleResponse handles the StageBlockFromURL response.
func (client *blockBlobClient) stageBlockFromURLHandleResponse(resp *azcore.Response) (BlockBlobStageBlockFromURLResponse, error) {
	result := BlockBlobStageBlockFromURLResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobStageBlockFromURLResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobStageBlockFromURLResponse{}, err
		}
		result.XMSContentCRC64 = xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockBlobStageBlockFromURLResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return BlockBlobStageBlockFromURLResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// stageBlockFromURLHandleError handles the StageBlockFromURL error response.
func (client *blockBlobClient) stageBlockFromURLHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Upload - The Upload Block Blob operation updates the content of an existing block blob. Updating an existing block blob overwrites any existing metadata
// on the blob. Partial updates are not supported with Put
// Blob; the content of the existing blob is overwritten with the content of the new blob. To perform a partial update of the content of a block blob, use
// the Put Block List operation.
// If the operation fails it returns the *StorageError error type.
func (client *blockBlobClient) Upload(ctx context.Context, contentLength int64, body azcore.ReadSeekCloser, blockBlobUploadOptions *BlockBlobUploadOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (BlockBlobUploadResponse, error) {
	req, err := client.uploadCreateRequest(ctx, contentLength, body, blockBlobUploadOptions, blobHTTPHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return BlockBlobUploadResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlockBlobUploadResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return BlockBlobUploadResponse{}, client.uploadHandleError(resp)
	}
	return client.uploadHandleResponse(resp)
}

// uploadCreateRequest creates the Upload request.
func (client *blockBlobClient) uploadCreateRequest(ctx context.Context, contentLength int64, body azcore.ReadSeekCloser, blockBlobUploadOptions *BlockBlobUploadOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if blockBlobUploadOptions != nil && blockBlobUploadOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*blockBlobUploadOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-blob-type", "BlockBlob")
	if blockBlobUploadOptions != nil && blockBlobUploadOptions.TransactionalContentMD5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(blockBlobUploadOptions.TransactionalContentMD5))
	}
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobHTTPHeaders.BlobContentType)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobHTTPHeaders.BlobContentEncoding)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobHTTPHeaders.BlobContentLanguage)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(blobHTTPHeaders.BlobContentMD5))
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobHTTPHeaders.BlobCacheControl)
	}
	if blockBlobUploadOptions != nil && blockBlobUploadOptions.Metadata != nil {
		for k, v := range blockBlobUploadOptions.Metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobHTTPHeaders.BlobContentDisposition)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if blockBlobUploadOptions != nil && blockBlobUploadOptions.Tier != nil {
		req.Header.Set("x-ms-access-tier", string(*blockBlobUploadOptions.Tier))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if blockBlobUploadOptions != nil && blockBlobUploadOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *blockBlobUploadOptions.RequestID)
	}
	if blockBlobUploadOptions != nil && blockBlobUploadOptions.BlobTagsString != nil {
		req.Header.Set("x-ms-tags", *blockBlobUploadOptions.BlobTagsString)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.SetBody(body, "application/octet-stream")
}

// uploadHandleResponse handles the Upload response.
func (client *blockBlobClient) uploadHandleResponse(resp *azcore.Response) (BlockBlobUploadResponse, error) {
	result := BlockBlobUploadResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockBlobUploadResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return BlockBlobUploadResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-version-id"); val != "" {
		result.VersionID = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return BlockBlobUploadResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return BlockBlobUploadResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// uploadHandleError handles the Upload error response.
func (client *blockBlobClient) uploadHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
