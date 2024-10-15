#language en
Feature:
  Scenario: invalid client
    When I get a account statement of the customer's id 6
    Then I will receive a error "cliente não encontrado"

  Scenario: valid client
    When I get a account statement of the customer's id 5
    Then I will see my statement with limit of 500000 and balance of 0
    And the lasts 0 transactions
