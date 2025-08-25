package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ApproveDecision represents the ApproveDecision schema from the OpenAPI specification
type ApproveDecision struct {
	Approvetime string `json:"approveTime,omitempty"` // The time at which approval was granted.
	Autoapproved bool `json:"autoApproved,omitempty"` // True when the request has been auto-approved.
	Expiretime string `json:"expireTime,omitempty"` // The time at which the approval expires.
	Invalidatetime string `json:"invalidateTime,omitempty"` // If set, denotes the timestamp at which the approval is invalidated.
	Signatureinfo SignatureInfo `json:"signatureInfo,omitempty"` // Information about the digital signature of the resource.
}

// AccessApprovalSettings represents the AccessApprovalSettings schema from the OpenAPI specification
type AccessApprovalSettings struct {
	Enrolledancestor bool `json:"enrolledAncestor,omitempty"` // Output only. This field is read only (not settable via UpdateAccessApprovalSettings method). If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Project or Folder (this field will always be unset for the organization since organizations do not have ancestors).
	Name string `json:"name,omitempty"` // The resource name of the settings. Format is one of: * "projects/{project}/accessApprovalSettings" * "folders/{folder}/accessApprovalSettings" * "organizations/{organization}/accessApprovalSettings"
	Notificationpubsubtopic string `json:"notificationPubsubTopic,omitempty"` // Optional. A pubsub topic to which notifications relating to approval requests should be sent.
	Preferredrequestexpirationdays int `json:"preferredRequestExpirationDays,omitempty"` // This preference is shared with Google personnel, but can be overridden if said personnel deems necessary. The approver ultimately can set the expiration at approval time.
	Ancestorhasactivekeyversion bool `json:"ancestorHasActiveKeyVersion,omitempty"` // Output only. This field is read only (not settable via UpdateAccessApprovalSettings method). If the field is true, that indicates that an ancestor of this Project or Folder has set active_key_version (this field will always be unset for the organization since organizations do not have ancestors).
	Invalidkeyversion bool `json:"invalidKeyVersion,omitempty"` // Output only. This field is read only (not settable via UpdateAccessApprovalSettings method). If the field is true, that indicates that there is some configuration issue with the active_key_version configured at this level in the resource hierarchy (e.g. it doesn't exist or the Access Approval service account doesn't have the correct permissions on it, etc.) This key version is not necessarily the effective key version at this level, as key versions are inherited top-down.
	Notificationemails []string `json:"notificationEmails,omitempty"` // A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email addresses are allowed.
	Activekeyversion string `json:"activeKeyVersion,omitempty"` // The asymmetric crypto key version to use for signing approval requests. Empty active_key_version indicates that a Google-managed key should be used for signing. This property will be ignored if set by an ancestor of this resource, and new non-empty values may not be set.
	Enrolledservices []EnrolledService `json:"enrolledServices,omitempty"` // A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the resource given by name against any of these services contained here will be required to have explicit approval. If name refers to an organization, enrollment can be done for individual services. If name refers to a folder or project, enrollment can only be done on an all or nothing basis. If a cloud_product is repeated in this list, the first entry will be honored and all following entries will be discarded. A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	Prefernobroadapprovalrequests bool `json:"preferNoBroadApprovalRequests,omitempty"` // This preference is communicated to Google personnel when sending an approval request but can be overridden if necessary.
}

// InvalidateApprovalRequestMessage represents the InvalidateApprovalRequestMessage schema from the OpenAPI specification
type InvalidateApprovalRequestMessage struct {
}

// SignatureInfo represents the SignatureInfo schema from the OpenAPI specification
type SignatureInfo struct {
	Customerkmskeyversion string `json:"customerKmsKeyVersion,omitempty"` // The resource name of the customer CryptoKeyVersion used for signing.
	Googlekeyalgorithm string `json:"googleKeyAlgorithm,omitempty"` // The hashing algorithm used for signature verification. It will only be present in the case of Google managed keys.
	Googlepublickeypem string `json:"googlePublicKeyPem,omitempty"` // The public key for the Google default signing, encoded in PEM format. The signature was created using a private key which may be verified using this public key.
	Serializedapprovalrequest string `json:"serializedApprovalRequest,omitempty"` // The ApprovalRequest that is serialized without the SignatureInfo message field. This data is used with the hashing algorithm to generate the digital signature, and it can be used for signature verification.
	Signature string `json:"signature,omitempty"` // The digital signature.
}

// ApproveApprovalRequestMessage represents the ApproveApprovalRequestMessage schema from the OpenAPI specification
type ApproveApprovalRequestMessage struct {
	Expiretime string `json:"expireTime,omitempty"` // The expiration time of this approval.
}

