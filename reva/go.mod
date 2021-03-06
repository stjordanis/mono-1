module github.com/owncloud/mono/reva

go 1.13

require (
	github.com/cs3org/reva v1.2.1-0.20200911111727-51649e37df2d
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/oklog/run v1.0.0
	github.com/owncloud/flaex v0.0.0-20200411150708-dce59891a203
	github.com/owncloud/mono/ocis-pkg v0.0.0-00010101000000-000000000000
	github.com/restic/calens v0.2.0
	github.com/spf13/viper v1.6.3
)

replace github.com/owncloud/mono/ocis-pkg => ../ocis-pkg
