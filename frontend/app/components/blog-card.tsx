import type { BlogPost } from "~/lib/api.server";

interface BlogCardProps {
  post: BlogPost;
}

export function BlogCard({ post }: BlogCardProps) {
  return (
    <article className="card card-blog">
      <div className="card-media">
        <img alt={post.title} src={post.heroImageUrl} />
      </div>
      <div className="card-body">
        <p className="eyebrow">
          {post.source.name} · {formatDate(post.publishedAt)}
        </p>
        <h2 className="card-title">
          <a href={post.externalUrl} rel="noreferrer" target="_blank">
            {post.title}
          </a>
        </h2>
        <p className="card-summary">{post.summary}</p>
        <div className="meta-row">
          <span>{post.readTimeMinutes} min read</span>
          <div className="tag-row">
            {post.tags.map((tag) => (
              <span className="tag-pill" key={tag}>
                {tag}
              </span>
            ))}
          </div>
        </div>
      </div>
    </article>
  );
}

function formatDate(value: string): string {
  return new Intl.DateTimeFormat("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric"
  }).format(new Date(value));
}
