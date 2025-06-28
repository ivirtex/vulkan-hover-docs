# VkPipelineRasterizationLineStateCreateInfo(3) Manual Page

## Name

VkPipelineRasterizationLineStateCreateInfo - Structure specifying parameters of a newly created pipeline line rasterization state



## [](#_c_specification)C Specification

Line segment rasterization options are controlled by the [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html) structure.

The `VkPipelineRasterizationLineStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPipelineRasterizationLineStateCreateInfo {
    VkStructureType            sType;
    const void*                pNext;
    VkLineRasterizationMode    lineRasterizationMode;
    VkBool32                   stippledLineEnable;
    uint32_t                   lineStippleFactor;
    uint16_t                   lineStipplePattern;
} VkPipelineRasterizationLineStateCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_line_rasterization
typedef VkPipelineRasterizationLineStateCreateInfo VkPipelineRasterizationLineStateCreateInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_line_rasterization
typedef VkPipelineRasterizationLineStateCreateInfo VkPipelineRasterizationLineStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `lineRasterizationMode` is a [VkLineRasterizationMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationMode.html) value selecting the style of line rasterization.
- `stippledLineEnable` enables [stippled line rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-lines-stipple).
- `lineStippleFactor` is the repeat factor used in stippled line rasterization.
- `lineStipplePattern` is the bit pattern used in stippled line rasterization.

## [](#_description)Description

If `stippledLineEnable` is `VK_FALSE`, the values of `lineStippleFactor` and `lineStipplePattern` are ignored.

Valid Usage

- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-02768)VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-02768  
  If `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_RECTANGULAR`, then the [`rectangularLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rectangularLines) feature **must** be enabled
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-02769)VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-02769  
  If `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_BRESENHAM`, then the [`bresenhamLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bresenhamLines) feature **must** be enabled
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-02770)VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-02770  
  If `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH`, then the [`smoothLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-smoothLines) feature **must** be enabled
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02771)VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02771  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_RECTANGULAR`, then the [`stippledRectangularLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-stippledRectangularLines) feature **must** be enabled
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02772)VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02772  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_BRESENHAM`, then the [`stippledBresenhamLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-stippledBresenhamLines) feature **must** be enabled
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02773)VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02773  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH`, then the [`stippledSmoothLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-stippledSmoothLines) feature **must** be enabled
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02774)VUID-VkPipelineRasterizationLineStateCreateInfo-stippledLineEnable-02774  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_DEFAULT`, then the [`stippledRectangularLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-stippledRectangularLines) feature **must** be enabled and [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`strictLines` **must** be `VK_TRUE`

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-sType-sType)VUID-VkPipelineRasterizationLineStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO`
- [](#VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-parameter)VUID-VkPipelineRasterizationLineStateCreateInfo-lineRasterizationMode-parameter  
  `lineRasterizationMode` **must** be a valid [VkLineRasterizationMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationMode.html) value

## [](#_see_also)See Also

[VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html), [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkLineRasterizationMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationLineStateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0