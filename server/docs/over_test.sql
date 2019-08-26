-- 用户表
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID' ,
  `user_name` varchar(150)  NOT NULL COMMENT '用户名称' ,
  `passwd_sha1` varchar(150)  NOT NULL COMMENT 'sha1加密后的密码' ,
  `mobile` varchar(30)  DEFAULT NULL COMMENT '手机',
  `delete_flag` int(1) NOT NULL DEFAULT 0 COMMENT '逻辑删除标志，1:删除,0:未删除',
  `created_at` datetime(3) NOT NULL DEFAULT current_timestamp(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT current_timestamp(3) on update current_timestamp(3) COMMENT '更新时间',
  primary key (`id`) using btree,
  unique key `user_name_idx` (`user_name`) using btree
) engine = InnoDB
  AUTO_INCREMENT = 20
  DEFAULT charset = utf8mb4 
  COLLATE=utf8mb4_unicode_ci 
  row_format = dynamic
  COMMENT '用户表';