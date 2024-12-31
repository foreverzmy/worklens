import { useBehaviorSubjectState } from '@service/useBehaviorSubjectState';
import { useStream } from '@service/useStream';
import type { RepoReference } from '@typings/message';
import { Card } from 'primereact/card';
import { Tag } from 'primereact/tag';
import type { FC } from 'react';
import { scan } from 'rxjs';
import { useRepoPath } from '../../context';

export const RepoInfo: FC = () => {
	const repoPath = useRepoPath();
	const [remotes, { sub$Ref: remoteSub$Ref }] = useBehaviorSubjectState<any[]>(
		[],
	);
	const [head, { sub$Ref: headSub$Ref }] =
		useBehaviorSubjectState<RepoReference | null>(null);

	const [branches, { sub$Ref: branchSub$Ref }] = useBehaviorSubjectState<
		RepoReference,
		RepoReference[]
	>(
		[],
		scan((acc, cur) => [...acc, cur], [] as any[]),
	);

	useStream('/repo/info', {
		remote: remoteSub$Ref.current,
		branch: branchSub$Ref.current,
		head: headSub$Ref.current,
	});

	return (
		<div className="m-4 flex flex-col gap-4">
			<Card title="仓库信息">
				<p className="m-0">{repoPath}</p>
				<p>
					{remotes.map((remote) => (
						<p key={remote.name}>{remote.name}</p>
					))}
				</p>
			</Card>
			<Card title="分支信息">
				<p>
					{branches.map((branch) => (
						<p key={branch.name}>
							{branch.name} {branch.name === head?.name && <Tag value="head" />}
						</p>
					))}
				</p>
			</Card>
		</div>
	);
};
