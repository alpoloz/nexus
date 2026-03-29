export interface BlogSource {
  id: string;
  name: string;
  siteUrl: string;
}

export interface BlogPost {
  id: string;
  title: string;
  summary: string;
  externalUrl: string;
  heroImageUrl: string;
  publishedAt: string;
  readTimeMinutes: number;
  source: BlogSource;
  tags: string[];
}

export interface BlogPostListResponse {
  items: BlogPost[];
  nextCursor: string | null;
}

export interface Startup {
  id: string;
  name: string;
  description: string;
  sector: string;
  fundingStage: string;
  fundingAmount: string;
  teamSize: string;
  location: string;
  logoUrl: string;
  websiteUrl: string;
  careersUrl: string;
  tags: string[];
}

export interface StartupListResponse {
  items: Startup[];
  total: number;
  limit: number;
  offset: number;
}

export function getApiBaseUrl(): string {
  return process.env.API_BASE_URL ?? "http://localhost:8080";
}

export async function fetchApi<T>(path: string, request?: Request): Promise<T> {
  const headers = new Headers();
  if (request) {
    const acceptLanguage = request.headers.get("accept-language");
    if (acceptLanguage) {
      headers.set("accept-language", acceptLanguage);
    }
  }

  const response = await fetch(`${getApiBaseUrl()}${path}`, { headers });
  if (!response.ok) {
    throw new Response("Upstream API request failed", { status: response.status });
  }

  return (await response.json()) as T;
}

export function searchParamsToString(params: URLSearchParams): string {
  const query = params.toString();
  return query ? `?${query}` : "";
}
