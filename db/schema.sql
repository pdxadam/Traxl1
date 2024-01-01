CREATE SCHEMA IF NOT EXISTS dbTraxl;

-- ************************************** Traxl.users

CREATE TABLE dbTraxl.users
(
    PkUser         bigserial NOT NULL,
    UserName       text NOT NULL,
    PasswordHash   text NOT NULL,
    Name           text NOT NULL,
    Config         jsonb NOT NULL DEFAULT '{}'::JSONB,
    CreatedTs      timestamp NOT NULL DEFAULT NOW(),
    IsEnabled     boolean NOT NULL DEFAULT TRUE,
    CONSTRAINT PK_users PRIMARY KEY ( PkUser )
);


-- ************************************** gowebapp.exercises

CREATE TABLE dbTraxl.topics
(
    PkTopic   bigserial NOT NULL,
    TopicName text NOT NULL,
    FkUser    bigint NOT NULL,  
    CONSTRAINT PK_topics PRIMARY KEY ( PkTopic ),
    CONSTRAINT FK_User FOREIGN KEY ( FkUser ) REFERENCES dbTraxl.users ( PkUser )
);


-- ************************************** gowebapp.images

-- CREATE TABLE gowebapp.images
-- (
--     Image_ID     bigserial NOT NULL,
--     User_ID      bigserial NOT NULL,
--     Content_Type text NOT NULL DEFAULT 'image/png',
--     Image_Data   bytea NOT NULL,
--     CONSTRAINT PK_images PRIMARY KEY ( Image_ID, User_ID ),
--     CONSTRAINT FK_65 FOREIGN KEY ( User_ID ) REFERENCES gowebapp.users ( User_ID )
-- );

-- CREATE INDEX FK_67 ON gowebapp.images
--     (
--      User_ID
--         );


-- ************************************** gowebapp.sets

-- CREATE TABLE gowebapp.sets
-- (
--     Set_ID      bigserial NOT NULL,
--     Exercise_ID bigserial NOT NULL,
--     Weight      int NOT NULL DEFAULT 0,
--     CONSTRAINT PK_sets PRIMARY KEY ( Set_ID, Exercise_ID ),
--     CONSTRAINT FK_106 FOREIGN KEY ( Exercise_ID ) REFERENCES gowebapp.exercises ( Exercise_ID )
-- );

-- CREATE INDEX FK_108 ON gowebapp.sets
--     (
--      Exercise_ID
--         );

-- ************************************** db.instances

CREATE TABLE dbTraxl.instances
(
    PkInstance  bigserial NOT NULL,
    FkTopic    bigint NOT NULL,
    FkUser   bigint NOT NULL,
    Start_Date  timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT PK_Instance PRIMARY KEY ( PkInstance ),
    CONSTRAINT FK_Topic FOREIGN KEY ( FkTopic ) REFERENCES dbTraxl.topics ( PkTopic ),
    CONSTRAINT FK_User FOREIGN KEY ( FkUser ) REFERENCES dbTraxl.users ( PkUser )
);

-- CREATE INDEX FK_73 ON gowebapp.workouts
--     (
--      Set_ID,
--      Exercise_ID
--         );

-- CREATE INDEX FK_76 ON gowebapp.workouts
--     (
--      User_ID
--         );
