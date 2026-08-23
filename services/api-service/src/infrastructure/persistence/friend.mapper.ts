import { Friend } from 'src/domain/entities/friend.entity';
import { TypeOrmFriendEntity } from './typeorm-friend.entity';
import { UserMapper } from './user.mapper';

export class FriendMapper {
  static toDomain(entity: TypeOrmFriendEntity): Friend {
    return new Friend({
      id: entity.id,
      requesterId: entity.requesterId,
      addresseeId: entity.addresseeId,
      status: entity.status,
      createdAt: entity.createdAt,
      updatedAt: entity.updatedAt,
      requester: entity.requester
        ? UserMapper.toDomain(entity.requester)
        : undefined,
      addressee: entity.addressee
        ? UserMapper.toDomain(entity.addressee)
        : undefined,
    });
  }

  static toPersistence(domain: Partial<Friend>): Partial<TypeOrmFriendEntity> {
    return {
      id: domain.id,
      requesterId: domain.requesterId,
      addresseeId: domain.addresseeId,
      status: domain.status,
    };
  }
}
