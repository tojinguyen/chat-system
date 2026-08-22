import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import {
  ArrayMinSize,
  IsArray,
  IsEnum,
  IsNotEmpty,
  IsOptional,
  IsString,
  IsUUID,
  Length,
} from 'class-validator';
import { ConversationType } from 'src/domain/entities/conversation.entity';

export class CreateGroupConversationDto {
  @ApiProperty({ example: 'Backend Team Chat', description: 'Tên nhóm' })
  @IsNotEmpty({ message: 'Tên nhóm không được để trống' })
  @IsString()
  @Length(2, 100)
  name: string;

  @ApiPropertyOptional({
    enum: ConversationType,
    default: ConversationType.GROUP,
    description: 'Loại hội thoại: group, clan, region, global',
  })
  @IsOptional()
  @IsEnum(ConversationType)
  type?: ConversationType;

  @ApiPropertyOptional({
    example: 'https://example.com/group-icon.png',
    description: 'Ảnh đại diện nhóm',
  })
  @IsOptional()
  @IsString()
  iconUrl?: string;

  @ApiProperty({
    example: ['a1b2c3d4-e5f6-7890-abcd-ef1234567890'],
    description: 'Danh sách User IDs của các thành viên ban đầu',
  })
  @IsArray()
  @ArrayMinSize(1, { message: 'Nhóm phải có ít nhất 1 thành viên khác' })
  @IsUUID('4', { each: true, message: 'Mỗi memberId phải là UUID hợp lệ' })
  memberIds: string[];
}