// EnrolledService represents the EnrolledService schema from the OpenAPI specification
type EnrolledService struct {
	Cloudproduct string `json:"cloudProduct,omitempty"` // The product for which Access Approval will be enrolled. Allowed values are listed below (case-sensitive): * all * GA * App Engine * Artifact Registry * BigQuery * Certificate Authority Service * Cloud Bigtable * Cloud Key Management Service * Compute Engine * Cloud Composer * Cloud Dataflow * Cloud Dataproc * Cloud DLP * Cloud EKM * Cloud Firestore * Cloud HSM * Cloud Identity and Access Management * Cloud Logging * Cloud NAT * Cloud Pub/Sub * Cloud Spanner * Cloud SQL * Cloud Storage * Eventarc * Google Kubernetes Engine * Organization Policy Serivice * Persistent Disk * Resource Manager * Secret Manager * Speaker ID Note: These values are supported as input for legacy purposes, but will not be returned from the API. * all * ga-only * appengine.googleapis.com * artifactregistry.googleapis.com * bigquery.googleapis.com * bigtable.googleapis.com * container.googleapis.com * cloudkms.googleapis.com * cloudresourcemanager.googleapis.com * cloudsql.googleapis.com * compute.googleapis.com * dataflow.googleapis.com * dataproc.googleapis.com * dlp.googleapis.com * iam.googleapis.com * logging.googleapis.com * orgpolicy.googleapis.com * pubsub.googleapis.com * spanner.googleapis.com * secretmanager.googleapis.com * speakerid.googleapis.com * storage.googleapis.com Calls to UpdateAccessApprovalSettings using 'all' or any of the XXX.googleapis.com will be translated to the associated product name ('all', 'App Engine', etc.). Note: 'all' will enroll the resource in all products supported at both 'GA' and 'Preview' levels. More information about levels of support is available at https://cloud.google.com/access-approval/docs/supported-services
	Enrollmentlevel string `json:"enrollmentLevel,omitempty"` // The enrollment level of the service.
}

// AccessReason represents the AccessReason schema from the OpenAPI specification
type AccessReason struct {
	Detail string `json:"detail,omitempty"` // More detail about certain reason types. See comments for each type above.
	TypeField string `json:"type,omitempty"` // Type of access justification.
}

// DismissApprovalRequestMessage represents the DismissApprovalRequestMessage schema from the OpenAPI specification
type DismissApprovalRequestMessage struct {
}

// ApprovalRequest represents the ApprovalRequest schema from the OpenAPI specification
type ApprovalRequest struct {
	Requestedresourceproperties ResourceProperties `json:"requestedResourceProperties,omitempty"` // The properties associated with the resource of the request.
	Name string `json:"name,omitempty"` // The resource name of the request. Format is "{projects|folders|organizations}/{id}/approvalRequests/{approval_request}".
	Requesttime string `json:"requestTime,omitempty"` // The time at which approval was requested.
	Requestedresourcename string `json:"requestedResourceName,omitempty"` // The resource for which approval is being requested. The format of the resource name is defined at https://cloud.google.com/apis/design/resource_names. The resource name here may either be a "full" resource name (e.g. "//library.googleapis.com/shelves/shelf1/books/book2") or a "relative" resource name (e.g. "shelves/shelf1/books/book2") as described in the resource name specification.
	Dismiss DismissDecision `json:"dismiss,omitempty"` // A decision that has been made to dismiss an approval request.
	Requestedreason AccessReason `json:"requestedReason,omitempty"`
	Approve ApproveDecision `json:"approve,omitempty"` // A decision that has been made to approve access to a resource.
	Requestedduration string `json:"requestedDuration,omitempty"` // The requested access duration.
	Requestedexpiration string `json:"requestedExpiration,omitempty"` // The original requested expiration for the approval. Calculated by adding the requested_duration to the request_time.
	Requestedlocations AccessLocations `json:"requestedLocations,omitempty"` // Home office and physical location of the principal.
}

// ResourceProperties represents the ResourceProperties schema from the OpenAPI specification
type ResourceProperties struct {
	Excludesdescendants bool `json:"excludesDescendants,omitempty"` // Whether an approval will exclude the descendants of the resource being requested.
}

// DismissDecision represents the DismissDecision schema from the OpenAPI specification
type DismissDecision struct {
	Implicit bool `json:"implicit,omitempty"` // This field will be true if the ApprovalRequest was implicitly dismissed due to inaction by the access approval approvers (the request is not acted on by the approvers before the exiration time).
	Dismisstime string `json:"dismissTime,omitempty"` // The time at which the approval request was dismissed.
}

// AccessLocations represents the AccessLocations schema from the OpenAPI specification
type AccessLocations struct {
	Principalofficecountry string `json:"principalOfficeCountry,omitempty"` // The "home office" location of the principal. A two-letter country code (ISO 3166-1 alpha-2), such as "US", "DE" or "GB" or a region code. In some limited situations Google systems may refer refer to a region code instead of a country code. Possible Region Codes: * ASI: Asia * EUR: Europe * OCE: Oceania * AFR: Africa * NAM: North America * SAM: South America * ANT: Antarctica * ANY: Any location
	Principalphysicallocationcountry string `json:"principalPhysicalLocationCountry,omitempty"` // Physical location of the principal at the time of the access. A two-letter country code (ISO 3166-1 alpha-2), such as "US", "DE" or "GB" or a region code. In some limited situations Google systems may refer refer to a region code instead of a country code. Possible Region Codes: * ASI: Asia * EUR: Europe * OCE: Oceania * AFR: Africa * NAM: North America * SAM: South America * ANT: Antarctica * ANY: Any location
}

// ListApprovalRequestsResponse represents the ListApprovalRequestsResponse schema from the OpenAPI specification
type ListApprovalRequestsResponse struct {
	Approvalrequests []ApprovalRequest `json:"approvalRequests,omitempty"` // Approval request details.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to retrieve the next page of results, or empty if there are no more.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// AccessApprovalServiceAccount represents the AccessApprovalServiceAccount schema from the OpenAPI specification
type AccessApprovalServiceAccount struct {
	Accountemail string `json:"accountEmail,omitempty"` // Email address of the service account.
	Name string `json:"name,omitempty"` // The resource name of the Access Approval service account. Format is one of: * "projects/{project}/serviceAccount" * "folders/{folder}/serviceAccount" * "organizations/{organization}/serviceAccount"
}
