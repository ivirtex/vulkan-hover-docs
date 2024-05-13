# VkLineRasterizationModeKHR(3) Manual Page

## Name

VkLineRasterizationModeKHR - Line rasterization modes



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)::`lineRasterizationMode`
are:

``` c
// Provided by VK_KHR_line_rasterization
typedef enum VkLineRasterizationModeKHR {
    VK_LINE_RASTERIZATION_MODE_DEFAULT_KHR = 0,
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR = 1,
    VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR = 2,
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR = 3,
    VK_LINE_RASTERIZATION_MODE_DEFAULT_EXT = VK_LINE_RASTERIZATION_MODE_DEFAULT_KHR,
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_EXT = VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR,
    VK_LINE_RASTERIZATION_MODE_BRESENHAM_EXT = VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR,
    VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT = VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR,
} VkLineRasterizationModeKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_line_rasterization
typedef VkLineRasterizationModeKHR VkLineRasterizationModeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_LINE_RASTERIZATION_MODE_DEFAULT_KHR` is equivalent to
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR` if
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`strictLines`
  is `VK_TRUE`, otherwise lines are drawn as non-`strictLines`
  parallelograms. Both of these modes are defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-basic"
  target="_blank" rel="noopener">Basic Line Segment Rasterization</a>.

- `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR` specifies lines drawn as
  if they were rectangles extruded from the line

- `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR` specifies lines drawn by
  determining which pixel diamonds the line intersects and exits, as
  defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-bresenham"
  target="_blank" rel="noopener">Bresenham Line Segment Rasterization</a>.

- `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR` specifies lines
  drawn if they were rectangles extruded from the line, with alpha
  falloff, as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-smooth"
  target="_blank" rel="noopener">Smooth Lines</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html),
[VK_KHR_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html),
[VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html),
[vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLineRasterizationModeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
