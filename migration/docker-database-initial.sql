CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    status BOOLEAN,
    type VARCHAR(255),
    priority INTEGER,
    description VARCHAR(500),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tasks(name, status, type, priority, description) VALUES
('Comprar maça', TRUE, 'Compras', 1, 'Comprar maça para a salada'),
('Lavar a casa', TRUE, 'Tarefas domésticas', 3, 'Lavar a casa para festa');