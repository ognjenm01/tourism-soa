﻿using Explorer.API.Controllers.Tourist.Blog;
using Explorer.Blog.API.Dtos;
using Explorer.Blog.API.Public.Blog;
using Explorer.BuildingBlocks.Core.UseCases;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;
using Shouldly;

namespace Explorer.Blog.Tests.Integration.Blog;

[Collection("Sequential")]
public class BlogQueryTests : BaseBlogIntegrationTest
{
    public BlogQueryTests(BlogTestFactory factory) : base(factory)
    {
    }

   /* [Fact]
    public void Retrieves_all()
    {
        //Arrange
        using var scope = Factory.Services.CreateScope();
        var controller = CreateController(scope);

        //Act
        var result = ((ObjectResult)controller.GetAll(0, 0).Result)?.Value as PagedResult<BlogDto>;

        //Assert
        result.ShouldNotBeNull();
        result.Results.Count.ShouldBe(3);
        result.TotalCount.ShouldBe(3);
    }*/

    /*private static BlogController CreateController(IServiceScope scope)
    {
        return new BlogController(scope.ServiceProvider.GetRequiredService<IBlogService>())
        {
            ControllerContext = BuildContext("-1")
        };
    }*/
}