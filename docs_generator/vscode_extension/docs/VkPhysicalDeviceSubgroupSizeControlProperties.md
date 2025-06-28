# VkPhysicalDeviceSubgroupSizeControlProperties(3) Manual Page

## Name

VkPhysicalDeviceSubgroupSizeControlProperties - Structure describing the control subgroup size properties of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSubgroupSizeControlProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceSubgroupSizeControlProperties {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              minSubgroupSize;
    uint32_t              maxSubgroupSize;
    uint32_t              maxComputeWorkgroupSubgroups;
    VkShaderStageFlags    requiredSubgroupSizeStages;
} VkPhysicalDeviceSubgroupSizeControlProperties;
```

or the equivalent

```c++
// Provided by VK_EXT_subgroup_size_control
typedef VkPhysicalDeviceSubgroupSizeControlProperties VkPhysicalDeviceSubgroupSizeControlPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`minSubgroupSize` is the minimum subgroup size supported by this device. `minSubgroupSize` is at least one if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `minSubgroupSize` is a power-of-two. `minSubgroupSize` is less than or equal to `maxSubgroupSize`. `minSubgroupSize` is less than or equal to [`subgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSize).
- []()`maxSubgroupSize` is the maximum subgroup size supported by this device. `maxSubgroupSize` is at least one if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `maxSubgroupSize` is a power-of-two. `maxSubgroupSize` is greater than or equal to `minSubgroupSize`. `maxSubgroupSize` is greater than or equal to [`subgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSize).
- []()`maxComputeWorkgroupSubgroups` is the maximum number of subgroups supported by the implementation within a workgroup.
- []()`requiredSubgroupSizeStages` is a bitfield of what shader stages support having a required subgroup size specified.

If the `VkPhysicalDeviceSubgroupSizeControlProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

If [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations` includes [`VK_SUBGROUP_FEATURE_QUAD_BIT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroup-quad), `minSubgroupSize` **must** be greater than or equal to 4.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSubgroupSizeControlProperties-sType-sType)VUID-VkPhysicalDeviceSubgroupSizeControlProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES`

## [](#_see_also)See Also

[VK\_EXT\_subgroup\_size\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subgroup_size_control.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSubgroupSizeControlProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0