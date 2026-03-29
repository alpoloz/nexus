import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import type { BlogPostListResponse } from "~/lib/api.server";
import { fetchApi, searchParamsToString } from "~/lib/api.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const requestUrl = new URL(request.url);
  const data = await fetchApi<BlogPostListResponse>(
    `/api/blog-posts${searchParamsToString(requestUrl.searchParams)}`,
    request
  );

  return json(data);
}
