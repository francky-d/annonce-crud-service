CREATE TABLE categories(
   id INT NOT NULL AUTO_INCREMENT,
   libelle VARCHAR(255) NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP ,
   PRIMARY KEY (id)
);

CREATE TABLE  marques(
     id INT NOT NULL,
     libelle VARCHAR(255) NOT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMP ,
     PRIMARY KEY (id)
);

CREATE TABLE  models(
    id INT NOT NULL,
    libelle VARCHAR(255) NOT NULL,
    marque_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP ,
    PRIMARY KEY (id),
    FOREIGN KEY (marque_id) REFERENCES marques(id)
);

CREATE TABLE articles(
     id INT NOT NULL AUTO_INCREMENT,
     titre VARCHAR(255) NOT NULL,
     contenu TEXT NOT NULL,
     categorie_id INT NOT NULL,
     marque_id INT,
     model_id INT,
     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMP NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (categorie_id) REFERENCES categories(id),
     FOREIGN KEY (marque_id) REFERENCES marques(id),
     FOREIGN KEY (model_id)  REFERENCES models(id)
);


/*
 ------------------------
 Insertion here
 -----------------------
 */

INSERT INTO categories (libelle, created_at)
VALUES
    ('emploi', NOW() ) ,
    ('automobile', NOW() ) ,
    ('immobilier', NOW() ) ;

INSERT INTO marques (id, libelle, created_at)
VALUES
    (1, 'audi', NOW()) ,
    (2, 'bmw', NOW()),
    (3, 'citroen', NOW());

INSERT INTO models (id, libelle, marque_id, created_at)
VALUES
    (1, 'Cabriolet', 1, NOW()) ,
    (2, 'Q2', 1, NOW()) ,
    (3, 'Q3', 1, NOW()) ,
    (4, 'Q5', 1, NOW()) ,
    (5, 'Q7', 1, NOW()) ,
    (6, 'Q8', 1, NOW()) ,
    (7, 'R8', 1, NOW()) ,
    (8, 'Rs3', 1, NOW()) ,
    (9, 'Rs4', 1, NOW()) ,
    (10, 'Rs5', 1, NOW()) ,
    (11, 'Rs7', 1, NOW()) ,
    (12, 'S3', 1, NOW()) ,
    (13, 'S4', 1, NOW()) ,
    (14, 'S4 Avant', 1, NOW()) ,
    (15, 'S4 Cabriolet', 1, NOW()) ,
    (16, 'S5', 1, NOW()) ,
    (17, 'S7', 1, NOW()),
    (18, 'S8', 1, NOW()) ,
    (19, 'SQ5', 1, NOW()) ,
    (20, 'SQ7', 1, NOW()) ,
    (22, 'Tt', 1, NOW()) ,
    (23, 'Tts', 1, NOW()) ,
    (24, 'V8', 1, NOW()) ,
    (25, 'Q3', 1, NOW()) ,
    (26, 'Q3', 1, NOW()) ,

    (27, 'M3', 2, NOW()) ,
    (28, 'M4', 2, NOW()) ,
    (29, 'M5', 2, NOW()) ,
    (30, 'M535', 2, NOW()) ,
    (31, 'M6', 2, NOW()) ,
    (32, 'M635', 2, NOW()) ,
    (33, 'Serie 1', 2, NOW()) ,
    (34, 'Serie 2', 2, NOW()) ,
    (35, 'Serie 3', 2, NOW()) ,
    (36, 'Serie 4', 2, NOW()) ,
    (37, 'Serie 5', 2, NOW()) ,
    (38, 'Serie 6', 2, NOW()) ,
    (39, 'Serie 7', 2, NOW()) ,
    (40, 'Serie 8', 2, NOW()) ,

    (41, 'C1', 3, NOW()) ,
    (42, 'C15', 3, NOW()) ,
    (43, 'C2', 3, NOW()) ,
    (44, 'C25', 3, NOW()) ,
    (45, 'C25D', 3, NOW()) ,
    (46, 'C25E', 3, NOW()),
    (47, 'C25TD', 3, NOW()) ,
    (48, 'C3', 3, NOW()) ,
    (49, 'C3 Aircross', 3, NOW()) ,
    (50, 'C3 Picasso', 3, NOW()) ,
    (51, 'C4', 3, NOW()) ,
    (52, 'C4 Picasso', 3, NOW()) ,
    (53, 'C5', 3, NOW()) ,
    (54, 'C6', 3, NOW()) ,
    (55, 'C8', 3, NOW()) ,
    (56, 'Ds3', 3, NOW()) ,
    (57, 'Ds4', 3, NOW()) ,
    (58, 'Ds5', 3, NOW());

