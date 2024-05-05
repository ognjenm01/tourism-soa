using System.Net.Http.Headers;
using System.Text;
using Explorer.Blog.API.Dtos;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.Core.Domain;
using Explorer.Stakeholders.Infrastructure.Authentication;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace Explorer.API.Controllers.Tourist.Blog;

[Authorize(Policy = "touristPolicy")]
[Route("api/blog")]
public class BlogController : BaseApiController
{
    private static readonly string _blogServicePort = Environment.GetEnvironmentVariable("BLOGS_APP_PORT") ?? "49155";
    private readonly HttpClient _httpClient;
    private readonly HttpClient _followerHttpClient;
    public BlogController()
    {
        _httpClient = new()
        {   
            BaseAddress = new Uri("http://localhost:" + _blogServicePort + "/blogs")
            
            //BaseAddress = new Uri("http://blogs-module:" + _blogServicePort + "api/blogs")
        };

        _followerHttpClient = new()
        {
            BaseAddress = new Uri("http://localhost:4201/api/")
        };
    }

    [HttpGet]
    public async Task<ActionResult<PagedResult<BlogDto>>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        //Uri uri = new Uri($"http://localhost:8080/blogs");
        HttpResponseMessage response = await _httpClient.GetAsync(".");

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            List<BlogDto> blogs = JsonConvert.DeserializeObject<List<BlogDto>>(responseContent);
            
            var pagedResult = new PagedResult<BlogDto>(blogs, totalCount: blogs.Count);

            return CreateResponse(Result.Ok(pagedResult));
        }

        return BadRequest();
    }

    [HttpGet("{id:int}")]
    public async Task<ActionResult<BlogDto>> Get(int id)
    {
        //Uri uri = new Uri($"http://localhost:8080/blogs/{id}");
        HttpResponseMessage response = await _httpClient.GetAsync($"/{id}");

        //Ognjen debug
       // HttpResponseMessage response = await _httpClient.GetAsync($"blogs/{id}");

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            BlogDto blog = JsonConvert.DeserializeObject<BlogDto>(responseContent);
            
            //TODO Prebaci ovo u go kasnije
            FollowRequestDto req = new FollowRequestDto(ClaimsPrincipalExtensions.PersonId(User),blog.CreatorId);
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(req), Encoding.UTF8, "application/json");
            HttpResponseMessage followerResponse = await _followerHttpClient.PostAsJsonAsync("person/follow/debug", req);
            if(followerResponse.IsSuccessStatusCode)
                return CreateResponse(Result.Ok(blog));
            else
                return CreateResponse(Result.Fail("Not following that user!"));
        }

        return BadRequest();
    }

    [HttpPost]
    public async Task<ActionResult<BlogDto>> Create([FromBody] BlogDto blog)
    {
        string jsonBlog = JsonConvert.SerializeObject(blog);

       // Uri uri = new Uri("http://localhost:8080/blogs");

        var content = new StringContent(jsonBlog);
        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
        HttpResponseMessage response = await _httpClient.PostAsync(".", content);

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            
            BlogDto createdBlog = JsonConvert.DeserializeObject<BlogDto>(responseContent);

            return CreateResponse(Result.Ok(createdBlog));
        }
        else
        {
            return BadRequest();
        }
    }

    [HttpPut("{id:int}")]
    public async Task<ActionResult<BlogDto>> Update([FromBody] BlogDto blog)
    {
        string jsonBlog = JsonConvert.SerializeObject(blog);
        
        //Uri uri = new Uri("http://localhost:8080/blogs");
        var content = new StringContent(jsonBlog);
        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
        
        HttpResponseMessage responseMessage = await _httpClient.PutAsync(".",content);
        if (responseMessage.IsSuccessStatusCode)
        {
            string responseContent = await responseMessage.Content.ReadAsStringAsync();
            BlogDto updatedBlog = JsonConvert.DeserializeObject<BlogDto>(responseContent);
            return CreateResponse(Result.Ok(updatedBlog));
        }

        return BadRequest();
    }

    [HttpDelete("{id:int}")]
    public async Task<ActionResult> Delete(int id)
    {
        //Uri uri = new Uri($"http://localhost:8080/blogs/{id}");
        HttpResponseMessage response = await _httpClient.DeleteAsync($"/{id}");

        if (response.IsSuccessStatusCode)
        {
            return Ok();
        }
        else
        {
            return BadRequest();
        }
    }

    [HttpPost("rate")]
    public async Task<ActionResult<BlogDto>> AddRating([FromBody] BlogRatingDto blogRatingDto)
    {
        string jsonRating = JsonConvert.SerializeObject(blogRatingDto);
        //Uri uri = new Uri("http://localhost:8080/blogs/rate");
        
        var content = new StringContent(jsonRating);
        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
        HttpResponseMessage response = await _httpClient.PostAsync("/rate", content);

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            
            BlogDto blog = JsonConvert.DeserializeObject<BlogDto>(responseContent);

            return CreateResponse(Result.Ok(blog));
        }
        
        return BadRequest();
        
    }
}