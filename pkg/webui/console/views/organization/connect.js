// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { connect } from 'react-redux'

import withRequest from '@ttn-lw/lib/components/with-request'

import { selectApplicationSiteName } from '@ttn-lw/lib/selectors/env'

import {
  checkFromState,
  mayViewOrganizationInformation,
  mayViewOrEditOrganizationApiKeys,
  mayViewOrEditOrganizationCollaborators,
  mayEditBasicOrganizationInformation,
} from '@console/lib/feature-checks'

import {
  getOrganization,
  stopOrganizationEventsStream,
  getOrganizationsRightsList,
} from '@console/store/actions/organizations'

import {
  selectSelectedOrganization,
  selectOrganizationFetching,
  selectOrganizationError,
  selectOrganizationRightsFetching,
  selectOrganizationRightsError,
} from '@console/store/selectors/organizations'

const mapStateToProps = (state, props) => ({
  orgId: props.match.params.orgId,
  fetching: selectOrganizationFetching(state) || selectOrganizationRightsFetching(state),
  organization: selectSelectedOrganization(state),
  error: selectOrganizationError(state) || selectOrganizationRightsError(state),
  siteName: selectApplicationSiteName(),
  mayViewInformation: checkFromState(mayViewOrganizationInformation, state),
  mayViewOrEditApiKeys: checkFromState(mayViewOrEditOrganizationApiKeys, state),
  mayViewOrEditCollaborators: checkFromState(mayViewOrEditOrganizationCollaborators, state),
  mayEditInformation: checkFromState(mayEditBasicOrganizationInformation, state),
})

const mapDispatchToProps = dispatch => ({
  stopStream: id => dispatch(stopOrganizationEventsStream(id)),
  loadData: id => {
    dispatch(getOrganization(id, 'name,description,administrative_contact,technical_contact'))
    dispatch(getOrganizationsRightsList(id))
  },
})

export default Organization =>
  connect(
    mapStateToProps,
    mapDispatchToProps,
  )(
    withRequest(
      ({ orgId, loadData }) => loadData(orgId),
      ({ fetching, organization }) => fetching || !Boolean(organization),
    )(Organization),
  )
