namespace DotnerStarterApi.Data;

public interface IRepository<T> where T : class
{
    public Task<IEnumerable<T>> GetAsync(int? skip, int? take);

    public Task<T> CreateAsync(T entity);
    
    public T Create(T entity);

    public Task<T?> GetByIdAsync(int id);
    
    public Task<List<T>> CreateAsync(List<T> entities);
    
    public Task UpdateAsync(List<T> entities);
    
}