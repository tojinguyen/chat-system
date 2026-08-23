import { ApiPropertyOptional } from '@nestjs/swagger';
import { IsOptional, IsString, Length, Matches } from 'class-validator';

export class UpdateUserDto {
  @ApiPropertyOptional({ example: 'John Doe', description: 'Họ và tên' })
  @IsOptional()
  @IsString()
  @Length(2, 100)
  name?: string;

  @ApiPropertyOptional({
    example: '+84901234567',
    description: 'Số điện thoại',
  })
  @IsOptional()
  @IsString()
  @Matches(/^[+0-9]{8,20}$/, { message: 'Số điện thoại không hợp lệ' })
  phone?: string;

  @ApiPropertyOptional({ example: 'Ho Chi Minh City', description: 'Địa chỉ' })
  @IsOptional()
  @IsString()
  @Length(2, 255)
  address?: string;
}
