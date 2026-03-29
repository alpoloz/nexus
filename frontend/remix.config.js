/** @type {import('@remix-run/dev').AppConfig} */
export default {
  appDirectory: "app",
  ignoredRouteFiles: ["**/*.css"],
  future: {
    v3_fetcherPersist: true,
    v3_relativeSplatPath: true,
    v3_throwAbortReason: true,
    v3_lazyRouteDiscovery: true
  }
};
