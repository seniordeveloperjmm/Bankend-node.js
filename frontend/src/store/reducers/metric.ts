import { resourceAttributesQueryToPromQL } from 'lib/resourceAttributes';
import { GetResourceAttributeQueriesFromURL } from 'store/actions/metrics/setResourceAttributeQueries';
import {
	GET_INITIAL_APPLICATION_ERROR,
	GET_INITIAL_APPLICATION_LOADING,
	GET_INTIAL_APPLICATION_DATA,
	GET_SERVICE_LIST_ERROR,
	GET_SERVICE_LIST_LOADING_START,
	GET_SERVICE_LIST_SUCCESS,
	MetricsActions,
	RESET_INITIAL_APPLICATION_DATA,
	SET_RESOURCE_ATTRIBUTE_QUERIES,
} from 'types/actions/metrics';
import InitialValueTypes from 'types/reducer/metrics';

const InitialValue: InitialValueTypes = {
	error: false,
	errorMessage: '',
	loading: true,
	metricsApplicationLoading: true,
	services: [],
	dbOverView: [],
	externalService: [],
	topEndPoints: [],
	externalAverageDuration: [],
	externalError: [],
	serviceOverview: [],
	resourceAttributeQueries: GetResourceAttributeQueriesFromURL() || [],
	resourceAttributePromQLQuery: resourceAttributesQueryToPromQL(
		GetResourceAttributeQueriesFromURL() || [],
	),
};

const metrics = (
	state = InitialValue,
	action: MetricsActions,
): InitialValueTypes => {
	switch (action.type) {
		case GET_SERVICE_LIST_ERROR: {
			const { errorMessage } = action.payload;

			return {
				...state,
				error: true,
				errorMessage,
				loading: false,
			};
		}

		case GET_SERVICE_LIST_LOADING_START: {
			return {
				...state,
				loading: true,
			};
		}

		case GET_SERVICE_LIST_SUCCESS: {
			return {
				...state,
				loading: false,
				services: action.payload,
			};
		}

		case GET_INITIAL_APPLICATION_LOADING: {
			return {
				...state,
				metricsApplicationLoading: true,
			};
		}
		case GET_INITIAL_APPLICATION_ERROR: {
			return {
				...state,
				metricsApplicationLoading: false,
				errorMessage: action.payload.errorMessage,
				error: true,
			};
		}

		case RESET_INITIAL_APPLICATION_DATA: {
			return {
				...InitialValue,
				// ...state.resourceAttributeQueries,
				// resourceAttributeQueries: state.resourceAttributeQueries,
			};
		}

		case GET_INTIAL_APPLICATION_DATA: {
			const {
				// dbOverView,
				topEndPoints,
				serviceOverview,
				// externalService,
				// externalAverageDuration,
				// externalError,
			} = action.payload;

			return {
				...state,
				// dbOverView,
				topEndPoints,
				serviceOverview,
				// externalService,
				// externalAverageDuration,
				// externalError,
				metricsApplicationLoading: false,
			};
		}

		case SET_RESOURCE_ATTRIBUTE_QUERIES: {
			const { queries, promQLQuery } = action.payload;
			return {
				...state,
				resourceAttributeQueries: queries,
				resourceAttributePromQLQuery: promQLQuery,
			};
		}

		default:
			return state;
	}
};

export default metrics;
