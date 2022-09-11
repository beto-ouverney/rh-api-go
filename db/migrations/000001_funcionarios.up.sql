CREATE TABLE IF NOT EXISTS Funcionarios(
    id SERIAL PRIMARY KEY,
    nome VARCHAR(20),
    sobrenome VARCHAR(60),
    documento VARCHAR(11),
    setor VARCHAR(20),
    salario_bruto DECIMAL(10,2),
    data_admissao DATE,
    saude BOOLEAN,
    dental BOOLEAN,
    transporte BOOLEAN,
    dependente INT,
    pensao DECIMAL(10,2)
    )
