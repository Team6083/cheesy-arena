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
	{1, "C05", false, false, "機器人上的零件不能分離"},
	{2, "C05", true, false, "機器人上的零件不能分離"},
	{3, "C07", true, false, "不得接觸敵方機器人超過5秒"},
	{4, "C10", true, false, "每台機器人至多控制3個場地物件, 超過五秒視為犯規"},
	{5, "C11", true, false, "只能用指定姿勢將物件丟進場地內"},
	{6, "C12", true, false, "不得使用任意方式阻擋比賽進行"},
	{7, "C13", true, false, "不得進入對方的 Parking"},
	{8, "C14", true, false, "不得故意擋在對方的 Collection 前面, 超過五秒視為犯規"},
	{9, "C16", true, false, "2分30秒前機器人或機器人控制的物件觸碰到 Cube"},
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
