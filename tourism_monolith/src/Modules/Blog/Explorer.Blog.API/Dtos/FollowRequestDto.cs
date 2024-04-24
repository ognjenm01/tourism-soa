namespace Explorer.Blog.API.Dtos;

public class FollowRequestDto
{
    public long PersonId { get; set; }
    public long FollowerId { get; set; }

    public FollowRequestDto(long personId, long followerId)
    {
        PersonId = personId;
        FollowerId = followerId;
    }
}