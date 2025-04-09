CREATE TABLE users(
  id smallint unsigned not null auto_increment,
  email varchar(255) not null,
  telegram_id bigint unsigned unique,
  type enum ('email', 'telegram') not null,
  primary key (id)
);
