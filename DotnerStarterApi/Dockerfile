#image size 141.19MB
FROM mcr.microsoft.com/dotnet/sdk:10.0-alpine AS build
WORKDIR /app
COPY . .
RUN ls
RUN dotnet restore 
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/aspnet:10.0-alpine
EXPOSE 5000
WORKDIR /api
COPY --from=build /app/out .
RUN ls
CMD ["./DotnerStarterApi"]

#img size 181.39MB
# FROM mcr.microsoft.com/dotnet/sdk:10.0-noble AS build
# WORKDIR /app
# COPY . .
# RUN ls
# RUN dotnet restore 
# RUN dotnet publish -c Release -o out

# FROM mcr.microsoft.com/dotnet/aspnet:10.0-noble-chiseled-extra
# EXPOSE 5000
# WORKDIR /api
# COPY --from=build /app/out .
# ENTRYPOINT ["./DotnerStarterApi"]
