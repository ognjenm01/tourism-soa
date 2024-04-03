namespace Explorer.Blog.API.Dtos;

public class BlogRatingDto
{
    public long BlogId { get; set; }
    public long UserId { get; set; }
    //public string Username { get; set; }
    public DateTime CreationTime { get; set; }
    public string Rating { get; set; }
}