import { Conversation } from 'src/domain/entities/conversation.entity';
import { ConversationMember } from 'src/domain/entities/conversation-member.entity';
import { ReadReceipt } from 'src/domain/entities/read-receipt.entity';
import { TypeOrmConversationEntity } from './typeorm-conversation.entity';
import { TypeOrmConversationMemberEntity } from './typeorm-conversation-member.entity';
import { TypeOrmReadReceiptEntity } from './typeorm-read-receipt.entity';
import { UserMapper } from './user.mapper';

export class ConversationMapper {
  static toDomain(entity: TypeOrmConversationEntity): Conversation {
    return new Conversation({
      id: entity.id,
      name: entity.name,
      type: entity.type,
      iconUrl: entity.iconUrl,
      createdBy: entity.createdBy,
      createdAt: entity.createdAt,
      updatedAt: entity.updatedAt,
      members: entity.members?.map((m: TypeOrmConversationMemberEntity) =>
        ConversationMapper.toDomainMember(m),
      ),
      readReceipts: entity.readReceipts?.map((r: TypeOrmReadReceiptEntity) =>
        ConversationMapper.toDomainReceipt(r),
      ),
    });
  }

  static toDomainMember(
    entity: TypeOrmConversationMemberEntity,
  ): ConversationMember {
    return new ConversationMember({
      id: entity.id,
      conversationId: entity.conversationId,
      userId: entity.userId,
      role: entity.role,
      joinedAt: entity.joinedAt,
      user: entity.user ? UserMapper.toDomain(entity.user) : undefined,
    });
  }

  static toDomainReceipt(entity: TypeOrmReadReceiptEntity): ReadReceipt {
    return new ReadReceipt({
      conversationId: entity.conversationId,
      userId: entity.userId,
      lastMessageSeen: entity.lastMessageSeen,
      lastReadAt: entity.lastReadAt,
      user: entity.user ? UserMapper.toDomain(entity.user) : undefined,
    });
  }

  static toPersistence(
    domain: Partial<Conversation>,
  ): Partial<TypeOrmConversationEntity> {
    return {
      id: domain.id,
      name: domain.name,
      type: domain.type,
      iconUrl: domain.iconUrl,
      createdBy: domain.createdBy,
    };
  }
}
