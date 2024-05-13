# VkPhysicalDeviceImageProcessingPropertiesQCOM(3) Manual Page

## Name

VkPhysicalDeviceImageProcessingPropertiesQCOM - Structure containing
image processing properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageProcessingPropertiesQCOM` structure is defined
as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-weightfilter-phases"></span> `maxWeightFilterPhases`
  is the maximum value that **can** be specified for
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)::`numPhases`
  in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-weightimage-filterphases"
  target="_blank" rel="noopener">weight image sampling</a> operations.

- <span id="limits-weightfilter-maxdimension"></span>
  `maxWeightFilterDimension` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html)
  describing the largest dimensions (`width` and `height`) that **can**
  be specified for
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)::`filterSize`.

- <span id="limits-blockmatch-maxblocksize"></span>
  `maxBlockMatchRegion` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) describing
  the largest dimensions (`width` and `height`) that **can** be
  specified for `blockSize` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-blockmatch"
  target="_blank" rel="noopener">block matching</a> operations.

- <span id="limits-boxfilter-maxblocksize"></span>
  `maxBoxFilterBlockSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) describing
  the maximum dimensions (`width` and `height`) that **can** be
  specified for `blocksize` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-boxfilter"
  target="_blank" rel="noopener">box filter sampling</a> operations.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceImageProcessingPropertiesQCOM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the image processing information of a physical
device.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceImageProcessingPropertiesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceImageProcessingPropertiesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceImageProcessingPropertiesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_PROPERTIES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_image_processing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageProcessingPropertiesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
