# Design System Strategy: The Architectural Career Portal

## 1. Overview & Creative North Star: "The Digital Curator"
To elevate this career portal beyond a generic job board, we are adopting the **"Digital Curator"** North Star. This system rejects the cluttered, high-anxiety layouts of traditional recruitment platforms. Instead, it treats career management as a high-end editorial experience. 

We achieve "High-Trust" not through rigid borders, but through **Tonal Certainty**—a layout that feels structurally sound because of its mathematical spacing and sophisticated layering. We move away from the "template" look by using intentional asymmetry in dashboard widgets and dramatic typographic scales that command attention.

## 2. Colors & Atmospheric Depth
The palette is a study in slate, ink, and atmospheric blues. It is designed to feel calm, professional, and expensive.

### The "No-Line" Rule
**Strict Directive:** 1px solid borders for sectioning are prohibited. To define boundaries between the sidebar, main content, and dashboard widgets, use background color shifts only.
*   **Sidebar:** `surface-container` (#e8eff3).
*   **Main Canvas:** `background` (#f7f9fb).
*   **High-Priority Widgets:** `surface-container-lowest` (#ffffff).
*   **Secondary Content:** `surface-container-low` (#f0f4f7).

### Surface Hierarchy & Nesting
Treat the UI as physical layers of fine paper. A card (`surface-container-lowest`) should sit on a section of `surface-container-low`. This "nested depth" creates a sense of organization that feels organic rather than forced.

### The "Glass & Gradient" Rule
To prevent the UI from feeling "flat," primary CTAs and hero headers should utilize a subtle linear gradient: `primary` (#565e74) to `primary-dim` (#4a5268) at a 135-degree angle. For floating overlays (like profile menus), use `surface-container-lowest` at 80% opacity with a `20px` backdrop-blur to create a "frosted glass" effect.

## 3. Typography: Editorial Authority
We utilize a pairing of **Manrope** (Display/Headlines) and **Inter** (Body/Labels) to balance character with utility.

*   **Display-LG (Manrope, 3.5rem):** Use sparingly for welcome messages (e.g., "Hello, Alex.") to create a high-end magazine feel.
*   **Headline-SM (Manrope, 1.5rem):** The standard for dashboard widget titles. The geometric nature of Manrope conveys "modern tech."
*   **Title-MD (Inter, 1.125rem):** Used for job titles and primary navigation links. Inter provides the "high-trust" legibility required for dense information.
*   **Label-SM (Inter, 0.6875rem):** Used for metadata (e.g., "Posted 2d ago"). Always use `on-surface-variant` (#566166) to create a clear visual hierarchy.

## 4. Elevation & Depth: Tonal Layering
Traditional shadows are often a crutch for poor spacing. In this system, depth is earned through tone.

*   **The Layering Principle:** Place a "Job Card" (`surface-container-lowest`) on the "Dashboard Canvas" (`background`). The subtle shift from #f7f9fb to #ffffff is enough to signal interactivity without visual noise.
*   **Ambient Shadows:** If a card must float (e.g., a dragged kanban item), use a `24px` blur, `0px` offset-y, and `4%` opacity using the `on-surface` color (#2a3439). It should look like a soft glow, not a drop shadow.
*   **The Ghost Border:** If a boundary is required for accessibility in input fields, use `outline-variant` (#a9b4b9) at **15% opacity**. It should be felt, not seen.

## 5. Components: The Executive Toolkit

### Buttons
*   **Primary:** Gradient from `primary` to `primary-dim`. Corner radius `md` (0.375rem). Text is `on-primary`.
*   **Secondary:** No background. `Ghost Border` (15% opacity `outline`). This keeps the UI light.
*   **Tertiary:** Text-only using `primary` color. Reserved for "Cancel" or "Back" actions.

### Cards & Lists (The "No-Divider" Mandate)
Forbid the use of horizontal rules (`<hr>`). 
*   **Separation:** Use `spacing-6` (1.5rem) of vertical white space to separate list items.
*   **Hover State:** On hover, transition the background from `transparent` to `surface-container-high` (#e1e9ee).

### The "Nexus" Sidebar
*   **Width:** Fixed 280px.
*   **Style:** `surface-container` background. 
*   **Active State:** The active nav item should use a "pill" shape (`rounded-full`) in `primary-container` (#dae2fd) with `on-primary-container` text. Avoid using high-contrast "active" colors; keep it tonal.

### Input Fields
*   **Surface:** `surface-container-low`.
*   **Interaction:** On focus, transition the background to `surface-container-lowest` and apply a `1px` "Ghost Border" at 40% opacity.

## 6. Do’s and Don’ts

### Do:
*   **Do** use asymmetrical margins. If the sidebar is 280px, give the right-side content a wider margin (`spacing-12`) than the top margin (`spacing-8`) to create an editorial "breathing" effect.
*   **Do** use `on-surface-variant` for helper text to keep the interface from feeling "heavy" with too much black text.
*   **Do** leverage the `xl` (0.75rem) roundedness for large layout containers to soften the professional tone.

### Don't:
*   **Don't** use pure black (#000000). Always use `on-surface` (#2a3439) for text to maintain the "Slate Grey" sophisticated vibe.
*   **Don't** use icons as the primary way to communicate. In a high-trust system, text labels are mandatory.
*   **Don't** use 100% opacity primary colors for large backgrounds. It creates "visual fatigue." Always lean on the `container` or `fixed` variants for large surface areas.