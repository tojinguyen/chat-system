import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { IFriendRepository } from 'src/domain/repositories/friend.repository.interface';
import { Friend, FriendStatus } from 'src/domain/entities/friend.entity';
import { TypeOrmFriendEntity } from './typeorm-friend.entity';
import { FriendMapper } from './friend.mapper';

@Injectable()
export class TypeOrmFriendRepository implements IFriendRepository {
  constructor(
    @InjectRepository(TypeOrmFriendEntity)
    private readonly repo: Repository<TypeOrmFriendEntity>,
  ) {}

  async findById(id: string): Promise<Friend | null> {
    const raw = await this.repo.findOne({
      where: { id },
      relations: {
        requester: true,
        addressee: true,
      },
    });
    if (!raw) return null;
    return FriendMapper.toDomain(raw);
  }

  async findByUsers(userA: string, userB: string): Promise<Friend | null> {
    const raw = await this.repo.findOne({
      where: [
        { requesterId: userA, addresseeId: userB },
        { requesterId: userB, addresseeId: userA },
      ],
      relations: {
        requester: true,
        addressee: true,
      },
    });
    if (!raw) return null;
    return FriendMapper.toDomain(raw);
  }

  async findPendingRequests(userId: string): Promise<Friend[]> {
    const raws = await this.repo.find({
      where: {
        addresseeId: userId,
        status: FriendStatus.PENDING,
      },
      relations: {
        requester: true,
      },
      order: { createdAt: 'DESC' },
    });
    return raws.map((raw) => FriendMapper.toDomain(raw));
  }

  async findFriendsList(userId: string): Promise<Friend[]> {
    const raws = await this.repo.find({
      where: [
        { requesterId: userId, status: FriendStatus.ACCEPTED },
        { addresseeId: userId, status: FriendStatus.ACCEPTED },
      ],
      relations: {
        requester: true,
        addressee: true,
      },
      order: { updatedAt: 'DESC' },
    });
    return raws.map((raw) => FriendMapper.toDomain(raw));
  }

  async create(
    friendData: Omit<Friend, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<Friend> {
    const entity = this.repo.create(FriendMapper.toPersistence(friendData));
    const saved = await this.repo.save(entity);
    return this.findById(saved.id) as Promise<Friend>;
  }

  async updateStatus(id: string, status: FriendStatus): Promise<Friend | null> {
    await this.repo.update(id, { status });
    return this.findById(id);
  }

  async delete(id: string): Promise<boolean> {
    const res = await this.repo.delete(id);
    return (res.affected ?? 0) > 0;
  }
}
