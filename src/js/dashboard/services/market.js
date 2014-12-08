angular.module('Dashboard').factory('Market', ['$http', MarketService]);

function MarketService($http) {
	return {
		list: function(success, error) {
			error = error || Function();
			$http.get('/api/market')
				.success(success)
				.error(error);
		},
		compare: function(from, to, data, success, error) {
			error = error || Function();
			$http.post('/api/market/' + encodeURIComponent(from) + '/' + (encodeURIComponent(to) || "ANY"), data)
				.success(success)
				.error(error);
		},
		findGood: function(good, from, range, success, error) {
			error = error || Function();
			$http.post('/api/market/' + encodeURIComponent(good), {
				range: range,
				from: from
			})
				.success(success)
				.error(error);
		}
	}
}
