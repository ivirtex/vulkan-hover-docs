# VkSubgroupFeatureFlagBits(3) Manual Page

## Name

VkSubgroupFeatureFlagBits - Bitmask describing what group operations are
supported with subgroup scope



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations`
and
[VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Properties.html)::`subgroupSupportedOperations`
to specify supported <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
target="_blank" rel="noopener">group operations</a> with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-subgroup"
target="_blank" rel="noopener">subgroup scope</a> are:

``` c
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
  // Provided by VK_NV_shader_subgroup_partitioned
    VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV = 0x00000100,
  // Provided by VK_KHR_shader_subgroup_rotate
    VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR = 0x00000200,
  // Provided by VK_KHR_shader_subgroup_rotate
    VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR = 0x00000400,
} VkSubgroupFeatureFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- <span id="features-subgroup-basic"></span>
  `VK_SUBGROUP_FEATURE_BASIC_BIT` specifies the device will accept
  SPIR-V shader modules containing the `GroupNonUniform` capability.

- <span id="features-subgroup-vote"></span>
  `VK_SUBGROUP_FEATURE_VOTE_BIT` specifies the device will accept SPIR-V
  shader modules containing the `GroupNonUniformVote` capability.

- <span id="features-subgroup-arithmetic"></span>
  `VK_SUBGROUP_FEATURE_ARITHMETIC_BIT` specifies the device will accept
  SPIR-V shader modules containing the `GroupNonUniformArithmetic`
  capability.

- <span id="features-subgroup-ballot"></span>
  `VK_SUBGROUP_FEATURE_BALLOT_BIT` specifies the device will accept
  SPIR-V shader modules containing the `GroupNonUniformBallot`
  capability.

- <span id="features-subgroup-shuffle"></span>
  `VK_SUBGROUP_FEATURE_SHUFFLE_BIT` specifies the device will accept
  SPIR-V shader modules containing the `GroupNonUniformShuffle`
  capability.

- <span id="features-subgroup-shuffle-relative"></span>
  `VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT` specifies the device will
  accept SPIR-V shader modules containing the
  `GroupNonUniformShuffleRelative` capability.

- <span id="features-subgroup-clustered"></span>
  `VK_SUBGROUP_FEATURE_CLUSTERED_BIT` specifies the device will accept
  SPIR-V shader modules containing the `GroupNonUniformClustered`
  capability.

- <span id="features-subgroup-quad"></span>
  `VK_SUBGROUP_FEATURE_QUAD_BIT` specifies the device will accept SPIR-V
  shader modules containing the `GroupNonUniformQuad` capability.

- <span id="features-subgroup-partitioned"></span>
  `VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV` specifies the device will
  accept SPIR-V shader modules containing the
  `GroupNonUniformPartitionedNV` capability.

- <span id="features-subgroup-rotate"></span>
  `VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR` specifies the device will accept
  SPIR-V shader modules containing the `GroupNonUniformRotateKHR`
  capability.

- <span id="features-subgroup-rotate-clustered"></span>
  `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR` specifies the device
  will accept SPIR-V shader modules that use the `ClusterSize` operand
  to `OpGroupNonUniformRotateKHR`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubgroupFeatureFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
