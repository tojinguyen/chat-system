import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { Exclude } from 'class-transformer';
import { User } from 'src/domain/entities/user.entity';

export class UserResponseDto {
  @ApiProperty({ example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890' })
  id: string;

  @ApiProperty({ example: 'john_doe' })
  username: string;

  @ApiProperty({ example: 'John Doe' })
  name: string;

  @ApiPropertyOptional({ example: '+84901234567' })
  phone?: string;

  @ApiPropertyOptional({ example: 'Ho Chi Minh City' })
  address?: string;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  createdAt: Date;

  @ApiProperty({ example: '2026-08-22T08:00:00.000Z' })
  updatedAt: Date;

  @Exclude()
  passwordHash: string;

  constructor(user: Partial<User>) {
    this.id = user.id!;
    this.username = user.username!;
    this.name = user.name!;
    this.phone = user.phone;
    this.address = user.address;
    this.createdAt = user.createdAt!;
    this.updatedAt = user.updatedAt!;
  }
}
