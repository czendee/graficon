CREATE USER banwire;

CREATE DATABASE gs_ivr_tokenization OWNER banwire;


CREATE TABLE banwirecard(
id_card bigserial,
token    varchar(100),   --constrain unique
last_digits   varchar(4),        
bin         varchar(6),
valid_thru    timestamp, 
score     integer,
is_banned          boolean,
banned_at         timestamp, 
created_at         timestamp,
last_update_at timestamp ,
id_customer      bigint,
brand   varchar(20) ,
type_card varchar(15)
);


CREATE TABLE banwirecustomer(
id_customer bigserial,
reference    varchar(100),   --constrain unique
created_at         timestamp,
last_update_at timestamp 
);

CREATE TABLE banwirepayment(
id_payment bigserial,
token    varchar(100),   --constrain unique
created_at         timestamp,
amount      bigint,
);

---mysql

CREATE TABLE banwiredash01graph01(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint,  
numsecuecial    bigint,     
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
    PRIMARY KEY (id_record)
);


CREATE TABLE banwiredash01graph02(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint,       
numsecuecial    bigint,
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
    PRIMARY KEY (id_record)
);


CREATE TABLE banwiredash01graph03(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint,   
numsecuecial    bigint,    
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
    PRIMARY KEY (id_record)
);



CREATE TABLE banwiredash02graph01(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint,  
numsecuecial    bigint,     
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
    PRIMARY KEY (id_record)
);


CREATE TABLE banwiredash02graph02(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint, 
numsecuecial    bigint,      
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
    PRIMARY KEY (id_record)
);


CREATE TABLE banwiredash02graph03(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint,
numsecuecial    bigint,       
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
    PRIMARY KEY (id_record)
);
-----postgres

CREATE TABLE banwiredash01graph01(
id_record bigserial,
id_grupoddatos    bigint, 
numsecuecial    bigint,      
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
);

CREATE TABLE banwiredash01graph02(
id_record bigserial,
id_grupoddatos    bigint,  
numsecuecial    bigint,     
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
);

CREATE TABLE banwiredash01graph03(
id_record bigserial,
id_grupoddatos    bigint,       
numsecuecial    bigint,
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
);