# VkPhysicalDeviceLineRasterizationProperties(3) Manual Page

## Name

VkPhysicalDeviceLineRasterizationProperties - Structure describing line rasterization properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceLineRasterizationProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceLineRasterizationProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           lineSubPixelPrecisionBits;
} VkPhysicalDeviceLineRasterizationProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_line_rasterization
typedef VkPhysicalDeviceLineRasterizationProperties VkPhysicalDeviceLineRasterizationPropertiesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_line_rasterization
typedef VkPhysicalDeviceLineRasterizationProperties VkPhysicalDeviceLineRasterizationPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`lineSubPixelPrecisionBits` is the number of bits of subpixel precision in framebuffer coordinates xf and yf when rasterizing [line segments](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines).

If the `VkPhysicalDeviceLineRasterizationProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceLineRasterizationProperties-sType-sType)VUID-VkPhysicalDeviceLineRasterizationProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES`

## [](#_see_also)See Also

[VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html), [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceLineRasterizationProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0