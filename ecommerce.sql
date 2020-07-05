create database ecommerce;

create table product (
id varchar(20),
product_code varchar(20),
product_name varchar(50),
category_id varchar(20),
price int,
created_at datetime,
update_at datetime,
primary key(id));

create table category (
category_id varchar(20),
category_name varchar(20),
created_at datetime,
update_at datetime,
primary key(category_id));

create table purchase_order (
id_product varchar(20),
order_date date,
created_at datetime,
update_at datetime);

create table purchase_order_item (
order_id varchar(20),
qty int,
product_id varchar(20),
created_at datetime,
update_at datetime,
primary key(order_id));

select * from product ;
select * from category;
select * from purchase_order;
select * from purchase_order_item;
SELECT * FROM product WHERE id ='P01';

select month(o.order_date), year(o.order_date), count(poi.qty), 
sum(p.price*poi.qty) from purchase_order o join purchase_order_item poi on poi.product_id=o.id_product 
join product p on o.id_product=p.id group by 1;


select day(b.order_date), month(b.order_date), year(b.order_date),
a.product_name, a.price, (c.qty), (c.qty*a.price) from product a join purchase_order b on a.id = b.id_product 
join purchase_order_item c on b.id_product=c.product_id;


select day(o.order_date), month(o.order_date), year(o.order_date), p.product_name, p.price, (poi.qty), 
(p.price*poi.qty) from purchase_order o join purchase_order_item poi on poi.product_id=o.id_product 
join product p on o.id_product=p.id group by 4, 3, 2, 1, 5 order by 1;


insert into product values 
('P01', 'SAV', 'Sasa', 'C01', 5000 ,'2020-02-22', '2020-02-22'),
('P02', 'XAC', 'CleanNClear', 'C02', 23000, '2020-02-22', '2020-02-22'),
('P03', 'QSX', 'Tropicana', 'C03', 15000, '2020-02-22', '2020-02-22');

insert into purchase_order values 
('C01', 'Bumbu Dapur', '2020-02-22', '2020-02-22'),
('C02', 'Shampoo', '2020-02-22', '2020-02-22'),
('C03', 'Minyak', '2020-02-22', '2020-02-22');

insert into purchase_order values 
('P01', '2020-02-22', '2020-02-22', '2020-02-22'),
('P02', '2020-02-22', '2020-02-22', '2020-02-22'),
('P03', '2020-02-22', '2020-02-22', '2020-02-22'),
('P01', '2020-02-23', '2020-02-23', '2020-02-23'),
('P02', '2020-02-23', '2020-02-23', '2020-02-23'),
('P03', '2020-02-23', '2020-02-23', '2020-02-23');

insert into purchase_order_item values 
('OID1', 2, 'P01', '2020-02-22', '2020-02-22'),
('OID2', 2, 'P02', '2020-02-22', '2020-02-22'),
('OID3', 3, 'P03','2020-02-22', '2020-02-22'),
('OID4', 1, 'P01','2020-02-23', '2020-02-23'),
('OID5', 3, 'P02','2020-02-23', '2020-02-23'),
('OID6', 3, 'P03','2020-02-23', '2020-02-23');


