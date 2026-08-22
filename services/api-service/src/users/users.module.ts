import { Module } from '@nestjs/common';
import { UserService } from './users.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { USER_REPOSITORY } from 'src/domain/repositories/user.repository.interface';
import { TypeOrmUserRepository } from 'src/infrastructure/persistence/typeorm-user.repository';
import { TypeOrmUserEntity } from 'src/infrastructure/persistence/typeorm-user.entity';

@Module({
  imports: [TypeOrmModule.forFeature([TypeOrmUserEntity])],
  providers: [
    UserService,
    {
      provide: USER_REPOSITORY,
      useClass: TypeOrmUserRepository,
    },
  ],
  exports: [UserService],
})
export class UsersModule {}
