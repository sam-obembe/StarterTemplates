// See https://aka.ms/new-console-template for more information


using Microsoft.Extensions.Hosting;

HostApplicationBuilder builder = Host.CreateApplicationBuilder(args);

var app = builder.Build();

Console.WriteLine("Running");



app.Run();