# VK\_EXT\_shader\_subgroup\_ballot(3) Manual Page

## Name

VK\_EXT\_shader\_subgroup\_ballot - device extension



## [](#_registered_extension_number)Registered Extension Number

65

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_shader\_ballot](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_shader_ballot.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-new-features)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_subgroup_ballot%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_subgroup_ballot%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-11-28

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_shader_ballot`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_ballot.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Neil Henning, Codeplay
- Daniel Koch, NVIDIA Corporation

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_shader_ballot`

This extension provides the ability for a group of invocations, which execute in parallel, to do limited forms of cross-invocation communication via a group broadcast of an invocation value, or broadcast of a bit array representing a predicate value from each invocation in the group.

This extension provides access to a number of additional built-in shader variables in Vulkan:

- `SubgroupEqMaskKHR`, containing the subgroup mask of the current subgroup invocation,
- `SubgroupGeMaskKHR`, containing the subgroup mask of the invocations greater than or equal to the current invocation,
- `SubgroupGtMaskKHR`, containing the subgroup mask of the invocations greater than the current invocation,
- `SubgroupLeMaskKHR`, containing the subgroup mask of the invocations less than or equal to the current invocation,
- `SubgroupLtMaskKHR`, containing the subgroup mask of the invocations less than the current invocation,
- `SubgroupLocalInvocationId`, containing the index of an invocation within a subgroup, and
- `SubgroupSize`, containing the maximum number of invocations in a subgroup.

Additionally, this extension provides access to the new SPIR-V instructions:

- `OpSubgroupBallotKHR`,
- `OpSubgroupFirstInvocationKHR`, and
- `OpSubgroupReadInvocationKHR`,

When using GLSL source-based shader languages, the following variables and shader functions from GL\_ARB\_shader\_ballot can map to these SPIR-V built-in decorations and instructions:

- `in uint64_t gl_SubGroupEqMaskARB;` → `SubgroupEqMaskKHR`,
- `in uint64_t gl_SubGroupGeMaskARB;` → `SubgroupGeMaskKHR`,
- `in uint64_t gl_SubGroupGtMaskARB;` → `SubgroupGtMaskKHR`,
- `in uint64_t gl_SubGroupLeMaskARB;` → `SubgroupLeMaskKHR`,
- `in uint64_t gl_SubGroupLtMaskARB;` → `SubgroupLtMaskKHR`,
- `in uint gl_SubGroupInvocationARB;` → `SubgroupLocalInvocationId`,
- `uniform uint gl_SubGroupSizeARB;` → `SubgroupSize`,
- `ballotARB`() → `OpSubgroupBallotKHR`,
- `readFirstInvocationARB`() → `OpSubgroupFirstInvocationKHR`, and
- `readInvocationARB`() → `OpSubgroupReadInvocationKHR`.

## [](#_deprecated_by_vulkan_1_2)Deprecated by Vulkan 1.2

Most of the functionality in this extension is superseded by the core Vulkan 1.1 [subgroup operations](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html). However, Vulkan 1.1 required the `OpGroupNonUniformBroadcast` “Id” to be constant. This restriction was removed in Vulkan 1.2 with the addition of the [`subgroupBroadcastDynamicId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroupBroadcastDynamicId) feature.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_SUBGROUP_BALLOT_EXTENSION_NAME`
- `VK_EXT_SHADER_SUBGROUP_BALLOT_SPEC_VERSION`

## [](#_new_built_in_variables)New Built-In Variables

- [`SubgroupEqMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgeq)
- [`SubgroupGeMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgge)
- [`SubgroupGtMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sggt)
- [`SubgroupLeMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgle)
- [`SubgroupLtMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sglt)
- [`SubgroupLocalInvocationId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgli)
- [`SubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`SubgroupBallotKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-SubgroupBallotKHR)

## [](#_version_history)Version History

- Revision 1, 2016-11-28 (Daniel Koch)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_subgroup_ballot)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0