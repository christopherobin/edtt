/**
 * Systems Controller
 */
angular.module('Dashboard').controller('SystemsCtrl', ['$scope', '$modal', 'Systems', 'Market', SystemsCtrl]);

function SystemsCtrl($scope, $modal, Systems, Market) {
    $scope.activeSystem = null;
    $scope.activeStation = null;
    $scope.systems = null;
    $scope.stations = null;
    $scope.station = null;
    $scope.market = null;
    $scope.galactic = null;

    Systems.list(function(systems) {
        $scope.systems = systems;
    });

    Market.list(function(market) {
        $scope.galactic = market;
    });

    $scope.$watch('systems', function(newVal) {
        if (!newVal) return;
        if (!$scope.activeSystem && newVal.length > 0) {
            $scope.activeSystem = newVal[0];
        }
    });

    $scope.$watch('activeSystem', function(newVal, oldVal) {
        if (newVal !== oldVal) {
            if (!newVal) {
                $scope.stations = [];
                return;
            }

            Systems.system(newVal.name, function(system) {
                $scope.stations = system.stations;
            });
        }
    });

    $scope.$watch('stations', function(newVal) {
        if (newVal && newVal.length > 0) {
            $scope.activeStation = newVal[0];
        } else {
            $scope.activeStation = null;
        }
    });

    $scope.$watch('activeStation', function(newVal, oldVal) {
        if (newVal !== oldVal) {
            if (!newVal) {
                $scope.station = null;
                $scope.market = {};
            } else {
                Systems.station($scope.activeSystem.name, newVal, function(station) {
                    $scope.station = station;
                    $scope.market = station.market;
                });
            }
        }
    });

    $scope.$watch('station.services', function(newVal, oldVal) {
        if (typeof newVal === 'undefined' || typeof oldVal === 'undefined') return;
        Systems.setServices($scope.activeSystem.name, $scope.activeStation, newVal, function(station) {

        });
    }, true);


    $scope.$on('editable', function(scope, row, name) {
        var data = {};
        data[name.toLowerCase()] = parseInt(row[name], 10);
        Systems.setEntry($scope.activeSystem.name, $scope.activeStation, row.Name, data, function(entry) {
            //$scope.station.market[row.Name] = entry;
        });
    });

    $scope.compare = function(category, entry) {
        return {
            'sell': entry.Sell ? entry.Sell - $scope.galactic[category][entry.Name].galactic_avg : 0,
            'buy': entry.Buy ? entry.Buy - $scope.galactic[category][entry.Name].galactic_avg : 0
        }
    };

    $scope.selectSystem = function(system) {
        $scope.activeSystem = system;
        $scope.activeStation = null;
    };

    $scope.selectStation = function(event, station) {
        $scope.activeStation = station;
        event.preventDefault();
    };

    $scope.createSystem = function(event) {
        var modalInstance = $modal.open({
            templateUrl: 'createSystem.html',
            controller: 'CreateSystemCtrl'
        });

        modalInstance.result.then(function (system) {
            Systems.createSystem(system, function() {
                Systems.list(function(systems) {
                    $scope.systems = systems;
                });
            });
        }, function () {
            //$log.info('Modal dismissed at: ' + new Date());
        });

        event.preventDefault();
    };

    $scope.createStation = function(event) {
        var modalInstance = $modal.open({
            templateUrl: 'createStation.html',
            controller: 'CreateStationCtrl'
        });

        modalInstance.result.then(function (station) {
            Systems.createStation($scope.activeSystem.name, station, function() {
                Systems.system($scope.activeSystem.name, function(system) {
                    $scope.stations = system.stations;
                });
            });
        }, function () {
            //$log.info('Modal dismissed at: ' + new Date());
        });

        event.preventDefault();
    };
}