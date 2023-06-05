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

import React, { useCallback, useState } from 'react'
import { defineMessages } from 'react-intl'

import Form from '@ttn-lw/components/form'
import Notification from '@ttn-lw/components/notification'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'
import toast from '@ttn-lw/components/toast'
import ModalButton from '@ttn-lw/components/button/modal-button'
import RightsGroup from '@ttn-lw/components/rights-group'

import UserInput from '@console/containers/user-input'
import composeOption from '@console/containers/user-input/util'

import Yup from '@ttn-lw/lib/yup'
import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { userId as collaboratorIdRegexp } from '@ttn-lw/lib/regexp'

const collaboratorOrganizationSchema = Yup.object().shape({
  organization_id: Yup.string().matches(collaboratorIdRegexp, sharedMessages.validateAlphanum),
})

const collaboratorUserSchema = Yup.object().shape({
  user_id: Yup.string().matches(collaboratorIdRegexp, sharedMessages.validateAlphanum),
})

const validationSchema = Yup.object().shape({
  collaborator: Yup.object()
    .shape({
      ids: Yup.object().when(['organization_ids'], {
        is: organizationIds => Boolean(organizationIds),
        then: collaboratorOrganizationSchema,
        otherwise: collaboratorUserSchema,
      }),
    })
    .required(sharedMessages.validateRequired)
    .nullable(),
  rights: Yup.array().min(1, sharedMessages.validateRights),
})

const m = defineMessages({
  collaboratorIdPlaceholder: 'Type to choose a collaborator',
})

const encodeCollaborator = collaboratorOption =>
  collaboratorOption
    ? {
        ids: {
          [`${collaboratorOption.icon}_ids`]: {
            [`${collaboratorOption.icon}_id`]: collaboratorOption.value,
          },
        },
      }
    : null

const decodeCollaborator = collaborator =>
  collaborator && collaborator.ids ? composeOption(collaborator) : null

const CollaboratorForm = props => {
  const {
    collaborator,
    collaboratorId,
    deleteDisabled,
    rights,
    pseudoRights,
    error: passedError,
    update,
    onSubmit,
    onSubmitSuccess,
    onSubmitFailure,
    onDelete,
    onDeleteSuccess,
    onDeleteFailure,
    isCollaboratorUser,
    isCollaboratorAdmin,
    isCollaboratorCurrentUser,
  } = props

  const collaboratorType = isCollaboratorUser ? 'user' : 'organization'

  const [submitError, setSubmitError] = useState(undefined)
  const error = submitError || passedError

  const handleSubmit = useCallback(
    async (values, { resetForm, setSubmitting }) => {
      const { collaborator, rights } = values

      const composedCollaborator = {
        ...collaborator,
        rights,
      }

      setSubmitError(undefined)

      try {
        await onSubmit(composedCollaborator)

        resetForm({ values })
        onSubmitSuccess()
      } catch (error) {
        setSubmitting(false)
        setSubmitError(error)
        onSubmitFailure(error)
      }
    },
    [onSubmit, onSubmitFailure, onSubmitSuccess],
  )
  const handleDelete = useCallback(async () => {
    const collaborator_ids = {
      [`${collaboratorType}_ids`]: {
        [`${collaboratorType}_id`]: collaboratorId,
      },
    }
    const updatedCollaborator = {
      ids: collaborator_ids,
    }
    setSubmitError(undefined)

    try {
      await onDelete(updatedCollaborator)
      toast({
        message: sharedMessages.collaboratorDeleteSuccess,
        type: toast.types.SUCCESS,
      })
      onDeleteSuccess()
    } catch (error) {
      setSubmitError(error)
      onDeleteFailure(error)
    }
  }, [collaboratorId, collaboratorType, onDelete, onDeleteFailure, onDeleteSuccess])

  const initialValues = React.useMemo(() => {
    if (!collaborator) {
      return {
        collaborator: '',
        rights: [...pseudoRights],
      }
    }

    return {
      collaborator: { ids: collaborator.ids, name: collaborator.name },
      rights: [...collaborator.rights],
    }
  }, [collaborator, pseudoRights])

  let warning = null
  if (update) {
    if (isCollaboratorCurrentUser) {
      warning = isCollaboratorAdmin ? (
        <Notification small warning content={sharedMessages.collaboratorWarningAdminSelf} />
      ) : (
        <Notification small warning content={sharedMessages.collaboratorWarningSelf} />
      )
    } else if (isCollaboratorAdmin) {
      warning = <Notification small warning content={sharedMessages.collaboratorWarningAdmin} />
    }
  }

  return (
    <Form
      error={error}
      onSubmit={handleSubmit}
      initialValues={initialValues}
      validationSchema={validationSchema}
    >
      {warning}
      <UserInput
        name="collaborator"
        title={sharedMessages.collaborator}
        placeholder={m.collaboratorIdPlaceholder}
        isClearable
        required
        autoFocus={!update}
        disabled={update}
        encode={encodeCollaborator}
        decode={decodeCollaborator}
      />
      <Form.Field
        name="rights"
        title={sharedMessages.rights}
        required
        component={RightsGroup}
        rights={rights}
        pseudoRight={pseudoRights}
        entityTypeMessage={sharedMessages.collaborator}
      />
      <SubmitBar>
        <Form.Submit
          component={SubmitButton}
          message={update ? sharedMessages.saveChanges : sharedMessages.addCollaborator}
        />
        {update && (
          <ModalButton
            type="button"
            icon="delete"
            disabled={deleteDisabled}
            danger
            naked
            message={
              deleteDisabled
                ? sharedMessages.removeCollaboratorLast
                : isCollaboratorCurrentUser
                ? sharedMessages.removeCollaboratorSelf
                : sharedMessages.removeCollaborator
            }
            modalData={{
              message: isCollaboratorCurrentUser
                ? sharedMessages.collaboratorModalWarningSelf
                : {
                    values: { collaboratorId },
                    ...sharedMessages.collaboratorModalWarning,
                  },
            }}
            onApprove={handleDelete}
          />
        )}
      </SubmitBar>
    </Form>
  )
}

CollaboratorForm.propTypes = {
  collaborator: PropTypes.collaborator,
  collaboratorId: PropTypes.string,
  deleteDisabled: PropTypes.bool,
  error: PropTypes.error,
  isCollaboratorAdmin: PropTypes.bool.isRequired,
  isCollaboratorCurrentUser: PropTypes.bool.isRequired,
  isCollaboratorUser: PropTypes.bool.isRequired,
  onDelete: PropTypes.func,
  onDeleteFailure: PropTypes.func,
  onDeleteSuccess: PropTypes.func,
  onSubmit: PropTypes.func.isRequired,
  onSubmitFailure: PropTypes.func,
  onSubmitSuccess: PropTypes.func.isRequired,
  pseudoRights: PropTypes.rights,
  rights: PropTypes.rights.isRequired,
  update: PropTypes.bool,
}

CollaboratorForm.defaultProps = {
  deleteDisabled: false,
  onSubmitFailure: () => null,
  onDelete: () => null,
  onDeleteSuccess: () => null,
  onDeleteFailure: () => null,
  pseudoRights: [],
  error: undefined,
  collaborator: undefined,
  update: false,
  collaboratorId: undefined,
}

export default CollaboratorForm
