namespace Explorer.Blog.API.Dtos;

public class FollowRequestDto
{
    public long personId { get; set; }
    public long followerId { get; set; }

    public FollowRequestDto(long personId, long followerId)
    {
        this.personId = personId;
        this.followerId = followerId;
    }
}