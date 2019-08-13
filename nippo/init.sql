create table reports(
  id int primary key auto_increment,
  title varchar(255) unique,
  body text,
  created_at datetime
)
