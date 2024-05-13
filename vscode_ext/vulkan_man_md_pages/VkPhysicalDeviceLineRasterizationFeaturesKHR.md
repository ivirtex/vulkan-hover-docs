# VkPhysicalDeviceLineRasterizationFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceLineRasterizationFeaturesKHR - Structure describing the
line rasterization features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceLineRasterizationFeaturesKHR` structure is defined
as:

``` c
// Provided by VK_KHR_line_rasterization
typedef struct VkPhysicalDeviceLineRasterizationFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rectangularLines;
    VkBool32           bresenhamLines;
    VkBool32           smoothLines;
    VkBool32           stippledRectangularLines;
    VkBool32           stippledBresenhamLines;
    VkBool32           stippledSmoothLines;
} VkPhysicalDeviceLineRasterizationFeaturesKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_line_rasterization
typedef VkPhysicalDeviceLineRasterizationFeaturesKHR VkPhysicalDeviceLineRasterizationFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-rectangularLines"></span> `rectangularLines`
  indicates whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines"
  target="_blank" rel="noopener">rectangular line rasterization</a>.

- <span id="features-bresenhamLines"></span> `bresenhamLines` indicates
  whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-bresenham"
  target="_blank" rel="noopener">Bresenham-style line rasterization</a>.

- <span id="features-smoothLines"></span> `smoothLines` indicates
  whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-smooth"
  target="_blank" rel="noopener">smooth line rasterization</a>.

- <span id="features-stippledRectangularLines"></span>
  `stippledRectangularLines` indicates whether the implementation
  supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-stipple"
  target="_blank" rel="noopener">stippled line rasterization</a> with
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR` lines.

- <span id="features-stippledBresenhamLines"></span>
  `stippledBresenhamLines` indicates whether the implementation supports
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-stipple"
  target="_blank" rel="noopener">stippled line rasterization</a> with
  `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR` lines.

- <span id="features-stippledSmoothLines"></span> `stippledSmoothLines`
  indicates whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-stipple"
  target="_blank" rel="noopener">stippled line rasterization</a> with
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR` lines.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceLineRasterizationFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceLineRasterizationFeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceLineRasterizationFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceLineRasterizationFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceLineRasterizationFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLineRasterizationFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
