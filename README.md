# go-design-patterns
simple examples or each design pattern

## 2. Adapter
- 機能は似ているがインタフェースが異なるクラス同士を結合させる
- 継承パターンと委譲パターンの二つある
  - Goでは継承がないので委譲パターンしか実装できない
- 既存のクラスに一皮被せて利用したいとき
  - 既存に触れないでおくことができる
- バージョンアップと互換性に

## 8. Abstract Factory
- 抽象的な工場では、抽象的な部品を組み合わせて抽象的な製品を作る
- インタフェースだけを使って部品を組み立て、製品にまとめる
- 具体的な工場を新たに追加するのは簡単
- 部品を新たに追加するのは困難

## 9. Bridge
- 機能の階層と実装の階層を分けた上で結合させる
- 分けておけば、それぞれのクラス階層を独立に拡張することができる
  - 機能を足しても実装に影響はないし、逆も然り
  - 追加した機能は、全ての実装で利用できる