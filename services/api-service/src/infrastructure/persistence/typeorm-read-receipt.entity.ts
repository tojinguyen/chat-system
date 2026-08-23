import {
  Entity,
  PrimaryColumn,
  Column,
  UpdateDateColumn,
  ManyToOne,
  JoinColumn,
} from 'typeorm';
import { TypeOrmConversationEntity } from './typeorm-conversation.entity';
import { TypeOrmUserEntity } from './typeorm-user.entity';

@Entity('conversation_read_receipts')
export class TypeOrmReadReceiptEntity {
  @PrimaryColumn({ type: 'uuid' })
  conversationId: string;

  @ManyToOne(
    () => TypeOrmConversationEntity,
    (conversation: TypeOrmConversationEntity) => conversation.readReceipts,
    { onDelete: 'CASCADE' },
  )
  @JoinColumn({ name: 'conversationId' })
  conversation?: TypeOrmConversationEntity;

  @PrimaryColumn({ type: 'uuid' })
  userId: string;

  @ManyToOne(() => TypeOrmUserEntity, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'userId' })
  user?: TypeOrmUserEntity;

  @Column({ type: 'varchar', length: 64 })
  lastMessageSeen: string;

  @UpdateDateColumn({ type: 'timestamp with time zone' })
  lastReadAt: Date;
}
