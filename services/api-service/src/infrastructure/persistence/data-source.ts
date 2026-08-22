import { DataSource } from 'typeorm';
import { config } from 'dotenv';
import { TypeOrmUserEntity } from './typeorm-user.entity';

// Nạp biến môi trường từ file .env
config();

export default new DataSource({
  type: 'postgres',
  host: process.env.DB_HOST || 'localhost',
  port: parseInt(process.env.DB_PORT || '5432', 10),
  username: process.env.DB_USERNAME || 'postgres',
  password: process.env.DB_PASSWORD || 'postgrespassword',
  database: process.env.DB_DATABASE || 'chat_db',
  entities: [TypeOrmUserEntity],
  migrations: ['src/infrastructure/persistence/migrations/*.ts'],
  synchronize: false, // Luôn để false khi dùng migration
});
