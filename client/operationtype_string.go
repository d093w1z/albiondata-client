// Code generated by "stringer -type=OperationType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opUnused-0]
	_ = x[opPing-1]
	_ = x[opJoin-2]
	_ = x[opVersionedOperation-3]
	_ = x[opCreateAccount-4]
	_ = x[opLogin-5]
	_ = x[opCreateGuestAccount-6]
	_ = x[opSendCrashLog-7]
	_ = x[opSendTraceRoute-8]
	_ = x[opSendVfxStats-9]
	_ = x[opSendGamePingInfo-10]
	_ = x[opCreateCharacter-11]
	_ = x[opDeleteCharacter-12]
	_ = x[opSelectCharacter-13]
	_ = x[opAcceptPopups-14]
	_ = x[opRedeemKeycode-15]
	_ = x[opGetGameServerByCluster-16]
	_ = x[opGetShopPurchaseUrl-17]
	_ = x[opGetReferralSeasonDetails-18]
	_ = x[opGetReferralLink-19]
	_ = x[opGetShopTilesForCategory-20]
	_ = x[opMove-21]
	_ = x[opAttackStart-22]
	_ = x[opCastStart-23]
	_ = x[opCastCancel-24]
	_ = x[opTerminateToggleSpell-25]
	_ = x[opChannelingCancel-26]
	_ = x[opAttackBuildingStart-27]
	_ = x[opInventoryDestroyItem-28]
	_ = x[opInventoryMoveItem-29]
	_ = x[opInventoryRecoverItem-30]
	_ = x[opInventoryRecoverAllItems-31]
	_ = x[opInventorySplitStack-32]
	_ = x[opInventorySplitStackInto-33]
	_ = x[opGetClusterData-34]
	_ = x[opChangeCluster-35]
	_ = x[opConsoleCommand-36]
	_ = x[opChatMessage-37]
	_ = x[opReportClientError-38]
	_ = x[opRegisterToObject-39]
	_ = x[opUnRegisterFromObject-40]
	_ = x[opCraftBuildingChangeSettings-41]
	_ = x[opCraftBuildingTakeMoney-42]
	_ = x[opRepairBuildingChangeSettings-43]
	_ = x[opRepairBuildingTakeMoney-44]
	_ = x[opActionBuildingChangeSettings-45]
	_ = x[opHarvestStart-46]
	_ = x[opHarvestCancel-47]
	_ = x[opTakeSilver-48]
	_ = x[opActionOnBuildingStart-49]
	_ = x[opActionOnBuildingCancel-50]
	_ = x[opInstallResourceStart-51]
	_ = x[opInstallResourceCancel-52]
	_ = x[opInstallSilver-53]
	_ = x[opBuildingFillNutrition-54]
	_ = x[opBuildingChangeRenovationState-55]
	_ = x[opBuildingBuySkin-56]
	_ = x[opBuildingClaim-57]
	_ = x[opBuildingGiveup-58]
	_ = x[opBuildingNutritionSilverStorageDeposit-59]
	_ = x[opBuildingNutritionSilverStorageWithdraw-60]
	_ = x[opBuildingNutritionSilverRewardSet-61]
	_ = x[opConstructionSiteCreate-62]
	_ = x[opPlaceableObjectPlace-63]
	_ = x[opPlaceableObjectPlaceCancel-64]
	_ = x[opPlaceableObjectPickup-65]
	_ = x[opFurnitureObjectUse-66]
	_ = x[opFarmableHarvest-67]
	_ = x[opFarmableFinishGrownItem-68]
	_ = x[opFarmableDestroy-69]
	_ = x[opFarmableGetProduct-70]
	_ = x[opFarmableFill-71]
	_ = x[opTearDownConstructionSite-72]
	_ = x[opPlaceholder1-73]
	_ = x[opAuctionCreateOffer-74]
	_ = x[opAuctionCreateRequest-75]
	_ = x[opAuctionGetOffers-76]
	_ = x[opAuctionGetRequests-77]
	_ = x[opAuctionBuyOffer-78]
	_ = x[opAuctionAbortAuction-79]
	_ = x[opAuctionModifyAuction-80]
	_ = x[opAuctionAbortOffer-81]
	_ = x[opAuctionAbortRequest-82]
	_ = x[opAuctionSellRequest-83]
	_ = x[opAuctionGetFinishedAuctions-84]
	_ = x[opAuctionGetFinishedAuctionsCount-85]
	_ = x[opAuctionFetchAuction-86]
	_ = x[opAuctionGetMyOpenOffers-87]
	_ = x[opAuctionGetMyOpenRequests-88]
	_ = x[opAuctionGetMyOpenAuctions-89]
	_ = x[opAuctionGetItemAverageStats-90]
	_ = x[opAuctionGetItemAverageValue-91]
	_ = x[opContainerOpen-92]
	_ = x[opContainerClose-93]
	_ = x[opContainerManageSubContainer-94]
	_ = x[opRespawn-95]
	_ = x[opSuicide-96]
	_ = x[opJoinGuild-97]
	_ = x[opLeaveGuild-98]
	_ = x[opCreateGuild-99]
	_ = x[opInviteToGuild-100]
	_ = x[opDeclineGuildInvitation-101]
	_ = x[opKickFromGuild-102]
	_ = x[opInstantJoinGuild-103]
	_ = x[opDuellingChallengePlayer-104]
	_ = x[opDuellingAcceptChallenge-105]
	_ = x[opDuellingDenyChallenge-106]
	_ = x[opChangeClusterTax-107]
	_ = x[opClaimTerritory-108]
	_ = x[opGiveUpTerritory-109]
	_ = x[opChangeTerritoryAccessRights-110]
	_ = x[opGetMonolithInfo-111]
	_ = x[opGetClaimInfo-112]
	_ = x[opGetAttackInfo-113]
	_ = x[opGetTerritorySeasonPoints-114]
	_ = x[opGetAttackSchedule-115]
	_ = x[opGetMatches-116]
	_ = x[opGetMatchDetails-117]
	_ = x[opJoinMatch-118]
	_ = x[opLeaveMatch-119]
	_ = x[opChangeChatSettings-120]
	_ = x[opLogoutStart-121]
	_ = x[opLogoutCancel-122]
	_ = x[opClaimOrbStart-123]
	_ = x[opClaimOrbCancel-124]
	_ = x[opMatchLootChestOpeningStart-125]
	_ = x[opMatchLootChestOpeningCancel-126]
	_ = x[opDepositToGuildAccount-127]
	_ = x[opWithdrawalFromAccount-128]
	_ = x[opChangeGuildPayUpkeepFlag-129]
	_ = x[opChangeGuildTax-130]
	_ = x[opGetMyTerritories-131]
	_ = x[opMorganaCommand-132]
	_ = x[opGetServerInfo-133]
	_ = x[opSubscribeToCluster-134]
	_ = x[opAnswerMercenaryInvitation-135]
	_ = x[opGetCharacterEquipment-136]
	_ = x[opGetCharacterSteamAchievements-137]
	_ = x[opGetCharacterStats-138]
	_ = x[opGetKillHistoryDetails-139]
	_ = x[opLearnMasteryLevel-140]
	_ = x[opReSpecAchievement-141]
	_ = x[opChangeAvatar-142]
	_ = x[opGetRankings-143]
	_ = x[opGetRank-144]
	_ = x[opGetGvgSeasonRankings-145]
	_ = x[opGetGvgSeasonRank-146]
	_ = x[opGetGvgSeasonHistoryRankings-147]
	_ = x[opGetGvgSeasonGuildMemberHistory-148]
	_ = x[opKickFromGvGMatch-149]
	_ = x[opGetCrystalLeagueDailySeasonPoints-150]
	_ = x[opGetChestLogs-151]
	_ = x[opGetAccessRightLogs-152]
	_ = x[opGetGuildAccountLogs-153]
	_ = x[opGetGuildAccountLogsLargeAmount-154]
	_ = x[opInviteToPlayerTrade-155]
	_ = x[opPlayerTradeCancel-156]
	_ = x[opPlayerTradeInvitationAccept-157]
	_ = x[opPlayerTradeAddItem-158]
	_ = x[opPlayerTradeRemoveItem-159]
	_ = x[opPlayerTradeAcceptTrade-160]
	_ = x[opPlayerTradeSetSilverOrGold-161]
	_ = x[opSendMiniMapPing-162]
	_ = x[opStuck-163]
	_ = x[opBuyRealEstate-164]
	_ = x[opClaimRealEstate-165]
	_ = x[opGiveUpRealEstate-166]
	_ = x[opChangeRealEstateOutline-167]
	_ = x[opGetMailInfos-168]
	_ = x[opGetMailCount-169]
	_ = x[opReadMail-170]
	_ = x[opSendNewMail-171]
	_ = x[opDeleteMail-172]
	_ = x[opMarkMailUnread-173]
	_ = x[opClaimAttachmentFromMail-174]
	_ = x[opApplyToGuild-175]
	_ = x[opAnswerGuildApplication-176]
	_ = x[opRequestGuildFinderFilteredList-177]
	_ = x[opUpdateGuildRecruitmentInfo-178]
	_ = x[opRequestGuildRecruitmentInfo-179]
	_ = x[opRequestGuildFinderNameSearch-180]
	_ = x[opRequestGuildFinderRecommendedList-181]
	_ = x[opRegisterChatPeer-182]
	_ = x[opSendChatMessage-183]
	_ = x[opSendModeratorMessage-184]
	_ = x[opJoinChatChannel-185]
	_ = x[opLeaveChatChannel-186]
	_ = x[opSendWhisperMessage-187]
	_ = x[opSay-188]
	_ = x[opPlayEmote-189]
	_ = x[opStopEmote-190]
	_ = x[opGetClusterMapInfo-191]
	_ = x[opAccessRightsChangeSettings-192]
	_ = x[opMount-193]
	_ = x[opMountCancel-194]
	_ = x[opBuyJourney-195]
	_ = x[opSetSaleStatusForEstate-196]
	_ = x[opResolveGuildOrPlayerName-197]
	_ = x[opGetRespawnInfos-198]
	_ = x[opMakeHome-199]
	_ = x[opLeaveHome-200]
	_ = x[opResurrectionReply-201]
	_ = x[opAllianceCreate-202]
	_ = x[opAllianceDisband-203]
	_ = x[opAllianceGetMemberInfos-204]
	_ = x[opAllianceInvite-205]
	_ = x[opAllianceAnswerInvitation-206]
	_ = x[opAllianceCancelInvitation-207]
	_ = x[opAllianceKickGuild-208]
	_ = x[opAllianceLeave-209]
	_ = x[opAllianceChangeGoldPaymentFlag-210]
	_ = x[opAllianceGetDetailInfo-211]
	_ = x[opGetIslandInfos-212]
	_ = x[opAbandonMyIsland-213]
	_ = x[opBuyMyIsland-214]
	_ = x[opBuyGuildIsland-215]
	_ = x[opAbandonGuildIsland-216]
	_ = x[opUpgradeMyIsland-217]
	_ = x[opUpgradeGuildIsland-218]
	_ = x[opMoveMyIsland-219]
	_ = x[opMoveGuildIsland-220]
	_ = x[opTerritoryFillNutrition-221]
	_ = x[opTeleportBack-222]
	_ = x[opPartyInvitePlayer-223]
	_ = x[opPartyRequestJoin-224]
	_ = x[opPartyAnswerInvitation-225]
	_ = x[opPartyAnswerJoinRequest-226]
	_ = x[opPartyLeave-227]
	_ = x[opPartyKickPlayer-228]
	_ = x[opPartyMakeLeader-229]
	_ = x[opPartyChangeLootSetting-230]
	_ = x[opPartyMarkObject-231]
	_ = x[opPartySetRole-232]
	_ = x[opSetGuildCodex-233]
	_ = x[opExitEnterStart-234]
	_ = x[opExitEnterCancel-235]
	_ = x[opQuestGiverRequest-236]
	_ = x[opGoldMarketGetBuyOffer-237]
	_ = x[opGoldMarketGetBuyOfferFromSilver-238]
	_ = x[opGoldMarketGetSellOffer-239]
	_ = x[opGoldMarketGetSellOfferFromSilver-240]
	_ = x[opGoldMarketBuyGold-241]
	_ = x[opGoldMarketSellGold-242]
	_ = x[opGoldMarketCreateSellOrder-243]
	_ = x[opGoldMarketCreateBuyOrder-244]
	_ = x[opGoldMarketGetInfos-245]
	_ = x[opGoldMarketCancelOrder-246]
	_ = x[opGoldMarketGetAverageInfo-247]
	_ = x[opTreasureChestUsingStart-248]
	_ = x[opTreasureChestUsingCancel-249]
	_ = x[opUseLootChest-250]
	_ = x[opUseShrine-251]
	_ = x[opUseHellgateShrine-252]
	_ = x[opUseSiegeBanner-253]
	_ = x[opGetSiegeBannerInfo-254]
	_ = x[opLaborerStartJob-255]
	_ = x[opLaborerTakeJobLoot-256]
	_ = x[opLaborerDismiss-257]
	_ = x[opLaborerMove-258]
	_ = x[opLaborerBuyItem-259]
	_ = x[opLaborerUpgrade-260]
	_ = x[opBuyPremium-261]
	_ = x[opRealEstateGetAuctionData-262]
	_ = x[opRealEstateBidOnAuction-263]
	_ = x[opFriendInvite-264]
	_ = x[opFriendAnswerInvitation-265]
	_ = x[opFriendCancelnvitation-266]
	_ = x[opFriendRemove-267]
	_ = x[opInventoryStack-268]
	_ = x[opInventorySort-269]
	_ = x[opInventoryDropAll-270]
	_ = x[opInventoryAddToStacks-271]
	_ = x[opEquipmentItemChangeSpell-272]
	_ = x[opExpeditionRegister-273]
	_ = x[opExpeditionRegisterCancel-274]
	_ = x[opJoinExpedition-275]
	_ = x[opDeclineExpeditionInvitation-276]
	_ = x[opVoteStart-277]
	_ = x[opVoteDoVote-278]
	_ = x[opRatingDoRate-279]
	_ = x[opEnteringExpeditionStart-280]
	_ = x[opEnteringExpeditionCancel-281]
	_ = x[opActivateExpeditionCheckPoint-282]
	_ = x[opArenaRegister-283]
	_ = x[opArenaAddInvite-284]
	_ = x[opArenaRegisterCancel-285]
	_ = x[opArenaLeave-286]
	_ = x[opJoinArenaMatch-287]
	_ = x[opDeclineArenaInvitation-288]
	_ = x[opEnteringArenaStart-289]
	_ = x[opEnteringArenaCancel-290]
	_ = x[opArenaCustomMatch-291]
	_ = x[opUpdateCharacterStatement-292]
	_ = x[opBoostFarmable-293]
	_ = x[opGetStrikeHistory-294]
	_ = x[opUseFunction-295]
	_ = x[opUsePortalEntrance-296]
	_ = x[opResetPortalBinding-297]
	_ = x[opQueryPortalBinding-298]
	_ = x[opClaimPaymentTransaction-299]
	_ = x[opChangeUseFlag-300]
	_ = x[opClientPerformanceStats-301]
	_ = x[opExtendedHardwareStats-302]
	_ = x[opClientLowMemoryWarning-303]
	_ = x[opTerritoryClaimStart-304]
	_ = x[opTerritoryClaimCancel-305]
	_ = x[opDeliverCarriableObjectStart-306]
	_ = x[opDeliverCarriableObjectCancel-307]
	_ = x[opTerritoryUpgradeWithPowerCrystal-308]
	_ = x[opRequestAppStoreProducts-309]
	_ = x[opVerifyProductPurchase-310]
	_ = x[opQueryGuildPlayerStats-311]
	_ = x[opQueryAllianceGuildStats-312]
	_ = x[opTrackAchievements-313]
	_ = x[opSetAchievementsAutoLearn-314]
	_ = x[opDepositItemToGuildCurrency-315]
	_ = x[opWithdrawalItemFromGuildCurrency-316]
	_ = x[opAuctionSellSpecificItemRequest-317]
	_ = x[opFishingStart-318]
	_ = x[opFishingCasting-319]
	_ = x[opFishingCast-320]
	_ = x[opFishingCatch-321]
	_ = x[opFishingPull-322]
	_ = x[opFishingGiveLine-323]
	_ = x[opFishingFinish-324]
	_ = x[opFishingCancel-325]
	_ = x[opCreateGuildAccessTag-326]
	_ = x[opDeleteGuildAccessTag-327]
	_ = x[opRenameGuildAccessTag-328]
	_ = x[opFlagGuildAccessTagGuildPermission-329]
	_ = x[opAssignGuildAccessTag-330]
	_ = x[opRemoveGuildAccessTagFromPlayer-331]
	_ = x[opModifyGuildAccessTagEditors-332]
	_ = x[opRequestPublicAccessTags-333]
	_ = x[opChangeAccessTagPublicFlag-334]
	_ = x[opUpdateGuildAccessTag-335]
	_ = x[opSteamStartMicrotransaction-336]
	_ = x[opSteamFinishMicrotransaction-337]
	_ = x[opSteamIdHasActiveAccount-338]
	_ = x[opCheckEmailAccountState-339]
	_ = x[opLinkAccountToSteamId-340]
	_ = x[opInAppConfirmPaymentGooglePlay-341]
	_ = x[opInAppConfirmPaymentAppleAppStore-342]
	_ = x[opInAppPurchaseRequest-343]
	_ = x[opInAppPurchaseFailed-344]
	_ = x[opCharacterSubscriptionInfo-345]
	_ = x[opAccountSubscriptionInfo-346]
	_ = x[opBuyGvgSeasonBooster-347]
	_ = x[opChangeFlaggingPrepare-348]
	_ = x[opOverCharge-349]
	_ = x[opOverChargeEnd-350]
	_ = x[opRequestTrusted-351]
	_ = x[opChangeGuildLogo-352]
	_ = x[opPartyFinderRegisterForUpdates-353]
	_ = x[opPartyFinderUnregisterForUpdates-354]
	_ = x[opPartyFinderEnlistNewPartySearch-355]
	_ = x[opPartyFinderDeletePartySearch-356]
	_ = x[opPartyFinderChangePartySearch-357]
	_ = x[opPartyFinderChangeRole-358]
	_ = x[opPartyFinderApplyForGroup-359]
	_ = x[opPartyFinderAcceptOrDeclineApplyForGroup-360]
	_ = x[opPartyFinderGetEquipmentSnapshot-361]
	_ = x[opPartyFinderRegisterApplicants-362]
	_ = x[opPartyFinderUnregisterApplicants-363]
	_ = x[opPartyFinderFulltextSearch-364]
	_ = x[opPartyFinderRequestEquipmentSnapshot-365]
	_ = x[opGetPersonalSeasonTrackerData-366]
	_ = x[opGetPersonalSeasonPastRewardData-367]
	_ = x[opUseConsumableFromInventory-368]
	_ = x[opClaimPersonalSeasonReward-369]
	_ = x[opEasyAntiCheatMessageToServer-370]
	_ = x[opXignCodeMessageToServer-371]
	_ = x[opBattlEyeMessageToServer-372]
	_ = x[opSetNextTutorialState-373]
	_ = x[opAddPlayerToMuteList-374]
	_ = x[opRemovePlayerFromMuteList-375]
	_ = x[opProductShopUserEvent-376]
	_ = x[opGetVanityUnlocks-377]
	_ = x[opBuyVanityUnlocks-378]
	_ = x[opGetMountSkins-379]
	_ = x[opSetMountSkin-380]
	_ = x[opSetWardrobe-381]
	_ = x[opChangeCustomization-382]
	_ = x[opChangePlayerIslandData-383]
	_ = x[opGetGuildChallengePoints-384]
	_ = x[opSmartQueueJoin-385]
	_ = x[opSmartQueueLeave-386]
	_ = x[opSmartQueueSelectSpawnCluster-387]
	_ = x[opUpgradeHideout-388]
	_ = x[opInitHideoutAttackStart-389]
	_ = x[opInitHideoutAttackCancel-390]
	_ = x[opHideoutFillNutrition-391]
	_ = x[opHideoutGetInfo-392]
	_ = x[opHideoutGetOwnerInfo-393]
	_ = x[opHideoutSetTribute-394]
	_ = x[opHideoutUpgradeWithPowerCrystal-395]
	_ = x[opHideoutDeclareHQ-396]
	_ = x[opHideoutUndeclareHQ-397]
	_ = x[opHideoutGetHQRequirements-398]
	_ = x[opHideoutBoost-399]
	_ = x[opHideoutBoostConstruction-400]
	_ = x[opOpenWorldAttackScheduleStart-401]
	_ = x[opOpenWorldAttackScheduleCancel-402]
	_ = x[opOpenWorldAttackConquerStart-403]
	_ = x[opOpenWorldAttackConquerCancel-404]
	_ = x[opGetOpenWorldAttackDetails-405]
	_ = x[opGetNextOpenWorldAttackScheduleTime-406]
	_ = x[opRecoverVaultFromHideout-407]
	_ = x[opGetGuildEnergyDrainInfo-408]
	_ = x[opChannelingUpdate-409]
	_ = x[opUseCorruptedShrine-410]
	_ = x[opRequestEstimatedMarketValue-411]
	_ = x[opLogFeedback-412]
	_ = x[opGetInfamyInfo-413]
	_ = x[opGetPartySmartClusterQueuePriority-414]
	_ = x[opSetPartySmartClusterQueuePriority-415]
	_ = x[opClientAntiAutoClickerInfo-416]
	_ = x[opClientBotPatternDetectionInfo-417]
	_ = x[opClientAntiGatherClickerInfo-418]
	_ = x[opLoadoutCreate-419]
	_ = x[opLoadoutRead-420]
	_ = x[opLoadoutReadHeaders-421]
	_ = x[opLoadoutUpdate-422]
	_ = x[opLoadoutDelete-423]
	_ = x[opLoadoutOrderUpdate-424]
	_ = x[opLoadoutEquip-425]
	_ = x[opBatchUseItemCancel-426]
	_ = x[opEnlistFactionWarfare-427]
	_ = x[opGetFactionWarfareWeeklyReport-428]
	_ = x[opClaimFactionWarfareWeeklyReport-429]
	_ = x[opGetFactionWarfareCampaignData-430]
	_ = x[opClaimFactionWarfareItemReward-431]
	_ = x[opSendMemoryConsumption-432]
	_ = x[opPickupCarriableObjectStart-433]
	_ = x[opPickupCarriableObjectCancel-434]
	_ = x[opSetSavingChestLogsFlag-435]
	_ = x[opGetSavingChestLogsFlag-436]
	_ = x[opRegisterGuestAccount-437]
	_ = x[opResendGuestAccountVerificationEmail-438]
	_ = x[opDoSimpleActionStart-439]
	_ = x[opDoSimpleActionCancel-440]
	_ = x[opGetGvgSeasonContributionByActivity-441]
	_ = x[opGetGvgSeasonContributionByCrystalLeague-442]
	_ = x[opGetGuildMightCategoryContribution-443]
	_ = x[opGetGuildMightCategoryOverview-444]
	_ = x[opGetPvpChallengeData-445]
	_ = x[opClaimPvpChallengeWeeklyReward-446]
	_ = x[opGetPersonalMightStats-447]
	_ = x[opAuctionGetLoadoutOffers-448]
	_ = x[opAuctionBuyLoadoutOffer-449]
	_ = x[opAccountDeletionRequest-450]
	_ = x[opAccountReactivationRequest-451]
	_ = x[opGetModerationEscalationDefiniton-452]
	_ = x[opEventBasedPopupAddSeen-453]
	_ = x[opGetItemKillHistory-454]
	_ = x[opGetVanityConsumables-455]
	_ = x[opEquipKillEmote-456]
	_ = x[opChangeKillEmotePlayOnKnockdownSetting-457]
	_ = x[opBuyVanityConsumableCharges-458]
	_ = x[opReclaimVanityItem-459]
	_ = x[opGetArenaRankings-460]
	_ = x[opGetCrystalLeagueStatistics-461]
	_ = x[opSendOptionsLog-462]
	_ = x[opSendControlsOptionsLog-463]
	_ = x[opMistsUseImmediateReturnExit-464]
	_ = x[opMistsUseStaticEntrance-465]
	_ = x[opMistsUseCityRoadsEntrance-466]
	_ = x[opChangeNewGuildMemberMail-467]
	_ = x[opGetNewGuildMemberMail-468]
	_ = x[opChangeGuildFactionAllegiance-469]
	_ = x[opGetGuildFactionAllegiance-470]
	_ = x[opGuildBannerChange-471]
	_ = x[opGuildGetOptionalStats-472]
	_ = x[opGuildSetOptionalStats-473]
	_ = x[opGetPlayerInfoForStalk-474]
	_ = x[opPayGoldForCharacterTypeChange-475]
	_ = x[opQuickSellAuctionQueryAction-476]
	_ = x[opQuickSellAuctionSellAction-477]
	_ = x[opFcmTokenToServer-478]
	_ = x[opApnsTokenToServer-479]
	_ = x[opDeathRecap-480]
	_ = x[opAuctionFetchFinishedAuctions-481]
	_ = x[opAbortAuctionFetchFinishedAuctions-482]
	_ = x[opRequestLegendaryEvenHistory-483]
	_ = x[opPartyAnswerStartHuntRequest-484]
	_ = x[opHuntAbort-485]
	_ = x[opUseFindTrackSpellFromItemPrepare-486]
	_ = x[opInteractWithTrackStart-487]
	_ = x[opInteractWithTrackCancel-488]
	_ = x[opTerritoryRaidStart-489]
	_ = x[opTerritoryRaidCancel-490]
	_ = x[opTerritoryClaimRaidedRawEnergyCrystalResult-491]
	_ = x[opGvGSeasonPlayerGuildParticipationDetails-492]
	_ = x[opDailyMightBonus-493]
	_ = x[opClaimDailyMightBonus-494]
	_ = x[opGetFortificationGroupInfo-495]
	_ = x[opUpgradeFortificationGroup-496]
}

