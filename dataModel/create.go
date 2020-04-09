package dataModel

const (
	Create_table_t_start_giirl =`create table if not exists t_start_girl
(
	id int auto_increment
		primary key,
	device_id varchar(255) not null,
	uid int not null,
	vote int default 0 null comment '得票数',
	rate_send_story int default 0 null comment '发story任务次',
	rate_voice_chat int default 0 null comment '语音聊天任务次',
	rate_words_chat int default 0 null comment '短信聊天任务次',
	rate_group_chat int default 0 null comment '群聊任务次',
	rate_thumb int default 0 null comment '点赞任务次',
	points int default 0 null comment '积分数',
	stage int default 1 not null comment '1:初选；2:决赛',
	first_name varchar(50) not null,
	last_name varchar(50) not null,
	university varchar(100) not null,
	birthday varchar(30) not null,
	contact_details varchar(30) not null comment '联系方式',
	slogan varchar(500) not null comment '口号',
	create_time bigint default 0 null comment '创建时间',
	constraint t_start_girl_device_id_uindex
		unique (device_id)
);`
)
