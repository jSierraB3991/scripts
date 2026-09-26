-- 1. CREACIÓN DE TABLAS (Sintaxis para SQLite)

CREATE TABLE usuarios (
    id_usuario INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    telefono TEXT,
    fecha_registro DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE productos (
    id_producto INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    categoria TEXT,
    precio REAL NOT NULL
);

CREATE TABLE stock (
    id_stock INTEGER PRIMARY KEY AUTOINCREMENT,
    id_producto INTEGER UNIQUE NOT NULL,
    cantidad_disponible INTEGER NOT NULL DEFAULT 0,
    ubicacion_almacen TEXT,
    FOREIGN KEY (id_producto) REFERENCES productos(id_producto)
);

CREATE TABLE compras (
    id_compra INTEGER PRIMARY KEY AUTOINCREMENT,
    id_usuario INTEGER NOT NULL,
    id_producto INTEGER NOT NULL,
    cantidad INTEGER NOT NULL,
    precio_unitario REAL NOT NULL,
    total REAL NOT NULL,
    fecha_compra DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_usuario) REFERENCES usuarios(id_usuario),
    FOREIGN KEY (id_producto) REFERENCES productos(id_producto)
);

-- 2. INSERCIÓN DE DATOS DE PRUEBA

INSERT INTO usuarios (nombre, email, telefono, fecha_registro) VALUES
('Carlos Mendoza', 'carlos.mendoza@email.com', '+573001234567', '2026-01-10 10:30:00'),
('Ana María Gómez', 'ana.gomez@email.com', '+573109876543', '2026-01-15 14:20:00'),
('Luis Fernando Torres', 'luis.torres@email.com', '+573204567890', '2026-02-01 09:00:00'),
('Sofia Rodriguez', 'sofia.rodriguez@email.com', '+573016549870', '2026-02-12 16:45:00'),
('David Morales', 'david.morales@email.com', '+573153216549', '2026-03-05 11:15:00');

INSERT INTO productos (nombre, categoria, precio) VALUES
('Laptop Pro 15"', 'Electrónica', 1200.00),
('Teclado Mecánico RGB', 'Accesorios', 85.50),
('Mouse Inalámbrico', 'Accesorios', 35.00),
('Monitor 27" 4K', 'Electrónica', 350.00),
('Silla Ergonómica', 'Mobiliario', 210.00);

INSERT INTO stock (id_producto, cantidad_disponible, ubicacion_almacen) VALUES
(1, 15, 'Pasillo A - Estante 1'),
(2, 45, 'Pasillo A - Estante 2'),
(3, 80, 'Pasillo A - Estante 2'),
(4, 10, 'Pasillo B - Estante 1'),
(5, 8,  'Pasillo C - Estante 3');

INSERT INTO compras (id_usuario, id_producto, cantidad, precio_unitario, total, fecha_compra) VALUES
(1, 1, 1, 1200.00, 1200.00, '2026-02-10 11:00:00'),
(1, 3, 2, 35.00, 70.00, '2026-02-10 11:05:00'),
(2, 2, 1, 85.50, 85.50, '2026-02-18 15:30:00'),
(3, 4, 1, 350.00, 350.00, '2026-03-01 10:15:00'),
(4, 5, 2, 210.00, 420.00, '2026-03-12 17:00:00'),
(5, 3, 1, 35.00, 35.00, '2026-03-20 12:40:00');