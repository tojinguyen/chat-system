import { Injectable } from '@nestjs/common';
import { RegisterDto } from 'src/auth/dto/register.dto';

@Injectable()
export class UserService {
  async findByUserName(identifier: string) {
    return null;
  }

  async findByEmail(email: string) {
    return null;
  }

  async create(user: RegisterDto) {
    return null;
  }
}
