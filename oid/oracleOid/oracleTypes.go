// 由gen.go生成的oracle模式的oid列表

package oracleOid

import (
	"kingbase.com/gokb/oid"
)

const (
	T_bool                           oid.Oid = 16;
	T_bytea                          oid.Oid = 17;
	T_char                           oid.Oid = 18;
	T_name                           oid.Oid = 19;
	T_int8                           oid.Oid = 20;
	T_int2                           oid.Oid = 21;
	T_int2vector                     oid.Oid = 22;
	T_int4                           oid.Oid = 23;
	T_regproc                        oid.Oid = 24;
	T_text                           oid.Oid = 25;
	T_oid                            oid.Oid = 26;
	T_tid                            oid.Oid = 27;
	T_xid                            oid.Oid = 28;
	T_cid                            oid.Oid = 29;
	T_oidvector                      oid.Oid = 30;
	T_pg_ddl_command                 oid.Oid = 32;
	T_pg_type                       oid.Oid = 71;
	T_pg_attribute                  oid.Oid = 75;
	T_pg_proc                       oid.Oid = 81;
	T_pg_class                      oid.Oid = 83;
	T_json                           oid.Oid = 114;
	T_xml                            oid.Oid = 142;
	T__xml                           oid.Oid = 143;
	T_pg_node_tree                   oid.Oid = 194;
	T__json                          oid.Oid = 199;
	T_pg_collation                  oid.Oid = 210;
	T_table_am_handler               oid.Oid = 269;
	T_index_am_handler               oid.Oid = 325;
	T_point                          oid.Oid = 600;
	T_lseg                           oid.Oid = 601;
	T_path                           oid.Oid = 602;
	T_box                            oid.Oid = 603;
	T_polygon                        oid.Oid = 604;
	T_line                           oid.Oid = 628;
	T__line                          oid.Oid = 629;
	T_cidr                           oid.Oid = 650;
	T__cidr                          oid.Oid = 651;
	T_float4                         oid.Oid = 700;
	T_float8                         oid.Oid = 701;
	T_unknown                        oid.Oid = 705;
	T_circle                         oid.Oid = 718;
	T__circle                        oid.Oid = 719;
	T_macaddr8                       oid.Oid = 774;
	T__macaddr8                      oid.Oid = 775;
	T_money                          oid.Oid = 790;
	T__money                         oid.Oid = 791
	T_macaddr                        oid.Oid = 829;
	T_inet                           oid.Oid = 869;
	T__bool                          oid.Oid = 1000;
	T__bytea                         oid.Oid = 1001;
	T__char                          oid.Oid = 1002;
	T__name                          oid.Oid = 1003;
	T__int2                          oid.Oid = 1005;
	T__int2vector                    oid.Oid = 1006;
	T__int4                          oid.Oid = 1007;
	T__regproc                       oid.Oid = 1008;
	T__text                          oid.Oid = 1009;
	T__tid                           oid.Oid = 1010;
	T__xid                           oid.Oid = 1011;
	T__cid                           oid.Oid = 1012;
	T__oidvector                     oid.Oid = 1013;
	T__bpchar                        oid.Oid = 1014;
	T__varchar                       oid.Oid = 1015;
	T__int8                          oid.Oid = 1016;
	T__point                         oid.Oid = 1017;
	T__lseg                          oid.Oid = 1018;
	T__path                          oid.Oid = 1019;
	T__box                           oid.Oid = 1020;
	T__float4                        oid.Oid = 1021;
	T__float8                        oid.Oid = 1022;
	T__polygon                       oid.Oid = 1027;
	T__oid                           oid.Oid = 1028;
	T_aclitem                        oid.Oid = 1033;
	T__aclitem                       oid.Oid = 1034;
	T__macaddr                       oid.Oid = 1040;
	T__inet                          oid.Oid = 1041;
	T_bpchar                         oid.Oid = 1042;
	T_varchar                        oid.Oid = 1043;
	//T_date                           oid.Oid = 1082;
	T_time                           oid.Oid = 1083;
	T_timestamp                      oid.Oid = 1114;
	T__timestamp                     oid.Oid = 1115;
	//T__date                          oid.Oid = 1182;
	T__time                          oid.Oid = 1183;
	T_timestamptz                    oid.Oid = 1184;
	T__timestamptz                   oid.Oid = 1185;
	T_interval                       oid.Oid = 1186;
	T__interval                      oid.Oid = 1187;
	T__numeric                       oid.Oid = 1231;
	T_pg_database                   oid.Oid = 1248;
	T__cstring                       oid.Oid = 1263;
	T_timetz                         oid.Oid = 1266;
	T__timetz                        oid.Oid = 1270;
	T_bit                            oid.Oid = 1560;
	T__bit                           oid.Oid = 1561;
	T_varbit                         oid.Oid = 1562;
	T__varbit                        oid.Oid = 1563;
	T_numeric                        oid.Oid = 1700;
	T_refcursor                      oid.Oid = 1790;
	T_pg_sysprivilege               oid.Oid = 2030;
	T__refcursor                     oid.Oid = 2201;
	T_regprocedure                   oid.Oid = 2202;
	T_regoper                        oid.Oid = 2203;
	T_regoperator                    oid.Oid = 2204;
	T_regclass                       oid.Oid = 2205;
	T_regtype                        oid.Oid = 2206;
	T__regprocedure                  oid.Oid = 2207;
	T__regoper                       oid.Oid = 2208;
	T__regoperator                   oid.Oid = 2209;
	T__regclass                      oid.Oid = 2210;
	T__regtype                       oid.Oid = 2211;
	T_record                         oid.Oid = 2249;
	T_cstring                        oid.Oid = 2275;
	T_any                            oid.Oid = 2276;
	T_anyarray                       oid.Oid = 2277;
	T_void                           oid.Oid = 2278;
	T_trigger                        oid.Oid = 2279;
	T_language_handler               oid.Oid = 2280;
	T_internal                       oid.Oid = 2281;
	T_opaque                         oid.Oid = 2282;
	T_anyelement                     oid.Oid = 2283;
	T__record                        oid.Oid = 2287;
	T_anynonarray                    oid.Oid = 2776;
	T_pg_authid                     oid.Oid = 2842;
	T_pg_auth_members               oid.Oid = 2843;
	T__txid_snapshot                 oid.Oid = 2949;
	T_uuid                           oid.Oid = 2950;
	T__uuid                          oid.Oid = 2951;
	T_txid_snapshot                  oid.Oid = 2970;
	T_fdw_handler                    oid.Oid = 3115;
	T_pg_lsn                         oid.Oid = 3220;
	T__pg_lsn                        oid.Oid = 3221;
	T_tsm_handler                    oid.Oid = 3310;
	T_pg_ndistinct                   oid.Oid = 3361;
	T_binary                         oid.Oid = 3383;
	T_varbinary                      oid.Oid = 3384;
	T__binary                        oid.Oid = 3385;
	T__varbinary                     oid.Oid = 3386;
	T_pg_dependencies                oid.Oid = 3402;
	T_anyenum                        oid.Oid = 3500;
	T_tsvector                       oid.Oid = 3614;
	T_tsquery                        oid.Oid = 3615;
	T_gtsvector                      oid.Oid = 3642;
	T__tsvector                      oid.Oid = 3643;
	T__gtsvector                     oid.Oid = 3644;
	T__tsquery                       oid.Oid = 3645;
	T_regconfig                      oid.Oid = 3734;
	T__regconfig                     oid.Oid = 3735;
	T_regdictionary                  oid.Oid = 3769;
	T__regdictionary                 oid.Oid = 3770;
	T_jsonb                          oid.Oid = 3802;
	T__jsonb                         oid.Oid = 3807;
	T_anyrange                       oid.Oid = 3831;
	T_event_trigger                  oid.Oid = 3838;
	T_int4range                      oid.Oid = 3904;
	T__int4range                     oid.Oid = 3905;
	T_numrange                       oid.Oid = 3906;
	T__numrange                      oid.Oid = 3907;
	T_tsrange                        oid.Oid = 3908;
	T__tsrange                       oid.Oid = 3909;
	T_tstzrange                      oid.Oid = 3910;
	T__tstzrange                     oid.Oid = 3911;
	T_daterange                      oid.Oid = 3912;
	T__daterange                     oid.Oid = 3913;
	T_int8range                      oid.Oid = 3926;
	T__int8range                     oid.Oid = 3927;
	T_pg_shseclabel                 oid.Oid = 4066;
	T_jsonpath                       oid.Oid = 4072;
	T__jsonpath                      oid.Oid = 4073;
	T_regnamespace                   oid.Oid = 4089;
	T__regnamespace                  oid.Oid = 4090;
	T_regrole                        oid.Oid = 4096;
	T__regrole                       oid.Oid = 4097;
	T_datetime                       oid.Oid = 4189;
	T__datetime                      oid.Oid = 4199;
	T_byteawithoutorderwithequalcol  oid.Oid = 4594;
	T_byteawithoutordercol           oid.Oid = 4595;
	T__byteawithoutorderwithequalcol oid.Oid = 4596;
	T__byteawithoutordercol          oid.Oid = 4597;
	T_anyset                         oid.Oid = 4600;
	T_pg_mcv_list                    oid.Oid = 5017;
	T_pg_subscription               oid.Oid = 6101;
	T_pg_sharedobject               oid.Oid = 6128;
	T_hash16                         oid.Oid = 6133;
	T__hash16                        oid.Oid = 6134;
	T_hash32                         oid.Oid = 6135;
	T__hash32                        oid.Oid = 6136;
	T_rowid                          oid.Oid = 6144;
	T__rowid                         oid.Oid = 6145;
	T_urowid                         oid.Oid = 6146;
	T__urowid                        oid.Oid = 6147;
	T__nclob                         oid.Oid = 6646;
	T__clob                          oid.Oid = 6647;
	T__blob                          oid.Oid = 6648;
	T_mysql_time                     oid.Oid = 6691;
	T__mysql_date                    oid.Oid = 6692;
	T_mysql_date                     oid.Oid = 6693;
	T__mysql_time                    oid.Oid = 6788;
	T_dsinterval                     oid.Oid = 7000;
	T__dsinterval                    oid.Oid = 7001;
	T_yminterval                     oid.Oid = 7002;
	T__yminterval                    oid.Oid = 7003;
	T_uint4                          oid.Oid = 7082;
	T__uint4                         oid.Oid = 7083;
	T_uint8                          oid.Oid = 7084;
	T__uint8                         oid.Oid = 7085;
	T_blob                           oid.Oid = 8013;
	T_clob                           oid.Oid = 8014;
	T_nclob                          oid.Oid = 8015;
	T_bpcharbyte                     oid.Oid = 8016;
	T__bpcharbyte                    oid.Oid = 8017;
	T_varcharbyte                    oid.Oid = 8018;
	T__varcharbyte                   oid.Oid = 8019;
	T_date                           oid.Oid = 8020;
	T__date                          oid.Oid = 8021;
	T_nestedtable                    oid.Oid = 8022;
	T_associative_array              oid.Oid = 8023;
	T_varray                         oid.Oid = 8024;
	T_pg_privilege                  oid.Oid = 8043;
	T_pg_audit_userlog              oid.Oid = 8047;
	T_pg_audit_blocklog             oid.Oid = 8049;
	T_kdb_job                        oid.Oid = 8075;
	T_bfile                          oid.Oid = 8078;
	T__bfile                         oid.Oid = 8079;
	T_tinyint                        oid.Oid = 8100;
	T__tinyint                       oid.Oid = 8101;
)

