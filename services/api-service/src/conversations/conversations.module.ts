import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConversationsController } from './conversations.controller';
import { ConversationsService } from './conversations.service';
import { TypeOrmConversationEntity } from 'src/infrastructure/persistence/typeorm-conversation.entity';
import { TypeOrmConversationMemberEntity } from 'src/infrastructure/persistence/typeorm-conversation-member.entity';
import { TypeOrmReadReceiptEntity } from 'src/infrastructure/persistence/typeorm-read-receipt.entity';
import { CONVERSATION_REPOSITORY } from 'src/domain/repositories/conversation.repository.interface';
import { TypeOrmConversationRepository } from 'src/infrastructure/persistence/typeorm-conversation.repository';
import { UsersModule } from 'src/users/users.module';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      TypeOrmConversationEntity,
      TypeOrmConversationMemberEntity,
      TypeOrmReadReceiptEntity,
    ]),
    UsersModule,
  ],
  controllers: [ConversationsController],
  providers: [
    ConversationsService,
    {
      provide: CONVERSATION_REPOSITORY,
      useClass: TypeOrmConversationRepository,
    },
  ],
  exports: [ConversationsService],
})
export class ConversationsModule {}
