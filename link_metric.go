package bitly

// LinkMetricService is a receiver for functions making requests against the
// Bitly API link_metric endpoint.
type LinkMetricService struct {
	client *Client
}

// LinkMetric holds data returned from any of the Bitly API link_metric
// endpoints.
type LinkMetric struct {
	AggregateLink   []byte `json:"aggregate_link"`
	LinkClicks      uint64 `json:"link_clicks"`
	Shares          int64  `json:"shares"`
	TotalShares     int64  `json:"total_shares"`
	TzOffset        int8   `json:"tz_offset"`
	Unit            []byte `json:"unit"`
	Units           int8   `json:"units"`
	UnitReferenceTs int64  `json:"unit_reference_ts"`

	Countries        []map[string]interface{} `json:"countries"`
	Entries          []map[string]interface{} `json:"entries"`
	Referrers        []map[string]interface{} `json:"referrers"`
	ReferringDomains []map[string]interface{} `json:"referring_domains"`
}

// LinkMetricOptions embeds the generally available requests parameters, and
// holds any link_metric specific query parameters.
type LinkMetricOptions struct {
	Options
}

// Request the number of clicks on a bitlink.
//
// Bitly API docs: http://dev.bitly.com/link_metrics.html#v3_link_clicks
func (s *LinkMetricService) Clicks(link string, opts LinkMetricOptions) LinkMetric {
	return LinkMetric{}
}
