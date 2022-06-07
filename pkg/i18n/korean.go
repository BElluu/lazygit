package i18n

const koreanIntroPopupMessage = `
lazygit!를 이용해주셔서 감사합니다. Seriously you rock. Three things to share with you:

 1) lazygit의 기능에 대해 알아보려면 다음 비디오를 참조하세요.
      https://youtu.be/CPLdltN7wgE

 2) 다음 사이트에서 최신 릴리스 노트를 읽어보세요.:
      https://github.com/jesseduffield/lazygit/releases

 3) 만약 당신이 Git을 사용한다면, 그것은 당신을 프로그래머로 만들 것입니다! 
	 당신의 도움으로 우리는 lazygit을 더 좋게 만들 수 있습니다, 그러니 기여자가 되는 것을 고려해보세요. 그리고 재미에 참여하세요:
      https://github.com/jesseduffield/lazygit
    또한 오른쪽 하단의 기부 버튼을 클릭하여 저를 후원하고 작업할 내용을 알려주실 수 있습니다.
    또는 저장소에 스타를 눌러 사랑을 공유할 수도 있습니다!
`

// exporting this so we can use it in tests
func koreanTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "패널을 렌더링 할 공간이 부족합니다.",
		DiffTitle:                           "Diff",
		FilesTitle:                          "파일",
		BranchesTitle:                       "브랜치",
		CommitsTitle:                        "커밋",
		StashTitle:                          "Stash",
		UnstagedChanges:                     `Staged되지 않은 변경 내용`,
		StagedChanges:                       `Staged된 변경 내용`,
		MainTitle:                           "메인",
		MergeConfirmTitle:                   "병합",
		StagingTitle:                        "메인 패널 (Staging)",
		MergingTitle:                        "메인 패널 (Merging)",
		NormalTitle:                         "메인 패널 (Normal)",
		LogTitle:                            "로그",
		CommitMessage:                       "커밋 메시지",
		CredentialsUsername:                 "사용자 이름",
		CredentialsPassword:                 "패스워드",
		CredentialsPassphrase:               "SSH키의 passphrase 입력",
		PassUnameWrong:                      "패스워드, passphrase 또는 사용자 이름이 잘못되었습니다.",
		CommitChanges:                       "커밋 변경내용",
		AmendLastCommit:                     "마지맛 커밋 수정",
		AmendLastCommitTitle:                "마지막 커밋 수정",
		SureToAmend:                         "마지막 커밋을 수정하시겠습니까? 그런 다음 커밋 패널에서 커밋 메시지를 변경할 수 있습니다.",
		NoCommitToAmend:                     "amend 가능한 커밋이 없습니다.",
		CommitChangesWithEditor:             "Git 편집기를 사용하여 변경 내용을 커밋합니다.",
		StatusTitle:                         "상태",
		LcNavigate:                          "이동",
		LcMenu:                              "메뉴",
		LcExecute:                           "실행",
		LcToggleStaged:                      "Staged 전환",
		LcToggleStagedAll:                   "모든 변경을 Staged/unstaged으로 전환",
		LcToggleTreeView:                    "파일 트리뷰로 전환",
		LcOpenMergeTool:                     "git mergetool를 열기",
		LcRefresh:                           "새로고침",
		LcPush:                              "푸시",
		LcPull:                              "업데이트",
		LcScroll:                            "스크롤",
		MergeConflictsTitle:                 "병합 충돌 내용",
		LcCheckout:                          "체크아웃",
		LcFileFilter:                        "파일을 필터하기 (Staged/unstaged)",
		FilterStagedFiles:                   "Staged된 파일만 표시",
		FilterUnstagedFiles:                 "Stage되지 않은 파일만 표시",
		ResetCommitFilterState:              "필터 리셋",
		NoChangedFiles:                      "변경된 파일이 없습니다.",
		NoFilesDisplay:                      "표시할 파일이 없습니다",
		NotAFile:                            "파일이 아닙니다.",
		PullWait:                            "업데이트 중...",
		PushWait:                            "푸시 중...",
		FetchWait:                           "패치 중...",
		LcSoftReset:                         "소프트 리셋",
		AlreadyCheckedOutBranch:             "브랜치가 이미 체크아웃 되었습니다",
		SureForceCheckout:                   "강제로 체크아웃하시겠습니까? 모든 로컬 변경 사항을 잃게 됩니다.",
		ForceCheckoutBranch:                 "브랜치 강제 체크아웃",
		BranchName:                          "브랜치 이름",
		NewBranchNameBranchOff:              "새 브랜치 이름 (Branch is off of '{{.branchName}}')",
		CantDeleteCheckOutBranch:            "체크아웃하는 브랜치는 삭제할 수 없습니다!",
		DeleteBranch:                        "브랜치 삭제",
		DeleteBranchMessage:                 "정말로 브랜치 '{{.selectedBranchName}}' 를 삭제하시겠습니까?",
		ForceDeleteBranchMessage:            "'{{.selectedBranchName}}'는 완전히 병합되지 않았습니다. 정말 삭제하시겠습니까?",
		LcRebaseBranch:                      "체크아웃된 브랜치를 이 브랜치에 리베이스",
		CantRebaseOntoSelf:                  "브랜치를 자기 자신에게 리베이스할 수는 없습니다.",
		CantMergeBranchIntoItself:           "브랜치를 자기 자신에게 병합할 수는 없습니다.",
		LcForceCheckout:                     "강제 체크아웃",
		LcCheckoutByName:                    "이름으로 체크아웃",
		LcNewBranch:                         "새 브랜치 생성",
		LcDeleteBranch:                      "브랜치 삭제",
		NoBranchesThisRepo:                  "저장소에 브랜치가 존재하지 않습니다.",
		CommitMessageConfirm:                "{{.keyBindClose}}: 닫기, {{.keyBindNewLine}}: 개행, {{.keyBindConfirm}}: 확인",
		CommitWithoutMessageErr:             "커밋 메시지를 입력하세요.",
		CloseConfirm:                        "{{.keyBindClose}}: 닫기/취소, {{.keyBindConfirm}}: 확인",
		LcClose:                             "닫기",
		LcQuit:                              "종료",
		LcSquashDown:                        "squash down",
		LcFixupCommit:                       "fixup commit",
		NoCommitsThisBranch:                 "이 브랜치에 커밋이 없습니다.",
		OnlySquashTopmostCommit:             "Can only squash topmost commit",
		YouNoCommitsToSquash:                "You have no commits to squash with",
		Fixup:                               "Fixup",
		SureFixupThisCommit:                 "Are you sure you want to 'fixup' this commit? It will be merged into the commit below",
		SureSquashThisCommit:                "Are you sure you want to squash this commit into the commit below?",
		Squash:                              "Squash",
		LcPickCommit:                        "pick commit (when mid-rebase)",
		LcRevertCommit:                      "커밋 되돌리기",
		LcRewordCommit:                      "커밋메시지 변경",
		LcDeleteCommit:                      "커밋 삭제",
		LcMoveDownCommit:                    "커밋을 1개 아래로 이동",
		LcMoveUpCommit:                      "커밋을 1개 위로 이동",
		LcEditCommit:                        "커밋을 편집",
		LcAmendToCommit:                     "amend commit with staged changes",
		LcResetCommitAuthor:                 "reset commit author",
		SureResetCommitAuthor:               "The author field of this commit will be updated to match the configured user. This also renews the author timestamp. Continue?",
		LcRenameCommitEditor:                "에디터에서 커밋메시지 수정",
		Error:                               "오류",
		LcSelectHunk:                        "hunk를 선택",
		LcNavigateConflicts:                 "navigate conflicts",
		LcPickHunk:                          "pick hunk",
		LcPickAllHunks:                      "pick all hunks",
		LcUndo:                              "되돌리기",
		LcUndoReflog:                        "되돌리기 (reflog) (실험적)",
		LcRedoReflog:                        "다시 실행 (reflog) (실험적)",
		LcPop:                               "pop",
		LcDrop:                              "drop",
		LcApply:                             "적용",
		NoStashEntries:                      "Stash가 존재하지 않습니다.",
		StashDrop:                           "Stash를 삭제",
		SureDropStashEntry:                  "정말로 Stash를 삭제하시겠습니까?",
		StashPop:                            "Stash를 pop",
		SurePopStashEntry:                   "정말로 Stash를 pop하시겠습니까?",
		StashApply:                          "Stash 적용",
		SureApplyStashEntry:                 "정말로 Stash를 적용하시겠습니까?",
		NoTrackedStagedFilesStash:           "You have no tracked/staged files to stash",
		StashChanges:                        "변경을 Stash",
		OpenConfig:                          "설정 파일 열기",
		EditConfig:                          "설정 파일 수정",
		ForcePush:                           "강제 푸시",
		ForcePushPrompt:                     "브랜치가 원격 브랜치에서 분기하고 있습니다. 'esc'를 눌러 취소하거나, 'enter'를 눌러 강제로 푸시하세요.",
		ForcePushDisabled:                   "브랜치가 원격 브랜치에서 분기하고 있습니다. force push가 비활성화 되었습니다.",
		UpdatesRejectedAndForcePushDisabled: "업데이트가 거부되었으며 강제 푸시를 비활성화했습니다.",
		LcCheckForUpdate:                    "업데이트 확인",
		CheckingForUpdates:                  "업데이트 확인 중...",
		UpdateAvailableTitle:                "새로운 업데이트 사용가능!",
		UpdateAvailable:                     "버전 {{.newVersion}} 을(를) 설치하시겠습니까?",
		UpdateInProgressWaitingStatus:       "업데이트 중",
		UpdateCompletedTitle:                "업데이트 완료!",
		UpdateCompleted:                     "업데이트 설치에 성공했습니다. lazygit를 재시작해주세요.",
		FailedToRetrieveLatestVersionErr:    "버전 정보를 받아오는데 실패했습니다.",
		OnLatestVersionErr:                  "이미 최신 버전을 사용하고 있습니다.",
		MajorVersionErr:                     "새 버전 ({{.newVersion}}) 에 현재 버전({{.currentVersion}}) 과 비교할 때 호환되지 않는 변경 사항이 있습니다.",
		CouldNotFindBinaryErr:               "{{.url}} 에서 바이너리를 찾을 수 없습니다.",
		UpdateFailedErr:                     "업데이트 실패: {{.errMessage}}",
		ConfirmQuitDuringUpdateTitle:        "현재 업데이트 중입니다.",
		ConfirmQuitDuringUpdate:             "현재 업데이트를 진행 중입니다.종료하시겠습니까?",
		MergeToolTitle:                      "병합 도구",
		MergeToolPrompt:                     "정말로 `git mergetool`을 여시겠습니까?",
		IntroPopupMessage:                   koreanIntroPopupMessage,
		GitconfigParseErr:                   `따옴표로 묶이지 않은 '\' 문자가 있어서 Gogit이 gitconfig 파일을 분석하지 못했습니다. 이를 제거하면 문제가 해결됩니다.`,
		LcEditFile:                          `파일 편집`,
		LcOpenFile:                          `파일 닫기`,
		LcIgnoreFile:                        `.gitignore에 추가`,
		LcRefreshFiles:                      `파일 새로고침`,
		LcMergeIntoCurrentBranch:            `현재 브랜치에 병합`,
		ConfirmQuit:                         `정말로 종료하시겠습니까?`,
		SwitchRepo:                          `최근에 사용한 저장소로 전환`,
		LcAllBranchesLogGraph:               `모든 브랜치 로그 표시`,
		UnsupportedGitService:               `지원되지 않는 Git 서비스입니다.`,
		LcCreatePullRequest:                 `풀 리퀘스트 생성`,
		LcCopyPullRequestURL:                `풀 리퀘스트 URL을 클립보드에 복사`,
		NoBranchOnRemote:                    `브랜치가 원격에 없습니다. 원격에 먼저 푸시해야합니다.`,
		LcFetch:                             `fetch`,
		NoAutomaticGitFetchTitle:            `자동 git 업데이트) 없음`,
		NoAutomaticGitFetchBody:             `Lazygit은 private 저장소에서 "git fetch"를 사용할 수 없습니다. 파일 패널에서 'f'를 사용하여 "git fetch"를 수동으로 실행하세요.`,
		FileEnter:                           `stage individual hunks/lines for file, or collapse/expand for directory`,
		FileStagingRequirements:             `추적된 파일에 대해 개별 라인만 stage할 수 있습니다.`,
		StageSelection:                      `선택한 행을 staged / unstaged`,
		ResetSelection:                      `변경을 삭제 (git reset)`,
		ToggleDragSelect:                    `드래그 선택 전환`,
		ToggleSelectHunk:                    `toggle select hunk`,
		ToggleSelectionForPatch:             `line(s)을 패치에 추가/삭제`,
		TogglePanel:                         `패널 전환`,
		ReturnToFilesPanel:                  `파일 목록으로 돌아가기`,
		FastForward:                         `fast-forward this branch from its upstream`,
		Fetching:                            "fetching and fast-forwarding {{.from}} -> {{.to}} ...",
		FoundConflicts:                      "Conflicts! To abort press 'esc', otherwise press 'enter'",
		FoundConflictsTitle:                 "Auto-merge failed",
		PickHunk:                            "pick hunk",
		PickAllHunks:                        "pick all hunks",
		ViewMergeRebaseOptions:              "view merge/rebase options",
		NotMergingOrRebasing:                "You are currently neither rebasing nor merging",
		RecentRepos:                         "최근에 사용한 저장소",
		MergeOptionsTitle:                   "Merge Options",
		RebaseOptionsTitle:                  "Rebase Options",
		CommitMessageTitle:                  "커밋메시지",
		LocalBranchesTitle:                  "브랜치",
		SearchTitle:                         "검색",
		TagsTitle:                           "태그",
		MenuTitle:                           "메뉴",
		RemotesTitle:                        "원격",
		RemoteBranchesTitle:                 "원격 브랜치",
		PatchBuildingTitle:                  "메인 패널 (Patch Building)",
		InformationTitle:                    "정보",
		SecondaryTitle:                      "Secondary",
		ReflogCommitsTitle:                  "Reflog",
		GlobalTitle:                         "글로벌 키 바인딩",
		ConflictsResolved:                   "모든 병합 충돌이 해결되었습니다. 계속 할까요?",
		RebasingTitle:                       "리베이스 중",
		ConfirmRebase:                       "정말로 '{{.checkedOutBranch}}' 을(를) '{{.selectedBranch}}'에 리베이스 하시겠습니까?",
		ConfirmMerge:                        "정말로 '{{.selectedBranch}}' 을(를) '{{.checkedOutBranch}}'에 병합하시겠습니까?",
		FwdNoUpstream:                       "Cannot fast-forward a branch with no upstream",
		FwdNoLocalUpstream:                  "Cannot fast-forward a branch whose remote is not registered locally",
		FwdCommitsToPush:                    "Cannot fast-forward a branch with commits to push",
		ErrorOccurred:                       "오류가 발생했습니다! issue를 작성해 주세요: ",
		NoRoom:                              "Not enough room",
		YouAreHere:                          "현재 위치",
		LcRewordNotSupported:                "rewording commits while interactively rebasing is not currently supported",
		LcCherryPickCopy:                    "커밋을 복사 (cherry-pick)",
		LcCherryPickCopyRange:               "커밋을 범위로 복사 (cherry-pick)",
		LcPasteCommits:                      "커밋을 붙여넣기 (cherry-pick)",
		SureCherryPick:                      "정말로 복사한 커밋을 이 브랜치에 체리픽하시겠습니까?",
		CherryPick:                          "체리픽",
		CannotRebaseOntoFirstCommit:         "첫 번째 커밋에 대해 대화식으로 리베이스할 수 없습니다.",
		CannotSquashOntoSecondCommit:        "두 번째 커밋을 squash/fixup할 수 없습니다.",
		Donate:                              "후원",
		AskQuestion:                         "질문하기",
		PrevLine:                            "이전 줄 선택",
		NextLine:                            "다음 줄 선택",
		PrevHunk:                            "이전 hunk를 선택",
		NextHunk:                            "다음 hunk를 선택",
		PrevConflict:                        "이전 충돌을 선택",
		NextConflict:                        "다음 충돌을 선택",
		SelectPrevHunk:                      "이전 hunk를 선택",
		SelectNextHunk:                      "다음 hunk를 선택",
		ScrollDown:                          "아래로 스크롤",
		ScrollUp:                            "위로 스크롤",
		LcScrollUpMainPanel:                 "메인 패널을 위로 스크롤",
		LcScrollDownMainPanel:               "메인 패널을 아래로로 스크롤",
		AmendCommitTitle:                    "Amend Commit",
		AmendCommitPrompt:                   "Are you sure you want to amend this commit with your staged files?",
		DeleteCommitTitle:                   "커밋 삭제",
		DeleteCommitPrompt:                  "정말로 선택한 커밋을 삭제하시겠습니까?",
		SquashingStatus:                     "squashing",
		FixingStatus:                        "fixing up",
		DeletingStatus:                      "deleting",
		MovingStatus:                        "moving",
		RebasingStatus:                      "rebasing",
		AmendingStatus:                      "amending",
		CherryPickingStatus:                 "cherry-picking",
		UndoingStatus:                       "undoing",
		RedoingStatus:                       "redoing",
		CheckingOutStatus:                   "checking out",
		CommittingStatus:                    "committing",
		CommitFiles:                         "Commit files",
		SubCommitsDynamicTitle:              "커밋 (%s)",
		CommitFilesDynamicTitle:             "Diff files (%s)",
		RemoteBranchesDynamicTitle:          "원격브랜치 (%s)",
		LcViewItemFiles:                     "view selected item's files",
		CommitFilesTitle:                    "커밋 파일",
		LcCheckoutCommitFile:                "checkout file",
		LcDiscardOldFileChange:              "discard this commit's changes to this file",
		DiscardFileChangesTitle:             "파일 변경 사항 버리기",
		DiscardFileChangesPrompt:            "Are you sure you want to discard this commit's changes to this file? If this file was created in this commit, it will be deleted",
		DisabledForGPG:                      "Feature not available for users using GPG",
		CreateRepo:                          "Git 저장소가 아닙니다. 저장소를 생성하시겠습니까? (y/n): ",
		AutoStashTitle:                      "Autostash?",
		AutoStashPrompt:                     "You must stash and pop your changes to bring them across. Do this automatically? (enter/esc)",
		StashPrefix:                         "Auto-stashing changes for ",
		LcViewDiscardOptions:                "view 'discard changes' options",
		LcCancel:                            "취소",
		LcDiscardAllChanges:                 "모든 변경사항 버리기",
		LcDiscardUnstagedChanges:            "discard unstaged changes",
		LcDiscardAllChangesToAllFiles:       "nuke working tree",
		LcDiscardAnyUnstagedChanges:         "discard unstaged changes",
		LcDiscardUntrackedFiles:             "discard untracked files",
		LcHardReset:                         "hard reset",
		LcViewResetOptions:                  `view reset options`,
		LcCreateFixupCommit:                 `create fixup commit for this commit`,
		LcSquashAboveCommits:                `squash all 'fixup!' commits above selected commit (autosquash)`,
		SquashAboveCommits:                  `Squash all 'fixup!' commits above selected commit (autosquash)`,
		SureSquashAboveCommits:              `Are you sure you want to squash all fixup! commits above {{.commit}}?`,
		CreateFixupCommit:                   `Create fixup commit`,
		SureCreateFixupCommit:               `Are you sure you want to create a fixup! commit for commit {{.commit}}?`,
		LcExecuteCustomCommand:              "execute custom command",
		CustomCommand:                       "Custom Command:",
		LcCommitChangesWithoutHook:          "commit changes without pre-commit hook",
		SkipHookPrefixNotConfigured:         "You have not configured a commit message prefix for skipping hooks. Set `git.skipHookPrefix = 'WIP'` in your config",
		LcResetTo:                           `reset to`,
		PressEnterToReturn:                  "엔터를 눌러 lazygit으로 돌아갑니다.",
		LcViewStashOptions:                  "Stash 옵션 보기",
		LcStashAllChanges:                   "변경사항을 Stash",
		LcStashStagedChanges:                "stash staged changes",
		LcStashOptions:                      "Stash 옵션",
		NotARepository:                      "Error: must be run inside a git repository",
		LcJump:                              "패널로 이동",
		LcScrollLeftRight:                   "좌우로 스크롤",
		LcScrollLeft:                        "우 스크롤",
		LcScrollRight:                       "좌 스크롤",
		DiscardPatch:                        "patch 버리기",
		DiscardPatchConfirm:                 "You can only build a patch from one commit/stash-entry at a time. Discard current patch?",
		CantPatchWhileRebasingError:         "You cannot build a patch or run patch commands while in a merging or rebasing state",
		LcToggleAddToPatch:                  "toggle file included in patch",
		LcToggleAllInPatch:                  "toggle all files included in patch",
		LcUpdatingPatch:                     "updating patch",
		ViewPatchOptions:                    "커스텀 Patch 옵션 보기",
		PatchOptionsTitle:                   "Patch 옵션",
		NoPatchError:                        "No patch created yet. To start building a patch, use 'space' on a commit file or enter to add specific lines",
		LcEnterFile:                         "enter file to add selected lines to the patch (or toggle directory collapsed)",
		ExitLineByLineMode:                  `line-by-line 모드 종료`,
		EnterUpstream:                       `'<remote> <branchname>'와 같은 형식으로 입력하세요.`,
		InvalidUpstream:                     "upstream의 형식이 잘못되었습니다.'<remote> <branchname>' 와 같은 형식으로 입력하세요.",
		ReturnToRemotesList:                 `원격목록으로 돌아가기`,
		LcAddNewRemote:                      `새로운 Remote 추가`,
		LcNewRemoteName:                     `새로운 Remote 이름:`,
		LcNewRemoteUrl:                      `새로운 Remote URL:`,
		LcEditRemoteName:                    `{{.remoteName}} 의 새로운 Remote 이름 입력:`,
		LcEditRemoteUrl:                     `{{.remoteName}} 의 새로운 Remote URL 입력:`,
		LcRemoveRemote:                      `Remote를 삭제`,
		LcRemoveRemotePrompt:                "정말로 Remote를 삭제하시겠습니까?",
		DeleteRemoteBranch:                  "원격 브랜치를 삭제",
		DeleteRemoteBranchMessage:           "정말로 원격 브랜치를 삭제하시겠습니까?",
		LcSetUpstream:                       "set as upstream of checked-out branch",
		SetUpstreamTitle:                    "Set upstream branch",
		SetUpstreamMessage:                  "Are you sure you want to set the upstream branch of '{{.checkedOut}}' to '{{.selected}}'",
		LcEditRemote:                        "Remote를 수정",
		LcTagCommit:                         "tag commit",
		TagMenuTitle:                        "태그 작성",
		TagNameTitle:                        "태그 이름:",
		TagMessageTitle:                     "태그 메시지: ",
		LcAnnotatedTag:                      "annotated tag",
		LcLightweightTag:                    "lightweight tag",
		LcDeleteTag:                         "태그 삭제",
		DeleteTagTitle:                      "태그 삭제",
		DeleteTagPrompt:                     "정말로 태그 '{{.tagName}}' 를 삭제하시겠습니까?",
		PushTagTitle:                        "원격에 태그 '{{.tagName}}' 를 푸시",
		LcPushTag:                           "태그를 push",
		LcCreateTag:                         "태그를 생성",
		CreateTagTitle:                      "태그 이름:",
		LcFetchRemote:                       "원격을 업데이트",
		FetchingRemoteStatus:                "원격을 업데이트 중",
		LcCheckoutCommit:                    "커밋을 체크아웃",
		SureCheckoutThisCommit:              "정말로 선택한 커밋을 체크아웃 하시겠습니까?",
		LcGitFlowOptions:                    "git-flow 옵션 보기",
		NotAGitFlowBranch:                   "This does not seem to be a git flow branch",
		NewGitFlowBranchPrompt:              "new {{.branchType}} name:",
		IgnoreTracked:                       "Ignore tracked file",
		IgnoreTrackedPrompt:                 "Are you sure you want to ignore a tracked file?",
		LcViewResetToUpstreamOptions:        "view upstream reset options",
		LcNextScreenMode:                    "다음 스크린 모드 (normal/half/fullscreen)",
		LcPrevScreenMode:                    "이전 스크린 모드",
		LcStartSearch:                       "검색 시작",
		Panel:                               "패널",
		Keybindings:                         "키 바인딩",
		LcRenameBranch:                      "브랜치 이름 변경",
		NewBranchNamePrompt:                 "새로운 브랜치 이름 입력",
		RenameBranchWarning:                 "This branch is tracking a remote. This action will only rename the local branch name, not the name of the remote branch. Continue?",
		LcOpenMenu:                          "매뉴 열기",
		LcResetCherryPick:                   "reset cherry-picked (copied) commits selection",
		LcNextTab:                           "이전 탭",
		LcPrevTab:                           "다음 탭",
		LcCantUndoWhileRebasing:             "리베이스중에는 되돌릴 수 없습니다.",
		LcCantRedoWhileRebasing:             "리베이스중에는 다시 실행할 수 없습니다.",
		MustStashWarning:                    "Pulling a patch out into the index requires stashing and unstashing your changes. If something goes wrong, you'll be able to access your files from the stash. Continue?",
		MustStashTitle:                      "Must stash",
		ConfirmationTitle:                   "확인 패널",
		LcPrevPage:                          "이전 페이지",
		LcNextPage:                          "다음 페이지",
		LcGotoTop:                           "맨 위로 스크롤 ",
		LcGotoBottom:                        "맨 아래로 스크롤 ",
		LcFilteringBy:                       "filtering by",
		ResetInParentheses:                  "(reset)",
		LcOpenFilteringMenu:                 "view filter-by-path options",
		LcFilterBy:                          "filter by",
		LcExitFilterMode:                    "stop filtering by path",
		LcFilterPathOption:                  "enter path to filter by",
		EnterFileName:                       "Enter path:",
		FilteringMenuTitle:                  "Filtering",
		MustExitFilterModeTitle:             "Command not available",
		MustExitFilterModePrompt:            "Command not available in filtered mode. Exit filtered mode?",
		LcDiff:                              "Diff",
		LcEnterRefToDiff:                    "enter ref to diff",
		LcEnteRefName:                       "ref 입력:",
		LcExitDiffMode:                      "Diff 모드 종료",
		DiffingMenuTitle:                    "Diff",
		LcSwapDiff:                          "reverse diff direction",
		LcOpenDiffingMenu:                   "Diff 메뉴 열기",
		// the actual view is the extras view which I intend to give more tabs in future but for now we'll only mention the command log part
		LcOpenExtrasMenu:                    "명령어 로그 메뉴 열기",
		LcShowingGitDiff:                    "showing output for:",
		LcCommitDiff:                        "커밋의 iff",
		LcCopyCommitShaToClipboard:          "커밋 SHA를 클립보드에 복사",
		LcCommitSha:                         "커밋 SHA",
		LcCommitURL:                         "커밋 URL",
		LcCopyCommitMessageToClipboard:      "커밋 메시지를 클립보드에 복사",
		LcCommitMessage:                     "커밋 메시지",
		LcCommitAuthor:                      "커밋 작성자",
		LcCopyCommitAttributeToClipboard:    "커밋 attribute 복사",
		LcCopyBranchNameToClipboard:         "브랜치명을 클립보드에 복사",
		LcCopyFileNameToClipboard:           "파일명을 클립보드에 복사",
		LcCopyCommitFileNameToClipboard:     "커밋한 파일명을 클립보드에 복사",
		LcCopySelectedTexToClipboard:        "선택한 텍스트를 클립보드에 복사",
		LcCommitPrefixPatternError:          "Error in commitPrefix pattern",
		NoFilesStagedTitle:                  "파일이 Staged 되지 않았습니다.",
		NoFilesStagedPrompt:                 "파일이 Staged 되지 않았습니다. 모든 파일을 커밋하시겠습니까?",
		BranchNotFoundTitle:                 "브랜치를 찾을 수 없습니다.",
		BranchNotFoundPrompt:                "브랜치를 찾을 수 없습니다. 새로운 브랜치를 생성합니다.",
		UnstageLinesTitle:                   "선택한 라인을 unstaged",
		UnstageLinesPrompt:                  "정말로 선택한 라인을 삭제 (git reset) 하시겠습니까? 이 조작은 취소할 수 없습니다.\n이 경고를 비활성화하려면 설정 파일의 'gui.skipUnstageLineWarning' 를 true 로 설정하세요.",
		LcCreateNewBranchFromCommit:         "커밋에서 새 브랜치를 만듭니다.",
		LcBuildingPatch:                     "building patch",
		LcViewCommits:                       "커밋 보기",
		MinGitVersionError:                  "lazygit 실행을 위해서는 Git 2.0 이후의 버전(2014년 이후의)이 필요합니다. Git를 갱신해 주세요.아니면 lazygit이 이전 버전과 더 잘 호환되도록 https://github.com/jesseduffield/lazygit/issues 에 issue를 작성해 주세요.",
		LcRunningCustomCommandStatus:        "커스텀 명령어 실행",
		LcSubmoduleStashAndReset:            "stash uncommitted submodule changes and update",
		LcAndResetSubmodules:                "and reset submodules",
		LcEnterSubmodule:                    "서브모듈 열기",
		LcCopySubmoduleNameToClipboard:      "서브모듈 이름을 클립보드에 복사",
		RemoveSubmodule:                     "서브모듈 삭제",
		LcRemoveSubmodule:                   "서브모듈 삭제",
		RemoveSubmodulePrompt:               "정말로 서브모듈 '%s'및 해당 디렉토리를 제거하시겠습니까? 이것은 되돌릴 수 없습니다.",
		LcResettingSubmoduleStatus:          "서브모듈를 리셋",
		LcNewSubmoduleName:                  "새로운 서브모듈이름 :",
		LcNewSubmoduleUrl:                   "새로운 서브모듈의 URL:",
		LcNewSubmodulePath:                  "새로운 서브모듈의 경로",
		LcAddSubmodule:                      "새로운 서브모듈 추가",
		LcAddingSubmoduleStatus:             "새로운 서브모듈 추가",
		LcUpdateSubmoduleUrl:                "서브모듈 '%s' 의 URL을 업데이트",
		LcUpdatingSubmoduleUrlStatus:        "updating URL",
		LcEditSubmoduleUrl:                  "서브모듈의 URL을 수정",
		LcInitializingSubmoduleStatus:       "서브모듈 초기화",
		LcInitSubmodule:                     "서브모듈 초기화",
		LcSubmoduleUpdate:                   "서브모듈 업데이트",
		LcUpdatingSubmoduleStatus:           "서브모듈 업데이트",
		LcBulkInitSubmodules:                "서브모듈 일괄 초기화",
		LcBulkUpdateSubmodules:              "서브모듈 일괄 업데이트",
		LcBulkDeinitSubmodules:              "bulk deinit submodules",
		LcViewBulkSubmoduleOptions:          "view bulk submodule options",
		LcBulkSubmoduleOptions:              "bulk submodule options",
		LcRunningCommand:                    "running command",
		SubCommitsTitle:                     "Sub-commits",
		SubmodulesTitle:                     "서브모듈",
		NavigationTitle:                     "List Panel Navigation",
		SuggestionsCheatsheetTitle:          "추천",
		SuggestionsTitle:                    "추천 (press %s to focus)",
		ExtrasTitle:                         "명령어 로그",
		PushingTagStatus:                    "pushing tag",
		PullRequestURLCopiedToClipboard:     "풀 리퀘스트의 URL을 클립보드에 복사했습니다.",
		CommitDiffCopiedToClipboard:         "커밋의 Diff를 클립보드에 복사했습니다.",
		CommitSHACopiedToClipboard:          "커밋의 SHA를 클립보드에 복사했습니다.",
		CommitURLCopiedToClipboard:          "커밋의 URL를 클립보드에 복사했습니다.",
		CommitMessageCopiedToClipboard:      "커밋 메시지를 클립보드에 복사했습니다.",
		CommitAuthorCopiedToClipboard:       "커밋 작성자를 클립보드에 복사했습니다.",
		LcCopiedToClipboard:                 "클립보드에 복사했습니다.",
		ErrCannotEditDirectory:              "디렉토리는 편집할 수 없습니다.",
		ErrStageDirWithInlineMergeConflicts: "병합 충돌이 발생한 파일을 포함하는 디렉토리는 Staged/untaged할 수 없습니다. 병합 충돌을 먼저 해결하세요.",
		ErrRepositoryMovedOrDeleted:         "저장소를 찾을 수 없습니다. 이미 삭제되었거나 이동되었을 가능성이 있습니다. ¯\\_(ツ)_/¯",
		CommandLog:                          "명령어 로그",
		ToggleShowCommandLog:                "명령어 로그 표시 여부 전환",
		FocusCommandLog:                     "명령어 로그에 포커스",
		CommandLogHeader:                    "명령어 로그표시 여부는 '%s' 으로 전환할 수 있습니다.\n",
		RandomTip:                           "랜덤 Tip",
		SelectParentCommitForMerge:          "병합을 위한 상위 커밋 선택",
		ToggleWhitespaceInDiffView:          "공백문자를 Diff 뷰에서 표시 여부 전환",
		IgnoringWhitespaceInDiffView:        "공백문자를 Diff 뷰에서 무시",
		ShowingWhitespaceInDiffView:         "공백문자를 Diff 뷰에서 표시",
		IncreaseContextInDiffView:           "diff 보기의 변경 사항 주위에 표시되는 컨텍스트의 크기를 늘리기",
		DecreaseContextInDiffView:           "diff 보기의 변경 사항 주위에 표시되는 컨텍스트 크기 줄이기",
		CreatePullRequest:                   "풀 리퀘스트 생성",
		CreatePullRequestOptions:            "풀 리퀘스트 생성 옵션",
		LcCreatePullRequestOptions:          "풀 리퀘스트 생성 옵션",
		LcDefaultBranch:                     "기본 브랜치",
		LcSelectBranch:                      "브랜치를 선택",
		SelectConfigFile:                    "설정파일 선택",
		NoConfigFileFoundErr:                "설정 파일을 찾지 못했습니다.",
		LcLoadingFileSuggestions:            "파일 제안 로딩 중",
		LcLoadingCommits:                    "커밋 로딩",
		MustSpecifyOriginError:              "Must specify a remote if specifying a branch",
		GitOutput:                           "Git output:",
		GitCommandFailed:                    "Git command failed. Check command log for details (open with %s)",
		AbortTitle:                          "%s 중지",
		AbortPrompt:                         "정말로 실행중인 %s 를 중지할까요?",
		LcOpenLogMenu:                       "로그 메뉴 열기",
		LogMenuTitle:                        "커밋 로그 옵션",
		ToggleShowGitGraphAll:               "toggle show whole git graph (pass the `--all` flag to `git log`)",
		ShowGitGraph:                        "커밋 그래프 표시",
		SortCommits:                         "커밋 정렬",
		CantChangeContextSizeError:          "Cannot change context while in patch building mode because we were too lazy to support it when releasing the feature. If you really want it, please let us know!",
		LcOpenCommitInBrowser:               "브라우저에서 커밋 열기",
		LcViewBisectOptions:                 "bisect 옵션 보기",
		ConfirmRevertCommit:                 "Are you sure you want to revert {{.selectedCommit}}?",
		RewordInEditorTitle:                 "커밋 메시지를 에디터에서 수정",
		RewordInEditorPrompt:                "Are you sure you want to reword this commit in your editor?",
		HardResetAutostashPrompt:            "Are you sure you want to hard reset to '%s'? An auto-stash will be performed if necessary.",
		CheckoutPrompt:                      "Are you sure you want to checkout '%s'?",
		UpstreamGone:                        "(upstream gone)",
		Actions: Actions{
			// TODO: combine this with the original keybinding descriptions (those are all in lowercase atm)
			CheckoutCommit:                    "커밋 체크아웃",
			CheckoutTag:                       "태그 체크아웃",
			CheckoutBranch:                    "브랜치 체크아웃",
			ForceCheckoutBranch:               "브랜치 Force 체크아웃",
			DeleteBranch:                      "브랜치 삭제",
			Merge:                             "병합",
			RebaseBranch:                      "브랜치 리베이스",
			RenameBranch:                      "브랜치 이름 변경",
			SetUnsetUpstream:                  "Set/unset upstream",
			CreateBranch:                      "브랜치 생성",
			CherryPick:                        "(Cherry-pick) 커밋 붙여넣기",
			CheckoutFile:                      "체크아웃 파일",
			DiscardOldFileChange:              "Discard old file change",
			SquashCommitDown:                  "Squash commit down",
			FixupCommit:                       "커밋 Fixup",
			RewordCommit:                      "커밋 Reword",
			DropCommit:                        "커밋 Drop",
			EditCommit:                        "커밋 수정",
			AmendCommit:                       "커밋 Amend",
			ResetCommitAuthor:                 "커밋 작성자 Reset",
			RevertCommit:                      "커밋 Revert",
			CreateFixupCommit:                 "fixup 커밋 생성",
			SquashAllAboveFixupCommits:        "Squash all above fixup commits",
			CreateLightweightTag:              "Create lightweight tag",
			CreateAnnotatedTag:                "Create annotated tag",
			CopyCommitMessageToClipboard:      "커밋 메시지를 클립보드에 복사",
			CopyCommitDiffToClipboard:         "커밋 diff를 클립보드에 복사",
			CopyCommitSHAToClipboard:          "커밋 SHA를 클립보드에 복사",
			CopyCommitURLToClipboard:          "커밋 URL를 클립보드에 복사",
			CopyCommitAuthorToClipboard:       "커밋 작성자를 클립보드에 복사",
			CopyCommitAttributeToClipboard:    "클립보드에 복사",
			MoveCommitUp:                      "Move commit up",
			MoveCommitDown:                    "Move commit down",
			CustomCommand:                     "Custom command",
			DiscardAllChangesInDirectory:      "Discard all changes in directory",
			DiscardUnstagedChangesInDirectory: "Discard unstaged changes in directory",
			DiscardAllChangesInFile:           "Discard all changes in file",
			DiscardAllUnstagedChangesInFile:   "Discard all unstaged changes in file",
			StageFile:                         "Stage file",
			StageResolvedFiles:                "Stage files whose merge conflicts were resolved",
			UnstageFile:                       "Unstage file",
			UnstageAllFiles:                   "Unstage all files",
			StageAllFiles:                     "Stage all files",
			IgnoreFile:                        "Ignore file",
			Commit:                            "커밋",
			EditFile:                          "파일 수정",
			Push:                              "푸시",
			Pull:                              "업데이트(Pull)",
			OpenFile:                          "파일 열기",
			StashAllChanges:                   "Stash all changes",
			StashAllChangesKeepIndex:          "Stash all changes and keep index",
			StashStagedChanges:                "Stash staged changes",
			StashUnstagedChanges:              "Stash unstaged changes",
			GitFlowFinish:                     "Git flow finish",
			GitFlowStart:                      "Git Flow start",
			CopyToClipboard:                   "Copy to clipboard",
			CopySelectedTextToClipboard:       "Copy selected text to clipboard",
			RemovePatchFromCommit:             "Remove patch from commit",
			MovePatchToSelectedCommit:         "Move patch to selected commit",
			MovePatchIntoIndex:                "Move patch into index",
			MovePatchIntoNewCommit:            "Move patch into new commit",
			DeleteRemoteBranch:                "Delete remote branch",
			SetBranchUpstream:                 "Set branch upstream",
			AddRemote:                         "Add remote",
			RemoveRemote:                      "Remove remote",
			UpdateRemote:                      "Update remote",
			ApplyPatch:                        "Apply patch",
			Stash:                             "Stash",
			RemoveSubmodule:                   "서브모듈 삭제",
			ResetSubmodule:                    "서브모듈 Reset",
			AddSubmodule:                      "서브모듈 추가",
			UpdateSubmoduleUrl:                "서브모듈 URL 업데이트",
			InitialiseSubmodule:               "서브모듈 초기화",
			BulkInitialiseSubmodules:          "Bulk initialise submodules",
			BulkUpdateSubmodules:              "Bulk update submodules",
			BulkDeinitialiseSubmodules:        "Bulk deinitialise submodules",
			UpdateSubmodule:                   "서브모듈 업데이트",
			DeleteTag:                         "태그 삭제",
			PushTag:                           "태그 푸시g",
			NukeWorkingTree:                   "Nuke working tree",
			DiscardUnstagedFileChanges:        "unstaged 파일 변경사항 버리기",
			RemoveUntrackedFiles:              "untracked 파일 삭제",
			RemoveStagedFiles:                 "staged 파일 삭제",
			SoftReset:                         "Soft reset",
			MixedReset:                        "Mixed reset",
			HardReset:                         "Hard reset",
			FastForwardBranch:                 "Fast forward branch",
			Undo:                              "되돌리기",
			Redo:                              "다시 실행",
			CopyPullRequestURL:                "풀 리퀘스트 URL 복사",
			OpenMergeTool:                     "병합 도구 열기",
			OpenCommitInBrowser:               "브라우저에서 커밋 열기",
			OpenPullRequest:                   "브라우저에서 풀 리퀘스트 열기",
			StartBisect:                       "Start bisect",
			ResetBisect:                       "Reset bisect",
			BisectSkip:                        "Bisect skip",
			BisectMark:                        "Bisect mark",
		},
		Bisect: Bisect{
			Mark:                        "mark %s as %s",
			MarkStart:                   "mark %s as %s (start bisect)",
			Skip:                        "%s 를 스킵",
			ResetTitle:                  "'git bisect' 를 리셋",
			ResetPrompt:                 "정말로 'git bisect' 를 리셋하시겠습니까?",
			ResetOption:                 "bisect를 리셋",
			BisectMenuTitle:             "bisect",
			CompleteTitle:               "Bisect 완료",
			CompletePrompt:              "Bisect complete! The following commit introduced the change:\n\n%s\n\nDo you want to reset 'git bisect' now?",
			CompletePromptIndeterminate: "Bisect complete! Some commits were skipped, so any of the following commits may have introduced the change:\n\n%s\n\nDo you want to reset 'git bisect' now?",
		},
	}
}
