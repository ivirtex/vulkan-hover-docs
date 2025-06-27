# VkPhysicalDeviceImageProcessing2PropertiesQCOM(3) Manual Page

## Name

VkPhysicalDeviceImageProcessing2PropertiesQCOM - Structure containing
image processing2 properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageProcessing2PropertiesQCOM` structure is
defined as:

``` c
// Provided by VK_QCOM_image_processing2
typedef struct VkPhysicalDeviceImageProcessing2PropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         maxBlockMatchWindow;
} VkPhysicalDeviceImageProcessing2PropertiesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-blockmatch-maxWindowExtent"></span>
  `maxBlockMatchWindow` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) describing
  the largest dimensions (`width` and `height`) that **can** be
  specified for the block match window.

If the `VkPhysicalDeviceImageProcessing2PropertiesQCOM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the image processing2 information of a physical
device.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceImageProcessing2PropertiesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceImageProcessing2PropertiesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceImageProcessing2PropertiesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_PROPERTIES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_image_processing2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing2.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageProcessing2PropertiesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
