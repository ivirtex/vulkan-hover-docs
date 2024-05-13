# VkImageViewSampleWeightCreateInfoQCOM(3) Manual Page

## Name

VkImageViewSampleWeightCreateInfoQCOM - Structure describing weight
sampling parameters for image view



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a `VkImageViewSampleWeightCreateInfoQCOM`
structure, then that structure includes a parameter specifying the
parameters for weight image views used in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-weightimage"
target="_blank" rel="noopener">weight image sampling</a>.

The `VkImageViewSampleWeightCreateInfoQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_image_processing
typedef struct VkImageViewSampleWeightCreateInfoQCOM {
    VkStructureType    sType;
    const void*        pNext;
    VkOffset2D         filterCenter;
    VkExtent2D         filterSize;
    uint32_t           numPhases;
} VkImageViewSampleWeightCreateInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `filterCenter` is a [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html) describing the
  location of the weight filter origin.

- `filterSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) specifying weight
  filter dimensions.

- `numPhases` is number of sub-pixel filter phases.

## <a href="#_description" class="anchor"></a>Description

The `filterCenter` specifies the origin or center of the filter kernel,
as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-weightimage-filteroperation"
target="_blank" rel="noopener">Weight Sampling Operation</a>. The
`numPhases` describes the number of sub-pixel filter phases as described
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-weightimage-filterphases"
target="_blank" rel="noopener">Weight Sampling Phases</a>.

Valid Usage

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06958"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06958"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06958  
  `filterSize.width` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-weightfilter-maxdimension"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceImageProcessingPropertiesQCOM</code>::<code>maxWeightFilterDimension.width</code></a>

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06959"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06959"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06959  
  `filterSize.height` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-weightfilter-maxdimension"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceImageProcessingPropertiesQCOM</code>::<code>maxWeightFilterDimension.height</code></a>

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06960"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06960"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06960  
  `filterCenter.x` **must** be less than or equal to (filterSize.width -
  1)

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06961"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06961"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06961  
  `filterCenter.y` **must** be less than or equal to
  (filterSize.height - 1)

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06962"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06962"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06962  
  `numPhases` **must** be a power of two squared value (i.e., 1, 4, 16,
  64, 256, etc.)

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06963"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06963"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06963  
  `numPhases` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-weightfilter-phases"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceImageProcessingPropertiesQCOM</code>::<code>maxWeightFilterPhases</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewSampleWeightCreateInfoQCOM-sType-sType"
  id="VUID-VkImageViewSampleWeightCreateInfoQCOM-sType-sType"></a>
  VUID-VkImageViewSampleWeightCreateInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_VIEW_SAMPLE_WEIGHT_CREATE_INFO_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_image_processing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewSampleWeightCreateInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
