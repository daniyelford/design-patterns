<?php
interface Specification {
    public function isSatisfiedBy(array $user): bool;
}

class AgeAbove18 implements Specification {
    public function isSatisfiedBy(array $user): bool {
        return $user['age'] > 18;
    }
}

class HasEmail implements Specification {
    public function isSatisfiedBy(array $user): bool {
        return !empty($user['email']);
    }
}

class AndSpec implements Specification {
    private Specification $a; private Specification $b;
    public function __construct(Specification $a, Specification $b){
        $this->a = $a; $this->b = $b;
    }
    public function isSatisfiedBy(array $user): bool {
        return $this->a->isSatisfiedBy($user) && $this->b->isSatisfiedBy($user);
    }
}

// استفاده
$user = ['age'=>20, 'email'=>'test@example.com'];
$spec = new AndSpec(new AgeAbove18(), new HasEmail());
echo $spec->isSatisfiedBy($user) ? "Ok" : "Rejected";
