import { DataSource } from 'typeorm';
import { config } from 'dotenv';
import { TypeOrmUserEntity } from './typeorm-user.entity';
import { TypeOrmFriendEntity } from './typeorm-friend.entity';
import { TypeOrmConversationEntity } from './typeorm-conversation.entity';
import { TypeOrmConversationMemberEntity } from './typeorm-conversation-member.entity';
import { TypeOrmReadReceiptEntity } from './typeorm-read-receipt.entity';

config();

export default new DataSource({
  type: 'postgres',
  host: process.env.DB_HOST || 'localhost',
  port: parseInt(process.env.DB_PORT || '5432', 10),
  username: process.env.DB_USERNAME || 'postgres',
  password: process.env.DB_PASSWORD || 'postgrespassword',
  database: process.env.DB_DATABASE || 'chat_db',
  entities: [
    TypeOrmUserEntity,
    TypeOrmFriendEntity,
    TypeOrmConversationEntity,
    TypeOrmConversationMemberEntity,
    TypeOrmReadReceiptEntity,
  ],
  migrations: ['src/infrastructure/persistence/migrations/*.ts'],
  synchronize: false,
});
