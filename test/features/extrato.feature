#language en
Feature:
  Scenario: invalid client
    When I get a account stattement of the customer's id 6
    Then I will receive a error "cliente não encontrado"

  Scenario: valid stattement
    When I get a account stattement of the customer's id 5
    Then I will see my limit of 100000 and balance of -1000
    And the lasts 10 transactions
