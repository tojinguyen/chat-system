import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsUUID } from 'class-validator';

export class CreateDirectConversationDto {
  @ApiProperty({
    example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
    description: 'ID của người muốn chat cùng',
  })
  @IsNotEmpty({ message: 'partnerId không được để trống' })
  @IsUUID('4', { message: 'partnerId phải là định dạng UUID' })
  partnerId: string;
}
