import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { FriendsController } from './friends.controller';
import { FriendsService } from './friends.service';
import { TypeOrmFriendEntity } from 'src/infrastructure/persistence/typeorm-friend.entity';
import { FRIEND_REPOSITORY } from 'src/domain/repositories/friend.repository.interface';
import { TypeOrmFriendRepository } from 'src/infrastructure/persistence/typeorm-friend.repository';
import { UsersModule } from 'src/users/users.module';

@Module({
  imports: [TypeOrmModule.forFeature([TypeOrmFriendEntity]), UsersModule],
  controllers: [FriendsController],
  providers: [
    FriendsService,
    {
      provide: FRIEND_REPOSITORY,
      useClass: TypeOrmFriendRepository,
    },
  ],
  exports: [FriendsService],
})
export class FriendsModule {}
