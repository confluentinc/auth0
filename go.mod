module github.com/confluentinc/auth0

go 1.12

require (
	github.com/PuerkitoBio/rehttp v0.0.0-20180310210549-11cf6ea5d3e9
	github.com/aybabtme/iocontrol v0.0.0-20150809002002-ad15bcfc95a0 // indirect
	github.com/benbjohnson/clock v1.0.0 // indirect
	github.com/confluentinc/cc-structs/kafka/auth v0.342.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace (
	// Override since git.apache.org is down.  The docs say to fetch from github.
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/confluentinc/cc-structs v0.0.0-20190320051855-97f99cb204dc => github.com/confluentinc/cc-structs v0.342.0
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
	golang.org/x/sys => golang.org/x/sys v0.0.0-20180810173357-98c5dad5d1a0
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
