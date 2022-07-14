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
	{1, "C1", true, false, "不得進入自己的 Collection"},
	{2, "C6", false, false, "機器人不能分離"},
	{3, "C6", true, false, "機器人不能分離"},
	{4, "C7", false, false, "不得接觸敵方機器人超過5秒"},
	{5, "C10", true, false, "只能使用 Catching 將物件丟進場地內"},
	{6, "C11", false, false, "每台機器人至多控制3個場地物件，超過五秒視為犯規"},
	{7, "C12", true, false, "機器人或機器人控制的物件觸碰到Trap或是在endgame時間外碰到cube"},
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
