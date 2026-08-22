import { Inject, Injectable } from '@nestjs/common';
import { User } from '../domain/entities/user.entity';
import {
  USER_REPOSITORY,
  type IUserRepository,
} from 'src/domain/repositories/user.repository.interface';

@Injectable()
export class UserService {
  constructor(
    @Inject(USER_REPOSITORY)
    private readonly userRepository: IUserRepository,
  ) {}

  async findByUserName(username: string): Promise<User | null> {
    return await this.userRepository.findByUsername(username);
  }

  async create(
    userData: Omit<User, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<User> {
    return await this.userRepository.create(userData);
  }
}
