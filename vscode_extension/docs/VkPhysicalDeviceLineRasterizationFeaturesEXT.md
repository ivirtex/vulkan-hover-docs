# VkPhysicalDeviceLineRasterizationFeatures(3) Manual Page

## Name

VkPhysicalDeviceLineRasterizationFeatures - Structure describing the line rasterization features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceLineRasterizationFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceLineRasterizationFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rectangularLines;
    VkBool32           bresenhamLines;
    VkBool32           smoothLines;
    VkBool32           stippledRectangularLines;
    VkBool32           stippledBresenhamLines;
    VkBool32           stippledSmoothLines;
} VkPhysicalDeviceLineRasterizationFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_line_rasterization
typedef VkPhysicalDeviceLineRasterizationFeatures VkPhysicalDeviceLineRasterizationFeaturesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_line_rasterization
typedef VkPhysicalDeviceLineRasterizationFeatures VkPhysicalDeviceLineRasterizationFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`rectangularLines` indicates whether the implementation supports [rectangular line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines).
- []()`bresenhamLines` indicates whether the implementation supports [Bresenham-style line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-bresenham).
- []()`smoothLines` indicates whether the implementation supports [smooth line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-smooth).
- []()`stippledRectangularLines` indicates whether the implementation supports [stippled line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-stipple) with `VK_LINE_RASTERIZATION_MODE_RECTANGULAR` lines.
- []()`stippledBresenhamLines` indicates whether the implementation supports [stippled line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-stipple) with `VK_LINE_RASTERIZATION_MODE_BRESENHAM` lines.
- []()`stippledSmoothLines` indicates whether the implementation supports [stippled line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-stipple) with `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH` lines.

If the `VkPhysicalDeviceLineRasterizationFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceLineRasterizationFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceLineRasterizationFeatures-sType-sType)VUID-VkPhysicalDeviceLineRasterizationFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES`

## [](#_see_also)See Also

[VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html), [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceLineRasterizationFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0