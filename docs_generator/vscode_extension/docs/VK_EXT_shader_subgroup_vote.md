# VK\_EXT\_shader\_subgroup\_vote(3) Manual Page

## Name

VK\_EXT\_shader\_subgroup\_vote - device extension



## [](#_registered_extension_number)Registered Extension Number

66

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_subgroup\_vote](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_subgroup_vote.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-new-features)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_subgroup_vote%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_subgroup_vote%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-11-28

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_shader_group_vote`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_group_vote.txt)

**Contributors**

- Neil Henning, Codeplay
- Daniel Koch, NVIDIA Corporation

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_subgroup_vote`

This extension provides new SPIR-V instructions:

- `OpSubgroupAllKHR`,
- `OpSubgroupAnyKHR`, and
- `OpSubgroupAllEqualKHR`.

to compute the composite of a set of boolean conditions across a group of shader invocations that are running concurrently (a *subgroup*). These composite results may be used to execute shaders more efficiently on a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html).

When using GLSL source-based shader languages, the following shader functions from GL\_ARB\_shader\_group\_vote can map to these SPIR-V instructions:

- `anyInvocationARB`() → `OpSubgroupAnyKHR`,
- `allInvocationsARB`() → `OpSubgroupAllKHR`, and
- `allInvocationsEqualARB`() → `OpSubgroupAllEqualKHR`.

The subgroup across which the boolean conditions are evaluated is implementation-dependent, and this extension provides no guarantee over how individual shader invocations are assigned to subgroups. In particular, a subgroup has no necessary relationship with the compute shader *local workgroup* — any pair of shader invocations in a compute local workgroup may execute in different subgroups as used by these instructions.

Compute shaders operate on an explicitly specified group of threads (a local workgroup), but many implementations will also group non-compute shader invocations and execute them concurrently. When executing code like

```c++
if (condition) {
  result = do_fast_path();
} else {
  result = do_general_path();
}
```

where `condition` diverges between invocations, an implementation might first execute `do_fast_path`() for the invocations where `condition` is true and leave the other invocations dormant. Once `do_fast_path`() returns, it might call `do_general_path`() for invocations where `condition` is `false` and leave the other invocations dormant. In this case, the shader executes **both** the fast and the general path and might be better off just using the general path for all invocations.

This extension provides the ability to avoid divergent execution by evaluating a condition across an entire subgroup using code like:

```c++
if (allInvocationsARB(condition)) {
  result = do_fast_path();
} else {
  result = do_general_path();
}
```

The built-in function `allInvocationsARB`() will return the same value for all invocations in the group, so the group will either execute `do_fast_path`() or `do_general_path`(), but never both. For example, shader code might want to evaluate a complex function iteratively by starting with an approximation of the result and then refining the approximation. Some input values may require a small number of iterations to generate an accurate result (`do_fast_path`) while others require a larger number (`do_general_path`). In another example, shader code might want to evaluate a complex function (`do_general_path`) that can be greatly simplified when assuming a specific value for one of its inputs (`do_fast_path`).

## [](#_deprecated_by_vulkan_1_1)Deprecated by Vulkan 1.1

All functionality in this extension is superseded by the core Vulkan 1.1 [subgroup operations](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html).

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_SUBGROUP_VOTE_EXTENSION_NAME`
- `VK_EXT_SHADER_SUBGROUP_VOTE_SPEC_VERSION`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`SubgroupVoteKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-SubgroupVoteKHR)

## [](#_version_history)Version History

- Revision 1, 2016-11-28 (Daniel Koch)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_subgroup_vote)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0