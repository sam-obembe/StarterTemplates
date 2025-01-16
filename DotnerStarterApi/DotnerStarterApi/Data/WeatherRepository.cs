namespace DotnerStarterApi.Data;

public class WeatherRepository:IRepository<WeatherEntity>
{
    public Task<IEnumerable<WeatherEntity>> GetAsync(int? skip, int? take)
    {
        throw new NotImplementedException();
    }

    public Task<WeatherEntity> CreateAsync(WeatherEntity entity)
    {
        throw new NotImplementedException();
    }

    public WeatherEntity Create(WeatherEntity entity)
    {
        throw new NotImplementedException();
    }

    public Task<WeatherEntity?> GetByIdAsync(int id)
    {
        throw new NotImplementedException();
    }

    public Task<List<WeatherEntity>> CreateAsync(List<WeatherEntity> entities)
    {
        throw new NotImplementedException();
    }

    public Task UpdateAsync(List<WeatherEntity> entities)
    {
        throw new NotImplementedException();
    }
}