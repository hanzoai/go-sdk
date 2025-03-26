// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"github.com/stainless-sdks/Hanzo-AI-go/option"
)

// AudioService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioService] method instead.
type AudioService struct {
	Options        []option.RequestOption
	Speech         *AudioSpeechService
	Transcriptions *AudioTranscriptionService
}

// NewAudioService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAudioService(opts ...option.RequestOption) (r *AudioService) {
	r = &AudioService{}
	r.Options = opts
	r.Speech = NewAudioSpeechService(opts...)
	r.Transcriptions = NewAudioTranscriptionService(opts...)
	return
}
