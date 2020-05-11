module github.com/padraicbc/ponzu

go 1.13

require (
	github.com/RoaringBitmap/roaring v0.4.23 // indirect
	github.com/blevesearch/bleve v1.0.7
	github.com/cznic/b v0.0.0-20181122101859-a26611c4d92d // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/cznic/strutil v0.0.0-20181122101858-275e90344537 // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/glycerine/go-unsnap-stream v0.0.0-20190901134440-81cf024a9e0a // indirect
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/gorilla/schema v1.1.0
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nilslice/email v0.1.0
	github.com/nilslice/jwt v1.0.0
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/padraicbc/ponzu/content v0.0.0-00010101000000-000000000000
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
	github.com/tidwall/gjson v1.6.0
	github.com/tidwall/sjson v1.1.1
	github.com/tinylib/msgp v1.1.2 // indirect
	go.etcd.io/bbolt v1.3.4
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79
	golang.org/x/net v0.0.0-20200505041828-1ed23360d12c // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/text v0.3.2
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

// Used only in builds, could add the import and forget about it here
// but just as easy do the 	_ "github.com/padraicbc/ponzu/content" import in main now
replace github.com/padraicbc/ponzu/content => ./content
