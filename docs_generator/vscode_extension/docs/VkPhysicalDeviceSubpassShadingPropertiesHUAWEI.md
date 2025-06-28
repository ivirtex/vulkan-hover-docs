# VkPhysicalDeviceSubpassShadingPropertiesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceSubpassShadingPropertiesHUAWEI - Structure describing subpass shading properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSubpassShadingPropertiesHUAWEI` structure is defined as:

```c++
// Provided by VK_HUAWEI_subpass_shading
typedef struct VkPhysicalDeviceSubpassShadingPropertiesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxSubpassShadingWorkgroupSizeAspectRatio;
} VkPhysicalDeviceSubpassShadingPropertiesHUAWEI;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxSubpassShadingWorkgroupSizeAspectRatio` indicates the maximum ratio between the width and height of the portion of the subpass shading shader workgroup size. `maxSubpassShadingWorkgroupSizeAspectRatio` **must** be a power-of-two value, and **must** be less than or equal to max(`WorkgroupSize.x` / `WorkgroupSize.y`, `WorkgroupSize.y` / `WorkgroupSize.x`).

## [](#_description)Description

If the `VkPhysicalDeviceSubpassShadingPropertiesHUAWEI` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSubpassShadingPropertiesHUAWEI-sType-sType)VUID-VkPhysicalDeviceSubpassShadingPropertiesHUAWEI-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_SHADING_PROPERTIES_HUAWEI`

## [](#_see_also)See Also

[VK\_HUAWEI\_subpass\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_subpass_shading.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSubpassShadingPropertiesHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0