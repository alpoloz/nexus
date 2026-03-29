import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import { Form, useFetcher, useLoaderData } from "@remix-run/react";
import { useEffect, useState } from "react";
import { AppShell } from "~/components/app-shell";
import { BlogCard } from "~/components/blog-card";
import type { BlogPost, BlogPostListResponse } from "~/lib/api.server";
import { fetchApi, searchParamsToString } from "~/lib/api.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const requestUrl = new URL(request.url);
  const params = new URLSearchParams(requestUrl.searchParams);

  if (!params.has("limit")) {
    params.set("limit", "12");
  }

  const response = await fetchApi<BlogPostListResponse>(
    `/api/blog-posts${searchParamsToString(params)}`,
    request
  );

  return json({
    filters: {
      q: params.get("q") ?? "",
      sourceId: params.get("sourceId") ?? ""
    },
    response
  });
}

export default function BlogsRoute() {
  const { filters, response } = useLoaderData<typeof loader>();
  const fetcher = useFetcher<BlogPostListResponse>();
  const [items, setItems] = useState<BlogPost[]>(response.items);
  const [nextCursor, setNextCursor] = useState<string | null>(response.nextCursor);

  useEffect(() => {
    setItems(response.items);
    setNextCursor(response.nextCursor);
  }, [response]);

  useEffect(() => {
    const nextPage = fetcher.data;
    if (!nextPage) {
      return;
    }

    setItems((current) => [...current, ...nextPage.items]);
    setNextCursor(nextPage.nextCursor);
  }, [fetcher.data]);

  function loadMore() {
    if (!nextCursor || fetcher.state !== "idle") {
      return;
    }

    const params = new URLSearchParams();
    params.set("limit", "12");
    params.set("cursor", nextCursor);
    if (filters.q) {
      params.set("q", filters.q);
    }
    if (filters.sourceId) {
      params.set("sourceId", filters.sourceId);
    }

    fetcher.load(`/resources/blog-posts?${params.toString()}`);
  }

  return (
    <AppShell activeTab="blogs">
      <section className="hero">
        <p className="page-kicker">Engineering Blogs</p>
        <h1>Engineering Insights</h1>
        <p className="page-copy">
          Imported writing from high-signal engineering teams, stored in Postgres and linked back
          to the original article.
        </p>
      </section>

      <Form className="toolbar" method="get">
        <label className="field">
          <span>Search</span>
          <input defaultValue={filters.q} name="q" placeholder="Search architecture, infra, scale..." />
        </label>
        <label className="field">
          <span>Source ID</span>
          <input defaultValue={filters.sourceId} name="sourceId" placeholder="airbnb-engineering" />
        </label>
        <button className="button-primary" type="submit">
          Apply
        </button>
      </Form>

      <div className="stack">
        {items.map((post) => (
          <BlogCard key={post.id} post={post} />
        ))}
      </div>

      {nextCursor ? (
        <div className="load-more-row">
          <button className="button-secondary" onClick={loadMore} type="button">
            {fetcher.state === "loading" ? "Loading..." : "Load More"}
          </button>
        </div>
      ) : null}
    </AppShell>
  );
}
