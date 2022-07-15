// Copyright 2020 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model of a game-specific rule.

package game

type Rule struct {
	Id             int
	RuleNumber     string
	IsTechnical    bool
	IsRankingPoint bool
	Description    string
}

// All rules from the 2022 game that carry point penalties.
var rules = []*Rule{
	{1, "C13", false, false, "機器人不能分離"},
	{2, "C13", true, false, "機器人不能分離"},
	{3, "C21", false, false, "不得接觸敵方機器人超過5秒"},
	{4, "C31", false, false, "每台機器人至多控制3個場地物件, 超過五秒視為犯規"},
	{5, "C32", true, false, "只能使用 Catching 將物件丟進場地內"},
	{6, "C33", true, false, "不得使用任意方式阻擋比賽進行"},
	{7, "C41", true, false, "機器人或機器人控制的物件觸碰到Trap或是在Endgame時間外碰到Cube"},
	{8, "C51", true, false, "不得進入自己的 Collection"},
	{9, "C52", true, false, "得進入敵方 Catching 超過5秒 "},
}
var ruleMap map[int]*Rule

// Returns the rule having the given ID, or nil if no such rule exists.
func GetRuleById(id int) *Rule {
	return GetAllRules()[id]
}

// Returns a slice of all defined rules that carry point penalties.
func GetAllRules() map[int]*Rule {
	if ruleMap == nil {
		ruleMap = make(map[int]*Rule, len(rules))
		for _, rule := range rules {
			ruleMap[rule.Id] = rule
		}
	}
	return ruleMap
}
