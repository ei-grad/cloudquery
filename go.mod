module github.com/cloudquery/cloudquery

go 1.17

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/cloudquery/cq-provider-sdk v0.13.2
	github.com/fatih/color v1.13.0
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/google/go-github/v35 v35.3.0
	github.com/hashicorp/go-hclog v1.2.1
	github.com/hashicorp/go-plugin v1.4.4
	github.com/hashicorp/go-version v1.5.0
	github.com/hashicorp/hcl/v2 v2.13.0
	github.com/jackc/pgx/v4 v4.16.1
	github.com/mattn/go-isatty v0.0.14
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/rs/zerolog v1.27.0
	github.com/spf13/afero v1.8.2
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.5
	github.com/thoas/go-funk v0.9.2
	github.com/vbauerster/mpb/v6 v6.0.4
	github.com/zclconf/go-cty v1.9.1
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/ProtonMail/go-crypto v0.0.0-20220623141421-5afb4c282135
	github.com/doug-martin/goqu/v9 v9.18.0
	github.com/driftprogramming/pgxpoolmock v1.1.0
	github.com/georgysavva/scany v0.3.0
	github.com/getsentry/sentry-go v0.13.0
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.8
	github.com/google/uuid v1.3.0
	github.com/hairyhenderson/go-fsimpl v0.0.0-20220626234318-d4f0b5a8cf3a
	github.com/hashicorp/go-getter v1.6.2
	github.com/jackc/pgconn v1.12.1
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/jeremywohl/flatten v1.0.1
	github.com/johannesboyne/gofakes3 v0.0.0-20220517215058-83a58ec253b6
	github.com/lib/pq v1.10.6
	github.com/modern-go/reflect2 v1.0.2
	github.com/olekukonko/tablewriter v0.0.5
	github.com/rudderlabs/analytics-go v3.3.2+incompatible
	github.com/savioxavier/termlink v1.1.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/pflag v1.0.5
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/xo/dburl v0.11.0
	google.golang.org/grpc v1.47.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloud.google.com/go v0.102.0 // indirect
	cloud.google.com/go/compute v1.6.1 // indirect
	cloud.google.com/go/iam v0.3.0 // indirect
	cloud.google.com/go/storage v1.22.1 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.21.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v0.8.3 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v0.3.0 // indirect
	github.com/Azure/azure-storage-blob-go v0.14.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.24 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go v1.44.42 // indirect
	github.com/aws/aws-sdk-go-v2 v1.16.5 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.15.11 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.6 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.13 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.26.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.7 // indirect
	github.com/aws/smithy-go v1.11.3 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/creasty/defaults v1.6.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elliotchance/orderedmap v1.4.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/googleapis/go-type-adapters v1.0.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jwalton/go-supportscolor v1.1.0 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lorenzosaino/go-sysctl v0.3.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/ryszard/goskiplist v0.0.0-20150312221310-2dfbae5fcf46 // indirect
	github.com/segmentio/backo-go v1.0.0 // indirect
	github.com/segmentio/stats/v4 v4.6.3 // indirect
	github.com/shabbyrobe/gocovmerge v0.0.0-20190829150210-3e036491d500 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.7.2 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tidwall/gjson v1.11.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/ulikunitz/xz v0.5.8 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	gocloud.dev v0.25.1-0.20220408200107-09b10f7359f7 // indirect
	golang.org/x/exp/typeparams v0.0.0-20220317015231-48e79f11773a // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	golang.org/x/oauth2 v0.0.0-20220524215830-622c5d57e401 // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220517211312-f3a8303e98df // indirect
	google.golang.org/api v0.82.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220527130721-00d5c0f3be58 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.3.0-0.dev.0.20220306074811-23e1086441d2 // indirect
)

//replace github.com/cloudquery/cq-provider-sdk => ../cq-provider-sdk
