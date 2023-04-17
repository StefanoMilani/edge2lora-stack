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

import React, { useState, useCallback } from 'react'
import { connect, useSelector } from 'react-redux'
import { push, replace } from 'connected-react-router'
import { defineMessages } from 'react-intl'

import toast from '@ttn-lw/components/toast'
import {
  composeContact,
  getAdministrativeContact,
  getTechnicalContact,
} from '@ttn-lw/components/contact-fields/utils'

import OAuthClientForm from '@account/components/oauth-client-form'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import PropTypes from '@ttn-lw/lib/prop-types'
import diff from '@ttn-lw/lib/diff'

import { mayEditBasicClientInformation, checkFromState } from '@account/lib/feature-checks'

import { deleteClient, updateClient } from '@account/store/actions/clients'

const m = defineMessages({
  deleteSuccess: 'OAuth client deleted',
  deleteFail: 'There was an error and the OAuth client could not be deleted',
  updateSuccess: 'OAuth client updated',
  updateFailure: 'There was an error updating this client',
})

const checkChanged = (changed, values) => {
  if ('redirect_uris' in changed) {
    return {
      ...changed,
      redirect_uris: values.redirect_uris,
    }
  } else if ('logout_redirect_uris' in changed) {
    return {
      ...changed,
      logout_redirect_uris: values.logout_redirect_uris,
    }
  } else if ('grants' in changed) {
    return {
      ...changed,
      grants: values.grants,
    }
  }

  return changed
}

const ClientAdd = props => {
  const {
    userId,
    isAdmin,
    rights,
    pseudoRights,
    navigateToOAuthClient,
    deleteOAuthClient,
    onDeleteSuccess,
    initialValues,
    updateOauthClient,
  } = props

  const mayEditBasicInformation = useSelector(state =>
    checkFromState(mayEditBasicClientInformation, state),
  )
  const [error, setError] = useState()
  const handleSubmit = useCallback(
    async (values, resetForm, setSubmitting) => {
      const { client_id } = values.ids
      setError(undefined)

      const {
        _administrative_contact_id,
        _administrative_contact_type,
        _technical_contact_id,
        _technical_contact_type,
      } = values

      const administrative_contact =
        _administrative_contact_id !== ''
          ? composeContact(_administrative_contact_type, _administrative_contact_id)
          : ''

      const technical_contact =
        _technical_contact_id !== ''
          ? composeContact(_technical_contact_type, _technical_contact_id)
          : ''

      const changed = diff(
        initialValues,
        { administrative_contact, technical_contact, ...values },
        {
          exclude: [
            '_administrative_contact_id',
            '_administrative_contact_type',
            '_technical_contact_id',
            '_technical_contact_type',
          ],
        },
      )

      // If there is a change in `redirect_uris`, `logout_redirect_uris` or `grants`,
      // copy all values so they don't get overwritten.
      const update = checkChanged(changed, values)
      if (technical_contact === '') {
        changed.technical_contact = null
      }
      if (administrative_contact === '') {
        changed.administrative_contact = null
      }
      const { owner_id, ...newClient } = update

      try {
        await updateOauthClient(client_id, newClient)
        resetForm({ values })
        toast({
          title: client_id,
          message: m.updateSuccess,
          type: toast.types.SUCCESS,
        })
      } catch (error) {
        setSubmitting(false)
        setError(error)
        toast({
          title: client_id,
          message: m.updateFailure,
          type: toast.types.ERROR,
        })
      }
    },
    [initialValues, updateOauthClient],
  )

  const handleDelete = useCallback(
    async (shouldPurge, clientId) => {
      setError(undefined)

      try {
        await deleteOAuthClient(clientId, shouldPurge)
        onDeleteSuccess()
        toast({
          title: clientId,
          message: m.deleteSuccess,
          type: toast.types.SUCCESS,
        })
      } catch (error) {
        setError(error)
        toast({
          title: clientId,
          message: m.deleteFail,
          type: toast.types.ERROR,
        })
      }
    },
    [deleteOAuthClient, onDeleteSuccess],
  )

  // Add technical and administrative contact to the initial values.
  const { administrative_contact, technical_contact, ...restInitialValues } = initialValues
  const technicalContact = getTechnicalContact(initialValues)
  const administrativeContact = getAdministrativeContact(initialValues)
  const composedInitialValues = {
    ...technicalContact,
    ...administrativeContact,
    ...restInitialValues,
  }

  return (
    <OAuthClientForm
      update
      initialValues={composedInitialValues}
      onSubmit={handleSubmit}
      onDelete={handleDelete}
      onDeleteSuccess={onDeleteSuccess}
      navigateToOAuthClient={navigateToOAuthClient}
      error={error}
      userId={userId}
      isAdmin={isAdmin}
      rights={rights}
      pseudoRights={pseudoRights}
      mayEditBasicInformation={mayEditBasicInformation}
    />
  )
}

ClientAdd.propTypes = {
  deleteOAuthClient: PropTypes.func.isRequired,
  initialValues: PropTypes.shape({
    grants: PropTypes.arrayOf(PropTypes.string),
    technical_contact: PropTypes.shape({
      user_ids: PropTypes.shape({
        user_id: PropTypes.string,
      }),
      organization_ids: PropTypes.shape({
        organization_id: PropTypes.string,
      }),
    }),
    administrative_contact: PropTypes.shape({
      user_ids: PropTypes.shape({
        user_id: PropTypes.string,
      }),
      organization_ids: PropTypes.shape({
        organization_id: PropTypes.string,
      }),
    }),
    _administrative_contact_id: PropTypes.string,
    _administrative_contact_type: PropTypes.string,
    _technical_contact_id: PropTypes.string,
    _technical_contact_type: PropTypes.string,
  }).isRequired,
  isAdmin: PropTypes.bool.isRequired,
  navigateToOAuthClient: PropTypes.func.isRequired,
  onDeleteSuccess: PropTypes.func.isRequired,
  pseudoRights: PropTypes.rights.isRequired,
  rights: PropTypes.rights,
  updateOauthClient: PropTypes.func.isRequired,
  userId: PropTypes.string.isRequired,
}

ClientAdd.defaultProps = {
  rights: undefined,
}

export default connect(null, dispatch => ({
  navigateToOAuthClient: clientId => dispatch(push(`/oauth-clients/${clientId}`)),
  deleteOAuthClient: (id, shouldPurge) => dispatch(attachPromise(deleteClient(id, shouldPurge))),
  onDeleteSuccess: () => dispatch(replace(`/oauth-clients`)),
  updateOauthClient: (id, patch) => dispatch(attachPromise(updateClient(id, patch))),
}))(ClientAdd)
