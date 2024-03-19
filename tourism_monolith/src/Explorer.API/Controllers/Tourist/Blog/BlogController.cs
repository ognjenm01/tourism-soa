﻿using System.Net.Http.Headers;
using AutoMapper;
using Explorer.Blog.API.Dtos;
using Explorer.Blog.API.Public.Blog;
using Explorer.BuildingBlocks.Core.UseCases;
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
    private readonly IBlogService _blogService;
    private readonly HttpClient _httpClient;

    public BlogController(IBlogService blogService)
    {
        _blogService = blogService;
        _httpClient = new HttpClient();
    }

    [HttpGet]
    public async Task<ActionResult<PagedResult<BlogDto>>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        Uri uri = new Uri($"http://localhost:8080/blogs");
        HttpResponseMessage response = await _httpClient.GetAsync(uri);

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            List<BlogDto> blogs = JsonConvert.DeserializeObject<List<BlogDto>>(responseContent);
            
            var pagedResult = new PagedResult<BlogDto>(blogs, totalCount: blogs.Count);

            return CreateResponse(Result.Ok(pagedResult));
        }

        return BadRequest();
    }

    [HttpGet("status")]
    public ActionResult<PagedResult<BlogDto>> GetWithStatuses([FromQuery] int page, [FromQuery] int pageSize)
    {
        var result = _blogService.GetWithStatuses(page, pageSize);
        return CreateResponse(result.ToResult());
    }

    [HttpGet("{id:int}")]
    public async Task<ActionResult<BlogDto>> Get(int id)
    {
        Uri uri = new Uri($"http://localhost:8080/blogs/{id}");
        HttpResponseMessage response = await _httpClient.GetAsync(uri);

        if (response.IsSuccessStatusCode)
        {
            string responseContent = await response.Content.ReadAsStringAsync();
            BlogDto blog = JsonConvert.DeserializeObject<BlogDto>(responseContent);
            return CreateResponse(Result.Ok(blog));
        }

        return BadRequest();
    }

    [HttpPost]
    public async Task<ActionResult<BlogDto>> Create([FromBody] BlogDto blog)
    {
        string jsonBlog = JsonConvert.SerializeObject(blog);

        Uri uri = new Uri("http://localhost:8080/blogs");

        var content = new StringContent(jsonBlog);
        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
        HttpResponseMessage response = await _httpClient.PostAsync(uri, content);

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
    public ActionResult<BlogDto> Update([FromBody] BlogDto blog)
    {
        var result = _blogService.Update(blog);
        return CreateResponse(result);
    }

    [HttpDelete("{id:int}")]
    public ActionResult Delete(int id)
    {
        var result = _blogService.Delete(id);
        return CreateResponse(result);
    }

    [HttpPost("rate")]
    public ActionResult<BlogDto> AddRating([FromBody] BlogRatingDto blogRatingDto)
    {
        var result = _blogService.AddRating(blogRatingDto, User.PersonId());
        return CreateResponse(result);
    }
}