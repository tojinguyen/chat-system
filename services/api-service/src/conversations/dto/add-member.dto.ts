import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsEnum, IsNotEmpty, IsOptional, IsUUID } from 'class-validator';
import { MemberRole } from 'src/domain/entities/conversation-member.entity';

export class AddMemberDto {
  @ApiProperty({
    example: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
    description: 'User ID của thành viên cần thêm',
  })
  @IsNotEmpty()
  @IsUUID('4')
  userId: string;

  @ApiPropertyOptional({
    enum: MemberRole,
    default: MemberRole.MEMBER,
    description: 'Vai trò trong nhóm: admin hoặc member',
  })
  @IsOptional()
  @IsEnum(MemberRole)
  role?: MemberRole;
}
