import { useStorage } from '@vueuse/core';
import { computed } from 'vue';

export type ReposSortBy = 'activity' | 'name';

interface UserConfig {
  isPipelineFeedOpen: boolean;
  redirectUrl: string;
  collapseLogGroupsByDefault: boolean;
  reposSortBy: ReposSortBy;
}

const config = useStorage<UserConfig>(
  'woodpecker:user-config',
  {
    isPipelineFeedOpen: false,
    redirectUrl: '',
    collapseLogGroupsByDefault: true,
    reposSortBy: 'activity',
  },
  undefined,
  { mergeDefaults: true },
);

export default () => ({
  setUserConfig<T extends keyof UserConfig>(key: T, value: UserConfig[T]): void {
    config.value = { ...config.value, [key]: value };
  },
  userConfig: computed(() => config.value),
});
