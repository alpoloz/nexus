import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import { AppShell } from "~/components/app-shell";
import { StartupCard } from "~/components/startup-card";
import type { StartupListResponse } from "~/lib/api.server";
import { fetchApi, searchParamsToString } from "~/lib/api.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const requestUrl = new URL(request.url);
  const params = new URLSearchParams(requestUrl.searchParams);

  if (!params.has("limit")) {
    params.set("limit", "20");
  }

  const response = await fetchApi<StartupListResponse>(
    `/api/startups${searchParamsToString(params)}`,
    request
  );

  return json({
    filters: {
      q: params.get("q") ?? "",
      sector: params.get("sector") ?? "",
      location: params.get("location") ?? "",
      stage: params.get("stage") ?? ""
    },
    response
  });
}

export default function StartupsRoute() {
  const { filters, response } = useLoaderData<typeof loader>();

  return (
    <AppShell activeTab="startups">
      <section className="hero">
        <p className="page-kicker">Startup List</p>
        <h1>Discover Startups</h1>
        <p className="page-copy">
          Imported startup records from Postgres, shaped for filtering and outbound career exploration.
        </p>
      </section>

      <Form className="toolbar toolbar-grid" method="get">
        <label className="field">
          <span>Search</span>
          <input defaultValue={filters.q} name="q" placeholder="Search startup name or focus..." />
        </label>
        <label className="field">
          <span>Sector</span>
          <input defaultValue={filters.sector} name="sector" placeholder="AI & Data" />
        </label>
        <label className="field">
          <span>Location</span>
          <input defaultValue={filters.location} name="location" placeholder="San Francisco, CA" />
        </label>
        <label className="field">
          <span>Stage</span>
          <input defaultValue={filters.stage} name="stage" placeholder="Series B" />
        </label>
        <button className="button-primary" type="submit">
          Apply
        </button>
      </Form>

      <div className="results-label">
        Showing {response.items.length} of {response.total} startups
      </div>

      <div className="stack">
        {response.items.map((startup) => (
          <StartupCard key={startup.id} startup={startup} />
        ))}
      </div>
    </AppShell>
  );
}
