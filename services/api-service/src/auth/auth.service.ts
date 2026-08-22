import { ConflictException, Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';
import { UserService } from 'src/users/users.service';
import { RegisterDto } from './dto/register.dto';
import { UserResponseDto } from 'src/users/dto/uset-response.dto';

@Injectable()
export class AuthService {
  constructor(
    private readonly userService: UserService,
    private readonly jwtService: JwtService,
  ) {}

  async register(dto: RegisterDto) {
    const existingUser = await this.userService.findByUserName(dto.userName);
    if (existingUser) {
      throw new ConflictException('Username already exists');
    }

    const saltRounds = 10;
    const passwordHash = await bcrypt.hash(dto.password, saltRounds);

    const newUser = await this.userService.create({
      userName: dto.userName,
      name: dto.name,
      phone: dto.phone,
      address: dto.address,
      passwordHash,
    });

    return new UserResponseDto(newUser);
  }
}
