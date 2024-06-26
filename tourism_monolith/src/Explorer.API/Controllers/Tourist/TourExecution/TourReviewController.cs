﻿using System.Text;
using System.Text.Json;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourExecution;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Tourist.TourExecution;

[Authorize(Policy = "touristPolicy")]
[Route("api/tourexecution/tourreview")]
public class TourReviewController : BaseApiController
{
    private readonly ITourReviewService _tourReviewService;
    private static readonly string _tourAppPort = Environment.GetEnvironmentVariable("TOURS_APP_PORT") ?? "8080";
    private static HttpClient httpReviewClient = new()
    {
        BaseAddress = new Uri("http://tours-module:" + _tourAppPort + "/api/tourreview"),
    };
    public TourReviewController(ITourReviewService tourReviewService)
    {
        _tourReviewService = tourReviewService;
    }

    [HttpGet]
    public async Task<string> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        //var result = _tourReviewService.GetPaged(page, pageSize);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpReviewClient.GetAsync(httpReviewClient.BaseAddress);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpGet("{id:int}")]
    public ActionResult<TourReviewDto> Get(int id)
    {
        var result = _tourReviewService.Get(id);
        return CreateResponse(result);
    }

    [HttpPost]
    public async Task<string> Create([FromBody] TourReviewDto tourReview)
    {
        tourReview.UserId = User.PersonId();
        //var result = _tourReviewService.Create(tourReview);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(tourReview), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpReviewClient.PostAsync(httpReviewClient.BaseAddress,  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpPut("{id:int}")]
    public ActionResult<TourReviewDto> Update([FromBody] TourReviewDto tourReview)
    {
        tourReview.UserId = User.PersonId();
        var result = _tourReviewService.Update(tourReview);
        return CreateResponse(result);
    }

    [HttpDelete("{id:int}")]
    public ActionResult Delete(int id)
    {
        var result = _tourReviewService.Delete(id);
        return CreateResponse(result);
    }

    [HttpGet("tour/{tourId:long}")]
    [AllowAnonymous]
    public ActionResult<PagedResult<TourReviewDto>> GetByTourId(long tourId, [FromQuery] int page,
        [FromQuery] int pageSize)
    {
        var result = _tourReviewService.GetByTourId(tourId, page, pageSize);
        return CreateResponse(result);
    }

    [HttpPost("averageRate")]
    [AllowAnonymous]
    public ActionResult<double> CalculateAverageRate([FromBody] List<TourReviewDto> reviews)
    {
        var result = _tourReviewService.CalculateAverageRate(reviews);
        return CreateResponse(result);
    }
}