import {
  BadRequestException,
  Inject,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { User } from '../domain/entities/user.entity';
import {
  USER_REPOSITORY,
  type IUserRepository,
} from 'src/domain/repositories/user.repository.interface';
import { UpdateUserDto } from './dto/update-user.dto';

@Injectable()
export class UserService {
  constructor(
    @Inject(USER_REPOSITORY)
    private readonly userRepository: IUserRepository,
  ) {}

  async findById(id: string): Promise<User | null> {
    return await this.userRepository.findById(id);
  }

  async findByUserName(username: string): Promise<User | null> {
    return await this.userRepository.findByUsername(username);
  }

  async searchUsers(keyword: string, limit = 20): Promise<User[]> {
    if (!keyword || keyword.trim() === '') {
      return [];
    }
    return await this.userRepository.searchUsers(keyword.trim(), limit);
  }

  async create(
    userData: Omit<User, 'id' | 'createdAt' | 'updatedAt'>,
  ): Promise<User> {
    return await this.userRepository.create(userData);
  }

  async updateProfile(userId: string, dto: UpdateUserDto): Promise<User> {
    const existing = await this.userRepository.findById(userId);
    if (!existing) {
      throw new NotFoundException('Không tìm thấy người dùng');
    }

    if (dto.phone && dto.phone !== existing.phone) {
      const phoneUser = await this.userRepository.findByPhone(dto.phone);
      if (phoneUser && phoneUser.id !== userId) {
        throw new BadRequestException('Số điện thoại đã được sử dụng');
      }
    }

    const updated = await this.userRepository.update(userId, dto);
    return updated!;
  }
}
