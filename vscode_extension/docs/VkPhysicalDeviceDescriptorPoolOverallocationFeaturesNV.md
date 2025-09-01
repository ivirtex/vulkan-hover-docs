# VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV - Structure describing feature to allow descriptor pool overallocation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_descriptor_pool_overallocation
typedef struct VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           descriptorPoolOverallocation;
} VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`descriptorPoolOverallocation` indicates that the implementation allows the application to opt into descriptor pool overallocation by creating the descriptor pool with `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV` and/or `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV` flags.

## [](#_description)Description

If the `VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV-sType-sType)VUID-VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_POOL_OVERALLOCATION_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_descriptor\_pool\_overallocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_descriptor_pool_overallocation.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0