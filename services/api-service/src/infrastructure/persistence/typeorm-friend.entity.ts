import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  ManyToOne,
  JoinColumn,
  Index,
  Unique,
} from 'typeorm';
import { TypeOrmUserEntity } from './typeorm-user.entity';
import { FriendStatus } from 'src/domain/entities/friend.entity';

@Entity('friends')
@Unique(['requesterId', 'addresseeId'])
export class TypeOrmFriendEntity {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Index()
  @Column({ type: 'uuid' })
  requesterId: string;

  @ManyToOne(() => TypeOrmUserEntity, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'requesterId' })
  requester?: TypeOrmUserEntity;

  @Index()
  @Column({ type: 'uuid' })
  addresseeId: string;

  @ManyToOne(() => TypeOrmUserEntity, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'addresseeId' })
  addressee?: TypeOrmUserEntity;

  @Column({
    type: 'enum',
    enum: FriendStatus,
    default: FriendStatus.PENDING,
  })
  status: FriendStatus;

  @CreateDateColumn({ type: 'timestamp with time zone' })
  createdAt: Date;

  @UpdateDateColumn({ type: 'timestamp with time zone' })
  updatedAt: Date;
}
