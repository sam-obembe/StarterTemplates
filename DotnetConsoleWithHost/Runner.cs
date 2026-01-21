using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;

namespace DotnetConsoleWithHost;

public class Runner : BackgroundService
{
    private readonly ILogger<Runner> _logger;
    private readonly IConfiguration _configuration;
    private readonly IHostApplicationLifetime _lifetime;

    public Runner(ILogger<Runner> logger, IConfiguration configuration, IHostApplicationLifetime lifetime)
    {
        _logger = logger;
        _configuration = configuration;
        _lifetime = lifetime;
    }

    protected override Task ExecuteAsync(CancellationToken stoppingToken)
    {
        _logger.LogInformation("Executing");
        _lifetime.StopApplication();
        return Task.CompletedTask;
    }
}
