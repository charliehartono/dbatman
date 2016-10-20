#!/usr/bin/env bash

TEST_CASES='
disable_query_logs
test_view_sp_list_fields
client_query
test_prepare_insert_update
test_fetch_seek
test_fetch_nobuffs
test_open_direct
test_fetch_null
test_ps_null_param
test_fetch_date
test_fetch_str
test_fetch_long
test_fetch_short
test_fetch_tiny
test_fetch_bigint
test_fetch_float
test_fetch_double
test_bind_result_ext
test_bind_result_ext1
test_select_direct
test_select_prepare
test_select
test_select_version
test_ps_conj_select
test_select_show_table
test_func_fields
test_long_data
test_insert
test_set_variable
test_select_show
test_prepare_noparam
test_bind_result
test_prepare_simple
test_prepare
test_null
test_debug_example
test_update
test_simple_update
test_simple_delete
test_double_compare
client_store_result
client_use_result
test_tran_bdb
test_tran_innodb
test_prepare_ext
test_prepare_syntax
test_field_names
test_field_flags
test_long_data_str
test_long_data_str1
test_long_data_bin
test_warnings
test_errors
test_prepare_resultset
test_stmt_close
test_prepare_field_result
test_multi_stmt
test_multi_statements
test_prepare_multi_statements
test_store_result
test_store_result1
test_store_result2
test_subselect
test_date
test_date_frac
test_temporal_param
test_date_date
test_date_time
test_date_ts
test_date_dt
test_prepare_alter
test_manual_sample
test_pure_coverage
test_buffers
test_ushort_bug
test_sshort_bug
test_stiny_bug
test_field_misc
test_set_option
test_prepare_grant
test_frm_bug
test_explain_bug
test_decimal_bug
test_nstmts
test_logs;
test_cuted_rows
test_fetch_offset
test_fetch_column
test_mem_overun
test_list_fields
test_free_result
test_free_store_result
test_sqlmode
test_ts
test_bug1115
test_bug1180
test_bug1500
test_bug1644
test_bug1946
test_bug2248
test_parse_error_and_bad_length
test_bug2247
test_subqueries
test_bad_union
test_distinct
test_subqueries_ref
test_union
test_bug3117
test_join
test_selecttmp
test_create_drop
test_rename
test_do_set
test_multi
test_insert_select
test_bind_nagative
test_derived
test_xjoin
test_bug3035
test_union2
test_bug1664
test_union_param
test_order_param
test_ps_i18n
test_bug3796
test_bug4026
test_bug4079
test_bug4236
test_bug4030
test_bug5126
test_bug4231
test_bug5399
test_bug5194
test_bug5315
test_bug6049
test_bug6058
test_bug6059
test_bug6046
test_bug6081
test_bug6096
test_datetime_ranges
test_bug4172
test_conversion
test_rewind
test_bug6761
test_view
test_view_where
test_view_2where
test_view_star
test_view_insert
test_left_join_view
test_view_insert_fields
test_basic_cursors
test_cursors_with_union
test_cursors_with_procedure
test_truncation
test_truncation_option
test_client_character_set
test_bug8330
test_bug7990
test_bug8378
test_bug8722
test_bug8880
test_bug9159
test_bug9520
test_bug9478
test_bug9643
test_bug10729
test_bug11111
test_bug9992
test_bug10736
test_bug10794
test_bug11172
test_bug11656
test_bug10214
test_bug21246
test_bug9735
test_bug11183
test_bug11037
test_bug10760
test_bug12001
test_bug11718
test_bug12925
test_bug11909
test_bug11901
test_bug11904
test_bug12243
test_bug14210
test_bug13488
test_bug13524
test_bug14845
test_opt_reconnect
test_bug15510
test_bug12744
test_bug16143
test_bug16144
test_bug15613
test_bug20152
test_bug14169
test_bug17667
test_bug15752
test_mysql_insert_id
test_bug19671
test_bug21206
test_bug21726
test_bug15518
test_bug23383
test_bug32265
test_bug21635
test_status
test_bug24179
test_ps_query_cache
test_bug28075
test_bug27876
test_bug28505
test_bug28934
test_bug27592
test_bug29687
test_bug29692
test_bug29306
test_change_user
test_bug30472
test_bug20023
test_bug45010
test_bug53371
test_bug31418
test_bug31669
test_bug28386
test_wl4166_1
test_wl4166_2
test_wl4166_3
test_wl4166_4
test_bug36004
test_wl4284_1
test_wl4435
test_wl4435_2
test_wl4435_3
test_bug38486
test_bug33831
test_bug40365
test_bug43560
test_bug36326
test_bug41078
test_bug44495
test_bug49972
test_bug42373
test_bug54041
test_bug47485
test_bug58036
test_bug57058
test_bug56976
test_bug11766854
test_bug54790
test_bug12337762
test_bug11754979
test_bug13001491
test_wl5968
test_wl5924
test_wl6587
test_wl5928
test_wl6797
test_wl6791
test_wl5768
test_bug17309863
test_bug17512527
test_bug20810928
test_wl8016
test_bug20645725
test_bug20444737
test_bug21104470
test_bug21293012
test_bug21199582
test_bug20821550
'

for i in ${TEST_CASES}; do
    echo ${i}
done