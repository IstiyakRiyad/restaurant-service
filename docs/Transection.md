# Transections


## Transection 1
BEGIN
Account 1 Update
Account 2 Update
COMMIT

## Transection 2
BEGIN
Account 2 Update
Account 1 Update
COMMIT


___

## Transection 1
BEGIN
Account 1 Update

## Transection 2
BEGIN
Account 2 Update

## Transection 1
Account 2 Update (waiting for fee)
COMMIT

## Transection 2
Account 1 Update (waiting for fee)
COMMIT (deadlock)


___
# Avoid Deadlock
compare the account id and always start transection with small id number. This will solve this problem












