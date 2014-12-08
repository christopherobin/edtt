/**
 * Systems Controller
 */
angular.module('Dashboard').controller('TradeCtrl', ['$scope', 'Systems', 'Market', TradeCtrl]);

function TradeCtrl($scope, Systems, Market) {
    $scope.systems = null;
    $scope.market = null;
    $scope.goods = null;

    Market.list(function(market) {
        $scope.market = market;

        var goods = [];
        for (var category in market) {
            for (var name in market[category]) {
                goods.push(market[category][name]);
            }
        }
        $scope.goods = goods;
    });

    Systems.list(function(systems) {
        $scope.systems = systems;
    });

    $scope.swapTrades = function() {
        var tmp = $scope.tradeFrom;

        if ($scope.tradeTo) {
            $scope.$broadcast('angucomplete-alt:selectResult', 'tradeFrom', $scope.tradeTo);
        } else {
            $scope.$broadcast('angucomplete-alt:clearInput', 'tradeFrom');
            $scope.tradeFrom = null;
        }

        if (tmp) {
            $scope.$broadcast('angucomplete-alt:selectResult', 'tradeTo', tmp);
        } else {
            $scope.$broadcast('angucomplete-alt:clearInput', 'tradeTo');
            $scope.tradeTo = null;
        }
    };

    $scope.findTrades = function(event) {
        Market.compare($scope.tradeFrom.title, $scope.tradeTo ? $scope.tradeTo.title : "ANY", { range: $scope.range }, function(routes) {
            routes.forEach(function(route) {
                route.trades.forEach(function(trade) {
                    if ($scope.cargoSize) {
                        var units = Math.min($scope.cargoSize, Math.floor($scope.funds / trade.Buy));
                        trade.RevenueTotal = units * trade.Revenue;
                        trade.UnitsTraded = units;
                    }
                });

                route.trades.sort(function(a, b) {
                    if (a.RevenueTotal) {
                        return b.RevenueTotal - a.RevenueTotal;
                    }
                    return b.Revenue - a.Revenue;
                });
            });

            routes.sort(function(a, b) {
                return a.distance - b.distance;
            });

            $scope.routes = routes;
        });

        event.preventDefault();
    };

    $scope.findGood = function(event) {
        Market.findGood($scope.find.good.title, $scope.find.from.title, $scope.find.range, function(goods) {
            $scope.found = goods;
        });

        event.preventDefault();
    }
}