import { Link } from "@remix-run/react";
import type { ReactNode } from "react";

interface AppShellProps {
  activeTab: "blogs" | "startups";
  children: ReactNode;
}

export function AppShell({ activeTab, children }: AppShellProps) {
  return (
    <div className="app-shell">
      <aside className="sidebar">
        <div className="brand-block">
          <p className="brand-title">Nexus Portal</p>
          <p className="brand-subtitle">The Digital Curator</p>
        </div>

        <nav className="sidebar-nav" aria-label="Primary">
          <Link className={navClass(activeTab === "blogs")} to="/blogs">
            Engineering Blogs
          </Link>
          <Link className={navClass(activeTab === "startups")} to="/startups">
            Startups
          </Link>
        </nav>
      </aside>

      <div className="content-shell">
        <header className="topbar">
          <div>
            <p className="topbar-label">Career Portal</p>
          </div>
        </header>
        <main className="page">{children}</main>
      </div>
    </div>
  );
}

function navClass(isActive: boolean): string {
  return isActive ? "sidebar-link sidebar-link-active" : "sidebar-link";
}
