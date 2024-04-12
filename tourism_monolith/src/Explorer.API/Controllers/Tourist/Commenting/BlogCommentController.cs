using Explorer.Blog.API.Dtos;
using Explorer.Blog.API.Public.Commenting;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.Infrastructure.Authentication;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using System.Net.Http.Headers;

namespace Explorer.API.Controllers.Tourist.Commenting;

[Authorize(Policy = "touristPolicy")]
[Route("api/blog/comments")]
public class BlogCommentController : BaseApiController
{
    private static readonly string _blogServicePort = Environment.GetEnvironmentVariable("BLOGS_APP_PORT") ?? "8080";
    private readonly HttpClient _httpClient;

    public BlogCommentController()
    {
        _httpClient = new()
        {
            BaseAddress = new Uri($"http://blogs-module:" + _blogServicePort + "/blogcomments")
        };
    }

    [HttpGet]
     public async Task<ActionResult<PagedResult<BlogCommentDto>>> GetAll([FromQuery] int page, [FromQuery] int pageSize, [FromQuery] long blogId)
    {
        //Uri uri = new Uri($"http://localhost:3333/api/blogcomments");
        HttpResponseMessage response = await _httpClient.GetAsync(".");

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            List<BlogCommentDto> blogComments = JsonConvert.DeserializeObject<List<BlogCommentDto>>(responseContent);
            
            var pagedResult = new PagedResult<BlogCommentDto>(blogComments, totalCount: blogComments.Count);

            return CreateResponse(Result.Ok(pagedResult));
        }

        return BadRequest();
    }

    [HttpGet("{id:int}")]
     public async Task<ActionResult<BlogCommentDto>> Get(int id)
    {
        //Uri uri = new Uri($"http://localhost:3333/api/blogcomments/{id}");
        HttpResponseMessage response = await _httpClient.GetAsync($"/{id}");

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            BlogCommentDto blogComment = JsonConvert.DeserializeObject<BlogCommentDto>(responseContent);
            return CreateResponse(Result.Ok(blogComment));
        }

        return BadRequest();
    }  
    

    [HttpPost]
     public async Task<ActionResult<BlogCommentDto>> Create([FromBody] BlogCommentDto blogComment)
    {
        string jsonBlogComment = JsonConvert.SerializeObject(blogComment);

        //Uri uri = new Uri("http://localhost:3333/api/blogcomments");

        var content = new StringContent(jsonBlogComment);
        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
        HttpResponseMessage response = await _httpClient.PostAsync(".", content);

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            
            BlogCommentDto createdBlogComment = JsonConvert.DeserializeObject<BlogCommentDto>(responseContent);

            return CreateResponse(Result.Ok(createdBlogComment));
        }
        else
        {
            return BadRequest();
        }
    }

    [HttpPut("{id:int}")]
     public async Task<ActionResult<BlogCommentDto>> Update([FromBody] BlogCommentDto blogComment)
    {
        string jsonBlogComment = JsonConvert.SerializeObject(blogComment);
        
        //Uri uri = new Uri("http://localhost:3333/api/blogcomments/{id}");
        var content = new StringContent(jsonBlogComment);
        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
        
        HttpResponseMessage responseMessage = await _httpClient.PutAsync(".", content);
        if (responseMessage.IsSuccessStatusCode)
        {
            string responseContent = await responseMessage.Content.ReadAsStringAsync();
            BlogCommentDto updatedBlogComment = JsonConvert.DeserializeObject<BlogCommentDto>(responseContent);
            return CreateResponse(Result.Ok(updatedBlogComment));
        }

        return BadRequest();
    }

    [HttpDelete("{id:int}")]
     public async Task<ActionResult> Delete(int id)
    {
        //Uri uri = new Uri($"http://localhost:3333/api/blogcomments/{id}");
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
}