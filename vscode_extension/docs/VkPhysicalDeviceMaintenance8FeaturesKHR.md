# VkPhysicalDeviceMaintenance8FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance8FeaturesKHR - Structure describing whether the implementation supports maintenance8 functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance8FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance8
typedef struct VkPhysicalDeviceMaintenance8FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance8;
} VkPhysicalDeviceMaintenance8FeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maintenance8` indicates that the implementation supports the following:
  
  - Allow copies between depth/stencil and “matching” color attachments
  - Allow `dstCache` in `vkMergePipelineCaches` to be implicitly synchronized.
  - Require src/dst sync scopes to work when doing queue family ownership transfers
  - Support `Offset` (as an alternative to `ConstOffset`) image operand in texture sampling and fetch operations
  - Use the SPIR-V definition of OpSRem and OpSMod, making these operations produce well-defined results for negative operands
  - Loosen layer restrictions when blitting from 3D images to other image types
  - Add space for an additional 64 access flags for use with VkMemoryBarrier2, VkBufferMemoryBarrier2, and VkImageMemoryBarrier2

## [](#_description)Description

If the `VkPhysicalDeviceMaintenance8FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMaintenance8FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance8FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceMaintenance8FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_8_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_maintenance8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance8.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance8FeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0