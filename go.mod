module github.com/pinax-network/firehose-antelope

go 1.21

toolchain go1.21.1

require (
	github.com/ShinyTrinkets/overseer v0.3.0
	github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883
	github.com/eoscanada/eos-go v0.10.3-0.20230609180026-5d4ee54fe15e
	github.com/lytics/ordpool v0.0.0-20130426221837-8d833f097fe7
	github.com/mitchellh/go-testing-interface v1.14.1
	github.com/pinax-network/firehose-antelope/types v0.0.0-20230317150217-417f23cbea8d
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	github.com/streamingfast/bstream v0.0.2-0.20230829131224-b9272048dc6a
	github.com/streamingfast/cli v0.0.4-0.20230608140431-f501fec4f1ce
	github.com/streamingfast/dbin v0.9.1-0.20220513054835-1abebbb944ad
	github.com/streamingfast/dmetrics v0.0.0-20230516031116-28fcfeb4b9ed
	github.com/streamingfast/dstore v0.1.1-0.20230620124109-3924b3b36c77
	github.com/streamingfast/firehose-core v0.1.9
	github.com/streamingfast/jsonpb v0.0.0-20210811021341-3670f0aa02d0
	github.com/streamingfast/logging v0.0.0-20230608130331-f22c91403091
	github.com/streamingfast/node-manager v0.0.2-0.20230406142433-692298a8b8d2
	github.com/streamingfast/pbgo v0.0.6-0.20221020131607-255008258d28
	github.com/streamingfast/sf-tools v0.0.0-20230414095737-2f7b6f257241
	github.com/stretchr/testify v1.8.3
	github.com/tidwall/gjson v1.14.3
	go.uber.org/atomic v1.10.0
	go.uber.org/zap v1.24.0
	google.golang.org/protobuf v1.30.0
)

replace (
	github.com/ShinyTrinkets/overseer => github.com/streamingfast/overseer v0.2.1-0.20210326144022-ee491780e3ef
	github.com/bytecodealliance/wasmtime-go/v4 => github.com/streamingfast/wasmtime-go/v4 v4.0.0-freemem3
	github.com/graph-gophers/graphql-go => github.com/streamingfast/graphql-go v0.0.0-20210204202750-0e485a040a3c
)

replace github.com/pinax-network/firehose-antelope/types => ./types

require (
	cloud.google.com/go v0.110.0 // indirect
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v0.12.0 // indirect
	cloud.google.com/go/monitoring v1.12.0 // indirect
	cloud.google.com/go/storage v1.30.1 // indirect
	cloud.google.com/go/trace v1.8.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.13.10 // indirect
	contrib.go.opencensus.io/exporter/zipkin v0.1.1 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-storage-blob-go v0.14.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v0.32.3 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.15.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.39.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator v0.0.0-20221018185641-36f91511cfd7 // indirect
	github.com/KimMachineGun/automemlimit v0.2.4 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/RoaringBitmap/roaring v0.9.4 // indirect
	github.com/ShinyTrinkets/meta-logger v0.2.0 // indirect
	github.com/abourget/llerrgroup v0.2.0 // indirect
	github.com/aws/aws-sdk-go v1.44.233 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.3.1 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bobg/go-generics/v2 v2.1.1 // indirect
	github.com/bufbuild/connect-go v1.8.0 // indirect
	github.com/bufbuild/connect-grpchealth-go v1.1.1 // indirect
	github.com/bufbuild/connect-grpcreflect-go v1.0.0 // indirect
	github.com/bufbuild/connect-opentelemetry-go v0.3.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cilium/ebpf v0.9.1 // indirect
	github.com/cncf/udpa/go v0.0.0-20220112060539-c52dc94e7fbe // indirect
	github.com/cncf/xds/go v0.0.0-20230105202645-06c439db220b // indirect
	github.com/containerd/cgroups v1.1.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/crackcomm/go-gitignore v0.0.0-20170627025303-887ab5e44cc3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/emicklei/go-restful/v3 v3.8.0 // indirect
	github.com/envoyproxy/go-control-plane v0.10.3 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.9.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/ipfs/boxo v0.8.0 // indirect
	github.com/ipfs/go-cid v0.4.1 // indirect
	github.com/ipfs/go-ipfs-api v0.6.0 // indirect
	github.com/jhump/protoreflect v1.12.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josephburnett/jd v1.7.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.16.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/libp2p/go-flow-metrics v0.1.0 // indirect
	github.com/libp2p/go-libp2p v0.27.8 // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/lytics/lifecycle v0.0.0-20130117214539-7b4c4028d422 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.1.17 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multiaddr v0.9.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.8.1 // indirect
	github.com/multiformats/go-multihash v0.2.1 // indirect
	github.com/multiformats/go-multistream v0.4.1 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/gomega v1.27.4 // indirect
	github.com/opencontainers/runtime-spec v1.0.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	github.com/paulbellamy/ratecounter v0.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.0 // indirect
	github.com/rs/cors v1.8.3 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sethvargo/go-retry v0.2.3 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/streamingfast/atm v0.0.0-20220131151839-18c87005e680 // indirect
	github.com/streamingfast/dauth v0.0.0-20230929180355-921f9c9be330 // indirect
	github.com/streamingfast/derr v0.0.0-20230515163924-8570aaa43fe1 // indirect
	github.com/streamingfast/dgrpc v0.0.0-20230623191332-4ed02c06bc6f // indirect
	github.com/streamingfast/dlauncher v0.0.0-20230607184145-76399faad89e // indirect
	github.com/streamingfast/dmetering v0.0.0-20230731155453-e1df53e362aa // indirect
	github.com/streamingfast/dtracing v0.0.0-20220305214756-b5c0e8699839 // indirect
	github.com/streamingfast/firehose v0.1.1-0.20230731171526-3428fc0ff8d9 // indirect
	github.com/streamingfast/index-builder v0.0.0-20221031203737-fa2e70f09dc2 // indirect
	github.com/streamingfast/merger v0.0.3-0.20230413191947-fe3ce68cb9f8 // indirect
	github.com/streamingfast/opaque v0.0.0-20210811180740-0c01d37ea308 // indirect
	github.com/streamingfast/relayer v0.0.2-0.20220909122435-e67fbc964fd9 // indirect
	github.com/streamingfast/sf-tracing v0.0.0-20230616174903-cd2ade641ca9 // indirect
	github.com/streamingfast/shutter v1.5.0 // indirect
	github.com/streamingfast/snapshotter v0.0.0-20230316190750-5bcadfde44d0 // indirect
	github.com/streamingfast/substreams v1.1.14 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf // indirect
	github.com/tetratelabs/wazero v1.1.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/whyrusleeping/tar-utils v0.0.0-20201201191210-20a61371de5b // indirect
	github.com/yourbasic/graph v0.0.0-20210606180040-8ecfec1c2869 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.9.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.36.4 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.15.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.15.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.15.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.15.1 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.15.1 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.15.1 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/automaxprocs v1.5.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.114.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230320184635-7606e756e683 // indirect
	google.golang.org/grpc v1.54.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/olivere/elastic.v3 v3.0.75 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.25.0 // indirect
	k8s.io/apimachinery v0.25.0 // indirect
	k8s.io/client-go v0.25.0 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	lukechampine.com/blake3 v1.1.7 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace github.com/streamingfast/firehose-core => github.com/pinax-network/firehose-core v0.0.0-20231009123906-83e51d167ea7
