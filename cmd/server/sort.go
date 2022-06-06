package server

import (
	"github.com/reaperhero/stock_dingding/model"
	"strings"
)

var (
	SortWithSubordinatePe sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.Pe > y.Pe {
			return true
		}
		if x.Pe < y.Pe {
			return false
		}
		return true
	}
	SortWithSubordinateMarkValue sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.TotalMarketValue > y.TotalMarketValue {
			return true
		}
		if x.TotalMarketValue < y.TotalMarketValue {
			return false
		}
		return true
	}
	SortWithSubordinateSixDaysChange sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.SixDaysUp > y.SixDaysUp {
			return true
		}
		if x.SixDaysUp < y.SixDaysUp {
			return false
		}
		return true
	}

	SortWithSubordinateThreeDaysChange sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.ThreeDaysUp > y.ThreeDaysUp {
			return true
		}
		if x.ThreeDaysUp < y.ThreeDaysUp {
			return false
		}
		return true
	}

	SortWithSubordinateIncrease sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.IncreasePrecent > y.IncreasePrecent {
			return true
		}
		if x.IncreasePrecent < y.IncreasePrecent {
			return false
		}
		if x.TotalMarketValue > y.TotalMarketValue {
			return true
		}
		if x.TotalMarketValue < y.TotalMarketValue {
			return false
		}
		return true
	}
)
