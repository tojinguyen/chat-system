import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { ILike, Repository } from 'typeorm';
import { IUserRepository } from 'src/domain/repositories/user.repository.interface';
import { User } from 'src/domain/entities/user.entity';
import { TypeOrmUserEntity } from './typeorm-user.entity';
import { UserMapper } from './user.mapper';

@Injectable()
export class TypeOrmUserRepository implements IUserRepository {
  constructor(
    @InjectRepository(TypeOrmUserEntity)
    private readonly repo: Repository<TypeOrmUserEntity>,
  ) {}

  async findById(id: string): Promise<User | null> {
    const raw = await this.repo.findOne({ where: { id } });
    if (!raw) return null;
    return UserMapper.toDomain(raw);
  }

  async findByUsername(username: string): Promise<User | null> {
    const raw = await this.repo.findOne({ where: { username } });
    if (!raw) return null;
    return UserMapper.toDomain(raw);
  }

  async findByPhone(phone: string): Promise<User | null> {
    const raw = await this.repo.findOne({ where: { phone } });
    if (!raw) return null;
    return UserMapper.toDomain(raw);
  }

  async searchUsers(keyword: string, limit = 20): Promise<User[]> {
    const raws = await this.repo.find({
      where: [
        { username: ILike(`%${keyword}%`) },
        { name: ILike(`%${keyword}%`) },
        { phone: ILike(`%${keyword}%`) },
      ],
      take: limit,
      order: { name: 'ASC' },
    });
    return raws.map((raw) => UserMapper.toDomain(raw));
  }

  async create(
    userData: Omit<User, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<User> {
    const entity = this.repo.create(UserMapper.toPersistence(userData));
    const saved = await this.repo.save(entity);
    return UserMapper.toDomain(saved);
  }

  async update(id: string, partial: Partial<User>): Promise<User | null> {
    await this.repo.update(id, UserMapper.toPersistence(partial));
    return this.findById(id);
  }
}
