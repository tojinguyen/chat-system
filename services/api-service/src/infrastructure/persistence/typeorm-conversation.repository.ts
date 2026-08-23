import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import { IConversationRepository } from 'src/domain/repositories/conversation.repository.interface';
import {
  Conversation,
  ConversationType,
} from 'src/domain/entities/conversation.entity';
import {
  ConversationMember,
  MemberRole,
} from 'src/domain/entities/conversation-member.entity';
import { ReadReceipt } from 'src/domain/entities/read-receipt.entity';
import { TypeOrmConversationEntity } from './typeorm-conversation.entity';
import { TypeOrmConversationMemberEntity } from './typeorm-conversation-member.entity';
import { TypeOrmReadReceiptEntity } from './typeorm-read-receipt.entity';
import { ConversationMapper } from './conversation.mapper';

@Injectable()
export class TypeOrmConversationRepository implements IConversationRepository {
  constructor(
    @InjectRepository(TypeOrmConversationEntity)
    private readonly convoRepo: Repository<TypeOrmConversationEntity>,
    @InjectRepository(TypeOrmConversationMemberEntity)
    private readonly memberRepo: Repository<TypeOrmConversationMemberEntity>,
    @InjectRepository(TypeOrmReadReceiptEntity)
    private readonly receiptRepo: Repository<TypeOrmReadReceiptEntity>,
    private readonly dataSource: DataSource,
  ) {}

  async findById(id: string): Promise<Conversation | null> {
    const raw = await this.convoRepo.findOne({
      where: { id },
      relations: {
        members: {
          user: true,
        },
        readReceipts: {
          user: true,
        },
      },
    });
    if (!raw) return null;
    return ConversationMapper.toDomain(raw);
  }

  async findDirectConversation(
    userA: string,
    userB: string,
  ): Promise<Conversation | null> {
    const result = await this.convoRepo
      .createQueryBuilder('c')
      .innerJoin('c.members', 'm1', 'm1.userId = :userA', { userA })
      .innerJoin('c.members', 'm2', 'm2.userId = :userB', { userB })
      .where('c.type = :type', { type: ConversationType.DIRECT })
      .getOne();

    if (!result) return null;
    return this.findById(result.id);
  }

  async findUserConversations(userId: string): Promise<Conversation[]> {
    const subQuery = this.memberRepo
      .createQueryBuilder('m')
      .select('m.conversationId')
      .where('m.userId = :userId', { userId });

    const rawConvos = await this.convoRepo
      .createQueryBuilder('c')
      .where(`c.id IN (${subQuery.getQuery()})`)
      .setParameters(subQuery.getParameters())
      .leftJoinAndSelect('c.members', 'members')
      .leftJoinAndSelect('members.user', 'user')
      .leftJoinAndSelect('c.readReceipts', 'receipts')
      .orderBy('c.updatedAt', 'DESC')
      .getMany();

    return rawConvos.map((c) => ConversationMapper.toDomain(c));
  }

  async create(
    conversationData: Omit<
      Conversation,
      'id' | 'createdAt' | 'updatedAt' | 'members' | 'readReceipts'
    >,
    memberUserIds: { userId: string; role: MemberRole }[],
  ): Promise<Conversation> {
    const queryRunner = this.dataSource.createQueryRunner();
    await queryRunner.connect();
    await queryRunner.startTransaction();

    try {
      const convoEntity = queryRunner.manager.create(
        TypeOrmConversationEntity,
        ConversationMapper.toPersistence(conversationData),
      );
      const savedConvo = await queryRunner.manager.save(convoEntity);

      const memberEntities = memberUserIds.map((m) =>
        queryRunner.manager.create(TypeOrmConversationMemberEntity, {
          conversationId: savedConvo.id,
          userId: m.userId,
          role: m.role,
        }),
      );
      await queryRunner.manager.save(memberEntities);

      await queryRunner.commitTransaction();
      return (await this.findById(savedConvo.id)) as Conversation;
    } catch (err) {
      await queryRunner.rollbackTransaction();
      throw err;
    } finally {
      await queryRunner.release();
    }
  }

  async addMember(
    conversationId: string,
    userId: string,
    role: MemberRole = MemberRole.MEMBER,
  ): Promise<ConversationMember> {
    const existing = await this.memberRepo.findOne({
      where: { conversationId, userId },
      relations: { user: true },
    });
    if (existing) {
      return ConversationMapper.toDomainMember(existing);
    }

    const newMember = this.memberRepo.create({
      conversationId,
      userId,
      role,
    });
    const saved = await this.memberRepo.save(newMember);
    const full = await this.memberRepo.findOne({
      where: { id: saved.id },
      relations: { user: true },
    });
    return ConversationMapper.toDomainMember(full!);
  }

  async removeMember(conversationId: string, userId: string): Promise<boolean> {
    const res = await this.memberRepo.delete({ conversationId, userId });
    return (res.affected ?? 0) > 0;
  }

  async findMember(
    conversationId: string,
    userId: string,
  ): Promise<ConversationMember | null> {
    const raw = await this.memberRepo.findOne({
      where: { conversationId, userId },
      relations: { user: true },
    });
    if (!raw) return null;
    return ConversationMapper.toDomainMember(raw);
  }

  async upsertReadReceipt(
    conversationId: string,
    userId: string,
    lastMessageSeen: string,
  ): Promise<ReadReceipt> {
    const existing = await this.receiptRepo.findOne({
      where: { conversationId, userId },
    });

    if (existing) {
      existing.lastMessageSeen = lastMessageSeen;
      existing.lastReadAt = new Date();
      const saved = await this.receiptRepo.save(existing);
      return ConversationMapper.toDomainReceipt(saved);
    } else {
      const created = this.receiptRepo.create({
        conversationId,
        userId,
        lastMessageSeen,
        lastReadAt: new Date(),
      });
      const saved = await this.receiptRepo.save(created);
      return ConversationMapper.toDomainReceipt(saved);
    }
  }

  async getReadReceipt(
    conversationId: string,
    userId: string,
  ): Promise<ReadReceipt | null> {
    const raw = await this.receiptRepo.findOne({
      where: { conversationId, userId },
      relations: { user: true },
    });
    if (!raw) return null;
    return ConversationMapper.toDomainReceipt(raw);
  }
}
