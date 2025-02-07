using Microsoft.EntityFrameworkCore;

namespace DotnerStarterApi.Data;

public class ProjectDbContext:DbContext
{
    public DbSet<WeatherEntity> Weather { get; set; }
    
    public ProjectDbContext(DbContextOptions<ProjectDbContext> options):base(options){}
}