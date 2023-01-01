create table if not exists articles (
	article_id integer unsigned auto_increment primary key,
	title varchar(100) not null,
	contents text not null,
	username varchar(100) not null,
	nice integer not null,
	created_at datetime
);

insert into articles (title, contents, username, nice, created_at) values
	('firstPost', 'This is my first blog', 'Takuro Matsuo', 2, now());

insert into articles (title, contents, username, nice) values
	('2nd', 'Second blog post', 'Takuro Matsuo', 4);
