CREATE TABLE IF NOT EXISTS blog_sources (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    site_url TEXT NOT NULL,
    logo_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blog_posts (
    id TEXT PRIMARY KEY,
    source_id TEXT NOT NULL REFERENCES blog_sources(id) ON DELETE RESTRICT,
    title TEXT NOT NULL,
    summary TEXT NOT NULL,
    external_url TEXT NOT NULL,
    hero_image_url TEXT NOT NULL,
    published_at TIMESTAMPTZ NOT NULL,
    read_time_minutes INTEGER NOT NULL DEFAULT 0,
    is_featured BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blog_tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS blog_post_tags (
    post_id TEXT NOT NULL REFERENCES blog_posts(id) ON DELETE CASCADE,
    tag_id TEXT NOT NULL REFERENCES blog_tags(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);

CREATE TABLE IF NOT EXISTS startup_companies (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    sector TEXT NOT NULL,
    funding_stage TEXT NOT NULL,
    funding_amount TEXT NOT NULL,
    team_size TEXT NOT NULL,
    location TEXT NOT NULL,
    logo_url TEXT NOT NULL,
    website_url TEXT NOT NULL,
    careers_url TEXT NOT NULL,
    is_featured BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS startup_tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS startup_company_tags (
    company_id TEXT NOT NULL REFERENCES startup_companies(id) ON DELETE CASCADE,
    tag_id TEXT NOT NULL REFERENCES startup_tags(id) ON DELETE CASCADE,
    PRIMARY KEY (company_id, tag_id)
);

CREATE INDEX IF NOT EXISTS idx_blog_posts_published_at ON blog_posts (published_at DESC, id DESC);
CREATE INDEX IF NOT EXISTS idx_blog_posts_source_id ON blog_posts (source_id);
CREATE INDEX IF NOT EXISTS idx_startup_companies_featured_name ON startup_companies (is_featured DESC, name ASC);
