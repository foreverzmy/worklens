import { useCallback, useEffect, useRef, useState } from 'react';
import { type BehaviorSubject, Subject } from 'rxjs';
import { useRepoPath } from '../context';
import { streamFetch } from './stream';

export type StreamMessageBehaviorSubjects = Record<
	string,
	BehaviorSubject<any>
>;

export const useStream = <S extends StreamMessageBehaviorSubjects>(
	path: string,
	subs$: S,
) => {
	const repoPath = useRepoPath();
	const [error, setError] = useState<Error | null>(null);
	const [loading, setLoading] = useState(false);
	const subs$Ref = useRef(subs$);

	const subjectRef = useRef(new Subject());
	const pollingRef = useRef(null);

	const cancel = useCallback(() => {
		if (pollingRef.current) {
			clearInterval(pollingRef.current);
		}
		subjectRef.current.complete();
		subjectRef.current = new Subject();
	}, []);

	const run = useCallback(() => {
		setLoading(true);
		streamFetch(repoPath, path).subscribe({
			complete: () => {
				setLoading(false);
			},
			error: (err) => {
				setLoading(false);
				setError(err);
			},
			next: (message) => {
				subs$Ref.current[message.type]?.next(message.data);
			},
		});
	}, [path, repoPath]);

	// biome-ignore lint/correctness/useExhaustiveDependencies: <explanation>
	useEffect(() => {
		run();
	}, []);

	return { loading, error, cancel, run, subs$Ref };
};
