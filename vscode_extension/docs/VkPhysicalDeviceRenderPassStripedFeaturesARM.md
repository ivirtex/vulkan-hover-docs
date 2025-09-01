# VkPhysicalDeviceRenderPassStripedFeaturesARM(3) Manual Page

## Name

VkPhysicalDeviceRenderPassStripedFeaturesARM - Structure describing whether striped rendering can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRenderPassStripedFeaturesARM` structure is defined as:

```c++
// Provided by VK_ARM_render_pass_striped
typedef struct VkPhysicalDeviceRenderPassStripedFeaturesARM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           renderPassStriped;
} VkPhysicalDeviceRenderPassStripedFeaturesARM;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceRenderPassStripedFeaturesARM` structure describe the following features:

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`renderPassStriped` indicates that striped rendering is supported by the implementation.

If the `VkPhysicalDeviceRenderPassStripedFeaturesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRenderPassStripedFeaturesARM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRenderPassStripedFeaturesARM-sType-sType)VUID-VkPhysicalDeviceRenderPassStripedFeaturesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_FEATURES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_render\_pass\_striped](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_render_pass_striped.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRenderPassStripedFeaturesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0