package bitly

type LinkMetricService struct {
	client *Client
}

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

type LinkMetricOptions struct {
	Options
}

func (s *LinkMetricService) Clicks(link string, opts LinkMetricOptions) LinkMetric {
	return LinkMetric{}
}
