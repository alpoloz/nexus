# API Contract

## Health

### `GET /healthz`

Returns service health for local orchestration and probes.

```json
{
  "status": "ok"
}
```

## Blog posts

### `GET /api/blog-posts`

Returns a cursor-paginated list of imported engineering blog posts.

#### Query params

- `cursor`: opaque pagination cursor based on `published_at` and `id`
- `limit`: default `12`, max `50`
- `q`: free-text search against title and summary
- `sourceId`: source id filter
- `tagId`: tag id filter

#### Response shape

```json
{
  "items": [
    {
      "id": "architecting-for-hyper-scale",
      "title": "Architecting for Hyper-Scale",
      "summary": "Load balancing and connection pooling strategies.",
      "externalUrl": "https://example.com/post",
      "heroImageUrl": "https://example.com/image.jpg",
      "publishedAt": "2026-03-01T00:00:00Z",
      "readTimeMinutes": 12,
      "source": {
        "id": "airbnb-engineering",
        "name": "Airbnb Engineering",
        "siteUrl": "https://airbnb.tech"
      },
      "tags": ["scalability", "distributed-systems"]
    }
  ],
  "nextCursor": "2026-03-01T00:00:00Z::1"
}
```

## Startups

### `GET /api/startups`

Returns a filterable list of startup records.

#### Query params

- `limit`: default `20`, max `50`
- `offset`: default `0`
- `q`: free-text search against name and description
- `sector`: sector filter
- `location`: location filter
- `stage`: funding stage filter
- `tagId`: tag id filter

#### Response shape

```json
{
  "items": [
    {
      "id": "luminary-ai",
      "name": "Luminary AI",
      "description": "Autonomous infrastructure monitoring.",
      "sector": "AI & Data",
      "fundingStage": "Series B",
      "fundingAmount": "$24M",
      "teamSize": "42",
      "location": "San Francisco, CA",
      "logoUrl": "https://example.com/logo.png",
      "websiteUrl": "https://luminary.example",
      "careersUrl": "https://luminary.example/careers",
      "tags": ["distributed-systems", "edge-computing"]
    }
  ],
  "total": 128,
  "limit": 20,
  "offset": 0
}
```

## Notes

- Blog cards should link directly to `externalUrl`.
- Startup records can remain list-only in MVP; detail endpoints can be added later without breaking the list contract.
- Future user-specific fields should be additive rather than replacing existing response fields.
