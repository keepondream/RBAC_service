-- 删除雪花算法方法时需要清除所有依赖的表才能够正确删除
drop SEQUENCE IF EXISTS global_id_seq;
drop FUNCTION IF EXISTS id_generator();