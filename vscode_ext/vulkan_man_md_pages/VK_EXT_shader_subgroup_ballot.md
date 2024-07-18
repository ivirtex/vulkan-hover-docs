# VK_EXT_shader_subgroup_ballot(3) Manual Page

## Name

VK_EXT_shader_subgroup_ballot - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

65

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_shader_ballot](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_shader_ballot.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-new-features"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_subgroup_ballot%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_subgroup_ballot%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-11-28

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_shader_ballot`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_ballot.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Neil Henning, Codeplay

- Daniel Koch, NVIDIA Corporation

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_KHR_shader_ballot`

This extension provides the ability for a group of invocations, which
execute in parallel, to do limited forms of cross-invocation
communication via a group broadcast of an invocation value, or broadcast
of a bit array representing a predicate value from each invocation in
the group.

This extension provides access to a number of additional built-in shader
variables in Vulkan:

- `SubgroupEqMaskKHR`, containing the subgroup mask of the current
  subgroup invocation,

- `SubgroupGeMaskKHR`, containing the subgroup mask of the invocations
  greater than or equal to the current invocation,

- `SubgroupGtMaskKHR`, containing the subgroup mask of the invocations
  greater than the current invocation,

- `SubgroupLeMaskKHR`, containing the subgroup mask of the invocations
  less than or equal to the current invocation,

- `SubgroupLtMaskKHR`, containing the subgroup mask of the invocations
  less than the current invocation,

- `SubgroupLocalInvocationId`, containing the index of an invocation
  within a subgroup, and

- `SubgroupSize`, containing the maximum number of invocations in a
  subgroup.

Additionally, this extension provides access to the new SPIR-V
instructions:

- `OpSubgroupBallotKHR`,

- `OpSubgroupFirstInvocationKHR`, and

- `OpSubgroupReadInvocationKHR`,

When using GLSL source-based shader languages, the following variables
and shader functions from GL_ARB_shader_ballot can map to these SPIR-V
built-in decorations and instructions:

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

## <a href="#_deprecated_by_vulkan_1_2" class="anchor"></a>Deprecated by Vulkan 1.2

Most of the functionality in this extension is superseded by the core
Vulkan 1.1
<a href="VkPhysicalDeviceSubgroupProperties.html" target="_blank"
rel="noopener">subgroup operations</a>. However, Vulkan 1.1 required the
`OpGroupNonUniformBroadcast` “Id” to be constant. This restriction was
removed in Vulkan 1.2 with the addition of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroupBroadcastDynamicId"
target="_blank"
rel="noopener"><code>subgroupBroadcastDynamicId</code></a> feature.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_SUBGROUP_BALLOT_EXTENSION_NAME`

- `VK_EXT_SHADER_SUBGROUP_BALLOT_SPEC_VERSION`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgeq"
  target="_blank" rel="noopener"><code>SubgroupEqMaskKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgge"
  target="_blank" rel="noopener"><code>SubgroupGeMaskKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sggt"
  target="_blank" rel="noopener"><code>SubgroupGtMaskKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgle"
  target="_blank" rel="noopener"><code>SubgroupLeMaskKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sglt"
  target="_blank" rel="noopener"><code>SubgroupLtMaskKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgli"
  target="_blank"
  rel="noopener"><code>SubgroupLocalInvocationId</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgs"
  target="_blank" rel="noopener"><code>SubgroupSize</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-SubgroupBallotKHR"
  target="_blank" rel="noopener"><code>SubgroupBallotKHR</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-11-28 (Daniel Koch)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_subgroup_ballot"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
