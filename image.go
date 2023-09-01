// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/apiquery"
	"github.com/hetzner/hetzner-go/internal/param"
	"github.com/hetzner/hetzner-go/internal/requestconfig"
	"github.com/hetzner/hetzner-go/internal/shared"
	"github.com/hetzner/hetzner-go/option"
)

// ImageService contains methods and other services that help with interacting with
// the hetzner API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
	Actions *ImageActionService
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r *ImageService) {
	r = &ImageService{}
	r.Options = opts
	r.Actions = NewImageActionService(opts...)
	return
}

// Returns a specific Image object.
func (r *ImageService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ImageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("images/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Image. You may change the description, convert a Backup Image to a
// Snapshot Image or change the Image labels. Only Images of type `snapshot` and
// `backup` can be updated.
//
// Note that when updating labels, the current set of labels will be replaced with
// the labels provided in the request body. So, for example, if you want to add a
// new label, you have to provide all existing labels plus the new label in the
// request body.
func (r *ImageService) Update(ctx context.Context, id int64, body ImageUpdateParams, opts ...option.RequestOption) (res *ImageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("images/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Image objects. You can select specific Image types only and sort the
// results by using URI parameters.
func (r *ImageService) List(ctx context.Context, query ImageListParams, opts ...option.RequestOption) (res *ImageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an Image. Only Images of type `snapshot` and `backup` can be deleted.
func (r *ImageService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("images/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/images/{id}
type ImageGetResponse struct {
	Image ImageGetResponseImage `json:"image"`
	JSON  imageGetResponseJSON
}

// imageGetResponseJSON contains the JSON metadata for the struct
// [ImageGetResponse]
type imageGetResponseJSON struct {
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageGetResponseImage struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ImageGetResponseImageArchitecture `json:"architecture,required"`
	// ID of Server the Image is bound to. Only set for Images of type `backup`.
	BoundTo int64 `json:"bound_to,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Information about the Server the Image was created from
	CreatedFrom ImageGetResponseImageCreatedFrom `json:"created_from,required,nullable"`
	// Point in time where the Image was deleted (in ISO-8601 format)
	Deleted string `json:"deleted,required,nullable"`
	// Point in time when the Image is considered to be deprecated (in ISO-8601 format)
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the Image
	Description string `json:"description,required"`
	// Size of the disk contained in the Image in GB
	DiskSize float64 `json:"disk_size,required"`
	// Size of the Image file in our storage in GB. For snapshot Images this is the
	// value relevant for calculating costs for the Image.
	ImageSize float64 `json:"image_size,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Unique identifier of the Image. This value is only set for system Images.
	Name string `json:"name,required,nullable"`
	// Flavor of operating system contained in the Image
	OsFlavor ImageGetResponseImageOsFlavor `json:"os_flavor,required"`
	// Operating system version
	OsVersion string `json:"os_version,required,nullable"`
	// Protection configuration for the Resource
	Protection ImageGetResponseImageProtection `json:"protection,required"`
	// Whether the Image can be used or if it's still being created or unavailable
	Status ImageGetResponseImageStatus `json:"status,required"`
	// Type of the Image
	Type ImageGetResponseImageType `json:"type,required"`
	// Indicates that rapid deploy of the Image is available
	RapidDeploy bool `json:"rapid_deploy"`
	JSON        imageGetResponseImageJSON
}

// imageGetResponseImageJSON contains the JSON metadata for the struct
// [ImageGetResponseImage]
type imageGetResponseImageJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	BoundTo      apijson.Field
	Created      apijson.Field
	CreatedFrom  apijson.Field
	Deleted      apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	DiskSize     apijson.Field
	ImageSize    apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	OsFlavor     apijson.Field
	OsVersion    apijson.Field
	Protection   apijson.Field
	Status       apijson.Field
	Type         apijson.Field
	RapidDeploy  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ImageGetResponseImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ImageGetResponseImageArchitecture string

const (
	ImageGetResponseImageArchitectureArm ImageGetResponseImageArchitecture = "arm"
	ImageGetResponseImageArchitectureX86 ImageGetResponseImageArchitecture = "x86"
)

// Information about the Server the Image was created from
type ImageGetResponseImageCreatedFrom struct {
	// ID of the Server the Image was created from
	ID int64 `json:"id,required"`
	// Server name at the time the Image was created
	Name string `json:"name,required"`
	JSON imageGetResponseImageCreatedFromJSON
}

// imageGetResponseImageCreatedFromJSON contains the JSON metadata for the struct
// [ImageGetResponseImageCreatedFrom]
type imageGetResponseImageCreatedFromJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageGetResponseImageCreatedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor of operating system contained in the Image
type ImageGetResponseImageOsFlavor string

const (
	ImageGetResponseImageOsFlavorAlma    ImageGetResponseImageOsFlavor = "alma"
	ImageGetResponseImageOsFlavorCentos  ImageGetResponseImageOsFlavor = "centos"
	ImageGetResponseImageOsFlavorDebian  ImageGetResponseImageOsFlavor = "debian"
	ImageGetResponseImageOsFlavorFedora  ImageGetResponseImageOsFlavor = "fedora"
	ImageGetResponseImageOsFlavorRocky   ImageGetResponseImageOsFlavor = "rocky"
	ImageGetResponseImageOsFlavorUbuntu  ImageGetResponseImageOsFlavor = "ubuntu"
	ImageGetResponseImageOsFlavorUnknown ImageGetResponseImageOsFlavor = "unknown"
)

// Protection configuration for the Resource
type ImageGetResponseImageProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   imageGetResponseImageProtectionJSON
}

// imageGetResponseImageProtectionJSON contains the JSON metadata for the struct
// [ImageGetResponseImageProtection]
type imageGetResponseImageProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageGetResponseImageProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Image can be used or if it's still being created or unavailable
type ImageGetResponseImageStatus string

const (
	ImageGetResponseImageStatusAvailable   ImageGetResponseImageStatus = "available"
	ImageGetResponseImageStatusCreating    ImageGetResponseImageStatus = "creating"
	ImageGetResponseImageStatusUnavailable ImageGetResponseImageStatus = "unavailable"
)

// Type of the Image
type ImageGetResponseImageType string

const (
	ImageGetResponseImageTypeApp       ImageGetResponseImageType = "app"
	ImageGetResponseImageTypeBackup    ImageGetResponseImageType = "backup"
	ImageGetResponseImageTypeSnapshot  ImageGetResponseImageType = "snapshot"
	ImageGetResponseImageTypeSystem    ImageGetResponseImageType = "system"
	ImageGetResponseImageTypeTemporary ImageGetResponseImageType = "temporary"
)

// Response to PUT https://api.hetzner.cloud/v1/images/{id}
type ImageUpdateResponse struct {
	Image ImageUpdateResponseImage `json:"image"`
	JSON  imageUpdateResponseJSON
}

// imageUpdateResponseJSON contains the JSON metadata for the struct
// [ImageUpdateResponse]
type imageUpdateResponseJSON struct {
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageUpdateResponseImage struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ImageUpdateResponseImageArchitecture `json:"architecture,required"`
	// ID of Server the Image is bound to. Only set for Images of type `backup`.
	BoundTo int64 `json:"bound_to,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Information about the Server the Image was created from
	CreatedFrom ImageUpdateResponseImageCreatedFrom `json:"created_from,required,nullable"`
	// Point in time where the Image was deleted (in ISO-8601 format)
	Deleted string `json:"deleted,required,nullable"`
	// Point in time when the Image is considered to be deprecated (in ISO-8601 format)
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the Image
	Description string `json:"description,required"`
	// Size of the disk contained in the Image in GB
	DiskSize float64 `json:"disk_size,required"`
	// Size of the Image file in our storage in GB. For snapshot Images this is the
	// value relevant for calculating costs for the Image.
	ImageSize float64 `json:"image_size,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Unique identifier of the Image. This value is only set for system Images.
	Name string `json:"name,required,nullable"`
	// Flavor of operating system contained in the Image
	OsFlavor ImageUpdateResponseImageOsFlavor `json:"os_flavor,required"`
	// Operating system version
	OsVersion string `json:"os_version,required,nullable"`
	// Protection configuration for the Resource
	Protection ImageUpdateResponseImageProtection `json:"protection,required"`
	// Whether the Image can be used or if it's still being created or unavailable
	Status ImageUpdateResponseImageStatus `json:"status,required"`
	// Type of the Image
	Type ImageUpdateResponseImageType `json:"type,required"`
	// Indicates that rapid deploy of the Image is available
	RapidDeploy bool `json:"rapid_deploy"`
	JSON        imageUpdateResponseImageJSON
}

// imageUpdateResponseImageJSON contains the JSON metadata for the struct
// [ImageUpdateResponseImage]
type imageUpdateResponseImageJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	BoundTo      apijson.Field
	Created      apijson.Field
	CreatedFrom  apijson.Field
	Deleted      apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	DiskSize     apijson.Field
	ImageSize    apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	OsFlavor     apijson.Field
	OsVersion    apijson.Field
	Protection   apijson.Field
	Status       apijson.Field
	Type         apijson.Field
	RapidDeploy  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ImageUpdateResponseImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ImageUpdateResponseImageArchitecture string

const (
	ImageUpdateResponseImageArchitectureArm ImageUpdateResponseImageArchitecture = "arm"
	ImageUpdateResponseImageArchitectureX86 ImageUpdateResponseImageArchitecture = "x86"
)

// Information about the Server the Image was created from
type ImageUpdateResponseImageCreatedFrom struct {
	// ID of the Server the Image was created from
	ID int64 `json:"id,required"`
	// Server name at the time the Image was created
	Name string `json:"name,required"`
	JSON imageUpdateResponseImageCreatedFromJSON
}

// imageUpdateResponseImageCreatedFromJSON contains the JSON metadata for the
// struct [ImageUpdateResponseImageCreatedFrom]
type imageUpdateResponseImageCreatedFromJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageUpdateResponseImageCreatedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor of operating system contained in the Image
type ImageUpdateResponseImageOsFlavor string

const (
	ImageUpdateResponseImageOsFlavorAlma    ImageUpdateResponseImageOsFlavor = "alma"
	ImageUpdateResponseImageOsFlavorCentos  ImageUpdateResponseImageOsFlavor = "centos"
	ImageUpdateResponseImageOsFlavorDebian  ImageUpdateResponseImageOsFlavor = "debian"
	ImageUpdateResponseImageOsFlavorFedora  ImageUpdateResponseImageOsFlavor = "fedora"
	ImageUpdateResponseImageOsFlavorRocky   ImageUpdateResponseImageOsFlavor = "rocky"
	ImageUpdateResponseImageOsFlavorUbuntu  ImageUpdateResponseImageOsFlavor = "ubuntu"
	ImageUpdateResponseImageOsFlavorUnknown ImageUpdateResponseImageOsFlavor = "unknown"
)

// Protection configuration for the Resource
type ImageUpdateResponseImageProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   imageUpdateResponseImageProtectionJSON
}

// imageUpdateResponseImageProtectionJSON contains the JSON metadata for the struct
// [ImageUpdateResponseImageProtection]
type imageUpdateResponseImageProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageUpdateResponseImageProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Image can be used or if it's still being created or unavailable
type ImageUpdateResponseImageStatus string

const (
	ImageUpdateResponseImageStatusAvailable   ImageUpdateResponseImageStatus = "available"
	ImageUpdateResponseImageStatusCreating    ImageUpdateResponseImageStatus = "creating"
	ImageUpdateResponseImageStatusUnavailable ImageUpdateResponseImageStatus = "unavailable"
)

// Type of the Image
type ImageUpdateResponseImageType string

const (
	ImageUpdateResponseImageTypeApp       ImageUpdateResponseImageType = "app"
	ImageUpdateResponseImageTypeBackup    ImageUpdateResponseImageType = "backup"
	ImageUpdateResponseImageTypeSnapshot  ImageUpdateResponseImageType = "snapshot"
	ImageUpdateResponseImageTypeSystem    ImageUpdateResponseImageType = "system"
	ImageUpdateResponseImageTypeTemporary ImageUpdateResponseImageType = "temporary"
)

// Response to GET https://api.hetzner.cloud/v1/images
type ImageListResponse struct {
	Images []ImageListResponseImage `json:"images,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON imageListResponseJSON
}

// imageListResponseJSON contains the JSON metadata for the struct
// [ImageListResponse]
type imageListResponseJSON struct {
	Images      apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageListResponseImage struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ImageListResponseImagesArchitecture `json:"architecture,required"`
	// ID of Server the Image is bound to. Only set for Images of type `backup`.
	BoundTo int64 `json:"bound_to,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Information about the Server the Image was created from
	CreatedFrom ImageListResponseImagesCreatedFrom `json:"created_from,required,nullable"`
	// Point in time where the Image was deleted (in ISO-8601 format)
	Deleted string `json:"deleted,required,nullable"`
	// Point in time when the Image is considered to be deprecated (in ISO-8601 format)
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the Image
	Description string `json:"description,required"`
	// Size of the disk contained in the Image in GB
	DiskSize float64 `json:"disk_size,required"`
	// Size of the Image file in our storage in GB. For snapshot Images this is the
	// value relevant for calculating costs for the Image.
	ImageSize float64 `json:"image_size,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Unique identifier of the Image. This value is only set for system Images.
	Name string `json:"name,required,nullable"`
	// Flavor of operating system contained in the Image
	OsFlavor ImageListResponseImagesOsFlavor `json:"os_flavor,required"`
	// Operating system version
	OsVersion string `json:"os_version,required,nullable"`
	// Protection configuration for the Resource
	Protection ImageListResponseImagesProtection `json:"protection,required"`
	// Whether the Image can be used or if it's still being created or unavailable
	Status ImageListResponseImagesStatus `json:"status,required"`
	// Type of the Image
	Type ImageListResponseImagesType `json:"type,required"`
	// Indicates that rapid deploy of the Image is available
	RapidDeploy bool `json:"rapid_deploy"`
	JSON        imageListResponseImageJSON
}

// imageListResponseImageJSON contains the JSON metadata for the struct
// [ImageListResponseImage]
type imageListResponseImageJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	BoundTo      apijson.Field
	Created      apijson.Field
	CreatedFrom  apijson.Field
	Deleted      apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	DiskSize     apijson.Field
	ImageSize    apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	OsFlavor     apijson.Field
	OsVersion    apijson.Field
	Protection   apijson.Field
	Status       apijson.Field
	Type         apijson.Field
	RapidDeploy  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ImageListResponseImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ImageListResponseImagesArchitecture string

const (
	ImageListResponseImagesArchitectureArm ImageListResponseImagesArchitecture = "arm"
	ImageListResponseImagesArchitectureX86 ImageListResponseImagesArchitecture = "x86"
)

// Information about the Server the Image was created from
type ImageListResponseImagesCreatedFrom struct {
	// ID of the Server the Image was created from
	ID int64 `json:"id,required"`
	// Server name at the time the Image was created
	Name string `json:"name,required"`
	JSON imageListResponseImagesCreatedFromJSON
}

// imageListResponseImagesCreatedFromJSON contains the JSON metadata for the struct
// [ImageListResponseImagesCreatedFrom]
type imageListResponseImagesCreatedFromJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageListResponseImagesCreatedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor of operating system contained in the Image
type ImageListResponseImagesOsFlavor string

const (
	ImageListResponseImagesOsFlavorAlma    ImageListResponseImagesOsFlavor = "alma"
	ImageListResponseImagesOsFlavorCentos  ImageListResponseImagesOsFlavor = "centos"
	ImageListResponseImagesOsFlavorDebian  ImageListResponseImagesOsFlavor = "debian"
	ImageListResponseImagesOsFlavorFedora  ImageListResponseImagesOsFlavor = "fedora"
	ImageListResponseImagesOsFlavorRocky   ImageListResponseImagesOsFlavor = "rocky"
	ImageListResponseImagesOsFlavorUbuntu  ImageListResponseImagesOsFlavor = "ubuntu"
	ImageListResponseImagesOsFlavorUnknown ImageListResponseImagesOsFlavor = "unknown"
)

// Protection configuration for the Resource
type ImageListResponseImagesProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   imageListResponseImagesProtectionJSON
}

// imageListResponseImagesProtectionJSON contains the JSON metadata for the struct
// [ImageListResponseImagesProtection]
type imageListResponseImagesProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageListResponseImagesProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Image can be used or if it's still being created or unavailable
type ImageListResponseImagesStatus string

const (
	ImageListResponseImagesStatusAvailable   ImageListResponseImagesStatus = "available"
	ImageListResponseImagesStatusCreating    ImageListResponseImagesStatus = "creating"
	ImageListResponseImagesStatusUnavailable ImageListResponseImagesStatus = "unavailable"
)

// Type of the Image
type ImageListResponseImagesType string

const (
	ImageListResponseImagesTypeApp       ImageListResponseImagesType = "app"
	ImageListResponseImagesTypeBackup    ImageListResponseImagesType = "backup"
	ImageListResponseImagesTypeSnapshot  ImageListResponseImagesType = "snapshot"
	ImageListResponseImagesTypeSystem    ImageListResponseImagesType = "system"
	ImageListResponseImagesTypeTemporary ImageListResponseImagesType = "temporary"
)

type ImageUpdateParams struct {
	// New description of Image
	Description param.Field[string] `json:"description"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// Destination Image type to convert to
	Type param.Field[ImageUpdateParamsType] `json:"type"`
}

func (r ImageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Destination Image type to convert to
type ImageUpdateParamsType string

const (
	ImageUpdateParamsTypeSnapshot ImageUpdateParamsType = "snapshot"
)

type ImageListParams struct {
	// Return only Images with the given architecture.
	Architecture param.Field[string] `query:"architecture"`
	// Can be used multiple times. Server ID linked to the Image. Only available for
	// Images of type `backup`
	BoundTo param.Field[string] `query:"bound_to"`
	// Can be used multiple times.
	IncludeDeprecated param.Field[bool] `query:"include_deprecated"`
	// Can be used to filter resources by labels. The response will only contain
	// resources matching the label selector.
	LabelSelector param.Field[string] `query:"label_selector"`
	// Can be used to filter resources by their name. The response will only contain
	// the resources matching the specified name
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[ImageListParamsSort] `query:"sort"`
	// Can be used multiple times. The response will only contain Images matching the
	// status.
	Status param.Field[ImageListParamsStatus] `query:"status"`
	// Can be used multiple times.
	Type param.Field[ImageListParamsType] `query:"type"`
}

// URLQuery serializes [ImageListParams]'s query parameters as `url.Values`.
func (r ImageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type ImageListParamsSort string

const (
	ImageListParamsSortID          ImageListParamsSort = "id"
	ImageListParamsSortIDAsc       ImageListParamsSort = "id:asc"
	ImageListParamsSortIDDesc      ImageListParamsSort = "id:desc"
	ImageListParamsSortName        ImageListParamsSort = "name"
	ImageListParamsSortNameAsc     ImageListParamsSort = "name:asc"
	ImageListParamsSortNameDesc    ImageListParamsSort = "name:desc"
	ImageListParamsSortCreated     ImageListParamsSort = "created"
	ImageListParamsSortCreatedAsc  ImageListParamsSort = "created:asc"
	ImageListParamsSortCreatedDesc ImageListParamsSort = "created:desc"
)

// Can be used multiple times. The response will only contain Images matching the
// status.
type ImageListParamsStatus string

const (
	ImageListParamsStatusAvailable ImageListParamsStatus = "available"
	ImageListParamsStatusCreating  ImageListParamsStatus = "creating"
)

// Can be used multiple times.
type ImageListParamsType string

const (
	ImageListParamsTypeSystem   ImageListParamsType = "system"
	ImageListParamsTypeSnapshot ImageListParamsType = "snapshot"
	ImageListParamsTypeBackup   ImageListParamsType = "backup"
	ImageListParamsTypeApp      ImageListParamsType = "app"
)
