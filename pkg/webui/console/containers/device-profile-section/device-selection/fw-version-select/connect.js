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

import { isUnknownHwVersion } from '@console/lib/device-utils'

import { selectDeviceModelFirmwareVersions } from '@console/store/selectors/device-repository'

const mapStateToProps = (state, props) => {
  const { brandId, modelId, hwVersion } = props

  const fwVersions = selectDeviceModelFirmwareVersions(state, brandId, modelId).filter(
    ({ supported_hardware_versions = [] }) =>
      (Boolean(hwVersion) && supported_hardware_versions.includes(hwVersion)) ||
      // Include firmware versions when there are no hardware versions configured in device repository
      // for selected end device model.
      isUnknownHwVersion(hwVersion),
  )

  return {
    versions: fwVersions,
  }
}

export default FirmwareVersionSelect => connect(mapStateToProps)(FirmwareVersionSelect)
