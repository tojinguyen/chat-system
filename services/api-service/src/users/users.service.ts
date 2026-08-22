import { Injectable } from '@nestjs/common';
import { User } from './entities/user.entity';

@Injectable()
export class UserService {
  async findByUserName(_identifier: string): Promise<User | null> {
    return Promise.resolve(null);
  }

  async create(
    userData: Omit<User, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<User> {
    return Promise.resolve({
      ...userData,
      id: 'mock-id',
      createdAt: new Date(),
      updatedAt: new Date(),
    });
  }
}
