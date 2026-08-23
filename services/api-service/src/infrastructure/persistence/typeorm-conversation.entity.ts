import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  OneToMany,
  ManyToOne,
  JoinColumn,
} from 'typeorm';
import { ConversationType } from 'src/domain/entities/conversation.entity';
import { TypeOrmConversationMemberEntity } from './typeorm-conversation-member.entity';
import { TypeOrmReadReceiptEntity } from './typeorm-read-receipt.entity';
import { TypeOrmUserEntity } from './typeorm-user.entity';

@Entity('conversations')
export class TypeOrmConversationEntity {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column({ nullable: true, length: 100 })
  name?: string;

  @Column({
    type: 'enum',
    enum: ConversationType,
    default: ConversationType.DIRECT,
  })
  type: ConversationType;

  @Column({ nullable: true })
  iconUrl?: string;

  @Column({ type: 'uuid', nullable: true })
  createdBy?: string;

  @ManyToOne(() => TypeOrmUserEntity, { nullable: true, onDelete: 'SET NULL' })
  @JoinColumn({ name: 'createdBy' })
  creator?: TypeOrmUserEntity;

  @CreateDateColumn({ type: 'timestamp with time zone' })
  createdAt: Date;

  @UpdateDateColumn({ type: 'timestamp with time zone' })
  updatedAt: Date;

  @OneToMany(
    () => TypeOrmConversationMemberEntity,
    (member: TypeOrmConversationMemberEntity) => member.conversation,
    { cascade: true },
  )
  members: TypeOrmConversationMemberEntity[];

  @OneToMany(
    () => TypeOrmReadReceiptEntity,
    (receipt: TypeOrmReadReceiptEntity) => receipt.conversation,
  )
  readReceipts: TypeOrmReadReceiptEntity[];
}
