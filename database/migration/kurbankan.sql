-- Qurban Management SaaS - PostgreSQL Schema

-- MASTER WILAYAH
CREATE TABLE provinces (
  id SERIAL PRIMARY KEY,
  code VARCHAR(10) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE regencies (
  id SERIAL PRIMARY KEY,
  code VARCHAR(10) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  province_code VARCHAR(10),
  FOREIGN KEY (province_code) REFERENCES provinces(code)
);

CREATE TABLE districts (
  id SERIAL PRIMARY KEY,
  code VARCHAR(15) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  regency_code VARCHAR(10),
  FOREIGN KEY (regency_code) REFERENCES regencies(code)
);

CREATE TABLE villages (
  id SERIAL PRIMARY KEY,
  code VARCHAR(20) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  district_code VARCHAR(15),
  FOREIGN KEY (district_code) REFERENCES districts(code)
);

-- USERS
CREATE TYPE user_platform_role AS ENUM ('owner','admin','support');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  platform_role user_platform_role DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- MOSQUES
CREATE TABLE mosques (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  name TEXT NOT NULL,
  address TEXT,
  photos TEXT,
  province_code VARCHAR(10),
  regency_code VARCHAR(10),
  district_code VARCHAR(15),
  village_code VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (province_code) REFERENCES provinces(code),
  FOREIGN KEY (regency_code) REFERENCES regencies(code),
  FOREIGN KEY (district_code) REFERENCES districts(code),
  FOREIGN KEY (village_code) REFERENCES villages(code)
);

CREATE TYPE mosque_member_role AS ENUM ('admin','committee','viewer');

CREATE TABLE mosque_members (
  id SERIAL PRIMARY KEY,
  mosque_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  role mosque_member_role DEFAULT 'viewer',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (mosque_id) REFERENCES mosques(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- PARTICIPANTS
CREATE TABLE participants (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  name TEXT NOT NULL,
  address TEXT,
  province_code VARCHAR(10),
  regency_code VARCHAR(10),
  district_code VARCHAR(15),
  village_code VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (province_code) REFERENCES provinces(code),
  FOREIGN KEY (regency_code) REFERENCES regencies(code),
  FOREIGN KEY (district_code) REFERENCES districts(code),
  FOREIGN KEY (village_code) REFERENCES villages(code)
);

-- QURBAN PERIOD
CREATE TABLE qurban_periods (
  id SERIAL PRIMARY KEY,
  year INTEGER NOT NULL,
  start_date DATE,
  end_date DATE,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);

-- OPTIONS
CREATE TYPE animal_type AS ENUM ('cow','goat');
CREATE TYPE scheme_type AS ENUM ('group','individual');

CREATE TABLE qurban_options (
  id SERIAL PRIMARY KEY,
  qurban_period_id INTEGER,
  animal_type animal_type,
  scheme_type scheme_type,
  price DECIMAL(12,2),
  slots INTEGER DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (qurban_period_id) REFERENCES qurban_periods(id)
);

-- TRANSACTIONS
CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  code VARCHAR(100) UNIQUE,
  qurban_period_id INTEGER,
  mosque_id INTEGER,
  qurban_option_id INTEGER,
  is_full BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (qurban_period_id) REFERENCES qurban_periods(id),
  FOREIGN KEY (mosque_id) REFERENCES mosques(id),
  FOREIGN KEY (qurban_option_id) REFERENCES qurban_options(id)
);

CREATE TYPE transaction_status AS ENUM ('pending','paid','cancelled');
CREATE TYPE payment_type AS ENUM ('VA');

CREATE TABLE transaction_items (
  id SERIAL PRIMARY KEY,
  transaction_id INTEGER,
  participant_id INTEGER,
  amount DECIMAL(12,2),
  status transaction_status DEFAULT 'pending',
  payment_type payment_type,
  external_id VARCHAR(255),
  paid_at TIMESTAMP NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (participant_id) REFERENCES participants(id)
);

-- QURBAN ANIMALS
CREATE TYPE animal_status AS ENUM ('purchased','arrived','slaughtered','processed','distributed');

CREATE TABLE qurban_animals (
  id SERIAL PRIMARY KEY,
  mosque_id INTEGER,
  qurban_period_id INTEGER,
  type animal_type,
  name VARCHAR(100),
  weight DECIMAL(6,2),
  price DECIMAL(12,2),
  status animal_status DEFAULT 'purchased',
  slaughtered_at TIMESTAMP NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (mosque_id) REFERENCES mosques(id),
  FOREIGN KEY (qurban_period_id) REFERENCES qurban_periods(id)
);

-- ANIMAL PARTICIPANTS
CREATE TABLE animal_participants (
  id SERIAL PRIMARY KEY,
  animal_id INTEGER,
  participant_id INTEGER,
  transaction_item_id INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (animal_id) REFERENCES qurban_animals(id),
  FOREIGN KEY (participant_id) REFERENCES participants(id),
  FOREIGN KEY (transaction_item_id) REFERENCES transaction_items(id)
);

-- BENEFICIARIES
CREATE TABLE beneficiaries (
  id SERIAL PRIMARY KEY,
  mosque_id INTEGER,
  name TEXT,
  address TEXT,
  phone VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (mosque_id) REFERENCES mosques(id)
);

-- DISTRIBUTION
CREATE TABLE distribution_batches (
  id SERIAL PRIMARY KEY,
  mosque_id INTEGER,
  qurban_period_id INTEGER,
  total_packages INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (mosque_id) REFERENCES mosques(id),
  FOREIGN KEY (qurban_period_id) REFERENCES qurban_periods(id)
);

CREATE TABLE distribution_items (
  id SERIAL PRIMARY KEY,
  batch_id INTEGER,
  beneficiary_id INTEGER,
  received_at TIMESTAMP NULL,
  notes TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (batch_id) REFERENCES distribution_batches(id),
  FOREIGN KEY (beneficiary_id) REFERENCES beneficiaries(id)
);

-- MEDIA
CREATE TYPE media_type AS ENUM ('photo','video');

CREATE TABLE animal_media (
  id SERIAL PRIMARY KEY,
  animal_id INTEGER,
  url TEXT,
  type media_type,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (animal_id) REFERENCES qurban_animals(id)
);

-- LOGS
CREATE TABLE animal_logs (
  id SERIAL PRIMARY KEY,
  animal_id INTEGER,
  status VARCHAR(50),
  note TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (animal_id) REFERENCES qurban_animals(id)
);

-- REPORTS
CREATE TABLE reports (
  id SERIAL PRIMARY KEY,
  mosque_id INTEGER,
  qurban_period_id INTEGER,
  generated_by INTEGER,
  snapshot_json JSONB,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (mosque_id) REFERENCES mosques(id),
  FOREIGN KEY (qurban_period_id) REFERENCES qurban_periods(id),
  FOREIGN KEY (generated_by) REFERENCES users(id)
);

-- INDEXES untuk performa query
CREATE INDEX idx_mosques_user_id ON mosques(user_id);
CREATE INDEX idx_mosques_deleted_at ON mosques(deleted_at);
CREATE INDEX idx_participants_user_id ON participants(user_id);
CREATE INDEX idx_transactions_mosque_id ON transactions(mosque_id);
CREATE INDEX idx_transactions_period_id ON transactions(qurban_period_id);
CREATE INDEX idx_transaction_items_transaction_id ON transaction_items(transaction_id);
CREATE INDEX idx_transaction_items_status ON transaction_items(status);
CREATE INDEX idx_qurban_animals_mosque_id ON qurban_animals(mosque_id);
CREATE INDEX idx_qurban_animals_period_id ON qurban_animals(qurban_period_id);
CREATE INDEX idx_reports_mosque_id ON reports(mosque_id);
CREATE INDEX idx_reports_period_id ON reports(qurban_period_id);
