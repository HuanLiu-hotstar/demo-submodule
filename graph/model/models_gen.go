// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type PlaybackResponse interface {
	IsPlaybackResponse()
}

type Headers struct {
	XHsUserToken        string  `json:"xHsUserToken"`
	XCountryCode        string  `json:"xCountryCode"`
	XHsRequestID        string  `json:"xHsRequestId"`
	XHsAsn              *string `json:"xHsAsn"`
	XHsCity             *string `json:"xHsCity"`
	XHsRegion           *string `json:"xHsRegion"`
	TrueClientIP        *string `json:"trueClientIp"`
	XHsUserTokenUpdated *string `json:"xHsUserTokenUpdated"`
	UserAgent           *string `json:"userAgent"`
	Referer             *string `json:"referer"`
	Origin              *string `json:"origin"`
}

type OsDetails struct {
	OsName    string `json:"osName"`
	OsVersion string `json:"osVersion"`
}

type PlatformDetails struct {
	AppName         string     `json:"appName"`
	AppVersion      string     `json:"appVersion"`
	DeviceID        string     `json:"deviceId"`
	Platform        string     `json:"platform"`
	PlatformVersion string     `json:"platformVersion"`
	OsDetails       *OsDetails `json:"osDetails"`
}

type PlayBackContent struct {
	ContentID            string                `json:"contentId"`
	ProfileRequestConfig string                `json:"profileRequestConfig"`
	PlaybackSets         []*PlaybackDataSource `json:"playbackSets"`
	Match                bool                  `json:"match"`
	DrmClass             string                `json:"drmClass"`
	DownloadDrmClass     string                `json:"downloadDrmClass"`
	IsInternalPlayback   bool                  `json:"isInternalPlayback"`
}

type PlaybackDataSource struct {
	TagsCombination string `json:"tagsCombination"`
	PlaybackURL     string `json:"playbackUrl"`
	PlaybackCdnType string `json:"playbackCdnType"`
	TokenAlgorithm  string `json:"tokenAlgorithm"`
}

type PlaybackErrorResponse struct {
	ErrorCode           string               `json:"errorCode"`
	Message             string               `json:"message"`
	ErrorDetails        *ErrorDetails        `json:"errorDetails"`
	AdditionalErrorInfo *AdditionalErrorInfo `json:"additionalErrorInfo"`
}

func (PlaybackErrorResponse) IsPlaybackResponse() {}

type PlaybackRequestBody struct {
	ContentID            string                  `json:"contentId"`
	Headers              *Headers                `json:"headers"`
	DeviceManufacturer   *string                 `json:"deviceManufacturer"`
	DeviceModel          *string                 `json:"deviceModel"`
	ProfileRequestConfig []*ProfileRequestConfig `json:"profileRequestConfig"`
	PlatformDetails      *PlatformDetails        `json:"platformDetails"`
	Language             *string                 `json:"language"`
	Resolution           string                  `json:"resolution"`
	DrmParameters        []*MapTypeInput         `json:"drmParameters"`
	ClientCapabilities   []*MapTypeInput         `json:"clientCapabilities"`
}

type PlaybackSuccessResponse struct {
	Message        string           `json:"message"`
	PlaybackData   *PlayBackContent `json:"playbackData"`
	AdditionalInfo []*MapType       `json:"additionalInfo"`
}

func (PlaybackSuccessResponse) IsPlaybackResponse() {}

type ProfileRequestConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AdditionalErrorInfo struct {
	Key   *string     `json:"key"`
	Value interface{} `json:"value"`
}

type ErrorDetail struct {
	Key   *string     `json:"key"`
	Value interface{} `json:"value"`
}

type ErrorDetails struct {
	ErrorDetail []*ErrorDetail `json:"errorDetail"`
}

type MapType struct {
	Key   *string   `json:"key"`
	Value []*string `json:"value"`
}

type MapTypeInput struct {
	Key   *string     `json:"key"`
	Value interface{} `json:"value"`
}
