import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsString } from 'class-validator';

export class LoginDto {
  @ApiProperty({ example: 'john_doe', description: 'Tên đăng nhập' })
  @IsString()
  @IsNotEmpty({ message: 'Username is required' })
  userName: string;

  @ApiProperty({ example: 'securepassword123', description: 'Mật khẩu' })
  @IsString()
  @IsNotEmpty({ message: 'Password is required' })
  password: string;
}

export class LoginResponseDto {
  accessToken: string;
  refreshToken: string;
}
