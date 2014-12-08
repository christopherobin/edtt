/**
 * Systems Controller
 */
angular.module('Dashboard').controller('CreateStationCtrl', ['$scope', '$modalInstance', 'Systems', CreateStationCtrl]);

function CreateStationCtrl($scope, $modalInstance, Systems) {
    $scope.ok = function () {
        $modalInstance.close($scope.station);
    };

    $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
    };
}