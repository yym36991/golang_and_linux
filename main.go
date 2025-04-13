package main

import (
	"fmt"
	"time"
)

func main() {
	startTs := time.Now().Unix()
	fmt.Println(startTs) // 输出: 0 false
}

/*
测试告警：
curl -X POST http://10.135.5.234:9500/store/v2/by_biz/sendsmscall \
-H "Content-Type: application/json" \
-d '{
    "caller": "wtable-health",
    "biz": 4,
    "alarmid": 10,
    "content": "This is an alarm message",
    "title": "Alarm Title",
    "id": 102,
    "clientIp": "10.0.0.1"
}'



iostat -x 1


修改集群的限流值：
curl -X POST "http://10.178.5.6:31001/set_proxy_config?cid=10" \
-H "Content-Type: application/json" \
-u your_username:your_usernamevBo6MDHf9YsGvuNE \
-d '{
    "org_addr": "10.162.5.214:8010",
    "cdn_addr": "wos10.58cdn.com.cn",
    "ret_sha2": 1,
    "ext_addr": "wos10.58.com",
    "white_appkey": [],
    "white_bucket": [],
    "dir_timeout": 0,
    "store_timeout": 0,
    "blacklist_agent": [],
    "black_dir": [],
    "fastcache_server": "fastcache-namec.58dns.org",
    "fastcache_bid": 196615,
    "url_mode":0,
    "fast_addr": "",
    "byte_read_limit": 30000000,
    "byte_write_limit": 30000000
}'


curl -X POST "http://10.178.5.6:31001/set_proxy_config?cid=10" \
-H "Content-Type: application/json" \
-u your_username:your_usernamevBo6MDHf9YsGvuNE \
-d '{
    "byte_read_limit": 300,
    "byte_write_limit": 300
}'





curl -X POST http://urltool.wos.58dns.org/api/v1/urls/disable \
-H "Content-Type: application/json" \
-u your_username:your_usernamevBo6MDHf9YsGvuNE \
-d '{"urls": ["http://example.com/url1", "http://example.com/url2"]}'

http://urltool.wos.58dns.org/api/v1/urls/disable






curl -X POST "http://10.178.5.6:31001/add_limit_app?cid=10" \
-H "Content-Type: application/json" \
-u your_username:your_usernamevBo6MDHf9YsGvuNE \
-d '{
    "appkey": "eFcBazYxWvem",
    "cid": "10",
    "byte_read_limit": 40000000,
    "byte_write_limit": 30000000
}'

dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", conf.DB.UserName, conf.DB.Password, conf.DB.Address, conf.DB.DBName)



curl -u username:usernameae93fa163e65b938 -X POST http://10.135.5.234:8005/get_general_service_info

curl -v -X POST http://10.186.23.96:9205/api/v1/urls/disable \
-H "Content-Type: application/x-www-form-urlencoded" \
-u "anquan:yYDcWqMOmPdr0dQ5" \
-d "urls=https://example.com/resource1,https://example.com/resource2"


https://wosmedia1.anjukestatic.com/ UcSrQpOnMUW/esf-7d96d11f/35202b277b3f16dee22b.mov 你好 能帮忙看一下这个视频是有什么问题吗


[DEBUG] single check nodes:[/commonservice/wos/alarmusage/10.135.5.234:32669 /commonservice/wos/alarmusage/10.186.22.162:32669 /commonservice/wos/alarmusage/ids /commonservice/wos/alarmusage/info]

etcdctl --endpoints="etcdn1.wtable.58dns.org:2879","etcdn2.wtable.58dns.org:2879","etcdn3.wtable.58dns.org:2879","etcdn4.wtable.58dns.org:2879","etcdn5.wtable.58dns.org:2879" --user="root" --password="VDOrKhfd1z0UvFZZ" get /commonservice/wos/dcserver/ --prefix


etcdctl --endpoints="etcdn1.wtable.58dns.org:2879","etcdn2.wtable.58dns.org:2879","etcdn3.wtable.58dns.org:2879","etcdn4.wtable.58dns.org:2879","etcdn5.wtable.58dns.org:2879" --user="root" --password="VDOrKhfd1z0UvFZZ" get /commonservice/wos/dcserver/10.178.5.6 --prefix


etcdctl --endpoints="etcdn1.wtable.58dns.org:2879","etcdn2.wtable.58dns.org:2879","etcdn3.wtable.58dns.org:2879","etcdn4.wtable.58dns.org:2879","etcdn5.wtable.58dns.org:2879" --user="root" --password="VDOrKhfd1z0UvFZZ" del /commonservice/wos/dcserver/10.178.5.6:32668


/commonservice/wos/dcserver/10.178.5.6:32668

$ pwd
/opt/storage-deploy/wos/urltool/1.0.27.yym-wos-common-1224/wos-urltool-1.0.27.yym-wos-common-1224/script



wangfeifei06,wangjiayuan,wangqiang61,yangguizeng,chenshouhao

插入数据wfs2机器ip信息
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/server/172.26.0.5/22991 '{"system":"wfs2","service_name":"deleteserver"}'
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/server/172.26.0.5/22993 '{"system":"wfs2","service_name":"deleteserver"}'

  etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/server/172.26.0.6/22991 '{"system":"wfs2","service_name":"deleteserver"}'
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/server/172.26.0.6/22993 '{"system":"wfs2","service_name":"deleteserver"}'

 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/server/172.26.0.6/22995 '{"system":"wfs2","service_name":"deleteserver"}'

 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/server/172.26.0.6/22997 '{"system":"wfs2","service_name":"deleteserver"}'


 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/deleteserver/10.135.5.234/22690


  etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/wfs2/deleteserver/ --prefix


插入数据wfs2 系统 服务名 信息
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wfs2/deleteserver/172.26.0.5:22991 '{"node_id":11,"address":"172.26.0.5:22991","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wfs2","startup_time":1734600246,"create_time":1734600240}'
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wfs2/deleteserver/172.26.0.5:22993 '{"node_id":11,"address":"172.26.0.5:22993","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wfs2","startup_time":1734600246,"create_time":1734600240}'

 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/wfs2/deleteserver/172.26.0.5:22991







  etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wfs2/deleteserver/172.26.0.6:22991 '{"node_id":16,"address":"172.26.0.6:22991","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wfs2","startup_time":1734600276,"create_time":1734600276}'






  etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wfs2/deleteserver/172.26.0.6:22993 '{"node_id":16,"address":"172.26.0.6:22993","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wfs2","startup_time":1734600276,"create_time":1734600267}'

 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wfs2/deleteserver/172.26.0.6:22995 '{"node_id":16,"address":"172.26.0.6:22995","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wfs2","startup_time":1734600277,"create_time":1734600268}'

 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wfs2/deleteserver/172.26.0.6:22997 '{"node_id":16,"address":"172.26.0.6:22997","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wfs2","startup_time":1734600277,"create_time":1734600270}'




etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/wfs2/deleteserver/172.26.0.6:22991


6、查看
etcdctl --endpoints=http://127.0.0.1:2379 --user=putong:putong del /commonservice/wfs2/deleteserver --prefix
7、删除
把上一条的get换成del

curl -X GET http://10.135.5.234:22690/get_service_info


curl -X GET http://10.135.5.234:32668/get_service_info




currentAddress=172.26.0.6:22991   node index=5,  redis = 1

currentAddress=172.26.0.6:22993, node index=6, redisIndex=2


查看ip对应的redis
etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/wfs2/redis/deleteserver/10.135.5.234

查看redis对应的ip列表
etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get  /commonservice/wfs2/redis/deleteserver/wfs2-delserver-es1.rdb.58dns.org:45470


etcdctl --endpoints=http://127.0.0.1:2379 --user=putong:putong get /commonservice/wfs2/redis/deleteserver --prefix

wos:
插入ip->服务名：
etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/dcserver/10.135.5.234/32668 '{"system":"wos","service_name":"dcserver"}'

获取ip对应的服务名：
etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/dcserver/10.135.5.234/32668

插入服务名->信息
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong put /commonservice/wos/dcserver/10.135.5.234:32668 '{"node_id":99,"address":"10.135.5.234:32668","version":"1.1.0","status":"offline","deploy_path":"/home/work/yangyanmeng/wos","startup_time":1733880326,"create_time":1733880325}'

 查看信息：
 etcdctl --endpoints=http://10.253.59.95:2379 --user=putong:putong get /commonservice/wos/dcserver/10.135.5.234:32668

etcdctl --endpoints=http://etcdv3n1.wtable-test.58dns.org:40379 --user=root:VovPrVJfucB@mS9K get /commonservice/wfs2/deleteserver/10.178.2.11:22690


etcd_servers = ["http://etcdv3n1.wtable-test.58dns.org:40379","http://etcdv3n2.wtable-test.58dns.org:40379","http://etcdv3n3.wtable-test.58dns.org:40379"]
        etcd_user   = "root"
        etcd_pwd    = "VovPrVJfucB@mS9K"


        SELECT count(*) FROM `db58_wos`.`t_bucket_file_info` where created_at < '2025-01-15 09:00:00' ORDER BY `created_at` asc LIMIT 1000;
SELECT `id`, `dump_id`, `c_id`, `app_id`, `create_time`, `bucket_id`, `upload_status`, `count`, `all_dump` FROM `db58_wos`.`t_wos_dump` ORDER BY `id` DESC LIMIT 1000;

*/
