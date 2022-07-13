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
	// {1, "G103", false, false, "BUMPERS must be in Bumper Zone (see R402) during the match."},
	{1, "G15", true, false, "不得進入 Collection 超過5秒"},
	{2, "O1", true, false, "友善溝通，禁止言語怒罵，例: 問候、歧視、嘲諷等"},
	{3, "O2", true, false, "Driver Station內限制最多三人(driver, driver, coach)"},
	{4, "G22", true, false, "機器人上 bumper 不能分離機器人"},
	{5, "G23", true, false, "不能將 bumper 舉起來"},
	{6, "G21", true, false, "機器人不能分離"},
	{7, "G18", true, false, "不得接觸敵方機器人超過5秒"},
	{8, "G17", true, false, "機器人倒下要遠離他"},
	{9, "G19", true, false, "不得刻意傷害、連接他人機器人"},
	{10, "H7", true, false, "只能使用 Catching 將物件丟進場地內"},
	{11, "O3", true, false, "比賽結果由裁判評分，請尊重裁判判決"},
	{12, "O4", false, false, "每台機器人至多攜帶3個場地物件,超過五秒視為犯規"},
	{13, "O5", true, false, "機器人或機器人上的物件觸碰到 Trap"},
	{14, "O6", true, false, "機器人或機器人上的物件在非 Endgame 時觸碰到 Cube"},
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
