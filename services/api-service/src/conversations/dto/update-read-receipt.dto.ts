import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsString } from 'class-validator';

export class UpdateReadReceiptDto {
  @ApiProperty({
    example: '018dc5e1-cb00-7adb-6f4f-8d556994111c',
    description: 'ID của tin nhắn cuối cùng đã đọc (UUIDv7/Snowflake)',
  })
  @IsNotEmpty({ message: 'lastMessageSeen không được để trống' })
  @IsString()
  lastMessageSeen: string;
}
