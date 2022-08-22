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

describe('Account App authorization management', () => {
  const user = {
    ids: { user_id: 'create-app-test-user' },
    primary_email_address: 'create-app-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  beforeEach(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
    cy.visitConsoleAuthorizationScreen({ user_id: user.ids.user_id, password: user.password })
    cy.findByRole('button', { name: /Authorize Console/ }).click()
  })

  it('succeeds showing authorized clients', () => {
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations`)

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 1)
    })
    cy.findByRole('cell', { name: 'console' }).should('be.visible')
  })

  it('succeeds viewing an auhtorized client', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console`)
    cy.findByText('Authorization settings').should('be.visible')
    cy.findByText('Active access tokens').should('be.visible')
    cy.findByText('General information').should('be.visible')
    cy.findByText('Actions').should('be.visible')
  })

  it('succeeds showing active authorization tokens', () => {
    const tokens = [
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
      {
        user_ids: {
          user_id: 'create-app-test-user',
        },
        user_session_id: '280eba34-a77c-4223-8bfb-d9c69c673441',
        client_ids: {
          client_id: 'console',
        },
        id: 'YAZ4CWZTBZ46MVEMYLGWB4G7CAV37MYW6GFXBGQ',
        rights: ['RIGHT_ALL'],
        created_at: '2022-08-22T08:42:39.031Z',
        expires_at: '2022-08-22T09:42:39.031Z',
      },
    ]
    cy.intercept('GET', '/api/v3/users/create-app-test-user/authorizations/console/tokens', {
      tokens,
    })
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/tokens`)
    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 1)
    })
  })

  it('succeeds revoking authorizations', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console`)
    cy.findByRole('button', { name: /Revoke authorization/ }).click()
    cy.findByTestId('modal-window')
      .should('be.visible')
      .within(() => {
        cy.findByText('Revoke authorization', { selector: 'h1' }).should('be.visible')
        cy.findByRole('button', { name: /Revoke authorization/ }).click()
      })
    cy.findByTestId('error-notification').should('not.exist')

    cy.location('pathname').should(
      'eq',
      `${Cypress.config('accountAppRootPath')}/client-authorizations`,
    )

    cy.findByRole('cell', { name: 'console' }).should('not.exist')
  })

  it('succeeds invalidating all tokens', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/tokens`)
    cy.findByRole('button', { name: /Invalidate all access tokens/ }).click()
    cy.findByTestId('error-notification').should('not.exist')

    cy.location('pathname').should(
      'eq',
      `${Cypress.config('accountAppRootPath')}/client-authorizations/console`,
    )

    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/tokens`)
    cy.findByRole('rowgroup').should('not.exist')
  })

  it('succeeds invalidating one token', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/tokens`)
    cy.findByRole('button', { name: /Invalidate this access token/ }).click()
    cy.findByTestId('error-notification').should('not.exist')

    cy.location('pathname').should(
      'eq',
      `${Cypress.config('accountAppRootPath')}/client-authorizations/console`,
    )
  })
})
