# VkLineRasterizationMode(3) Manual Page

## Name

VkLineRasterizationMode - Line rasterization modes



## [](#_c_specification)C Specification

Possible values of [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html)::`lineRasterizationMode`, specifying the line rasterization mode, are:

```c++
// Provided by VK_VERSION_1_4
typedef enum VkLineRasterizationMode {
    VK_LINE_RASTERIZATION_MODE_DEFAULT = 0,
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR = 1,
    VK_LINE_RASTERIZATION_MODE_BRESENHAM = 2,
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH = 3,
  // Provided by VK_EXT_line_rasterization
    VK_LINE_RASTERIZATION_MODE_DEFAULT_EXT = VK_LINE_RASTERIZATION_MODE_DEFAULT,
  // Provided by VK_EXT_line_rasterization
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_EXT = VK_LINE_RASTERIZATION_MODE_RECTANGULAR,
  // Provided by VK_EXT_line_rasterization
    VK_LINE_RASTERIZATION_MODE_BRESENHAM_EXT = VK_LINE_RASTERIZATION_MODE_BRESENHAM,
  // Provided by VK_EXT_line_rasterization
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT = VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH,
  // Provided by VK_KHR_line_rasterization
    VK_LINE_RASTERIZATION_MODE_DEFAULT_KHR = VK_LINE_RASTERIZATION_MODE_DEFAULT,
  // Provided by VK_KHR_line_rasterization
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR = VK_LINE_RASTERIZATION_MODE_RECTANGULAR,
  // Provided by VK_KHR_line_rasterization
    VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR = VK_LINE_RASTERIZATION_MODE_BRESENHAM,
  // Provided by VK_KHR_line_rasterization
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR = VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH,
} VkLineRasterizationMode;
```

or the equivalent

```c++
// Provided by VK_KHR_line_rasterization
typedef VkLineRasterizationMode VkLineRasterizationModeKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_line_rasterization
typedef VkLineRasterizationMode VkLineRasterizationModeEXT;
```

## [](#_description)Description

- `VK_LINE_RASTERIZATION_MODE_DEFAULT` is equivalent to `VK_LINE_RASTERIZATION_MODE_RECTANGULAR` if [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`strictLines` is `VK_TRUE`, otherwise lines are drawn as non-`strictLines` parallelograms. Both of these modes are defined in [Basic Line Segment Rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-basic).
- `VK_LINE_RASTERIZATION_MODE_RECTANGULAR` specifies lines drawn as if they were rectangles extruded from the line
- `VK_LINE_RASTERIZATION_MODE_BRESENHAM` specifies lines drawn by determining which pixel diamonds the line intersects and exits, as defined in [Bresenham Line Segment Rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-bresenham).
- `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH` specifies lines drawn if they were rectangles extruded from the line, with alpha falloff, as defined in [Smooth Lines](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-smooth).

## [](#_see_also)See Also

[VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html), [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html), [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineRasterizationModeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLineRasterizationMode).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0