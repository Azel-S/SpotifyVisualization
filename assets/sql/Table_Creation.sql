CREATE TABLE Tracks
(
release_year INT NOT NULL,
release_month INT,
release_day INT,
track_id VARCHAR2(25) NOT NULL,
acousticness FLOAT NOT NULL,
danceability FLOAT NOT NULL,
track_mode INT NOT NULL,
popularity INT NOT NULL,
energy FLOAT NOT NULL,
track_key INT NOT NULL,
speechiness FLOAT NOT NULL,
loudness FLOAT NOT NULL,
instrumentalness FLOAT NOT NULL,
liveness FLOAT NOT NULL,
tempo FLOAT NOT NULL,
duration INT NOT NULL,
valence FLOAT NOT NULL,
time_signature INT NOT NULL,
explicit NUMBER(1) NOT NULL,
title VARCHAR2(640) NOT NULL,
PRIMARY KEY (track_id)
);

CREATE TABLE Artists
(
name VARCHAR2(512) NOT NULL,
followers INT NOT NULL,
popularity INT NOT NULL,
artist_id VARCHAR2(24) NOT NULL,
PRIMARY KEY (artist_id)
);

CREATE TABLE Artist_To_Tracks
(
artist_id VARCHAR2(25) NOT NULL,
track_id VARCHAR2(25) NOT NULL,
PRIMARY KEY (artist_id, track_id),
FOREIGN KEY (artist_id) REFERENCES Artists(artist_id),
FOREIGN KEY (track_id) REFERENCES Tracks(track_id)
);

CREATE TABLE Countries
(
name VARCHAR2(255) NOT NULL,
region VARCHAR2(255),
subregion VARCHAR2(255),
life_exp_male FLOAT,
population INT,
GNI FLOAT,
internet_use FLOAT,
tourism FLOAT,
physicians FLOAT,
life_exp_female FLOAT,
women_in_parl FLOAT,
hi_income_eco INT,
green_gas_emi FLOAT,
tert_edu_female FLOAT,
tert_edu_male FLOAT,
PRIMARY KEY (name)
);

CREATE TABLE Country_To_Code
(
country VARCHAR2(255) NOT NULL,
code VARCHAR2(2) NOT NULL,
PRIMARY KEY (code),
FOREIGN KEY (country) REFERENCES Countries(name)
);

CREATE TABLE Artist_To_Genres
(
artist_id VARCHAR2(25) NOT NULL,
genre VARCHAR2(255) NOT NULL,
PRIMARY KEY (artist_id, genre),
FOREIGN KEY (artist_id) REFERENCES Artists(artist_id)
);

CREATE TABLE Track_To_Country
(
track_id VARCHAR2(25) NOT NULL,
code VARCHAR2(2) NOT NULL,
PRIMARY KEY (track_id, code),
FOREIGN KEY (track_id) REFERENCES Tracks(track_id),
FOREIGN KEY (code) REFERENCES Country_To_Code(code)
);