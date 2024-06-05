
CREATE TABLE pets_data (
                      id INT PRIMARY KEY AUTO_INCREMENT,
                      name VARCHAR(255) NOT NULL,
                      owner VARCHAR(255) NOT NULL,
                      species VARCHAR(255) NOT NULL,
                      birth DATE,
                      death DATE
)
CREATE TABLE pet_events (
                        id INT PRIMARY KEY AUTO_INCREMENT,
                        pet_id INT NOT NULL,
                        date DATE NOT NULL,
                        type VARCHAR(255) NOT NULL,
                        remark TEXT,
                        FOREIGN KEY (pet_id) REFERENCES pets_data(id)
)


