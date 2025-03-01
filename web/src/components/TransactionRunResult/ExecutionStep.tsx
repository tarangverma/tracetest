import {LinkOutlined} from '@ant-design/icons';
import {Tooltip} from 'antd';
import {capitalize} from 'lodash';
import React, {useState} from 'react';
import {Link} from 'react-router-dom';
import KeyValueRow from 'components/KeyValueRow';
import {TestState} from 'constants/TestRun.constants';
import {TTestRunState} from 'types/TestRun.types';
import Test from 'models/Test.model';
import TestRun from 'models/TestRun.model';
import * as S from './TransactionRunResult.styled';

const iconBasedOnResult = (state: TTestRunState, failedAssertions: number, index: number) => {
  if (state !== TestState.FAILED && state !== TestState.FINISHED) {
    return null;
  }

  if (state === TestState.FAILED || failedAssertions > 0) {
    return <S.IconFail />;
  }
  if (state === TestState.FINISHED || failedAssertions === 0) {
    return <S.IconSuccess />;
  }

  return index + 1;
};

interface IProps {
  index: number;
  test: Test;
  testRun?: TestRun;
  hasRunFailed: boolean;
}

const ExecutionStep = ({
  index,
  test: {name, trigger, id: testId},
  hasRunFailed,
  testRun: {id: runId, state, testVersion, passedAssertionCount, failedAssertionCount, outputs} = TestRun({
    state: hasRunFailed ? TestState.SKIPPED : TestState.WAITING,
  }),
}: IProps) => {
  const [toggleOutputs, setToggleOutputs] = useState(false);
  const stateIsFinished = ([TestState.FINISHED, TestState.FAILED] as string[]).includes(state);
  const toLink = runId ? `/test/${testId}/run/${runId}` : `/test/${testId}`;

  return (
    <S.Container data-cy={`transaction-execution-step-${name}`}>
      <S.Content>
        <S.ExecutionStepStatus>{iconBasedOnResult(state, failedAssertionCount, index)}</S.ExecutionStepStatus>
        <Link to={toLink} target="_blank">
          <S.Info>
            <S.ItemName>{`${name} v${testVersion}`}</S.ItemName>
            <S.TagContainer>
              <S.TextTag>{trigger.method}</S.TextTag>
              <S.EntryPointTag $isLight>{trigger.entryPoint}</S.EntryPointTag>
              {!stateIsFinished && <S.TextTag>{capitalize(state)}</S.TextTag>}
            </S.TagContainer>
          </S.Info>
        </Link>
        <S.AssertionResultContainer>
          {runId && (
            <>
              <Tooltip title="Passed assertions">
                <S.HeaderDetail>
                  <S.HeaderDot $passed />
                  {passedAssertionCount}
                </S.HeaderDetail>
              </Tooltip>
              <Tooltip title="Failed assertions">
                <S.HeaderDetail>
                  <S.HeaderDot $passed={false} />
                  {failedAssertionCount}
                </S.HeaderDetail>
              </Tooltip>
            </>
          )}
        </S.AssertionResultContainer>
        <S.ExecutionStepStatus>
          <Tooltip title="Go to Run">
            <S.ExecutionStepRunLink to={toLink} target="_blank" data-cy="execution-step-run-link">
              <LinkOutlined />
            </S.ExecutionStepRunLink>
          </Tooltip>
        </S.ExecutionStepStatus>
      </S.Content>

      <S.OutputsContainer>
        {!!outputs?.length && (
          <S.OutputsButton onClick={() => setToggleOutputs(prev => !prev)} type="link">
            {toggleOutputs ? 'Hide Outputs' : 'Show Outputs'}
          </S.OutputsButton>
        )}

        {toggleOutputs && (
          <S.OutputsContent>
            {outputs?.map?.(output => (
              <KeyValueRow key={output.name} keyName={output.name} value={output.value} />
            ))}
          </S.OutputsContent>
        )}
      </S.OutputsContainer>
    </S.Container>
  );
};

export default ExecutionStep;
