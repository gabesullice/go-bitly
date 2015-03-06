package bitly

// Options holds possible query string parameters generally
// available across the Bitly API.
type Options struct {
	Limit           uint        `url:"limit"`
	Link            []byte      `url:"link"`
	Rollup          bool        `url:"rollup"`
	Timezone        interface{} `url:"timezone"`
	Unit            []byte      `url:"unit"`
	Units           int8        `url:"units"`
	UnitReferenceTs int64       `url:"unit_reference_ts"`
}
