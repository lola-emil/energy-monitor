// composables/useThemeColors.ts
import { ref, onMounted } from 'vue';

type ThemeColors = {
  background: string;
  foreground: string;

  primary: string;
  primaryForeground: string;

  secondary: string;
  secondaryForeground: string;

  muted: string;
  mutedForeground: string;

  accent: string;
  accentForeground: string;

  destructive: string;
  destructiveForeground: string;

  border: string;
  input: string;
  ring: string;

  popover: string;
  popoverForeground: string;

  card: string;
  cardForeground: string;

  chart1: string;
  chart2: string;
  chart3: string;
  chart4: string;
  chart5: string;
};

export function useThemeColors() {
  const colors = ref<ThemeColors | null>(null);

  const readVar = (css: CSSStyleDeclaration, name: string) =>
    css.getPropertyValue(name).trim();

  const update = () => {
    if (typeof window === 'undefined') return;

    const css = getComputedStyle(document.documentElement);

    colors.value = {
      background: readVar(css, '--background'),
      foreground: readVar(css, '--foreground'),

      primary: readVar(css, '--primary'),
      primaryForeground: readVar(css, '--primary-foreground'),

      secondary: readVar(css, '--secondary'),
      secondaryForeground: readVar(css, '--secondary-foreground'),

      muted: readVar(css, '--muted'),
      mutedForeground: readVar(css, '--muted-foreground'),

      accent: readVar(css, '--accent'),
      accentForeground: readVar(css, '--accent-foreground'),

      destructive: readVar(css, '--destructive'),
      destructiveForeground: readVar(css, '--destructive-foreground'),

      border: readVar(css, '--border'),
      input: readVar(css, '--input'),
      ring: readVar(css, '--ring'),

      popover: readVar(css, '--popover'),
      popoverForeground: readVar(css, '--popover-foreground'),

      card: readVar(css, '--card'),
      cardForeground: readVar(css, '--card-foreground'),

      chart1: readVar(css, '--chart-1'),
      chart2: readVar(css, '--chart-2'),
      chart3: readVar(css, '--chart-3'),
      chart4: readVar(css, '--chart-4'),
      chart5: readVar(css, '--chart-5'),
    };
  };

  onMounted(update);

  return {
    colors,
    refresh: update,
  };
}