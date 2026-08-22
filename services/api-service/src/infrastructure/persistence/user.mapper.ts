import { User } from 'src/domain/entities/user.entity';
import { TypeOrmUserEntity } from './typeorm-user.entity';

export class UserMapper {
  static toDomain(entity: TypeOrmUserEntity): User {
    return new User({
      id: entity.id,
      username: entity.username,
      passwordHash: entity.passwordHash,
      name: entity.name,
      phone: entity.phone,
      address: entity.address,
      createdAt: entity.createdAt,
      updatedAt: entity.updatedAt,
      deletedAt: entity.deletedAt,
    });
  }

  static toPersistence(domain: Partial<User>): Partial<TypeOrmUserEntity> {
    return {
      id: domain.id,
      username: domain.username,
      passwordHash: domain.passwordHash,
      name: domain.name,
      phone: domain.phone,
      address: domain.address,
      deletedAt: domain.deletedAt,
    };
  }
}
