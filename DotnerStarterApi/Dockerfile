FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine AS build
WORKDIR /app
COPY . .
RUN ls
RUN dotnet restore 
RUN dotnet publish -c Release -o project

FROM mcr.microsoft.com/dotnet/aspnet:8.0-alpine
EXPOSE 8080
WORKDIR /api
COPY --from=build /app/project .
RUN ls
CMD ["./DotnetStarterApi"]