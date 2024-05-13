# VkPipelineRasterizationLineStateCreateInfoKHR(3) Manual Page

## Name

VkPipelineRasterizationLineStateCreateInfoKHR - Structure specifying
parameters of a newly created pipeline line rasterization state



## <a href="#_c_specification" class="anchor"></a>C Specification

Line segment rasterization options are controlled by the
[VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
structure.

The `VkPipelineRasterizationLineStateCreateInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_line_rasterization
typedef struct VkPipelineRasterizationLineStateCreateInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkLineRasterizationModeKHR    lineRasterizationMode;
    VkBool32                      stippledLineEnable;
    uint32_t                      lineStippleFactor;
    uint16_t                      lineStipplePattern;
} VkPipelineRasterizationLineStateCreateInfoKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_line_rasterization
typedef VkPipelineRasterizationLineStateCreateInfoKHR VkPipelineRasterizationLineStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `lineRasterizationMode` is a
  [VkLineRasterizationModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeKHR.html) value
  selecting the style of line rasterization.

- `stippledLineEnable` enables <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-stipple"
  target="_blank" rel="noopener">stippled line rasterization</a>.

- `lineStippleFactor` is the repeat factor used in stippled line
  rasterization.

- `lineStipplePattern` is the bit pattern used in stippled line
  rasterization.

## <a href="#_description" class="anchor"></a>Description

If `stippledLineEnable` is `VK_FALSE`, the values of `lineStippleFactor`
and `lineStipplePattern` are ignored.

Valid Usage

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02768"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02768"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02768  
  If `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rectangularLines"
  target="_blank" rel="noopener"><code>rectangularLines</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02769"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02769"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02769  
  If `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bresenhamLines"
  target="_blank" rel="noopener"><code>bresenhamLines</code></a> feature
  **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02770"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02770"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-02770  
  If `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-smoothLines"
  target="_blank" rel="noopener"><code>smoothLines</code></a> feature
  **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02771"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02771"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02771  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-stippledRectangularLines"
  target="_blank" rel="noopener"><code>stippledRectangularLines</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02772"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02772"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02772  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-stippledBresenhamLines"
  target="_blank" rel="noopener"><code>stippledBresenhamLines</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02773"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02773"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02773  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-stippledSmoothLines"
  target="_blank" rel="noopener"><code>stippledSmoothLines</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02774"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02774"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-stippledLineEnable-02774  
  If `stippledLineEnable` is `VK_TRUE` and `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_DEFAULT_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-stippledRectangularLines"
  target="_blank" rel="noopener"><code>stippledRectangularLines</code></a>
  feature **must** be enabled and
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`strictLines`
  **must** be `VK_TRUE`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-sType-sType"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-sType-sType"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_KHR`

- <a
  href="#VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-parameter"
  id="VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-parameter"></a>
  VUID-VkPipelineRasterizationLineStateCreateInfoKHR-lineRasterizationMode-parameter  
  `lineRasterizationMode` **must** be a valid
  [VkLineRasterizationModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeKHR.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html),
[VK_KHR_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkLineRasterizationModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationLineStateCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
