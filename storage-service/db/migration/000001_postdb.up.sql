create table if not exists post (post_id BIGSERIAL primary key, title varchar(250), content text, author varchar(250), created_at timestamp without time zone);

create table if not exists comment ( comment_id BIGSERIAL primary key, post_id int, content text, author varchar(250), created_at timestamp, constraint fk_post foreign key (post_id) references post (post_id) on delete cascade);