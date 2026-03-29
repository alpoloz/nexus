import type { Startup } from "~/lib/api.server";

interface StartupCardProps {
  startup: Startup;
}

export function StartupCard({ startup }: StartupCardProps) {
  return (
    <article className="card card-startup">
      <div className="startup-logo">
        <img alt={startup.name} src={startup.logoUrl} />
      </div>

      <div className="startup-content">
        <div className="startup-header">
          <div>
            <h2 className="card-title">{startup.name}</h2>
            <p className="eyebrow">{startup.sector}</p>
          </div>
          <span className="tag-pill">{startup.fundingStage}</span>
        </div>

        <p className="card-summary">{startup.description}</p>

        <div className="startup-facts">
          <span>{startup.fundingAmount}</span>
          <span>{startup.teamSize} team</span>
          <span>{startup.location}</span>
        </div>

        <div className="action-row">
          <a className="primary-link" href={startup.careersUrl} rel="noreferrer" target="_blank">
            View Careers
          </a>
          <a className="secondary-link" href={startup.websiteUrl} rel="noreferrer" target="_blank">
            Company Site
          </a>
        </div>
      </div>
    </article>
  );
}
