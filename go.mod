module github.com/confluentinc/auth0

go 1.12

require (
	cloud.google.com/go v0.35.1 // indirect
	github.com/PuerkitoBio/rehttp v0.0.0-20180310210549-11cf6ea5d3e9
	github.com/confluentinc/cc-structs v0.139.0
	github.com/confluentinc/cc-structs/kafka/flow v0.139.0
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/shurcooL/gofontwoff v0.0.0-20181114050219-180f79e6909d // indirect
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20181215221002-9d8641ddf2e1 // indirect
	github.com/shurcooL/issuesapp v0.0.0-20181229001453-b8198a402c58 // indirect
	github.com/shurcooL/notifications v0.0.0-20181111060504-bcc2b3082a7a // indirect
	github.com/shurcooL/octicon v0.0.0-20181222203144-9ff1a4cf27f4 // indirect
	github.com/shurcooL/reactions v0.0.0-20181222204718-145cd5e7f3d1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.3.0 // indirect
	golang.org/x/build v0.0.0-20190125014518-ad59fb13d315 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/genproto v0.0.0-20190123001331-8819c946db44 // indirect
	sourcegraph.com/sqs/pbtypes v1.0.0 // indirect
)

replace (
	// Override since git.apache.org is down.  The docs say to fetch from github.
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
	golang.org/x/sys => golang.org/x/sys v0.0.0-20180810173357-98c5dad5d1a0
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
