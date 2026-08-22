import { User } from 'src/domain/entities/user.entity';
import { TypeOrmUserEntity } from './typeorm-user.entity';

export class UserMapper {
  static toDomain(raw: TypeOrmUserEntity): User {
    return new User({
      id: raw.id,
      username: raw.username,
      passwordHash: raw.passwordHash,
      name: raw.name,
      phone: raw.phone,
      address: raw.address,
      createdAt: raw.createdAt,
      updatedAt: raw.updatedAt,
    });
  }

  static toPersistence(user: Partial<User>): Partial<TypeOrmUserEntity> {
    const entity = new TypeOrmUserEntity();
    if (user.id) entity.id = user.id;
    if (user.username) entity.username = user.username;
    if (user.passwordHash) entity.passwordHash = user.passwordHash;
    if (user.name) entity.name = user.name;
    if (user.phone !== undefined) entity.phone = user.phone;
    if (user.address !== undefined) entity.address = user.address;
    return entity;
  }
}
