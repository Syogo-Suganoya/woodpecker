import { createPinia, setActivePinia } from 'pinia';
import { beforeEach, describe, expect, it } from 'vitest';

import type { Repo } from '~/lib/api/types';

import useRepos from './useRepos';

function repo(id: number, fullName: string, created?: number): Repo {
  return {
    id,
    full_name: fullName,
    last_pipeline: created === undefined ? undefined : { created },
  } as Repo;
}

describe('useRepos', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
  });

  it('sorts repos by name', () => {
    const { sortReposByName } = useRepos();
    const repos = [repo(1, 'woodpecker/web'), repo(2, 'alice/repo'), repo(3, 'woodpecker/server')];

    expect(sortReposByName(repos).map((r) => r.id)).toStrictEqual([2, 3, 1]);
    // the input is not modified
    expect(repos.map((r) => r.id)).toStrictEqual([1, 2, 3]);
  });

  it('sorts repos by last activity', () => {
    const { sortReposByLastActivity } = useRepos();
    const repos = [repo(1, 'a/a', 10), repo(2, 'b/b'), repo(3, 'c/c', 20)];

    expect(sortReposByLastActivity(repos).map((r) => r.id)).toStrictEqual([3, 1, 2]);
  });
});
