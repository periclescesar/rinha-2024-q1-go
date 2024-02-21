#language en
Feature:

  @wip
  Scenario: invalid client
    When I make a debit of 100 to the customer's account with id 6 and description "debit of invalid client"
    Then I will receive a error "cliente não encontrado"

  @wip
  Scenario: first debit
    When I make a debit of 1000 to the customer's account with id 5 and description "first debit"
    Then I will see my limit of 100000 and balance of -1000

  @wip
  Scenario: debit bigger than limit
    When I make a debit of 100001 to the customer's account with id 5 and description "bigger debit"
    Then I will receive a error "saldo inconsistente"
