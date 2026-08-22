import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, MinLength } from 'class-validator';

export class RegisterDto {
  @ApiProperty({ example: 'john_doe', description: 'Tên đăng nhập' })
  @IsString()
  @IsNotEmpty({ message: 'Username is required' })
  @MinLength(3, { message: 'Username must be at least 3 characters' })
  userName: string;

  @ApiProperty({ example: 'securepassword123', description: 'Mật khẩu' })
  @IsString()
  @IsNotEmpty({ message: 'Password is required' })
  @MinLength(6, { message: 'Password must be at least 6 characters' })
  password: string;

  @ApiProperty({ example: 'John Doe', description: 'Tên hiển thị' })
  @IsString()
  @IsNotEmpty({ message: 'Name is required' })
  name: string;

  @ApiProperty({
    example: '0123456789',
    description: 'Số điện thoại',
    required: false,
  })
  @IsString()
  @IsOptional()
  phone?: string;

  @ApiProperty({
    example: 'Hồ Chí Minh City',
    description: 'Địa chỉ',
    required: false,
  })
  @IsString()
  @IsOptional()
  address?: string;
}
