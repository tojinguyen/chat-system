import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  ManyToOne,
  JoinColumn,
  Index,
  Unique,
} from 'typeorm';
import { MemberRole } from 'src/domain/entities/conversation-member.entity';
import { TypeOrmConversationEntity } from './typeorm-conversation.entity';
import { TypeOrmUserEntity } from './typeorm-user.entity';

@Entity('conversation_members')
@Unique(['conversationId', 'userId'])
export class TypeOrmConversationMemberEntity {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Index()
  @Column({ type: 'uuid' })
  conversationId: string;

  @ManyToOne(
    () => TypeOrmConversationEntity,
    (conversation: TypeOrmConversationEntity) => conversation.members,
    { onDelete: 'CASCADE' },
  )
  @JoinColumn({ name: 'conversationId' })
  conversation?: TypeOrmConversationEntity;

  @Index()
  @Column({ type: 'uuid' })
  userId: string;

  @ManyToOne(() => TypeOrmUserEntity, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'userId' })
  user?: TypeOrmUserEntity;

  @Column({
    type: 'enum',
    enum: MemberRole,
    default: MemberRole.MEMBER,
  })
  role: MemberRole;

  @CreateDateColumn({ type: 'timestamp with time zone' })
  joinedAt: Date;
}