const _OperationType_name = "opUnusedopPingopJoinopVersionedOperationopCreateAccountopLoginopCreateGuestAccountopSendCrashLogopSendTraceRouteopSendVfxStatsopSendGamePingInfoopCreateCharacteropDeleteCharacteropSelectCharacteropAcceptPopupsopRedeemKeycodeopGetGameServerByClusteropGetShopPurchaseUrlopGetReferralSeasonDetailsopGetReferralLinkopGetShopTilesForCategoryopMoveopAttackStartopCastStartopCastCancelopTerminateToggleSpellopChannelingCancelopAttackBuildingStartopInventoryDestroyItemopInventoryMoveItemopInventoryRecoverItemopInventoryRecoverAllItemsopInventorySplitStackopInventorySplitStackIntoopGetClusterDataopChangeClusteropConsoleCommandopChatMessageopReportClientErroropRegisterToObjectopUnRegisterFromObjectopCraftBuildingChangeSettingsopCraftBuildingTakeMoneyopRepairBuildingChangeSettingsopRepairBuildingTakeMoneyopActionBuildingChangeSettingsopHarvestStartopHarvestCancelopTakeSilveropActionOnBuildingStartopActionOnBuildingCancelopInstallResourceStartopInstallResourceCancelopInstallSilveropBuildingFillNutritionopBuildingChangeRenovationStateopBuildingBuySkinopBuildingClaimopBuildingGiveupopBuildingNutritionSilverStorageDepositopBuildingNutritionSilverStorageWithdrawopBuildingNutritionSilverRewardSetopConstructionSiteCreateopPlaceableObjectPlaceopPlaceableObjectPlaceCancelopPlaceableObjectPickupopFurnitureObjectUseopFarmableHarvestopFarmableFinishGrownItemopFarmableDestroyopFarmableGetProductopFarmableFillopTearDownConstructionSiteopPlaceholder1opAuctionCreateOfferopAuctionCreateRequestopAuctionGetOffersopAuctionGetRequestsopAuctionBuyOfferopAuctionAbortAuctionopAuctionModifyAuctionopAuctionAbortOfferopAuctionAbortRequestopAuctionSellRequestopAuctionGetFinishedAuctionsopAuctionGetFinishedAuctionsCountopAuctionFetchAuctionopAuctionGetMyOpenOffersopAuctionGetMyOpenRequestsopAuctionGetMyOpenAuctionsopAuctionGetItemAverageStatsopAuctionGetItemAverageValueopContainerOpenopContainerCloseopContainerManageSubContaineropRespawnopSuicideopJoinGuildopLeaveGuildopCreateGuildopInviteToGuildopDeclineGuildInvitationopKickFromGuildopInstantJoinGuildopDuellingChallengePlayeropDuellingAcceptChallengeopDuellingDenyChallengeopChangeClusterTaxopClaimTerritoryopGiveUpTerritoryopChangeTerritoryAccessRightsopGetMonolithInfoopGetClaimInfoopGetAttackInfoopGetTerritorySeasonPointsopGetAttackScheduleopGetMatchesopGetMatchDetailsopJoinMatchopLeaveMatchopChangeChatSettingsopLogoutStartopLogoutCancelopClaimOrbStartopClaimOrbCancelopMatchLootChestOpeningStartopMatchLootChestOpeningCancelopDepositToGuildAccountopWithdrawalFromAccountopChangeGuildPayUpkeepFlagopChangeGuildTaxopGetMyTerritoriesopMorganaCommandopGetServerInfoopSubscribeToClusteropAnswerMercenaryInvitationopGetCharacterEquipmentopGetCharacterSteamAchievementsopGetCharacterStatsopGetKillHistoryDetailsopLearnMasteryLevelopReSpecAchievementopChangeAvataropGetRankingsopGetRankopGetGvgSeasonRankingsopGetGvgSeasonRankopGetGvgSeasonHistoryRankingsopGetGvgSeasonGuildMemberHistoryopKickFromGvGMatchopGetCrystalLeagueDailySeasonPointsopGetChestLogsopGetAccessRightLogsopGetGuildAccountLogsopGetGuildAccountLogsLargeAmountopInviteToPlayerTradeopPlayerTradeCancelopPlayerTradeInvitationAcceptopPlayerTradeAddItemopPlayerTradeRemoveItemopPlayerTradeAcceptTradeopPlayerTradeSetSilverOrGoldopSendMiniMapPingopStuckopBuyRealEstateopClaimRealEstateopGiveUpRealEstateopChangeRealEstateOutlineopGetMailInfosopGetMailCountopReadMailopSendNewMailopDeleteMailopMarkMailUnreadopClaimAttachmentFromMailopApplyToGuildopAnswerGuildApplicationopRequestGuildFinderFilteredListopUpdateGuildRecruitmentInfoopRequestGuildRecruitmentInfoopRequestGuildFinderNameSearchopRequestGuildFinderRecommendedListopRegisterChatPeeropSendChatMessageopSendModeratorMessageopJoinChatChannelopLeaveChatChannelopSendWhisperMessageopSayopPlayEmoteopStopEmoteopGetClusterMapInfoopAccessRightsChangeSettingsopMountopMountCancelopBuyJourneyopSetSaleStatusForEstateopResolveGuildOrPlayerNameopGetRespawnInfosopMakeHomeopLeaveHomeopResurrectionReplyopAllianceCreateopAllianceDisbandopAllianceGetMemberInfosopAllianceInviteopAllianceAnswerInvitationopAllianceCancelInvitationopAllianceKickGuildopAllianceLeaveopAllianceChangeGoldPaymentFlagopAllianceGetDetailInfoopGetIslandInfosopAbandonMyIslandopBuyMyIslandopBuyGuildIslandopAbandonGuildIslandopUpgradeMyIslandopUpgradeGuildIslandopMoveMyIslandopMoveGuildIslandopTerritoryFillNutritionopTeleportBackopPartyInvitePlayeropPartyRequestJoinopPartyAnswerInvitationopPartyAnswerJoinRequestopPartyLeaveopPartyKickPlayeropPartyMakeLeaderopPartyChangeLootSettingopPartyMarkObjectopPartySetRoleopSetGuildCodexopExitEnterStartopExitEnterCancelopQuestGiverRequestopGoldMarketGetBuyOfferopGoldMarketGetBuyOfferFromSilveropGoldMarketGetSellOfferopGoldMarketGetSellOfferFromSilveropGoldMarketBuyGoldopGoldMarketSellGoldopGoldMarketCreateSellOrderopGoldMarketCreateBuyOrderopGoldMarketGetInfosopGoldMarketCancelOrderopGoldMarketGetAverageInfoopTreasureChestUsingStartopTreasureChestUsingCancelopUseLootChestopUseShrineopUseHellgateShrineopUseSiegeBanneropGetSiegeBannerInfoopLaborerStartJobopLaborerTakeJobLootopLaborerDismissopLaborerMoveopLaborerBuyItemopLaborerUpgradeopBuyPremiumopRealEstateGetAuctionDataopRealEstateBidOnAuctionopFriendInviteopFriendAnswerInvitationopFriendCancelnvitationopFriendRemoveopInventoryStackopInventorySortopInventoryDropAllopInventoryAddToStacksopEquipmentItemChangeSpellopExpeditionRegisteropExpeditionRegisterCancelopJoinExpeditionopDeclineExpeditionInvitationopVoteStartopVoteDoVoteopRatingDoRateopEnteringExpeditionStartopEnteringExpeditionCancelopActivateExpeditionCheckPointopArenaRegisteropArenaAddInviteopArenaRegisterCancelopArenaLeaveopJoinArenaMatchopDeclineArenaInvitationopEnteringArenaStartopEnteringArenaCancelopArenaCustomMatchopUpdateCharacterStatementopBoostFarmableopGetStrikeHistoryopUseFunctionopUsePortalEntranceopResetPortalBindingopQueryPortalBindingopClaimPaymentTransactionopChangeUseFlagopClientPerformanceStatsopExtendedHardwareStatsopClientLowMemoryWarningopTerritoryClaimStartopTerritoryClaimCancelopDeliverCarriableObjectStartopDeliverCarriableObjectCancelopTerritoryUpgradeWithPowerCrystalopRequestAppStoreProductsopVerifyProductPurchaseopQueryGuildPlayerStatsopQueryAllianceGuildStatsopTrackAchievementsopSetAchievementsAutoLearnopDepositItemToGuildCurrencyopWithdrawalItemFromGuildCurrencyopAuctionSellSpecificItemRequestopFishingStartopFishingCastingopFishingCastopFishingCatchopFishingPullopFishingGiveLineopFishingFinishopFishingCancelopCreateGuildAccessTagopDeleteGuildAccessTagopRenameGuildAccessTagopFlagGuildAccessTagGuildPermissionopAssignGuildAccessTagopRemoveGuildAccessTagFromPlayeropModifyGuildAccessTagEditorsopRequestPublicAccessTagsopChangeAccessTagPublicFlagopUpdateGuildAccessTagopSteamStartMicrotransactionopSteamFinishMicrotransactionopSteamIdHasActiveAccountopCheckEmailAccountStateopLinkAccountToSteamIdopInAppConfirmPaymentGooglePlayopInAppConfirmPaymentAppleAppStoreopInAppPurchaseRequestopInAppPurchaseFailedopCharacterSubscriptionInfoopAccountSubscriptionInfoopBuyGvgSeasonBoosteropChangeFlaggingPrepareopOverChargeopOverChargeEndopRequestTrustedopChangeGuildLogoopPartyFinderRegisterForUpdatesopPartyFinderUnregisterForUpdatesopPartyFinderEnlistNewPartySearchopPartyFinderDeletePartySearchopPartyFinderChangePartySearchopPartyFinderChangeRoleopPartyFinderApplyForGroupopPartyFinderAcceptOrDeclineApplyForGroupopPartyFinderGetEquipmentSnapshotopPartyFinderRegisterApplicantsopPartyFinderUnregisterApplicantsopPartyFinderFulltextSearchopPartyFinderRequestEquipmentSnapshotopGetPersonalSeasonTrackerDataopGetPersonalSeasonPastRewardDataopUseConsumableFromInventoryopClaimPersonalSeasonRewardopEasyAntiCheatMessageToServeropXignCodeMessageToServeropBattlEyeMessageToServeropSetNextTutorialStateopAddPlayerToMuteListopRemovePlayerFromMuteListopProductShopUserEventopGetVanityUnlocksopBuyVanityUnlocksopGetMountSkinsopSetMountSkinopSetWardrobeopChangeCustomizationopChangePlayerIslandDataopGetGuildChallengePointsopSmartQueueJoinopSmartQueueLeaveopSmartQueueSelectSpawnClusteropUpgradeHideoutopInitHideoutAttackStartopInitHideoutAttackCancelopHideoutFillNutritionopHideoutGetInfoopHideoutGetOwnerInfoopHideoutSetTributeopHideoutUpgradeWithPowerCrystalopHideoutDeclareHQopHideoutUndeclareHQopHideoutGetHQRequirementsopHideoutBoostopHideoutBoostConstructionopOpenWorldAttackScheduleStartopOpenWorldAttackScheduleCancelopOpenWorldAttackConquerStartopOpenWorldAttackConquerCancelopGetOpenWorldAttackDetailsopGetNextOpenWorldAttackScheduleTimeopRecoverVaultFromHideoutopGetGuildEnergyDrainInfoopChannelingUpdateopUseCorruptedShrineopRequestEstimatedMarketValueopLogFeedbackopGetInfamyInfoopGetPartySmartClusterQueuePriorityopSetPartySmartClusterQueuePriorityopClientAntiAutoClickerInfoopClientBotPatternDetectionInfoopClientAntiGatherClickerInfoopLoadoutCreateopLoadoutReadopLoadoutReadHeadersopLoadoutUpdateopLoadoutDeleteopLoadoutOrderUpdateopLoadoutEquipopBatchUseItemCancelopEnlistFactionWarfareopGetFactionWarfareWeeklyReportopClaimFactionWarfareWeeklyReportopGetFactionWarfareCampaignDataopClaimFactionWarfareItemRewardopSendMemoryConsumptionopPickupCarriableObjectStartopPickupCarriableObjectCancelopSetSavingChestLogsFlagopGetSavingChestLogsFlagopRegisterGuestAccountopResendGuestAccountVerificationEmailopDoSimpleActionStartopDoSimpleActionCancelopGetGvgSeasonContributionByActivityopGetGvgSeasonContributionByCrystalLeagueopGetGuildMightCategoryContributionopGetGuildMightCategoryOverviewopGetPvpChallengeDataopClaimPvpChallengeWeeklyRewardopGetPersonalMightStatsopAuctionGetLoadoutOffersopAuctionBuyLoadoutOfferopAccountDeletionRequestopAccountReactivationRequestopGetModerationEscalationDefinitonopEventBasedPopupAddSeenopGetItemKillHistoryopGetVanityConsumablesopEquipKillEmoteopChangeKillEmotePlayOnKnockdownSettingopBuyVanityConsumableChargesopReclaimVanityItemopGetArenaRankingsopGetCrystalLeagueStatisticsopSendOptionsLogopSendControlsOptionsLogopMistsUseImmediateReturnExitopMistsUseStaticEntranceopMistsUseCityRoadsEntranceopChangeNewGuildMemberMailopGetNewGuildMemberMailopChangeGuildFactionAllegianceopGetGuildFactionAllegianceopGuildBannerChangeopGuildGetOptionalStatsopGuildSetOptionalStatsopGetPlayerInfoForStalkopPayGoldForCharacterTypeChangeopQuickSellAuctionQueryActionopQuickSellAuctionSellActionopFcmTokenToServeropApnsTokenToServeropDeathRecapopAuctionFetchFinishedAuctionsopAbortAuctionFetchFinishedAuctionsopRequestLegendaryEvenHistoryopPartyAnswerStartHuntRequestopHuntAbortopUseFindTrackSpellFromItemPrepareopInteractWithTrackStartopInteractWithTrackCancelopTerritoryRaidStartopTerritoryRaidCancelopTerritoryClaimRaidedRawEnergyCrystalResultopGvGSeasonPlayerGuildParticipationDetailsopDailyMightBonusopClaimDailyMightBonusopGetFortificationGroupInfoopUpgradeFortificationGroup"

