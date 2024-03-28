using System.Text;
using System.Text.Json;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourAuthoring;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Author;

[Authorize(Policy = "personPolicy")]
[Route("api/author/keypoints")]
public class KeypointController : BaseApiController
{
    private readonly IKeypointService _keypointService;
    private static readonly string _tourAppPort = Environment.GetEnvironmentVariable("TOURS_APP_PORT") ?? "8080";
    private static HttpClient httpKeypointClient = new()
    {
        BaseAddress = new Uri("http://localhost:" + _tourAppPort + "/api/keypoints/"),
    };

    public KeypointController(IKeypointService keypointService)
    {
        _keypointService = keypointService;
    }

    [HttpGet]
    public ActionResult<PagedResult<KeypointDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        var result = _keypointService.GetPaged(page, pageSize);
        return CreateResponse(result);
    }

    [HttpGet("tour/{tourId:int}")]
    public async Task<string> GetByTour([FromQuery] int page, [FromQuery] int pageSize,
        [FromRoute] int tourId)
    {
        //var result = _keypointService.GetByTourId(page, pageSize, tourId);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpKeypointClient.GetAsync("tour/"+tourId);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpPost]
    public async Task<string> Create([FromBody] KeypointDto keypoint)
    {
        //var result = _keypointService.Create(keypoint);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(keypoint), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpKeypointClient.PostAsync("http://localhost:" + _tourAppPort + "/api/keypoints",  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpPost("/multiple")]
    public ActionResult<KeypointDto> CreateMultiple([FromBody] List<KeypointDto> keypoints)
    {
        var result = _keypointService.CreateMultiple(keypoints);
        return CreateResponse(result);
    }

    [HttpPut("{id:int}")]
    public async Task<string> Update([FromBody] KeypointDto keypoint)
    {
        //var result = _keypointService.Update(keypoint);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(keypoint), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpKeypointClient.PutAsync("http://localhost:" + _tourAppPort + "/api/keypoints",  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpDelete("{id:int}")]
    public async Task<string> Delete(int id)
    {
        //var result = _keypointService.Delete(id);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpKeypointClient.DeleteAsync(id.ToString());
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }
}