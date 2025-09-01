# VkSubgroupFeatureFlagBits(3) Manual Page

## Name

VkSubgroupFeatureFlagBits - Bitmask describing what group operations are supported with subgroup scope



## [](#_c_specification)C Specification

Bits which **can** be set in [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations` and [VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan11Properties.html)::`subgroupSupportedOperations` to specify supported [group operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-group-operations) with [subgroup scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-subgroup) are:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkSubgroupFeatureFlagBits {
    VK_SUBGROUP_FEATURE_BASIC_BIT = 0x00000001,
    VK_SUBGROUP_FEATURE_VOTE_BIT = 0x00000002,
    VK_SUBGROUP_FEATURE_ARITHMETIC_BIT = 0x00000004,
    VK_SUBGROUP_FEATURE_BALLOT_BIT = 0x00000008,
    VK_SUBGROUP_FEATURE_SHUFFLE_BIT = 0x00000010,
    VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT = 0x00000020,
    VK_SUBGROUP_FEATURE_CLUSTERED_BIT = 0x00000040,
    VK_SUBGROUP_FEATURE_QUAD_BIT = 0x00000080,
  // Provided by VK_VERSION_1_4
    VK_SUBGROUP_FEATURE_ROTATE_BIT = 0x00000200,
  // Provided by VK_VERSION_1_4
    VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT = 0x00000400,
  // Provided by VK_NV_shader_subgroup_partitioned
    VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV = 0x00000100,
  // Provided by VK_KHR_shader_subgroup_rotate
    VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR = VK_SUBGROUP_FEATURE_ROTATE_BIT,
  // Provided by VK_KHR_shader_subgroup_rotate
    VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR = VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT,
} VkSubgroupFeatureFlagBits;
```

## [](#_description)Description

- []()`VK_SUBGROUP_FEATURE_BASIC_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniform` capability.
- []()`VK_SUBGROUP_FEATURE_VOTE_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformVote` capability.
- []()`VK_SUBGROUP_FEATURE_ARITHMETIC_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformArithmetic` capability.
- []()`VK_SUBGROUP_FEATURE_BALLOT_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformBallot` capability.
- []()`VK_SUBGROUP_FEATURE_SHUFFLE_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformShuffle` capability.
- []()`VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformShuffleRelative` capability.
- []()`VK_SUBGROUP_FEATURE_CLUSTERED_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformClustered` capability.
- []()`VK_SUBGROUP_FEATURE_QUAD_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformQuad` capability.
- []()`VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformPartitionedNV` capability.
- []()`VK_SUBGROUP_FEATURE_ROTATE_BIT` specifies the device will accept SPIR-V shader modules containing the `GroupNonUniformRotateKHR` capability.
- []()`VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT` specifies the device will accept SPIR-V shader modules that use the `ClusterSize` operand to `OpGroupNonUniformRotateKHR`.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubgroupFeatureFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0