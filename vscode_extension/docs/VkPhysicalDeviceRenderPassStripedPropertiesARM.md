# VkPhysicalDeviceRenderPassStripedPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceRenderPassStripedPropertiesARM - Structure describing striped rendering limits of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRenderPassStripedPropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_render_pass_striped
typedef struct VkPhysicalDeviceRenderPassStripedPropertiesARM {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         renderPassStripeGranularity;
    uint32_t           maxRenderPassStripes;
} VkPhysicalDeviceRenderPassStripedPropertiesARM;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceRenderPassStripedPropertiesARM` structure describe the following limits:

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`renderPassStripeGranularity` indicates the minimum supported granularity of striped render pass regions.
- []()`maxRenderPassStripes` indicates the maximum number of stripes supported in striped rendering.

If the `VkPhysicalDeviceRenderPassStripedPropertiesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRenderPassStripedPropertiesARM-sType-sType)VUID-VkPhysicalDeviceRenderPassStripedPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_PROPERTIES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_render\_pass\_striped](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_render_pass_striped.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRenderPassStripedPropertiesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0