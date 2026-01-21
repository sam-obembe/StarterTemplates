// See https://aka.ms/new-console-template for more information


using DotnetConsoleWithHost;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

HostApplicationBuilder builder = Host.CreateApplicationBuilder(args);
builder.Services.AddHostedService<Runner>();

var app = builder.Build();
await app.RunAsync();