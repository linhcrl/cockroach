# PROTOBUF_TARGETS lists the indirect dependencies needed to compile proto dependencies
# of //pkg/server/serverpb:serverpb_proto target
PROTOBUF_TARGETS = [
    "//pkg/build:build_proto",
    "//pkg/clusterversion:clusterversion_proto",
    "//pkg/config/zonepb:zonepb_proto",
    "//pkg/geo/geopb:geopb_proto",
    "//pkg/gossip:gossip_proto",
    "//pkg/jobs/jobspb:jobspb_proto",
    "//pkg/kv/kvpb:kvpb_proto",
    "//pkg/kv/kvserver/concurrency/isolation:isolation_proto",
    "//pkg/kv/kvserver/concurrency/lock:lock_proto",
    "//pkg/kv/kvserver/kvserverpb:kvserverpb_proto",
    "//pkg/kv/kvserver/liveness/livenesspb:livenesspb_proto",
    "//pkg/kv/kvserver/loqrecovery/loqrecoverypb:loqrecoverypb_proto",
    "//pkg/kv/kvserver/readsummary/rspb:rspb_proto",
    "//pkg/kv/kvserver/kvflowcontrol/kvflowcontrolpb:kvflowcontrolpb_proto",
    "//pkg/multitenant/mtinfopb:mtinfopb_proto",
    "//pkg/multitenant/tenantcapabilities/tenantcapabilitiespb:tenantcapabilitiespb_proto",
    "//pkg/roachpb:roachpb_proto",
    "//pkg/rpc/rpcpb:rpcpb_proto",
    "//pkg/server/diagnostics/diagnosticspb:diagnosticspb_proto",
    "//pkg/server/serverpb:serverpb_proto",
    "//pkg/server/status/statuspb:statuspb_proto",
    "//pkg/settings:settings_proto",
    "//pkg/sql/appstatspb:appstatspb_proto",
    "//pkg/sql/catalog/catenumpb:catenumpb_proto",
    "//pkg/sql/catalog/catpb:catpb_proto",
    "//pkg/sql/catalog/descpb:descpb_proto",
    "//pkg/sql/catalog/fetchpb:fetchpb_proto",
    "//pkg/sql/contentionpb:contentionpb_proto",
    "//pkg/sql/execinfrapb:execinfrapb_proto",
    "//pkg/sql/sem/semenumpb:semenumpb_proto",
    "//pkg/sql/inverted:inverted_proto",
    "//pkg/sql/lex:lex_proto",
    "//pkg/sql/pgwire/pgerror:pgerror_proto",
    "//pkg/sql/schemachanger/scpb:scpb_proto",
    "//pkg/sql/sessiondatapb:sessiondatapb_proto",
    "//pkg/sql/sqlstats/insights:insights_proto",
    "//pkg/sql/types:types_proto",
    "//pkg/storage/enginepb:enginepb_proto",
    "//pkg/ts/catalog:catalog_proto",
    "//pkg/ts/tspb:tspb_proto",
    "//pkg/util/duration:duration_proto",
    "//pkg/util/hlc:hlc_proto",
    "//pkg/util/admission/admissionpb:admissionpb_proto",
    "//pkg/util/log/logpb:logpb_proto",
    "//pkg/util/metric:metric_proto",
    "//pkg/util/timeutil/pgdate:pgdate_proto",
    "//pkg/util/tracing/tracingpb:tracingpb_proto",
    "//pkg/util:util_proto",
    "@com_github_prometheus_client_model//io/prometheus/client:io_prometheus_client_proto",
    "@com_github_cockroachdb_errors//errorspb:errorspb_proto",
    "@com_github_gogo_protobuf//gogoproto:gogo_proto",
    "@com_google_protobuf//:any_proto",
    "@com_google_protobuf//:descriptor_proto",
    "@com_google_protobuf//:duration_proto",
    "@com_google_protobuf//:timestamp_proto",
    "@go_googleapis//google/api:annotations_proto",
    "@go_googleapis//google/api:http_proto",
    "@io_etcd_go_raft_v3//raftpb:raftpb_proto",
]