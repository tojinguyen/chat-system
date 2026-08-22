import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AuthModule } from './auth/auth.module';
import { UsersModule } from './users/users.module';
import { FriendsModule } from './friends/friends.module';
import { ConversationsModule } from './conversations/conversations.module';
import { TypeOrmUserEntity } from './infrastructure/persistence/typeorm-user.entity';
import { TypeOrmFriendEntity } from './infrastructure/persistence/typeorm-friend.entity';
import { TypeOrmConversationEntity } from './infrastructure/persistence/typeorm-conversation.entity';
import { TypeOrmConversationMemberEntity } from './infrastructure/persistence/typeorm-conversation-member.entity';
import { TypeOrmReadReceiptEntity } from './infrastructure/persistence/typeorm-read-receipt.entity';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: '.env',
    }),
    TypeOrmModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => ({
        type: 'postgres',
        host: configService.get<string>('DB_HOST', 'localhost'),
        port: configService.get<number>('DB_PORT', 5432),
        username: configService.get<string>('DB_USERNAME', 'postgres'),
        password: configService.get<string>('DB_PASSWORD', 'postgrespassword'),
        database: configService.get<string>('DB_DATABASE', 'chat_db'),
        entities: [
          TypeOrmUserEntity,
          TypeOrmFriendEntity,
          TypeOrmConversationEntity,
          TypeOrmConversationMemberEntity,
          TypeOrmReadReceiptEntity,
        ],
        synchronize: false,
        migrationsRun: true,
        migrations: [
          __dirname + '/infrastructure/persistence/migrations/*{.ts,.js}',
        ],
      }),
    }),
    AuthModule,
    UsersModule,
    FriendsModule,
    ConversationsModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
