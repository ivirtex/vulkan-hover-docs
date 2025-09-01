# VkPhysicalDeviceSubgroupProperties(3) Manual Page

## Name

VkPhysicalDeviceSubgroupProperties - Structure describing subgroup support for an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSubgroupProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceSubgroupProperties {
    VkStructureType           sType;
    void*                     pNext;
    uint32_t                  subgroupSize;
    VkShaderStageFlags        supportedStages;
    VkSubgroupFeatureFlags    supportedOperations;
    VkBool32                  quadOperationsInAllStages;
} VkPhysicalDeviceSubgroupProperties;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`subgroupSize` is the default number of invocations in each subgroup. `subgroupSize` is at least 1 if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `subgroupSize` is a power-of-two.
- []()`supportedStages` is a bitfield of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) describing the shader stages that [group operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-group-operations) with [subgroup scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-subgroup) are supported in. `supportedStages` will have the `VK_SHADER_STAGE_COMPUTE_BIT` bit set if any of the physical device’s queues support `VK_QUEUE_COMPUTE_BIT`.
- []()`supportedOperations` is a bitmask of [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlagBits.html) specifying the sets of [group operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-group-operations) with [subgroup scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-subgroup) supported on this device. `supportedOperations` will have the `VK_SUBGROUP_FEATURE_BASIC_BIT` bit set if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`.
- []()`quadOperationsInAllStages` is a boolean specifying whether [quad group operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-quad-operations) are available in all stages, or are restricted to fragment and compute stages.

If the `VkPhysicalDeviceSubgroupProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

If `supportedOperations` includes [`VK_SUBGROUP_FEATURE_QUAD_BIT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroup-quad), or the [`shaderSubgroupUniformControlFlow`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupUniformControlFlow) feature is enabled, `subgroupSize` **must** be greater than or equal to 4.

If the [`shaderQuadControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderQuadControl) feature is supported, `supportedOperations` **must** include [`VK_SUBGROUP_FEATURE_QUAD_BIT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroup-quad).

If [VK\_KHR\_shader\_subgroup\_rotate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_rotate.html) is supported, and the implementation advertises support with a [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html)::`specVersion` greater than or equal to 2, and the [`shaderSubgroupRotate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupRotate) feature is supported, `VK_SUBGROUP_FEATURE_ROTATE_BIT` **must** be returned in [VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan11Properties.html)::`subgroupSupportedOperations` and [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations`. If [VK\_KHR\_shader\_subgroup\_rotate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_rotate.html) is supported, and the implementation advertises support with a [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html)::`specVersion` greater than or equal to 2, and the [`shaderSubgroupRotateClustered`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupRotateClustered) feature is supported, `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT` **must** be returned in [VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan11Properties.html)::`subgroupSupportedOperations` and [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations`.

If Vulkan 1.4 is supported, `VK_SUBGROUP_FEATURE_ROTATE_BIT` and `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT` **must** be returned in [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations` and [VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan11Properties.html)::`subgroupSupportedOperations`

Note

`VK_SUBGROUP_FEATURE_ROTATE_BIT` and `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT` were added in version 2 of the [VK\_KHR\_shader\_subgroup\_rotate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_rotate.html) extension, after the initial release, so there are implementations that do not advertise these bits. Applications should use the [`shaderSubgroupRotate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupRotate) and [`shaderSubgroupRotateClustered`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupRotateClustered) features to determine and enable support. These bits are advertised here for consistency and for future dependencies.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSubgroupProperties-sType-sType)VUID-VkPhysicalDeviceSubgroupProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSubgroupProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0