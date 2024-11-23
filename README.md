# gencodec for golang 1.23 and above

Command gencodec generates marshaling methods for Go struct types.

The generated methods add features which json and other marshaling packages cannot offer.

	gencodec -dir . -type MyType -formats json,yaml,toml -out mytype_json.go

