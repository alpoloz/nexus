INSERT INTO blog_sources (id, name, site_url, logo_url)
VALUES
    ('airbnb-engineering', 'Airbnb Engineering', 'https://airbnb.tech', 'https://images.unsplash.com/photo-1454165804606-c3d57bc86b40'),
    ('uber-engineering', 'Uber Engineering', 'https://www.uber.com/blog/engineering', 'https://images.unsplash.com/photo-1516321318423-f06f85e504b3'),
    ('stripe-engineering', 'Stripe Engineering', 'https://stripe.com/blog/engineering', 'https://images.unsplash.com/photo-1518770660439-4636190af475')
ON CONFLICT (id) DO NOTHING;

INSERT INTO blog_tags (id, name)
VALUES
    ('distributed-systems', 'Distributed Systems'),
    ('performance', 'Performance'),
    ('developer-tools', 'Developer Tools'),
    ('scalability', 'Scalability')
ON CONFLICT (id) DO NOTHING;

INSERT INTO blog_posts (id, source_id, title, summary, external_url, hero_image_url, published_at, read_time_minutes, is_featured)
SELECT 'architecting-for-hyper-scale', id, 'Architecting for Hyper-Scale', 'Load balancing and connection pooling decisions for heavy concurrent traffic.', 'https://airbnb.tech', 'https://images.unsplash.com/photo-1516321497487-e288fb19713f', '2026-03-01T00:00:00Z', 12, TRUE
FROM blog_sources
WHERE id = 'airbnb-engineering'
ON CONFLICT (id) DO NOTHING;

INSERT INTO blog_posts (id, source_id, title, summary, external_url, hero_image_url, published_at, read_time_minutes, is_featured)
SELECT 'improving-mobile-build-performance', id, 'Improving Mobile Build Performance with Bazel', 'A closer look at faster build pipelines for large engineering organizations.', 'https://www.uber.com/blog/engineering', 'https://images.unsplash.com/photo-1498050108023-c5249f4df085', '2026-02-18T00:00:00Z', 8, FALSE
FROM blog_sources
WHERE id = 'uber-engineering'
ON CONFLICT (id) DO NOTHING;

INSERT INTO blog_posts (id, source_id, title, summary, external_url, hero_image_url, published_at, read_time_minutes, is_featured)
SELECT 'five-nines-api-availability', id, 'How We Handle Five-Nines API Availability', 'Multi-region fallback and operational discipline behind resilient APIs.', 'https://stripe.com/blog/engineering', 'https://images.unsplash.com/photo-1558494949-ef010cbdcc31', '2026-01-27T00:00:00Z', 15, TRUE
FROM blog_sources
WHERE id = 'stripe-engineering'
ON CONFLICT (id) DO NOTHING;

INSERT INTO blog_post_tags (post_id, tag_id)
SELECT bp.id, bt.id
FROM blog_posts bp
JOIN blog_tags bt ON bt.id IN ('distributed-systems', 'scalability')
WHERE bp.id = 'architecting-for-hyper-scale'
ON CONFLICT DO NOTHING;

INSERT INTO blog_post_tags (post_id, tag_id)
SELECT bp.id, bt.id
FROM blog_posts bp
JOIN blog_tags bt ON bt.id IN ('developer-tools', 'performance')
WHERE bp.id = 'improving-mobile-build-performance'
ON CONFLICT DO NOTHING;

INSERT INTO blog_post_tags (post_id, tag_id)
SELECT bp.id, bt.id
FROM blog_posts bp
JOIN blog_tags bt ON bt.id IN ('distributed-systems', 'scalability')
WHERE bp.id = 'five-nines-api-availability'
ON CONFLICT DO NOTHING;

INSERT INTO startup_tags (id, name)
VALUES
    ('edge-computing', 'Edge Computing'),
    ('health-data', 'Health Data'),
    ('fintech-infra', 'Fintech Infrastructure'),
    ('observability', 'Observability')
ON CONFLICT (id) DO NOTHING;

INSERT INTO startup_companies (id, name, description, sector, funding_stage, funding_amount, team_size, location, logo_url, website_url, careers_url, is_featured)
VALUES
    (
        'luminary-ai',
        'Luminary AI',
        'Distributed infrastructure monitoring with autonomous decision-making at the edge.',
        'AI & Data',
        'Series B',
        '$24M',
        '42',
        'San Francisco, CA',
        'https://images.unsplash.com/photo-1579546929518-9e396f3cc809',
        'https://example.com/luminary',
        'https://example.com/luminary/careers',
        TRUE
    ),
    (
        'velox-health',
        'Velox Health',
        'Wearable telemetry and cloud analytics to reduce patient readmission and improve clinical response.',
        'HealthTech',
        'Seed',
        '$12.5M',
        '28',
        'Boston, MA',
        'https://images.unsplash.com/photo-1576091160550-2173dba999ef',
        'https://example.com/velox',
        'https://example.com/velox/careers',
        FALSE
    ),
    (
        'ledgerforge',
        'LedgerForge',
        'Reliable fintech infrastructure for modern transaction orchestration and reconciliation.',
        'Fintech',
        'Series A',
        '$18M',
        '35',
        'New York, NY',
        'https://images.unsplash.com/photo-1556740749-887f6717d7e4',
        'https://example.com/ledgerforge',
        'https://example.com/ledgerforge/careers',
        TRUE
    )
ON CONFLICT (id) DO NOTHING;

INSERT INTO startup_company_tags (company_id, tag_id)
SELECT sc.id, st.id
FROM startup_companies sc
JOIN startup_tags st ON st.id IN ('edge-computing', 'observability')
WHERE sc.id = 'luminary-ai'
ON CONFLICT DO NOTHING;

INSERT INTO startup_company_tags (company_id, tag_id)
SELECT sc.id, st.id
FROM startup_companies sc
JOIN startup_tags st ON st.id IN ('health-data')
WHERE sc.id = 'velox-health'
ON CONFLICT DO NOTHING;

INSERT INTO startup_company_tags (company_id, tag_id)
SELECT sc.id, st.id
FROM startup_companies sc
JOIN startup_tags st ON st.id IN ('fintech-infra')
WHERE sc.id = 'ledgerforge'
ON CONFLICT DO NOTHING;
