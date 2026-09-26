  package pr211
//        ^^^^^ definition 0.1.test `sg/pr211`/
//              kind Package
//              display_name pr211
//              signature_documentation
//              > package pr211
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/Map#
  // Map is a generic map type.
  type Map[K comparable, V any] struct {
//     ^^^ definition 0.1.test `sg/pr211`/Map#
//         kind Struct
//         display_name Map
//         signature_documentation
//         > type Map struct{ entries []entry[K, V] }
//         documentation
//         > Map is a generic map type.
//         ^ definition local 0
//           kind Interface
//           display_name K
//           signature_documentation
//           > type parameter K comparable
//                       ^ definition local 1
//                         kind Interface
//                         display_name V
//                         signature_documentation
//                         > type parameter V any
// ⌄ enclosing_range_start 0.1.test `sg/pr211`/Map#entries.
   entries []entry[K, V]
// ^^^^^^^ definition 0.1.test `sg/pr211`/Map#entries.
//         kind Field
//         display_name entries
//         signature_documentation
//         > struct field entries []entry[K, V]
//           ^^^^^ reference 0.1.test `sg/pr211`/entry#
//                 ^ reference local 0
//                    ^ reference local 1
//                     ⌃ enclosing_range_end 0.1.test `sg/pr211`/Map#entries.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr211`/Map#
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/entry#
  type entry[K comparable, V any] struct {
//     ^^^^^ definition 0.1.test `sg/pr211`/entry#
//           kind Struct
//           display_name entry
//           signature_documentation
//           > type entry struct {
//           >     key   K
//           >     value V
//           > }
//           ^ definition local 2
//             kind Interface
//             display_name K
//             signature_documentation
//             > type parameter K comparable
//                         ^ definition local 3
//                           kind Interface
//                           display_name V
//                           signature_documentation
//                           > type parameter V any
// ⌄ enclosing_range_start 0.1.test `sg/pr211`/entry#key.
   key   K
// ^^^ definition 0.1.test `sg/pr211`/entry#key.
//     kind Field
//     display_name key
//     signature_documentation
//     > struct field key K
//       ^ reference local 2
//       ⌃ enclosing_range_end 0.1.test `sg/pr211`/entry#key.
// ⌄ enclosing_range_start 0.1.test `sg/pr211`/entry#value.
   value V
// ^^^^^ definition 0.1.test `sg/pr211`/entry#value.
//       kind Field
//       display_name value
//       signature_documentation
//       > struct field value V
//       ^ reference local 3
//       ⌃ enclosing_range_end 0.1.test `sg/pr211`/entry#value.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr211`/entry#
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/Set#
  // Set is a generic alias that partially instantiates Map.
  type Set[K comparable] = Map[K, bool]
//     ^^^ definition 0.1.test `sg/pr211`/Set#
//         kind TypeAlias
//         display_name Set
//         signature_documentation
//         > type Set[K comparable] = Map[K, bool]
//         documentation
//         > Set is a generic alias that partially instantiates Map.
//         ^ definition local 4
//           kind Interface
//           display_name K
//           signature_documentation
//           > type parameter K comparable
//                         ^^^ reference 0.1.test `sg/pr211`/Map#
//                             ^ reference local 4
//                                    ⌃ enclosing_range_end 0.1.test `sg/pr211`/Set#
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/OrderedSet#
  // Alias with a tighter constraint.
  type OrderedSet[K ~int | ~string] = Set[K]
//     ^^^^^^^^^^ definition 0.1.test `sg/pr211`/OrderedSet#
//                kind TypeAlias
//                display_name OrderedSet
//                signature_documentation
//                > type OrderedSet[K ~int | ~string] = Set[K]
//                documentation
//                > Alias with a tighter constraint.
//                ^ definition local 5
//                  kind Interface
//                  display_name K
//                  signature_documentation
//                  > type parameter K ~int | ~string
//                                    ^^^ reference 0.1.test `sg/pr211`/Set#
//                                        ^ reference local 5
//                                         ⌃ enclosing_range_end 0.1.test `sg/pr211`/OrderedSet#
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/StringSet#
  // Alias of an alias (chained).
  type StringSet = Set[string]
//     ^^^^^^^^^ definition 0.1.test `sg/pr211`/StringSet#
//               kind TypeAlias
//               display_name StringSet
//               signature_documentation
//               > type StringSet = Set[string]
//               documentation
//               > Alias of an alias (chained).
//                 ^^^ reference 0.1.test `sg/pr211`/Set#
//                           ⌃ enclosing_range_end 0.1.test `sg/pr211`/StringSet#
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/PairMap#
  // Alias with all params forwarded.
  type PairMap[K comparable, V any] = Map[K, V]
//     ^^^^^^^ definition 0.1.test `sg/pr211`/PairMap#
//             kind TypeAlias
//             display_name PairMap
//             signature_documentation
//             > type PairMap[K comparable, V any] = Map[K, V]
//             documentation
//             > Alias with all params forwarded.
//             ^ definition local 6
//               kind Interface
//               display_name K
//               signature_documentation
//               > type parameter K comparable
//                           ^ definition local 7
//                             kind Interface
//                             display_name V
//                             signature_documentation
//                             > type parameter V any
//                                    ^^^ reference 0.1.test `sg/pr211`/Map#
//                                        ^ reference local 6
//                                           ^ reference local 7
//                                            ⌃ enclosing_range_end 0.1.test `sg/pr211`/PairMap#
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/UseAliases().
  func UseAliases() {
//     ^^^^^^^^^^ definition 0.1.test `sg/pr211`/UseAliases().
//                kind Function
//                display_name UseAliases
//                signature_documentation
//                > func UseAliases()
   _ = Set[int]{}
//     ^^^ reference 0.1.test `sg/pr211`/Set#
   _ = OrderedSet[int]{}
//     ^^^^^^^^^^ reference 0.1.test `sg/pr211`/OrderedSet#
   _ = StringSet{}
//     ^^^^^^^^^ reference 0.1.test `sg/pr211`/StringSet#
   _ = PairMap[string, int]{}
//     ^^^^^^^ reference 0.1.test `sg/pr211`/PairMap#
  }
//⌃ enclosing_range_end 0.1.test `sg/pr211`/UseAliases().
  
