import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
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

  async findByUsername(username: string): Promise<User | null> {
    const raw = await this.repo.findOne({ where: { username } });
    if (!raw) return null;
    return UserMapper.toDomain(raw);
  }

  async create(
    userData: Omit<User, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<User> {
    const entity = this.repo.create(UserMapper.toPersistence(userData));
    const saved = await this.repo.save(entity);
    return UserMapper.toDomain(saved);
  }
}
