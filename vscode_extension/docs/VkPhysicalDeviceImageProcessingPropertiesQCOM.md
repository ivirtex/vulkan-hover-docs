# VkPhysicalDeviceImageProcessingPropertiesQCOM(3) Manual Page

## Name

VkPhysicalDeviceImageProcessingPropertiesQCOM - Structure containing image processing properties



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageProcessingPropertiesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_image_processing
typedef struct VkPhysicalDeviceImageProcessingPropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxWeightFilterPhases;
    VkExtent2D         maxWeightFilterDimension;
    VkExtent2D         maxBlockMatchRegion;
    VkExtent2D         maxBoxFilterBlockSize;
} VkPhysicalDeviceImageProcessingPropertiesQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxWeightFilterPhases` is the maximum value that **can** be specified for [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)::`numPhases` in [weight image sampling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage-filterphases) operations.
- []()`maxWeightFilterDimension` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) describing the largest dimensions (`width` and `height`) that **can** be specified for [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)::`filterSize`.
- []()`maxBlockMatchRegion` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) describing the largest dimensions (`width` and `height`) that **can** be specified for `blockSize` in [block matching](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-blockmatch) operations.
- []()`maxBoxFilterBlockSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) describing the maximum dimensions (`width` and `height`) that **can** be specified for `blocksize` in [box filter sampling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-boxfilter) operations.

## [](#_description)Description

If the `VkPhysicalDeviceImageProcessingPropertiesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the image processing information of a physical device.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageProcessingPropertiesQCOM-sType-sType)VUID-VkPhysicalDeviceImageProcessingPropertiesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_PROPERTIES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_image\_processing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageProcessingPropertiesQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0