var TypeName = map[oid.Oid]string{
	T_bool:                           "BOOL",
	T_bytea:                          "BYTEA",
	T_char:                           "CHAR",
	T_name:                           "NAME",
	T_int8:                           "INT8",
	T_int2:                           "INT2",
	T_int2vector:                     "INT2VECTOR",
	T_int4:                           "INT4",
	T_regproc:                        "REGPROC",
	T_text:                           "TEXT",
	T_oid:                            "OID",
	T_tid:                            "TID",
	T_xid:                            "XID",
	T_cid:                            "CID",
	T_oidvector:                      "OIDVECTOR",
	T_pg_ddl_command:                 "KB_DDL_COMMAND",
	T_pg_type:                       "KB_TYPE",
	T_pg_attribute:                  "KB_ATTRIBUTE",
	T_pg_proc:                       "KB_PROC",
	T_pg_class:                      "KB_CLASS",
	T_json:                           "JSON",
	T_xml:                            "XML",
	T__xml:                           "_XML",
	T_pg_node_tree:                   "KB_NODE_TREE",
	T__json:                          "_JSON",
	T_pg_collation:                  "KB_COLLATION",
	T_table_am_handler:               "TABLE_AM_HANDLER",
	T_index_am_handler:               "INDEX_AM_HANDLER",
	T_point:                          "POINT",
	T_lseg:                           "LSEG",
	T_path:                           "PATH",
	T_box:                            "BOX",
	T_polygon:                        "POLYGON",
	T_line:                           "LINE",
	T__line:                          "_LINE",
	T_cidr:                           "CIDR",
	T__cidr:                          "_CIDR",
	T_float4:                         "FLOAT4",
	T_float8:                         "FLOAT8",
	T_unknown:                        "UNKNOWN",
	T_circle:                         "CIRCLE",
	T__circle:                        "_CIRCLE",
	T_macaddr8:                       "MACADDR8",
	T__macaddr8:                      "_MACADDR8",
	T_money:                          "MONEY",
	T__money:                         "_MONEY",
	T_macaddr:                        "MACADDR",
	T_inet:                           "INET",
	T__bool:                          "_BOOL",
	T__bytea:                         "_BYTEA",
	T__char:                          "_CHAR",
	T__name:                          "_NAME",
	T__int2:                          "_INT2",
	T__int2vector:                    "_INT2VECTOR",
	T__int4:                          "_INT4",
	T__regproc:                       "_REGPROC",
	T__text:                          "_TEXT",
	T__tid:                           "_TID",
	T__xid:                           "_XID",
	T__cid:                           "_CID",
	T__oidvector:                     "_OIDVECTOR",
	T__bpchar:                        "_BPCHAR",
	T__varchar:                       "_VARCHAR",
	T__int8:                          "_INT8",
	T__point:                         "_POINT",
	T__lseg:                          "_LSEG",
	T__path:                          "_PATH",
	T__box:                           "_BOX",
	T__float4:                        "_FLOAT4",
	T__float8:                        "_FLOAT8",
	T__polygon:                       "_POLYGON",
	T__oid:                           "_OID",
	T_aclitem:                        "ACLITEM",
	T__aclitem:                       "_ACLITEM",
	T__macaddr:                       "_MACADDR",
	T__inet:                          "_INET",
	T_bpchar:                         "BPCHAR",
	T_varchar:                        "VARCHAR",
	//T_date:                           "DATE",
	T_time:                           "TIME",
	T_timestamp:                      "TIMESTAMP",
	T__timestamp:                     "_TIMESTAMP",
	//T__date:                          "_DATE",
	T__time:                          "_TIME",
	T_timestamptz:                    "TIMESTAMPTZ",
	T__timestamptz:                   "_TIMESTAMPTZ",
	T_interval:                       "INTERVAL",
	T__interval:                      "_INTERVAL",
	T__numeric:                       "_NUMERIC",
	T_pg_database:                   "KB_DATABASE",
	T__cstring:                       "_CSTRING",
	T_timetz:                         "TIMETZ",
	T__timetz:                        "_TIMETZ",
	T_bit:                            "BIT",
	T__bit:                           "_BIT",
	T_varbit:                         "VARBIT",
	T__varbit:                        "_VARBIT",
	T_numeric:                        "NUMERIC",
	T_refcursor:                      "REFCURSOR",
	T_pg_sysprivilege:               "KB_SYSPRIVILEGE",
	T__refcursor:                     "_REFCURSOR",
	T_regprocedure:                   "REGPROCEDURE",
	T_regoper:                        "REGOPER",
	T_regoperator:                    "REGOPERATOR",
	T_regclass:                       "REGCLASS",
	T_regtype:                        "REGTYPE",
	T__regprocedure:                  "_REGPROCEDURE",
	T__regoper:                       "_REGOPER",
	T__regoperator:                   "_REGOPERATOR",
	T__regclass:                      "_REGCLASS",
	T__regtype:                       "_REGTYPE",
	T_record:                         "RECORD",
	T_cstring:                        "CSTRING",
	T_any:                            "ANY",
	T_anyarray:                       "ANYARRAY",
	T_void:                           "VOID",
	T_trigger:                        "TRIGGER",
	T_language_handler:               "LANGUAGE_HANDLER",
	T_internal:                       "INTERNAL",
	T_opaque:                         "OPAQUE",
	T_anyelement:                     "ANYELEMENT",
	T__record:                        "_RECORD",
	T_anynonarray:                    "ANYNONARRAY",
	T_pg_authid:                     "KB_AUTHID",
	T_pg_auth_members:               "KB_AUTH_MEMBERS",
	T__txid_snapshot:                 "_TXID_SNAPSHOT",
	T_uuid:                           "UUID",
	T__uuid:                          "_UUID",
	T_txid_snapshot:                  "TXID_SNAPSHOT",
	T_fdw_handler:                    "FDW_HANDLER",
	T_pg_lsn:                         "KB_LSN",
	T__pg_lsn:                        "_KB_LSN",
	T_tsm_handler:                    "TSM_HANDLER",
	T_pg_ndistinct:                   "KB_NDISTINCT",
	T_binary:                         "BINARY",
	T_varbinary:                      "VARBINARY",
	T__binary:                        "_BINARY",
	T__varbinary:                     "_VARBINARY",
	T_pg_dependencies:                "KB_DEPENDENCIES",
	T_anyenum:                        "ANYENUM",
	T_tsvector:                       "TSVECTOR",
	T_tsquery:                        "TSQUERY",
	T_gtsvector:                      "GTSVECTOR",
	T__tsvector:                      "_TSVECTOR",
	T__gtsvector:                     "_GTSVECTOR",
	T__tsquery:                       "_TSQUERY",
	T_regconfig:                      "REGCONFIG",
	T__regconfig:                     "_REGCONFIG",
	T_regdictionary:                  "REGDICTIONARY",
	T__regdictionary:                 "_REGDICTIONARY",
	T_jsonb:                          "JSONB",
	T__jsonb:                         "_JSONB",
	T_anyrange:                       "ANYRANGE",
	T_event_trigger:                  "EVENT_TRIGGER",
	T_int4range:                      "INT4RANGE",
	T__int4range:                     "_INT4RANGE",
	T_numrange:                       "NUMRANGE",
	T__numrange:                      "_NUMRANGE",
	T_tsrange:                        "TSRANGE",
	T__tsrange:                       "_TSRANGE",
	T_tstzrange:                      "TSTZRANGE",
	T__tstzrange:                     "_TSTZRANGE",
	T_daterange:                      "DATERANGE",
	T__daterange:                     "_DATERANGE",
	T_int8range:                      "INT8RANGE",
	T__int8range:                     "_INT8RANGE",
	T_pg_shseclabel:                 "KB_SHSECLABEL",
	T_jsonpath:                       "JSONPATH",
	T__jsonpath:                      "_JSONPATH",
	T_regnamespace:                   "REGNAMESPACE",
	T__regnamespace:                  "_REGNAMESPACE",
	T_regrole:                        "REGROLE",
	T__regrole:                       "_REGROLE",
	T_datetime:                       "DATETIME",
	T__datetime:                      "_DATETIME",
	T_byteawithoutorderwithequalcol:  "BYTEAWITHOUTORDERWITHEQUALCOL",
	T_byteawithoutordercol:           "BYTEAWITHOUTORDERCOL",
	T__byteawithoutorderwithequalcol: "_BYTEAWITHOUTORDERWITHEQUALCOL",
	T__byteawithoutordercol:          "_BYTEAWITHOUTORDERCOL",
	T_anyset:                         "ANYSET",
	T_pg_mcv_list:                    "KB_MCV_LIST",
	T_pg_subscription:               "KB_SUBSCRIPTION",
	T_pg_sharedobject:               "KB_SHAREDOBJECT",
	T_hash16:                         "HASH16",
	T__hash16:                        "_HASH16",
	T_hash32:                         "HASH32",
	T__hash32:                        "_HASH32",
	T_rowid:                          "ROWID",
	T__rowid:                         "_ROWID",
	T_urowid:                         "UROWID",
	T__urowid:                        "_UROWID",
	T__nclob:                         "_NCLOB",
	T__clob:                          "_CLOB",
	T__blob:                          "_BLOB",
	T_mysql_time:                     "MYSQL_TIME",
	T__mysql_date:                    "_MYSQL_DATE",
	T_mysql_date:                     "MYSQL_DATE",
	T__mysql_time:                    "_MYSQL_TIME",
	T_dsinterval:                     "DSINTERVAL",
	T__dsinterval:                    "_DSINTERVAL",
	T_yminterval:                     "YMINTERVAL",
	T__yminterval:                    "_YMINTERVAL",
	T_uint4:                          "UINT4",
	T__uint4:                         "_UINT4",
	T_uint8:                          "UINT8",
	T__uint8:                         "_UINT8",
	T_blob:                           "BLOB",
	T_clob:                           "CLOB",
	T_nclob:                          "NCLOB",
	T_bpcharbyte:                     "BPCHARBYTE",
	T__bpcharbyte:                    "_BPCHARBYTE",
	T_varcharbyte:                    "VARCHARBYTE",
	T__varcharbyte:                   "_VARCHARBYTE",
	T_date:                           "DATE",
	T__date:                          "_DATE",
	T_nestedtable:                    "NESTEDTABLE",
	T_associative_array:              "ASSOCIATIVE_ARRAY",
	T_varray:                         "VARRAY",
	T_pg_privilege:                  "KB_PRIVILEGE",
	T_pg_audit_userlog:              "KB_AUDIT_USERLOG",
	T_pg_audit_blocklog:             "KB_AUDIT_BLOCKLOG",
	T_kdb_job:                        "KDB_JOB",
	T_bfile:                          "BFILE",
	T__bfile:                         "_BFILE",
	T_tinyint:                        "TINYINT",
	T__tinyint:                       "_TINYINT",
}
