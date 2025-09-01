# VkImageViewSampleWeightCreateInfoQCOM(3) Manual Page

## Name

VkImageViewSampleWeightCreateInfoQCOM - Structure describing weight sampling parameters for image view



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkImageViewSampleWeightCreateInfoQCOM` structure, then that structure includes a parameter specifying the parameters for weight image views used in [weight image sampling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage).

The `VkImageViewSampleWeightCreateInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_image_processing
typedef struct VkImageViewSampleWeightCreateInfoQCOM {
    VkStructureType    sType;
    const void*        pNext;
    VkOffset2D         filterCenter;
    VkExtent2D         filterSize;
    uint32_t           numPhases;
} VkImageViewSampleWeightCreateInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `filterCenter` is a [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html) describing the location of the weight filter origin.
- `filterSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) specifying weight filter dimensions.
- `numPhases` is number of sub-pixel filter phases.

## [](#_description)Description

The `filterCenter` specifies the origin or center of the filter kernel, as described in [Weight Sampling Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage-filteroperation). The `numPhases` describes the number of sub-pixel filter phases as described in [Weight Sampling Phases](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage-filterphases).

Valid Usage

- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06958)VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06958  
  `filterSize.width` **must** be less than or equal to [`VkPhysicalDeviceImageProcessingPropertiesQCOM`::`maxWeightFilterDimension.width`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-weightfilter-maxdimension)
- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06959)VUID-VkImageViewSampleWeightCreateInfoQCOM-filterSize-06959  
  `filterSize.height` **must** be less than or equal to [`VkPhysicalDeviceImageProcessingPropertiesQCOM`::`maxWeightFilterDimension.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-weightfilter-maxdimension)
- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06960)VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06960  
  `filterCenter.x` **must** be less than or equal to (filterSize.width - 1)
- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06961)VUID-VkImageViewSampleWeightCreateInfoQCOM-filterCenter-06961  
  `filterCenter.y` **must** be less than or equal to (filterSize.height - 1)
- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06962)VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06962  
  `numPhases` **must** be a power of two squared value (i.e., 1, 4, 16, 64, 256, etc.)
- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06963)VUID-VkImageViewSampleWeightCreateInfoQCOM-numPhases-06963  
  `numPhases` **must** be less than or equal to [`VkPhysicalDeviceImageProcessingPropertiesQCOM`::`maxWeightFilterPhases`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-weightfilter-phases)

Valid Usage (Implicit)

- [](#VUID-VkImageViewSampleWeightCreateInfoQCOM-sType-sType)VUID-VkImageViewSampleWeightCreateInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_SAMPLE_WEIGHT_CREATE_INFO_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_image\_processing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewSampleWeightCreateInfoQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0