var _OperationType_index = [...]uint16{0, 8, 14, 20, 40, 55, 62, 82, 96, 112, 126, 144, 161, 178, 195, 209, 224, 248, 268, 294, 311, 336, 342, 355, 366, 378, 400, 418, 439, 461, 480, 502, 528, 549, 574, 590, 605, 621, 634, 653, 671, 693, 722, 746, 776, 801, 831, 845, 860, 872, 895, 919, 941, 964, 979, 1002, 1033, 1050, 1065, 1081, 1120, 1160, 1194, 1218, 1240, 1268, 1291, 1311, 1328, 1353, 1370, 1390, 1404, 1430, 1444, 1464, 1486, 1504, 1524, 1541, 1562, 1584, 1603, 1624, 1644, 1672, 1705, 1726, 1750, 1776, 1802, 1830, 1858, 1873, 1889, 1918, 1927, 1936, 1947, 1959, 1972, 1987, 2011, 2026, 2044, 2069, 2094, 2117, 2135, 2151, 2168, 2197, 2214, 2228, 2243, 2269, 2288, 2300, 2317, 2328, 2340, 2360, 2373, 2387, 2402, 2418, 2446, 2475, 2498, 2521, 2547, 2563, 2581, 2597, 2612, 2632, 2659, 2682, 2713, 2732, 2755, 2774, 2793, 2807, 2820, 2829, 2851, 2869, 2898, 2930, 2948, 2983, 2997, 3017, 3038, 3070, 3091, 3110, 3139, 3159, 3182, 3206, 3234, 3251, 3258, 3273, 3290, 3308, 3333, 3347, 3361, 3371, 3384, 3396, 3412, 3437, 3451, 3475, 3507, 3535, 3564, 3594, 3629, 3647, 3664, 3686, 3703, 3721, 3741, 3746, 3757, 3768, 3787, 3815, 3822, 3835, 3847, 3871, 3897, 3914, 3924, 3935, 3954, 3970, 3987, 4011, 4027, 4053, 4079, 4098, 4113, 4144, 4167, 4183, 4200, 4213, 4229, 4249, 4266, 4286, 4300, 4317, 4341, 4355, 4374, 4392, 4415, 4439, 4451, 4468, 4485, 4509, 4526, 4540, 4555, 4571, 4588, 4607, 4630, 4663, 4687, 4721, 4740, 4760, 4787, 4813, 4833, 4856, 4882, 4907, 4933, 4947, 4958, 4977, 4993, 5013, 5030, 5050, 5066, 5079, 5095, 5111, 5123, 5149, 5173, 5187, 5211, 5234, 5248, 5264, 5279, 5297, 5319, 5345, 5365, 5391, 5407, 5436, 5447, 5459, 5473, 5498, 5524, 5554, 5569, 5585, 5606, 5618, 5634, 5658, 5678, 5699, 5717, 5743, 5758, 5776, 5789, 5808, 5828, 5848, 5873, 5888, 5912, 5935, 5959, 5980, 6002, 6031, 6061, 6095, 6120, 6143, 6166, 6191, 6210, 6236, 6264, 6297, 6329, 6343, 6359, 6372, 6386, 6399, 6416, 6431, 6446, 6468, 6490, 6512, 6547, 6569, 6601, 6630, 6655, 6682, 6704, 6732, 6761, 6786, 6810, 6832, 6863, 6897, 6919, 6940, 6967, 6992, 7013, 7036, 7048, 7063, 7079, 7096, 7127, 7160, 7193, 7223, 7253, 7276, 7302, 7343, 7376, 7407, 7440, 7467, 7504, 7534, 7567, 7595, 7622, 7652, 7677, 7702, 7724, 7745, 7771, 7793, 7811, 7829, 7844, 7858, 7871, 7892, 7916, 7941, 7957, 7974, 8004, 8020, 8044, 8069, 8091, 8107, 8128, 8147, 8179, 8197, 8217, 8243, 8257, 8283, 8313, 8344, 8373, 8403, 8430, 8466, 8491, 8516, 8534, 8554, 8583, 8596, 8611, 8646, 8681, 8708, 8739, 8768, 8783, 8796, 8816, 8831, 8846, 8866, 8880, 8900, 8922, 8953, 8986, 9017, 9048, 9071, 9099, 9128, 9152, 9176, 9198, 9235, 9256, 9278, 9314, 9355, 9390, 9421, 9442, 9473, 9496, 9521, 9545, 9569, 9597, 9631, 9655, 9675, 9697, 9713, 9752, 9780, 9799, 9817, 9845, 9861, 9885, 9914, 9938, 9965, 9991, 10014, 10044, 10071, 10090, 10113, 10136, 10159, 10190, 10219, 10247, 10265, 10284, 10296, 10326, 10361, 10390, 10419, 10430, 10464, 10488, 10513, 10533, 10554, 10598, 10640, 10657, 10679, 10706, 10733}

func (i OperationType) String() string {
	if i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
